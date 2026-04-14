package core

import (
	"errors"
	"fmt"

	"github.com/4strodev/wiring_graphs/pkg/container"
)

// A module is used to contain services and controllers allowing to load them in batch instead of having to manually
// set resolvers on a container
type Module struct {
	container   *container.Container
	Controllers []Controller
	Singletons  []any
	Transients  []any
	Imports     []*Module

	// Human readable name for the module. Needed for debugging
	Name string
}

// initDependencies initializes the dependency graph of the modules. Then
// the container can be accessed to resolve global dependencies before starting
// controllers
func (m *Module) initDependencies(ctner *container.Container) error {
	var err error
	m.container = ctner

	// Always initialize first imported modules because they export resolvers that could be
	// necessary for this module
	for _, module := range m.Imports {
		err := module.initDependencies(ctner)
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

// initControllers starts the controllers on the module graph
// if the dependency graph was not
func (m *Module) initControllers() error {
	if m.container == nil {
		fmt.Println("module failed", m.Name)
		return errors.New("dependencies not initialized")
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

	for _, module := range m.Imports {
		err := module.initControllers()
		if err != nil {
			return err
		}
	}

	return nil
}
