package main

import (
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/os/gcron"
	"github.com/gogf/gf/os/glog"
	_ "kingProject/router"
)

func main() {
	gcron.SetLogLevel(glog.LEVEL_ALL)
	g.Server().Run()
}
