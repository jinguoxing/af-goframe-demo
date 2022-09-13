package controller

import (
    "AF-Excel/adapter/controller/user"
    "github.com/gin-gonic/gin"
)

func Routes(r *gin.Engine){

   // container := r.GetContainer()
   // configService := container.MustMake(contract.ConfigKey).(contract.Config)


    user.RegisterRoutes(r)

}
