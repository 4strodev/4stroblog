package main

import (
	"fmt"
	"log"
	"log/slog"
	"os"
	"strconv"

	"github.com/4strodev/4stroblog/site/server/api"
	"github.com/4strodev/4stroblog/site/server/core"
	"github.com/4strodev/4stroblog/site/server/site"
	"github.com/4strodev/4stroblog/site/shared"
	"github.com/4strodev/wiring_graphs/pkg/container"
)

var appModule = core.Module{
	Imports: []*core.Module{
		&shared.SharedModule,
		&site.SiteModule,
		&api.ApiModule,
	},
}

func main() {
	port, err := strconv.ParseInt(os.Getenv("PORT"), 0, 32)
	if err != nil {
		panic(fmt.Errorf("cannot parse PORT env: %w", err))
	}

	cont := container.New()

	s := core.Server{Wiring: cont}

	s.AddModule(appModule)

	err = s.Init()
	if err != nil {
		log.Fatal(err)
	}

	logger, err := container.Resolve[*slog.Logger](cont)
	if err != nil {
		log.Fatal("no logger resolved: ", err)
	}

	err = s.Start(int(port))
	if err != nil {
		logger.Error(err.Error())
	}
}
