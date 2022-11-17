package main

import (
	"log"
	"route69/config"
	proxy "route69/proxyManager"
)

func main() {
	conf, err := config.ReadInConfig()

	if err != nil {
		log.Fatalln(err)
	}

	//	conf.PrintRoutes()

	proxyMan := proxy.NewProxyManager(conf)

	proxyMan.Start()
}
