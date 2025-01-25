package main

import (
	"log"
	"os"
	"strconv"

	"github.com/4strodev/4stroblog/site/server/api"
	"github.com/4strodev/4stroblog/site/server/core"
	"github.com/4strodev/4stroblog/site/server/site"
	"github.com/4strodev/4stroblog/site/shared"
	wiring "github.com/4strodev/wiring/pkg"
)

func main() {
	port, err := strconv.ParseInt(os.Getenv("PORT"), 0, 32)
	if err != nil {
		panic(err)
	}

	container := wiring.New()

	s := core.Server{Wiring: container}

	s.AddModule(shared.SharedModule)
	s.AddModule(site.SiteModule)
	s.AddModule(api.ApiModule)

	log.Fatal(s.Start(int(port)))
}
