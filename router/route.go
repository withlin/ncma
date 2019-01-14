package router

import (
	"github.com/WithLin/ncma/config"
	"github.com/WithLin/ncma/controller/album"
	"github.com/WithLin/ncma/controller/artist"
	"github.com/WithLin/ncma/controller/comment"
	"github.com/WithLin/ncma/controller/dj"
	"github.com/WithLin/ncma/controller/event"
	"github.com/WithLin/ncma/controller/fm"
	"github.com/WithLin/ncma/controller/follow"
	"github.com/WithLin/ncma/controller/like"
	"github.com/WithLin/ncma/controller/login"
	"github.com/WithLin/ncma/controller/lyric"
	"github.com/WithLin/ncma/controller/mv"
	"github.com/WithLin/ncma/controller/personal"
	"github.com/WithLin/ncma/controller/playlist"
	"github.com/WithLin/ncma/controller/recommend"
	"github.com/WithLin/ncma/controller/search"
	"github.com/WithLin/ncma/controller/send"
	"github.com/WithLin/ncma/controller/signin"
	"github.com/WithLin/ncma/controller/simi"
	"github.com/WithLin/ncma/controller/song"
	"github.com/WithLin/ncma/controller/top"
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

		//Comment
		api.GET("/comment/album",comment.CommentAlbum)
		api.GET("/comment/dj",comment.CommentDj)
		api.GET("/comment/hot",comment.CommentHot)
		api.GET("/comment/like",comment.CommentLike)
		api.GET("/comment/music",comment.CommentMusic)
		api.GET("/comment/mv",comment.CommentMV)
		api.GET("/comment/playlist",comment.CommentPlayList)
		api.GET("/comment/video",comment.CommentVidio)
		api.GET("/comment",comment.Comment)

		//signin
		api.GET("/daily_signin",signin.DailySignin)

		//dj
		api.GET("/dj/recommend",dj.DjRecommend)
		api.GET("/dj/catelist",dj.DjCatelist)
		api.GET("/dj/recommend/type",dj.DjRecommendType)
		api.GET("/dj/sub",dj.DjSub)
		api.GET("/dj/sublist",dj.DjSublist)
		api.GET("/dj/paygift",dj.DjPayGift)
		api.GET("/dj/detail",dj.DjDetail)
		api.GET("/dj/program",dj.DjProgram)
		api.GET("/dj/program/detail",dj.DjProgramDetail)
		api.GET("/dj/host",dj.DjHost)

		//event
		api.GET("/event",event.Event)

		//FmTrash
		api.GET("/fm_trash",fm.FmTrash)

		//follow
		api.GET("/follow",follow.Follow)

		//like
		api.GET("/like",like.Like)
		api.GET("/likelist",like.LikeList)

		//Lyric
		api.GET("/lyric",lyric.Lyric)

		//mv
		api.GET("/mv/detail",mv.MvDetail)
		api.GET("/mv/url",mv.MvUrl)
		api.GET("/mv/sub",mv.MvSub)
		api.GET("/mv/sublist",mv.MvList)
		api.GET("/mv/fist",mv.MvFist)

		//Personal
		api.GET("/personal_fm",personal.PersonalFm)
		api.GET("/personalized/mv",personal.PersonalizedMv)
		api.GET("/personalized/newsong",personal.PersonalizedNewsong)
		api.GET("/personalized/djprogram",personal.PersonalizedDjprogram)
		api.GET("/personalized/privatecontent",personal.PersonalizedPrivatecontent)
		api.GET("/personalized",personal.Personalized)

		//PlayList
		api.GET("/playlist/update",playlist.PlayListUpdate)
		api.GET("/playlist/catlist",playlist.PlayListCatlist)
		api.GET("/playlist/hot",playlist.PlayListHot)
		api.GET("/playlist/detail",playlist.PlayListDetail)
		api.GET("/playlist/create",playlist.PlayListCreate)
		api.GET("/playlist/subscribe",playlist.PlayListSubscribe)
		api.GET("/playlist/tracks",playlist.PlayListTracks)


		//recommend
		api.GET("/recommend/source",recommend.RecommendResource)
		api.GET("/recommend/songs",recommend.RecommendSongs)
		api.GET("/program/recommend",recommend.ProgramRecommend)

		//search
		api.GET("/search",search.Search)
		api.GET("/search/suggest",search.SearchSuggest)
		api.GET("/search/multimatch",search.SearchMultiMatch)
		api.GET("/search/hot",search.SearchHot)
		
		
		//send
		api.GET("/send/text",send.SendText)
		api.GET("/send/playlist",send.SendPlayList)

		//Simi
		api.GET("/simi/artist",simi.SimiArtist)
		api.GET("/simi/playlist",simi.SimiPlayList)
		api.GET("/simi/mv",simi.SimiMV)
		api.GET("/simi/song",simi.SimiSong)
		api.GET("/simi/user",simi.SimiUser)

		//song
		api.GET("/song/detail",song.SongDetail)
		api.GET("/song/url",song.SongUrl)

		//top
		api.GET("/top/playlist",top.TopPlayList)
		api.GET("/top/playlist/highquality",top.TopPlayListHighQuality)
		api.GET("/top/album",top.TopAlbum)
		api.GET("/top/artists",top.TopArtists)
		api.GET("/top/mv",top.TopMv)
		api.GET("/top/song",top.TopSong)
		api.GET("/toplist",top.TopList)
		api.GET("/toplist/artist",top.TopListArtist)
		api.GET("/toplist/detail",top.TopListDetail)
		api.GET("/top/list",top.TopArray)






	}
}