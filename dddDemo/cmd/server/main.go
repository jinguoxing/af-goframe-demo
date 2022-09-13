package main

import (
    "AF-Excel/adapter/controller"
    af_go_frame "af-go-frame"
    //"af-go-frame/core/container"
    "af-go-frame/core/transport/rest"
    "fmt"

    // "os"
   // "os/signal"
   // "syscall"
   // "time"
   //"net/http"
   // "context"
   // "log"
)


var (

    Name ="af_ddd_demo"
    // Version is the version of the compiled software.
    Version = "1.0"
)

func newApp(hs *rest.Server) *af_go_frame.App{

    return af_go_frame.New(
            af_go_frame.Name(Name),
            af_go_frame.Server(hs),
        )
}

func main(){

   // c := container.NewAFContainer()



     engine := controller.NewHttpEngine()

       // c.Bind(&httpx.AFKernelProvider{HttpEngine: engine})

   // gin.SetMode(gin.ReleaseMode)
    // 默认启动一个Web引擎
    //engine := gin.New()
    //// 设置了Engine
    ////r.SetContainer(container)
    //
    //// 默认注册recovery中间件
    //engine.Use(gin.Recovery())
    //
    //// 业务绑定路由操作
    ////Routes(engine)
    //controller.Routes(engine)

       httpSrv :=  rest.NewServer(engine,rest.Address(":8000"))


       app :=  af_go_frame.New(
            af_go_frame.Name(Name),
            af_go_frame.Server(httpSrv),
        )

        if err := app.Run(); err != nil {
            fmt.Errorf("error:",err)
        }

}

