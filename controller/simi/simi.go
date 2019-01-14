package simi

import (
	"fmt"
	"github.com/WithLin/ncma/utils"
	"github.com/gin-gonic/gin"
	"log"
)



// 相似歌手
func SimiArtist(c *gin.Context){

	artistid :=c.Query("id")
	params, encSecKey, encErr := utils.Encrypt(fmt.Sprintf(`{"artistid":"%s","csrf_token":""}`,artistid))
	if encErr != nil {
		log.Println(encErr)
	}
	res, resErr := utils.DoPostRequest(c.Request.Cookies(),"https://music.163.com/weapi/discovery/simiArtist", params, encSecKey)
	if resErr != nil {
		log.Println(resErr)
	}
	c.String(200, res)

}


// 相似歌手
func SimiMV(c *gin.Context){

	mvid :=c.Query("mvid")
	params, encSecKey, encErr := utils.Encrypt(fmt.Sprintf(`{"mvid":"%s","csrf_token":""}`,mvid))
	if encErr != nil {
		log.Println(encErr)
	}
	res, resErr := utils.DoPostRequest(c.Request.Cookies(),"https://music.163.com/weapi/discovery/simiMV", params, encSecKey)
	if resErr != nil {
		log.Println(resErr)
	}
	c.String(200, res)
}


// 相似歌单
func SimiPlayList(c *gin.Context){


	songid :=c.Query("mvid")
	limit :=c.DefaultQuery("limit","50")
	offset :=c.DefaultQuery("offset","0")

	params, encSecKey, encErr := utils.Encrypt(fmt.Sprintf(`{"limit":%s,"offset":%s,"songid":%s","csrf_token":""}`,
		limit,offset,songid))
	if encErr != nil {
		log.Println(encErr)
	}
	res, resErr := utils.DoPostRequest(c.Request.Cookies(),"https://music.163.com/weapi/discovery/simiPlaylist", params, encSecKey)
	if resErr != nil {
		log.Println(resErr)
	}
	c.String(200, res)
}



// 相似歌曲
func SimiSong(c *gin.Context){

	songid :=c.Query("id")
	limit :=c.DefaultQuery("limit","50")
	offset :=c.DefaultQuery("offset","0")

	params, encSecKey, encErr := utils.Encrypt(fmt.Sprintf(`{"limit":%s,"offset":%s,"songid":%s","csrf_token":""}`,
		limit,offset,songid))
	if encErr != nil {
		log.Println(encErr)
	}
	res, resErr := utils.DoPostRequest(c.Request.Cookies(),"https://music.163.com/weapi/v1/discovery/simiSong", params, encSecKey)
	if resErr != nil {
		log.Println(resErr)
	}
	c.String(200, res)
}



// 相似用户
func SimiUser(c *gin.Context){

	songid :=c.Query("id")
	limit :=c.DefaultQuery("limit","50")
	offset :=c.DefaultQuery("offset","0")

	params, encSecKey, encErr := utils.Encrypt(fmt.Sprintf(`{"limit":%s,"offset":%s,"songid":%s","csrf_token":""}`,
		limit,offset,songid))
	if encErr != nil {
		log.Println(encErr)
	}
	res, resErr := utils.DoPostRequest(c.Request.Cookies(),"https://music.163.com/weapi/discovery/simiUser", params, encSecKey)
	if resErr != nil {
		log.Println(resErr)
	}
	c.String(200, res)
}


