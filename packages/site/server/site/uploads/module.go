package uploads

import (
	"github.com/4strodev/4stroblog/site/server/core"
)

var UploadsModule = core.Module{
	Controllers: []core.Controller{
		&SiteUploadsController{},
	},
}
