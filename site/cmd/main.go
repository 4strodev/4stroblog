package main

import (
	"log"

	"github.com/4strodev/4stroblog/site/server/api"
	"github.com/4strodev/4stroblog/site/server/core"
	"github.com/4strodev/4stroblog/site/server/site"
)

func main() {
	s := core.Server{}
	s.AddController(&api.ApiController{})
	s.AddController(&site.SiteController{})
	log.Fatal(s.Start(3000))
}
