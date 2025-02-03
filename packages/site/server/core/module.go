package core

import (
	"errors"

	wiring "github.com/4strodev/wiring/pkg"
	"github.com/4strodev/wiring/pkg/extended"
)

// A module is used to contain services and controllers allowing to load them in batch instead of having to manually
// set resolvers on a container
type Module struct {
	container   wiring.Container
	Controllers []Controller
	Singletons  []any
	Transients  []any
	Imports     []*Module

	ExportSingletons []any
	ExportTransients []any
}

// initDependencies initialies the dependency graph of the modules. Then
// the container can be accessed to resolve global dependencies before starting
// controllers
func (m *Module) initDependencies(container wiring.Container) error {
	m.container = extended.Derived(container)

	for _, resolver := range m.Singletons {
		err := container.Singleton(resolver)
		if err != nil {
			return err
		}
	}

	for _, resolver := range m.Transients {
		err := container.Transient(resolver)
		if err != nil {
			return err
		}
	}

	for _, module := range m.Imports {
		err := module.initDependencies(m.container)
		if err != nil {
			return err
		}

		for _, resolver := range module.ExportSingletons {
			err := m.container.Singleton(resolver)
			if err != nil {
				return err
			}
		}

		for _, resolver := range module.ExportTransients {
			err := m.container.Transient(resolver)
			if err != nil {
				return err
			}
		}
	}

	for _, resolver := range m.ExportSingletons {
		err := container.Singleton(resolver)
		if err != nil {
			return err
		}
	}

	for _, resolver := range m.ExportTransients {
		err := container.Transient(resolver)
		if err != nil {
			return err
		}
	}

	return nil
}

// initControllers starts the controllers on the module graph
// if the dependency graph was not
func (m *Module) initControllers() error {
	if m.container == nil {
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
			return err
		}

		err = controller.Init(m.container)
		if err != nil {
			return err
		}
	}

	return nil
}
