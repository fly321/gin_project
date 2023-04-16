package main

import (
	"flyblog/config"
	"flyblog/router"
)

func main() {
	r := router.Router()
	r.Run(config.AutoConfig{}.GetConfigByString("server.address")) // listen and serve on 0.0.0.0:8080
}
