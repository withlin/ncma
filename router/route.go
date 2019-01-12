package router

import (
	"github.com/WithLin/ncma/config"
	"github.com/WithLin/ncma/controller/album"
	"github.com/WithLin/ncma/controller/login"
	"github.com/WithLin/ncma/controller/user"
	"github.com/gin-gonic/gin"
)

// Route
func Route(router *gin.Engine) {
	apiPrefix := config.Conf.GoConf.APIPrefix

	api := router.Group(apiPrefix)
	{
		//album
		api.GET("/album/:id",album.Album)
		api.GET("/login/cellphone/",login.LoginCellPhone)
		//login
		api.GET("/login/",login.LoginEmail)
		api.GET("/login/refresh",login.LoginRefresh)
		api.GET("/login/status",login.LoginStatus)
		api.GET("/logout",login.LogOut)
		//User
		api.GET("/user/detail",user.UserDetail)
		api.GET("/user/subcount",user.UserSubCount)
	}
}