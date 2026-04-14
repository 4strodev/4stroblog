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
	"github.com/4strodev/4stroblog/site/shared/config"
	"github.com/4strodev/4stroblog/site/shared/db"
	"github.com/4strodev/4stroblog/site/shared/i18n"
	"github.com/4strodev/4stroblog/site/shared/logger"
	"github.com/4strodev/4stroblog/site/shared/s3"
	"github.com/4strodev/wiring_graphs/pkg/container"
)

func main() {
	port, err := strconv.ParseInt(os.Getenv("PORT"), 0, 32)
	if err != nil {
		panic(fmt.Errorf("cannot parse PORT env: %w", err))
	}

	cont := container.New()

	s := core.Server{
		Container: cont,
		Modules: []*core.Module{
			&api.ApiModule,
			&site.SiteModule,
		},
		GlobalSingletons: []any{
			db.NewDb,
			s3.NewS3Client,
			logger.NewLogger,
			config.GetConfig,
			func(conf config.Config) (*i18n.TranslationService, error) {
				return i18n.NewTranslationsService(conf.I18n.Folder)
			},
		},
	}

	err = s.Init()
	if err != nil {
		log.Panic(err)
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
