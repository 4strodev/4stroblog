package core

import (
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
	Imports     []Module

	ExportSingletons []any
	ExportTransients []any
}

func (m *Module) init(container wiring.Container) error {
	m.container = extended.Derived(container)

	for _, resolver := range m.Singletons {
		err := m.container.Singleton(resolver)
		if err != nil {
			return err
		}
	}

	for _, resolver := range m.Transients {
		err := m.container.Transient(resolver)
		if err != nil {
			return err
		}
	}

	for _, module := range m.Imports {
		err := module.init(m.container)
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
