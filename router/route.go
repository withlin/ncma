package router

import (
	"github.com/WithLin/ncma/config"
	"github.com/WithLin/ncma/controller/album"
	"github.com/gin-gonic/gin"
)

// Route
func Route(router *gin.Engine) {
	apiPrefix := config.Conf.GoConf.APIPrefix

	api := router.Group(apiPrefix)
	{
		api.GET("/album/:id",album.Album)
	}
}