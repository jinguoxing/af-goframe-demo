package controller

import (
    "af-goframe-demo/dddDemo/adapter/controller/greeter"
    "github.com/gin-gonic/gin"
    "github.com/google/wire"
)


var _ IRouter = (*Router)(nil)

var RouterSet = wire.NewSet(wire.Struct(new(Router), "*"), wire.Bind(new(IRouter), new(*Router)))

type IRouter interface {

    Register (r *gin.Engine) error

}

type Router struct {

    GreeterApi *greeter.GreeterService
}

func (a *Router)Register(r *gin.Engine)error{

    a.RegisterApi(r)
    return nil
}


func (a *Router)RegisterApi(r *gin.Engine){


    r.GET("/greeter/sayhello", a.GreeterApi.SayHello)

}