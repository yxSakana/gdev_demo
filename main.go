package main

import (
	"github.com/yxSakana/gdev_demo/internal/router"
	"github.com/yxSakana/gdev_demo/settings"
	"log"
	"time"
)

//go:generate go env -w GO111MODULE=on
//go:generate go env -w GOPROXY=https://goproxy.cn,direct
//go:generate go mod tidy
//go:generate go mod download

func main() {
	log.Printf("starting at 4 sec....")
	time.Sleep(4 * time.Second)
	r := router.InitRouter()
	r.Run(settings.Settings.Server.Address)
}
