package controller

import (
    "github.com/gin-gonic/gin"
)

// NewHttpEngine 创建了一个绑定了路由的Web引擎
func NewHttpEngine() (*gin.Engine) {
    // 设置为Release，为的是默认在启动中不输出调试信息
    gin.SetMode(gin.ReleaseMode)
    // 默认启动一个Web引擎
    r := gin.New()
    // 设置了Engine
  //  r.SetContainer(container)

    // 默认注册recovery中间件
    r.Use(gin.Recovery())

    // 业务绑定路由操作
    Routes(r)
    // 返回绑定路由后的Web引擎
    return r
}
