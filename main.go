package main

import (
	"log"
	"web/router"
)

func main() {
	r := router.InitRouter()
	if err := r.Run(); err != nil {
		log.Fatal("服务器启动失败...")
	}
}
