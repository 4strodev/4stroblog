package api

import (
	"github.com/4strodev/4stroblog/site/server/api/controllers/user"
	"github.com/4strodev/4stroblog/site/server/core"
)

var ApiModule = core.Module{
	Imports: []*core.Module{
		&user.UserApiModule,
	},
}
