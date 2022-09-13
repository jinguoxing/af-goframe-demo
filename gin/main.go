package main

import (
    "github.com/gin-gonic/gin"
    "go.uber.org/zap"
    )


func main(){

    r := gin.New()

    logger, _ := zap.NewProduction()

    r.Use(gin.Logger())


    r.GET("/ping", func(c *gin.Context) {
        c.JSON(200,gin.H{"message":"pong"})
    })

    r.Run()

}
