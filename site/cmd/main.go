package main

import (
	"log"
	"os"
	"strconv"

	"github.com/4strodev/4stroblog/site/modules"
	"github.com/4strodev/4stroblog/site/server/api"
	"github.com/4strodev/4stroblog/site/server/core"
	"github.com/4strodev/4stroblog/site/server/site"
	wiring "github.com/4strodev/wiring/pkg"
)

func main() {
	port, err := strconv.ParseInt(os.Getenv("PORT"), 0, 32)
	if err != nil {
		panic(err)
	}

	container := wiring.New()
	modules.LoadServices(container)
	s := core.Server{Wiring: container}
	s.AddController(&api.ApiController{})
	s.AddController(&site.SiteController{})
	log.Fatal(s.Start(int(port)))
}
