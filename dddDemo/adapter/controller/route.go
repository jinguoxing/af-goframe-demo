package controller

import (
    "af-goframe-demo/dddDemo/adapter/controller/greeter"
    "github.com/gin-gonic/gin"
)


func Routes(r *gin.Engine){

    greeter.RegisterRoutes(r)

}
