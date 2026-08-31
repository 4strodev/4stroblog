package upload

import (
	"github.com/4strodev/4stroblog/site/features/uploads/application"
	"github.com/4strodev/4stroblog/site/features/uploads/domain"
	"github.com/4strodev/4stroblog/site/features/uploads/infrastructure"
	"github.com/4strodev/4stroblog/site/server/core"
	"gorm.io/gorm"
)

var UploadFeatureModule = core.Module{
	Singletons: []any{
		func(db *gorm.DB) domain.UploadsRepository {
			return &infrastructure.GormUploadsRepository{
				DB: db,
			}
		},
		application.NewUploadsService,
	},
	Name: "UploadFeatureModule",
}
