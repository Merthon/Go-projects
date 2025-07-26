package main

import (
	"log"
	"net/http"
	"time"
)

var startTime time.Time
func main() {
	// 启动时间
	startTime = time.Now() 
	//监听端口
	addr := ":8080" 

	//调用自定义的路由初始化函数
	mux := SetupRouter() 

	//启动http服务
	log.Printf("服务正在运行http://localhost%s\n", addr)
	if err := http.ListenAndServe(addr, mux); err != nil {
		log.Fatalf("服务启动失败: %s\n", err)
	}

}