package user

import (
	//"AF-Excel/interface/user"
	"github.com/gin-gonic/gin"
)

type UserApi struct{}

// RegisterRoutes 注册路由
func RegisterRoutes(r *gin.Engine) error {
	api := &UserApi{}
	//if !r.IsBind(user.UserKey) {
	//    r.Bind(&UserProvider{})
	//}

	// 登录
	r.GET("/user/login", api.Login)
	//// 登出
	//r.GET("/user/logout", auth.AuthMiddleware(), api.Logout)
	// 注册
	//r.POST("/user/register", api.Register)
	// 获取全部用户列表
	//r.GET("/user/list", api.ListAllUser)
	//// 注册验证
	//r.GET("/user/register/verify", api.Verify)

	return nil
}
