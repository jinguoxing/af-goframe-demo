package controller

import (
	"af-go-frame/core/logx/zapx"
	"af-go-frame/core/middleware/ginMiddleWare"
	"af-go-frame/core/transport/rest"
	"af-goframe-demo/dddDemo/infrastructure/conf"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/zeromicro/go-zero/core/logx"
	"time"

	_ "af-goframe-demo/dddDemo/cmd/server/docs"
)

func NewHttpServer(c *conf.Server, r IRouter) *rest.Server {

	engine := NewHttpEngine(r)

	httpSrv := rest.NewServer(engine, rest.Address(c.Http.Addr))

	return httpSrv
}

// NewHttpEngine 创建了一个绑定了路由的Web引擎
func NewHttpEngine(r IRouter) *gin.Engine {
	// 设置为Release，为的是默认在启动中不输出调试信息
	gin.SetMode(gin.ReleaseMode)
	// 默认启动一个Web引擎
	app := gin.New()

	// 默认注册recovery中间件
	app.Use(gin.Recovery())

	writer, err := zapx.NewZapWriter()
	logx.Must(err)
	logx.SetWriter(writer)

	app.Use(ginMiddleWare.GinZap(writer, time.RFC3339, false))
	app.Use(ginMiddleWare.RecoveryWithZap(writer, true))

	app.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	// 业务绑定路由操作
	r.Register(app)

	// 返回绑定路由后的Web引擎
	return app
}
