package middlewares

import (
	"log"
	"github.com/gin-gonic/gin"
	"time"
)

func Logger()gin.HandlerFunc{
	return func(c *gin.Context) {
		//开始时间
		start := time.Now()
		//处理请求
		c.Next()
		// 请求结束，计算耗时
        duration := time.Since(start)
		//获取请求信息
		method := c.Request.Method
        path   := c.Request.URL.Path
        status := c.Writer.Status()

		// 输出日志
        log.Printf("%s %s -> %d (%s)\n", method, path, status, duration)
	}
}