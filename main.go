package main

import (
	"github.com/gin-gonic/gin"
	_ "granada1/init"
	"granada1/server"
)

func main() {
	// 启动 http 服务
	r := gin.New()
	r=server.InitApiRouter(r)
	err := r.Run(":8080")
	if err != nil {
		return 
	}
}