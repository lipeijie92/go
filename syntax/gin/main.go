package main

import (
	"basic/syntax/gin/routers"
	"github.com/gin-gonic/gin"
)

func main() {
	// 创建一个默认的路由引擎
	r := gin.Default()
	routers.ApiRoutersInit(r)
	routers.AdminRoutersInit(r)
	r.Run(":8080")

}
