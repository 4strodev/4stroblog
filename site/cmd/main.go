package main

import (
	"log"
	"os"
	"strconv"

	"github.com/4strodev/4stroblog/site/server/api"
	"github.com/4strodev/4stroblog/site/server/core"
	"github.com/4strodev/4stroblog/site/server/site"
)

func main() {
	port, err := strconv.ParseInt(os.Getenv("PORT"), 0, 32)
	if err != nil {
		panic(err)
	}
	s := core.Server{}
	s.AddController(&api.ApiController{})
	s.AddController(&site.SiteController{})
	log.Fatal(s.Start(int(port)))
}
