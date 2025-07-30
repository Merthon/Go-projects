package main

import (
	"log"
	"time"

	"github.com/gin-gonic/gin"
)

var startTime time.Time
func main() {
	// 启动时间
	startTime = time.Now() 

	//创建Gin
	r := gin.New()
	//使用中间件
	r.Use(gin.Logger(), gin.Recovery())
	//加载模版
	r.LoadHTMLGlob("templates/*")
	// 静态资源
    r.Static("/static", "./static")

	//路由注册
    registerRoutes(r)
	
	//监听端口
	addr := ":8080" 

	// //调用自定义的路由初始化函数
	// mux := SetupRouter() 

	//启动服务
	log.Printf("服务正在运行 %s", addr)
    if err := r.Run(addr); err != nil {
        log.Fatalf("Server failed: %v", err)
    }

}