package top

import (
	"fmt"
	"github.com/WithLin/ncma/utils"
	"github.com/gin-gonic/gin"
	"log"
	"strconv"
)

// 新碟上架
func TopAlbum(c *gin.Context){

	area :=c.DefaultQuery("type","ALL")
	limit := c.DefaultQuery("limit","30")
	offset := c.DefaultQuery("offset","0")
	params, encSecKey, encErr := utils.Encrypt(fmt.Sprintf(`{"area":"%s","limit":"%s","offset":"%s","total":true,"csrf_token":""}`,
		area,limit,offset))
	if encErr != nil {
		log.Println(encErr)
	}
	res, resErr := utils.DoPostRequest(c.Request.Cookies(),"https://music.163.com/weapi/album/new", params, encSecKey)
	if resErr != nil {
		log.Println(resErr)
	}
	c.String(200,res)

}



// 热门歌手
func TopArtists(c *gin.Context){

	limit := c.DefaultQuery("limit","50")
	offset := c.DefaultQuery("offset","0")
	params, encSecKey, encErr := utils.Encrypt(fmt.Sprintf(`{"limit":"%s","offset":"%s","total":true,"csrf_token":""}`,limit,offset))
	if encErr != nil {
		log.Println(encErr)
	}
	res, resErr := utils.DoPostRequest(c.Request.Cookies(),"https://music.163.com/weapi/artist/top", params, encSecKey)
	if resErr != nil {
		log.Println(resErr)
	}
	c.String(200,res)

}



// 热门歌手
func TopList(c *gin.Context){

	topList :=[24]string{
			 "3779629", //云音乐新歌榜
			 "3778678", //云音乐热歌榜
			 "2884035", ///云音乐原创榜
			 "19723756", //云音乐飙升榜
			 "10520166", //云音乐电音榜
			 "180106", //UK排行榜周榜
			"60198", //美国Billboard周榜
			"21845217", //KTV嗨榜
				 "11641012", //iTunes榜
				"120001", //Hit FM Top榜
				 "60131", //日本Oricon周榜
				 "3733003", //韩国Melon排行榜周榜
				 "60255", //韩国Mnet排行榜周榜
				 "46772709", //韩国Melon原声周榜
				 "112504", //中国TOP排行榜(港台榜)
				 "64016", //中国TOP排行榜(内地榜)
				 "10169002", //香港电台中文歌曲龙虎榜
				"4395559", //华语金曲榜
				 "1899724", //中国嘻哈榜
				 "27135204", //法国 NRJ EuroHot 30周榜
				"112463", //台湾Hito排行榜
				 "3812895", //Beatport全球电子舞曲榜
				"71385702", //云音乐ACG音乐榜
				"991319590", //云音乐嘻哈榜
	}
	id :=c.Query("idx")
	index ,_:=strconv.Atoi(id)
	params, encSecKey, encErr := utils.Encrypt(fmt.Sprintf(`{"id":"%s","n":10000,"csrf_token":""}`,topList[index]))
	if encErr != nil {
		log.Println(encErr)
	}
	res, resErr := utils.DoPostRequest(c.Request.Cookies(),"https://music.163.com/weapi/v3/playlist/detail", params, encSecKey)
	if resErr != nil {
		log.Println(resErr)
	}
	c.String(200,res)

}



// MV排行榜
func TopMv(c *gin.Context){

	limit := c.DefaultQuery("limit","30")
	offset := c.DefaultQuery("offset","0")
	params, encSecKey, encErr := utils.Encrypt(fmt.Sprintf(`{"limit":"%s","offset":"%s","total":true,"csrf_token":""}`,limit,offset))
	if encErr != nil {
		log.Println(encErr)
	}
	res, resErr := utils.DoPostRequest(c.Request.Cookies(),"https://music.163.com/weapi/mv/toplist", params, encSecKey)
	if resErr != nil {
		log.Println(resErr)
	}
	c.String(200,res)

}


// 精品歌单
func TopPlayListHighQuality(c *gin.Context){

	cat := c.DefaultQuery("cat","全部")
	limit := c.DefaultQuery("limit","30")
	lasttime := c.DefaultQuery("before","0")
	params, encSecKey, encErr := utils.Encrypt(fmt.Sprintf(`{"cat":"%s","limit":"%s","lasttime":"%s","total":true,"csrf_token":""}`,
		cat,limit,lasttime))
	if encErr != nil {
		log.Println(encErr)
	}
	res, resErr := utils.DoPostRequest(c.Request.Cookies(),"https://music.163.com/weapi/playlist/highquality/list", params, encSecKey)
	if resErr != nil {
		log.Println(resErr)
	}
	c.String(200,res)

}



// 精品歌单
func TopPlayList(c *gin.Context){

	cat := c.DefaultQuery("cat","全部")
	order :=c.DefaultQuery("order","hot")
	limit := c.DefaultQuery("limit","50")
	offset := c.DefaultQuery("offset","0")
	params, encSecKey, encErr := utils.Encrypt(fmt.Sprintf(`{"cat":"%s","order":"%s","limit":"%s","offset":"%s","total":true,"csrf_token":""}`,
		cat,order,limit,offset))
	if encErr != nil {
		log.Println(encErr)
	}
	res, resErr := utils.DoPostRequest(c.Request.Cookies(),"https://music.163.com/weapi/playlist/list", params, encSecKey)
	if resErr != nil {
		log.Println(resErr)
	}
	c.String(200,res)

}


// 新歌速递
func TopSong(c *gin.Context){

	areaId := c.DefaultQuery("areaId","0")
	params, encSecKey, encErr := utils.Encrypt(fmt.Sprintf(`{"areaId":"%s","total":true,"csrf_token":""}`,
		areaId))
	if encErr != nil {
		log.Println(encErr)
	}
	res, resErr := utils.DoPostRequest(c.Request.Cookies(),"https://music.163.com/weapi/playlist/list", params, encSecKey)
	if resErr != nil {
		log.Println(resErr)
	}
	c.String(200,res)

}


// 歌手榜
func TopListArtist(c *gin.Context){

	params, encSecKey, encErr := utils.Encrypt(fmt.Sprintf(`{"type":1,"limit":100，"offset":0,"total":true,"csrf_token":""}`))
	if encErr != nil {
		log.Println(encErr)
	}
	res, resErr := utils.DoPostRequest(c.Request.Cookies(),"https://music.163.com/weapi/toplist/artist", params, encSecKey)
	if resErr != nil {
		log.Println(resErr)
	}
	c.String(200,res)

}


// 所有榜单内容摘要
func TopListDetail(c *gin.Context){

	params, encSecKey, encErr := utils.Encrypt(fmt.Sprintf(`{"csrf_token":""}`))
	if encErr != nil {
		log.Println(encErr)
	}
	res, resErr := utils.DoPostRequest(c.Request.Cookies(),"https://music.163.com/weapi/toplist/detail", params, encSecKey)
	if resErr != nil {
		log.Println(resErr)
	}
	c.String(200,res)

}



// 所有榜单介绍
func TopArray(c *gin.Context){
	params, encSecKey, encErr := utils.Encrypt(fmt.Sprintf(`{"csrf_token":""}`))
	if encErr != nil {
		log.Println(encErr)
	}
	res, resErr := utils.DoPostRequest(c.Request.Cookies(),"https://music.163.com/weapi/toplist", params, encSecKey)
	if resErr != nil {
		log.Println(resErr)
	}
	c.String(200,res)

}





