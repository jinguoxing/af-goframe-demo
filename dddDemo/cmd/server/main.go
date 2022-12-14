package main

import (
	af_go_frame "af-go-frame"
	"af-go-frame/core/config"
	"af-go-frame/core/config/file"
	"af-go-frame/core/transport/rest"
	"af-goframe-demo/dddDemo/infrastructure/conf"
	"flag"
	"go.uber.org/zap"
)

var (
	Name = "af_ddd_demo"
	// Version is the version of the compiled software.
	Version = "1.0"

	flagconf string
)

func newApp(hs *rest.Server) *af_go_frame.App {

	return af_go_frame.New(
		af_go_frame.Name(Name),
		af_go_frame.Server(hs),
	)
}

func init() {
	flag.StringVar(&flagconf, "conf", "./config/config.yaml", "config path, eg: -conf config.yaml")
}

// @title 这里写标题 [*必须]
// @version 1.0 [*必须]
// @description 这里写描述信息
// @termsOfService http://swagger.io/terms/

// @contact.name 这里写联系人信息
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host 127.0.0.1:8000
// @BasePath /v1
func main() {
	flag.Parse()

	c := config.New(
		config.WithSource(
			file.NewSource(flagconf),
		),
	)
	if err := c.Load(); err != nil {
		panic(err)
	}

	var bc conf.Bootstrap
	if err := c.Scan(&bc); err != nil {
		panic(err)
	}

	//var lc Conf
	//if err := c.Scan(&lc); err != nil {
	//	panic(err)
	//}
	//fmt.Println(lc)

	logger, _ := zap.NewProduction()
	app, cleanup, err := InitApp(bc.Server, bc.Data, logger)
	if err != nil {
		panic(err)
	}
	defer cleanup()

	// start and wait for stop signal
	if err := app.Run(); err != nil {
		panic(err)
	}

}
