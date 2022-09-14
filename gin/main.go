package main

import (
    "af-go-frame/core/logx/zapx"
    "af-go-frame/core/middleware/ginMiddleWare"
    "github.com/gin-gonic/gin"
    "github.com/zeromicro/go-zero/core/logx"
    "time"
)


func main(){

    r := gin.New()

   // logger, _ := zap.NewProduction()

    writer, err := zapx.NewZapWriter()
    logx.Must(err)
    logx.SetWriter(writer)


    r.Use(ginMiddleWare.GinZap(writer,time.RFC3339, false))
   // r.Use(ginMiddleWare.RecoveryWithZap(writer,true))


    r.GET("/ping", func(c *gin.Context) {
        c.JSON(200,gin.H{"message":"pong"})
    })

    // Example when panic happen.
    r.GET("/panic", func(c *gin.Context) {
        panic("An unexpected error happen!")
    })

    r.Run()

}
