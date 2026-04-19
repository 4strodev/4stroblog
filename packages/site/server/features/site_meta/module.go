package sitemeta

import (
	"github.com/4strodev/4stroblog/site/features/site_meta/domain"
	"github.com/4strodev/4stroblog/site/features/site_meta/infrastructure"
	"github.com/4strodev/4stroblog/site/server/core"
	"gorm.io/gorm"
)

var SiteMetaFeatureModule = core.Module{
	Singletons: []any{
		func(db *gorm.DB) domain.SiteMetaRepository {
			return infrastructure.GormSiteMetaRepository{DB: db}
		},
	},
	Name: "SiteMetaFeatureModule",
}
