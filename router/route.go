package router

import (
	"github.com/WithLin/ncma/config"
	"github.com/WithLin/ncma/controller/album"
	"github.com/WithLin/ncma/controller/artist"
	"github.com/WithLin/ncma/controller/login"
	"github.com/WithLin/ncma/controller/user"
	"github.com/WithLin/ncma/controller/video"
	"github.com/gin-gonic/gin"
)

// Route
func Route(router *gin.Engine) {
	apiPrefix := config.Conf.GoConf.APIPrefix

	api := router.Group(apiPrefix)
	{
		//album
		api.GET("/album/:id",album.Album)

		//login
		api.GET("/login/cellphone/",login.LoginCellPhone)
		api.GET("/login/",login.LoginEmail)
		api.GET("/login/refresh",login.LoginRefresh)
		api.GET("/login/status",login.LoginStatus)
		api.GET("/logout",login.LogOut)

		//User
		api.GET("/user/detail",user.UserDetail)
		api.GET("/user/subcount",user.UserSubCount)
		api.GET("/user/audio",user.UserAudio)
		api.GET("/user/cloud",user.UserCloud)
		api.GET("/user/cloudsearch",user.UserCloudSearch)
		api.GET("/user/dj",user.UserDj)
		api.GET("/user/event",user.UserEvent)
		api.GET("/user/followeds",user.UserFolloweds)
		api.GET("/user/follows",user.UserFollows)
		api.GET("/user/record",user.UserRecord)
		api.GET("/user/update",user.UserUpdate)
		api.GET("/user/playlist",user.UserPlayList)

		//video
		api.GET("/video/detail",video.VedioDetail)
		api.GET("/video/sub",video.VedioSub)
		api.GET("/video/url",video.VedioUrl)

		//artist
		api.GET("/artist/mv",artist.ArtistMv)
		api.GET("/artist/album",artist.ArtistAlbum)
		api.GET("/artist",artist.Artist)
		api.GET("/artist/list",artist.ArtistList)
		api.GET("/artist/sub",artist.ArtistSub)
		api.GET("/artist/sublist",artist.ArtistSubList)
		api.GET("/artists",artist.Artists)
	}
}