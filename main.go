package main

import (
	"im/conf"
	"im/router"
)

func main() {
	conf.Init()
	r := router.NewRouter()
	r.Run(conf.HttpPort)
}
