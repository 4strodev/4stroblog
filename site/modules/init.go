package modules

import (
	"github.com/4strodev/4stroblog/site/modules/session/application"
	"github.com/4strodev/4stroblog/site/shared/config"
	"github.com/4strodev/4stroblog/site/shared/db"
	"github.com/4strodev/wiring/pkg"
)

func LoadServices(container pkg.Container) {
	var err error

	err = container.Singleton(db.GetDbInstance)
	if err != nil {
		panic(err)
	}
	err = container.Singleton(config.GetConfig)
	if err != nil {
		panic(err)
	}
	err = container.Transient(application.NewSessionService)
	if err != nil {
		panic(err)
	}
}
