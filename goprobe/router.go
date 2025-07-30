package main

import (
	"goprobe/probe"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

// gatherData 收集所有探针数据
func gatherData(c *gin.Context) gin.H {
    sysInfo, _ := probe.GetSystemInfo()
    memInfo := probe.GetMemoryInfo()
    diskInfo, _ := probe.GetDiskInfo()
    uptime := time.Since(startTime).Round(time.Second).String()

    return gin.H{
        "System":    sysInfo,
        "Memory":    memInfo,
        "Disk":      diskInfo,
        "Uptime":    uptime,
        "ClientIP":  c.ClientIP(),
        "UserAgent": c.Request.UserAgent(),
    }
}

// registerRoutes 统一路由注册
func registerRoutes(r *gin.Engine) {
    // 首页 HTML 渲染
    r.GET("/", func(c *gin.Context) {
        data := gatherData(c)
        // 渲染主模板 base.html
        c.HTML(http.StatusOK, "base", data)
    })

    // JSON 接口
    r.GET("/api/status", func(c *gin.Context) {
        c.JSON(http.StatusOK, gatherData(c))
    })
}

/* // LoggingMiddleware 是简单的请求日志中间件
func LoggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("[%s] %s %s from %s\n", time.Now().Format("2006-01-02 15:04:05"), r.Method, r.URL.Path, r.RemoteAddr)
		next.ServeHTTP(w, r)
	})
}

// SetupRouter 返回一个 http.Handler，用于注册所有路由
func SetupRouter() http.Handler {
	mux := http.NewServeMux()
	// 注册首页路由
	mux.HandleFunc("/", IndexHandler)
	return LoggingMiddleware(mux) // 用中间件包裹
}

// IndexHandler 主页处理
func IndexHandler(w http.ResponseWriter, r *http.Request) {
	system, err := probe.GetSystemInfo()
	if err != nil {
		http.Error(w, "获取系统信息失败", http.StatusInternalServerError)
		return
	}

	memory := probe.GetMemoryInfo()

	disk, err := probe.GetDiskInfo()
	if err != nil {
		http.Error(w, "获取磁盘信息失败", http.StatusInternalServerError)
		return
	}

	// 计算运行时长
	uptime := time.Since(startTime).Round(time.Second).String()

	// 获取客户端 IP（注意可能有代理）
	clientIP := r.RemoteAddr
	if forwarded := r.Header.Get("X-Forwarded-For"); forwarded != "" {
		clientIP = forwarded
	}

	userAgent := r.UserAgent()

	data := IndexData{
		System: system,
		Memory: memory,
		Disk:   disk,
		Uptime:    uptime,
		ClientIP:  clientIP,
		UserAgent: userAgent,
	}

    tmpl, err := template.ParseFiles("templates/base.html", "templates/index.html")
		if err != nil {
			http.Error(w, "模板解析失败", http.StatusInternalServerError)
			return
		}
    tmpl.ExecuteTemplate(w, "base", data)
} */