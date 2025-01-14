package main

import (
	"log"
	"os"
	"strconv"

	"github.com/4strodev/4stroblog/site/modules/session/application"
	"github.com/4strodev/4stroblog/site/server/api"
	"github.com/4strodev/4stroblog/site/server/core"
	"github.com/4strodev/4stroblog/site/server/site"
	"github.com/4strodev/4stroblog/site/shared/config"
	"github.com/4strodev/4stroblog/site/shared/db"
	wiring "github.com/4strodev/wiring/pkg"
)

func main() {
	port, err := strconv.ParseInt(os.Getenv("PORT"), 0, 32)
	if err != nil {
		panic(err)
	}

	container := wiring.New()
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

	s := core.Server{Wiring: container}
	s.AddController(&api.ApiController{})
	s.AddController(&site.SiteController{})
	log.Fatal(s.Start(int(port)))
}
