package artist

import (
	"fmt"
	"github.com/WithLin/ncma/utils"
	"github.com/gin-gonic/gin"
	"log"
	"strings"
)

// 歌手专辑列表
func ArtistAlbum(c *gin.Context){
	uid :=c.Query("uid")
	limit := c.DefaultQuery("limit","30")
	offset := c.DefaultQuery("offset","0")
	params, encSecKey, encErr := utils.Encrypt(fmt.Sprintf(`{"uid":"%s","limit":"%s","offset":"%s","total":true,"csrf_token":""}`,uid,limit,offset))
	if encErr != nil {
		log.Println(encErr)
	}
	res, resErr := utils.DoPostRequest(c.Request.Cookies(),"https://music.163.com/weapi/artist/albums/"+uid, params, encSecKey)
	if resErr != nil {
		log.Println(resErr)
	}
	c.String(200,res)
}



// 歌手介绍
func Artist(c *gin.Context){
	uid :=c.Query("id")
	params, encSecKey, encErr := utils.Encrypt(fmt.Sprintf(`{"id":"%s","csrf_token":""}`,uid))
	if encErr != nil {
		log.Println(encErr)
	}
	res, resErr := utils.DoPostRequest(c.Request.Cookies(),"https://music.163.com/weapi/artist/introduction", params, encSecKey)
	if resErr != nil {
		log.Println(resErr)
	}
	c.String(200,res)
}



// 歌手分类

/*
    categoryCode 取值
    入驻歌手 5001
    华语男歌手 1001
    华语女歌手 1002
    华语组合/乐队 1003
    欧美男歌手 2001
    欧美女歌手 2002
    欧美组合/乐队 2003
    日本男歌手 6001
    日本女歌手 6002
    日本组合/乐队 6003
    韩国男歌手 7001
    韩国女歌手 7002
    韩国组合/乐队 7003
    其他男歌手 4001
    其他女歌手 4002
    其他组合/乐队 4003

    initial 取值 a-z/A-Z
*/
func ArtistList(c *gin.Context){
	uid :=c.DefaultQuery("cat","1001")
	initial := c.DefaultQuery("initial","")
	offset := c.DefaultQuery("offset","0")
	limit := c.DefaultQuery("limit","30")
	params, encSecKey, encErr := utils.Encrypt(fmt.Sprintf(`{"uid":"%s","initial":"%s","limit":"%s","offset":"%s","total":true,"csrf_token":""}`,
		uid,strings.ToUpper(initial),limit,offset))
	if encErr != nil {
		log.Println(encErr)
	}
	res, resErr := utils.DoPostRequest(c.Request.Cookies(),"https://music.163.com/weapi/artist/list", params, encSecKey)
	if resErr != nil {
		log.Println(resErr)
	}
	c.String(200,res)
}



// 歌手相关MV
func ArtistMv(c *gin.Context){

	artistId :=c.Query("id")
	limit := c.DefaultQuery("limit","30")
	offset := c.DefaultQuery("offset","0")
	params, encSecKey, encErr := utils.Encrypt(fmt.Sprintf(`{"artistId":"%s","limit":"%s","offset":"%s","total":true,"csrf_token":""}`,
		artistId,limit,offset))
	if encErr != nil {
		log.Println(encErr)
	}
	res, resErr := utils.DoPostRequest(c.Request.Cookies(),"https://music.163.com/weapi/artist/mvs", params, encSecKey)
	if resErr != nil {
		log.Println(resErr)
	}
	c.String(200,res)
}



// 收藏与取消收藏歌手
func ArtistSub(c *gin.Context){
	t := c.Query("t")
	if t=="1" {
		t = "sub"
	}else {
		t = "unsub"
	}
	artistId := c.Query("id")
	params, encSecKey, encErr := utils.Encrypt(fmt.Sprintf(`{"artistIds":"%s","t":"%s","csrf_token":""}`, artistId,t))
	if encErr != nil {
		log.Println(encErr)
	}
	res, resErr := utils.DoPostRequest(c.Request.Cookies(),"https://music.163.com/weapi/artist/"+t, params, encSecKey)
	if resErr != nil {
		log.Println(resErr)
	}
	c.String(200,res)
}



// 关注歌手列表
func ArtistSubList(c *gin.Context){

	limit := c.DefaultQuery("limit","25")
	offset := c.DefaultQuery("offset","0")
	params, encSecKey, encErr := utils.Encrypt(fmt.Sprintf(`{"limit":"%s","offset":"%s","total":true,"csrf_token":""}`,limit,offset))
	if encErr != nil {
		log.Println(encErr)
	}
	res, resErr := utils.DoPostRequest(c.Request.Cookies(),"https://music.163.com/weapi/artist/sublist", params, encSecKey)
	if resErr != nil {
		log.Println(resErr)
	}
	c.String(200,res)

}




// 歌手单曲
func Artists(c *gin.Context){
	uid :=c.Query("id")
	params, encSecKey, encErr := utils.Encrypt(fmt.Sprintf(`{"id":"%s","csrf_token":""}`,uid))
	if encErr != nil {
		log.Println(encErr)
	}
	res, resErr := utils.DoPostRequest(c.Request.Cookies(),"https://music.163.com/weapi/v1/artist/"+uid, params, encSecKey)
	if resErr != nil {
		log.Println(resErr)
	}
	c.String(200,res)
}