package core

import (
	"errors"
	"fmt"

	"github.com/4strodev/wiring_graphs/pkg/container"
	"github.com/gofiber/fiber/v3"
)

// A module is used to contain services and controllers allowing to load them in batch instead of having to manually
// set resolvers on a container
type Module struct {
	container   *container.Container
	Controllers []Controller
	Singletons  []any
	Transients  []any
	Imports     []*Module

	// Router optionally overrides the router used by this module's controllers.
	// It receives the router of the container given to the module and the module's
	// own container (to resolve dependencies such as middlewares). The returned router
	// is what controllers get when resolving fiber.Router. It is resolved lazily, the
	// first time a controller asks for it. When nil, the parent router is used.
	// Imported modules share the importer's container, so set Router only on
	// top-level modules.
	Router func(router fiber.Router, cont *container.Container) (fiber.Router, error)

	// Human readable name for the module. Needed for debugging
	Name string
}

// initDependencies initializes the dependency graph of the modules. Then
// the container can be accessed to resolve global dependencies before starting
// controllers
func (m *Module) initDependencies(ctner *container.Container) error {
	var err error
	m.container = ctner

	if m.Router != nil {
		// Override the router in a derived container, used by this module and its imports
		m.container = ctner.Derived()
		err = m.container.Singleton(func() (fiber.Router, error) {
			parentRouter, err := container.Resolve[fiber.Router](ctner)
			if err != nil {
				return nil, fmt.Errorf("error resolving parent router for %s: %w", m.Name, err)
			}
			router, err := m.Router(parentRouter, m.container)
			if err != nil {
				return nil, fmt.Errorf("error setting up router for %s: %w", m.Name, err)
			}
			return router, nil
		})
		if err != nil {
			return fmt.Errorf("error setting up router for %s: %w", m.Name, err)
		}
	}

	// Always initialize first imported modules because they export resolvers that could be
	// necessary for this module
	for _, module := range m.Imports {
		err := module.initDependencies(m.container)
		if err != nil {
			return err
		}
	}

	err = m.container.Singleton(m.Singletons...)
	if err != nil {
		return fmt.Errorf("error setting up Singleton dependencies for %s: %w", m.Name, err)
	}

	err = m.container.Dependencies(m.Transients...)
	if err != nil {
		return fmt.Errorf("error setting up Transient dependencies for %s: %w", m.Name, err)
	}

	return nil
}

// initControllers starts the controllers on the module graph.
// Imported modules go first so their specific routes are registered before
// the catch-all routes (e.g. "/site/*") of the modules that import them.
// Fiber matches routes in registration order.
func (m *Module) initControllers() error {
	if m.container == nil {
		fmt.Println("module failed", m.Name)
		return errors.New("dependencies not initialized")
	}

	for _, module := range m.Imports {
		err := module.initControllers()
		if err != nil {
			return err
		}
	}

	for _, controller := range m.Controllers {
		err := m.container.Fill(controller)
		if err != nil {
			return fmt.Errorf("error initializing controllers for %s: %w", m.Name, err)
		}

		err = controller.Init(m.container)
		if err != nil {
			return fmt.Errorf("error initializing controllers for %s: %w", m.Name, err)
		}
	}

	return nil
}
