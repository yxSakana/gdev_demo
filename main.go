package main

import (
	"github.com/yxSakana/gdev_demo/internal/router"
	"github.com/yxSakana/gdev_demo/settings"
)

func main() {
	r := router.InitRouter()
	r.Run(settings.Settings.Server.Address)
}
