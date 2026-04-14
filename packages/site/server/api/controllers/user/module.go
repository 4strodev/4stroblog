package user

import (
	"github.com/4strodev/4stroblog/site/server/core"
	"github.com/4strodev/4stroblog/site/server/features/user"
)

var UserApiModule = core.Module{
	Controllers: []core.Controller{&UserController{}},
	Imports: []*core.Module{
		&user.UserFeatureModule,
	},
	Name: "UserApiModule",
}
