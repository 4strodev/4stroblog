package uploads

import (
	"github.com/4strodev/4stroblog/site/server/core"
	"github.com/4strodev/4stroblog/site/server/features/upload"
)

var SiteUploadsModule = core.Module{
	Imports: []*core.Module{
		&upload.UploadFeatureModule,
	},
	Controllers: []core.Controller{
		&SiteUploadsController{},
	},
	Name: "SiteUploadsModule",
}
