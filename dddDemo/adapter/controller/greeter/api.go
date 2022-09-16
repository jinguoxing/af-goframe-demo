package greeter

import (

    "github.com/gin-gonic/gin"
)

// RegisterRoutes 注册路由
func RegisterRoutes(r *gin.Engine) error {

    api := &GreeterService{}

    r.GET("/greeter/sayhello", api.SayHello)


}
