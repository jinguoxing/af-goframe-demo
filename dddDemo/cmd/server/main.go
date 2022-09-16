package main

import (
    af_go_frame "af-go-frame"
    "af-go-frame/core/config"
    "af-go-frame/core/config/file"
    "af-goframe-demo/dddDemo/infrastructure/conf"
    "flag"

    "af-go-frame/core/transport/rest"
)


var (

    Name ="af_ddd_demo"
    // Version is the version of the compiled software.
    Version = "1.0"

    flagconf string
)

func newApp(hs *rest.Server) *af_go_frame.App{

    return af_go_frame.New(
            af_go_frame.Name(Name),
            af_go_frame.Server(hs),
        )
}

func init() {
    flag.StringVar(&flagconf, "conf", "./config/config.yaml", "config path, eg: -conf config.yaml")
}

func main(){
    flag.Parse()
   // c := container.NewAFContainer()

    c := config.New(
        config.WithSource(
            file.NewSource(flagconf),
        ),
    )
    if err := c.Load(); err != nil {
        panic(err)
    }

    var bc  conf.Bootstrap

    if err := c.Scan(&bc); err != nil {
        panic(err)
    }



    app,  cleanup, err := InitApp(bc.Server, bc.Data)
    if err != nil {
       panic(err)
    }
    defer cleanup()

    // start and wait for stop signal
    if err := app.Run(); err != nil {
       panic(err)
    }

}

