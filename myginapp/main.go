/* 应用入口，负责加载路由、中间件，并启动服务 */
package main

import (
    "github.com/Merthon/myginapp/routes"
)

func main() {
    r := routes.SetupRouter()
    // r.Run(":8080") // 默认监听 0.0.0.0:8080
    r.Run() // 也可不传端口参数
}
