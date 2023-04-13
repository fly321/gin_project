package main

import (
	"flyblog/router"
)

func main() {
	r := router.Router()
	r.Run() // listen and serve on 0.0.0.0:8080
}
