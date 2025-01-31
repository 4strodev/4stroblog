package shared

import (
	"github.com/4strodev/4stroblog/site/server/core"
	"github.com/4strodev/4stroblog/site/shared/config"
	"github.com/4strodev/4stroblog/site/shared/db"
	"github.com/4strodev/4stroblog/site/shared/logger"
	"github.com/4strodev/4stroblog/site/shared/s3"
)

var SharedModule = core.Module{
	Singletons: []any{
		config.GetConfig,
		logger.NewLogger,
		db.NewDb,
	},
	ExportSingletons: []any {
		db.NewDb,
		s3.NewS3Client,
		logger.NewLogger,
	},
}
