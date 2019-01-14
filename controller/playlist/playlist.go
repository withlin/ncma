package playlist

import (
	"fmt"
	"github.com/WithLin/ncma/utils"
	"github.com/gin-gonic/gin"
	"log"
)

// 全部歌单分类
func PlayListCatlist(c *gin.Context) {
	params, encSecKey, encErr := utils.Encrypt(fmt.Sprintf(`{"csrf_token":""}`))
	if encErr != nil {
		log.Println(encErr)
	}
	res, resErr := utils.DoPostRequest(c.Request.Cookies(),"https://music.163.com/weapi/playlist/catalogue", params, encSecKey)
	if resErr != nil {
		log.Println(resErr)
	}
	c.String(200, res)
}



// 创建歌单
func PlayListCreate(c *gin.Context) {
	name :=c.Query("name")
	params, encSecKey, encErr := utils.Encrypt(fmt.Sprintf(`{"name":"%s","csrf_token":""}`,name))
	if encErr != nil {
		log.Println(encErr)
	}
	res, resErr := utils.DoPostRequest(c.Request.Cookies(),"https://music.163.com/weapi/playlist/create", params, encSecKey)
	if resErr != nil {
		log.Println(resErr)
	}
	c.String(200, res)
}




// 歌单详情
func PlayListDetail(c *gin.Context) {
	id :=c.Query("id")
	s := c.DefaultQuery("s","8")
	params, encSecKey, encErr := utils.Encrypt(fmt.Sprintf(`{"id":"%s","s":%s,"n":10000,"csrf_token":""}`,id,s))
	if encErr != nil {
		log.Println(encErr)
	}
	res, resErr := utils.DoPostRequest(c.Request.Cookies(),"https://music.163.com/weapi/v3/playlist/detail", params, encSecKey)
	if resErr != nil {
		log.Println(resErr)
	}
	c.String(200, res)
}


// 热门歌单分类
func PlayListHot(c *gin.Context) {

	params, encSecKey, encErr := utils.Encrypt(fmt.Sprintf(`{"csrf_token":""}`))
	if encErr != nil {
		log.Println(encErr)
	}
	res, resErr := utils.DoPostRequest(c.Request.Cookies(),"https://music.163.com/weapi/playlist/hottags", params, encSecKey)
	if resErr != nil {
		log.Println(resErr)
	}
	c.String(200, res)
}



// 收藏与取消收藏歌单
func PlayListSubscribe(c *gin.Context) {
	t :=c.Query("t")
	if t=="1" {
		t = "subscribe"
	}else {
		t = "unsubscribe"
	}

	id :=c.Query("id")
	params, encSecKey, encErr := utils.Encrypt(fmt.Sprintf(`{"t":"%s","id":"%s","csrf_token":""}`,id))
	if encErr != nil {
		log.Println(encErr)
	}
	res, resErr := utils.DoPostRequest(c.Request.Cookies(),"https://music.163.com/weapi/playlist/"+t, params, encSecKey)
	if resErr != nil {
		log.Println(resErr)
	}
	c.String(200, res)
}


// 收藏单曲到歌单 从歌单删除歌曲
func PlayListTracks(c *gin.Context) {

	op :=c.Query("op")
	pid :=c.Query("pid")
	trackIds :=fmt.Sprintf(`[%s]`,c.Query("tracks"))
	params, encSecKey, encErr := utils.Encrypt(fmt.Sprintf(`{"op":"%s","pid":"%s","trackIds":"%s","csrf_token":""}`,
		op,pid,trackIds))
	if encErr != nil {
		log.Println(encErr)
	}
	res, resErr := utils.DoPostRequest(c.Request.Cookies(),"https://music.163.com/weapi/playlist/manipulate/tracks", params, encSecKey)
	if resErr != nil {
		log.Println(resErr)
	}
	c.String(200, res)
}


// 编辑歌单
func PlayListUpdate(c *gin.Context) {

	desc :=c.DefaultQuery("desc","")
	tags :=c.DefaultQuery("tags","")
	id :=c.Query("id")
	name :=c.Query("name")
	descUpdate :=fmt.Sprintf(`"/api/playlist/desc/update":"{"id":%s,"desc":"%s")}`,
		id,desc)
	tagsUpdate :=fmt.Sprintf(`"/api/playlist/tags/update":"{"id":%s,"tags":"%s")}`,
		id,tags)
	nameUpdate :=fmt.Sprintf(`"/api/playlist/update/name":"{"id":%s,"name":"%s")}`,
		id,name)


	params, encSecKey, encErr := utils.Encrypt(fmt.Sprintf(`{%s,%s,%s,"csrf_token":""}`,
		descUpdate,tagsUpdate,nameUpdate))
	if encErr != nil {
		log.Println(encErr)
	}
	res, resErr := utils.DoPostRequest(c.Request.Cookies(),"https://music.163.com/weapi/batch", params, encSecKey)
	if resErr != nil {
		log.Println(resErr)
	}
	c.String(200, res)
}