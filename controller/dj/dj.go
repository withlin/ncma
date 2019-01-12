package dj

import (
	"fmt"
	"github.com/WithLin/ncma/utils"
	"github.com/gin-gonic/gin"
	"log"
)


// 电台分类列表
func DjCatelist(c *gin.Context){
	params, encSecKey, encErr := utils.Encrypt(fmt.Sprintf(`{"csrf_token":""}`))
	if encErr != nil {
		log.Println(encErr)
	}
	res, resErr := utils.DoPostRequest(c.Request.Cookies(),"https://music.163.com/weapi/djradio/category/get", params, encSecKey)
	if resErr != nil {
		log.Println(resErr)
	}
	c.String(200, res)
}





// 电台详情
func DjDetail(c *gin.Context){

	cid :=c.Query("rid")

	params, encSecKey, encErr := utils.Encrypt(fmt.Sprintf(`{"cid":"%s","csrf_token":""}`,cid))
	if encErr != nil {
		log.Println(encErr)
	}
	res, resErr := utils.DoPostRequest(c.Request.Cookies(),"https://music.163.com/weapi/djradio/get", params, encSecKey)
	if resErr != nil {
		log.Println(resErr)
	}
	c.String(200, res)
}



// 热门电台
func DjHost(c *gin.Context){
	ty :=c.Query("type")
	limit :=c.Query("limit")
	offset :=c.Query("offset")


	params, encSecKey, encErr := utils.Encrypt(fmt.
		Sprintf(
			`{"cat":"%s","cateId":"%s","type":"%s","categoryId":"%s","category":"%s","limit":%s,"offset":%s,"csrf_token":""}`,
			ty,ty,ty,ty,ty,limit,offset))
	if encErr != nil {
		log.Println(encErr)
	}
	res, resErr := utils.DoPostRequest(c.Request.Cookies(),"https://music.163.com/weapi/djradio/hot/v1", params, encSecKey)
	if resErr != nil {
		log.Println(resErr)
	}
	c.String(200, res)
}


// 付费电台
func DjPayGift(c *gin.Context){

	limit :=c.DefaultQuery("limit","20")
	offset :=c.DefaultQuery("offset","0")
	params, encSecKey, encErr := utils.Encrypt(fmt.Sprintf(`{"limit":%s,"offset":%s,"csrf_token":""}`,
		limit,offset))
	if encErr != nil {
		log.Println(encErr)
	}
	res, resErr := utils.DoPostRequest(c.Request.Cookies(),"https://music.163.com/weapi/djradio/home/paygift/list?_nmclfl=1", params, encSecKey)
	if resErr != nil {
		log.Println(resErr)
	}
	c.String(200, res)

}


// 电台节目详情
func DjProgramDetail(c *gin.Context){
	id :=c.Query("id")
	params, encSecKey, encErr := utils.Encrypt(fmt.Sprintf(`{"id":%s,"csrf_token":""}`,
		id))
	if encErr != nil {
		log.Println(encErr)
	}
	res, resErr := utils.DoPostRequest(c.Request.Cookies(),"https://music.163.com/weapi/dj/program/detail", params, encSecKey)
	if resErr != nil {
		log.Println(resErr)
	}
	c.String(200, res)
}



// 电台节目列表
func DjProgram(c *gin.Context){

	radioId :=c.Query("rid")
	limit :=c.DefaultQuery("limit","30")
	offset :=c.DefaultQuery("offset","0")
	asc :=c.Query("asc")
	params, encSecKey, encErr := utils.Encrypt(fmt.Sprintf(`{"radioId":%s,"limit":%s,"offset":%s,"asc":%s,"csrf_token":""}`,
		radioId,limit,offset,asc))
	if encErr != nil {
		log.Println(encErr)
	}
	res, resErr := utils.DoPostRequest(c.Request.Cookies(),"https://music.163.com/weapi/dj/program/byradio", params, encSecKey)
	if resErr != nil {
		log.Println(resErr)
	}
	c.String(200, res)
}



// 精选电台分类

/*
    有声书 10001
    知识技能 453050
    商业财经 453051
    人文历史 11
    外语世界 13
    亲子宝贝 14
    创作|翻唱 2001
    音乐故事 2
    3D|电子 10002
    相声曲艺 8
    情感调频 3
    美文读物 6
    脱口秀 5
    广播剧 7
    二次元 3001
    明星做主播 1
    娱乐|影视 4
    科技科学 453052
    校园|教育 4001
    旅途|城市 12
*/
func DjRecommendType(c *gin.Context){
	ty :=c.Query("type")
	params, encSecKey, encErr := utils.Encrypt(fmt.Sprintf(`{"cateId":%s,"csrf_token":""}`,
		ty))
	if encErr != nil {
		log.Println(encErr)
	}
	res, resErr := utils.DoPostRequest(c.Request.Cookies(),"https://music.163.com/weapi/djradio/recommend", params, encSecKey)
	if resErr != nil {
		log.Println(resErr)
	}
	c.String(200, res)
}



// 精选电台
func DjRecommend(c *gin.Context){
	params, encSecKey, encErr := utils.Encrypt(fmt.Sprintf(`{"csrf_token":""}`))
	if encErr != nil {
		log.Println(encErr)
	}
	res, resErr := utils.DoPostRequest(c.Request.Cookies(),"https://music.163.com/weapi/djradio/recommend/v1", params, encSecKey)
	if resErr != nil {
		log.Println(resErr)
	}
	c.String(200, res)
}



// 订阅与取消电台
func DjSub(c *gin.Context){

	t :=c.Query("t")
	if t=="1" {
		t = "sub"
	}else {
		t = "unsub"
	}

	id := c.Query("rid")


	params, encSecKey, encErr := utils.Encrypt(fmt.Sprintf(`{"t":"%s","id":"%s","csrf_token":""}`,
		t,id))
	if encErr != nil {
		log.Println(encErr)
	}
	res, resErr := utils.DoPostRequest(c.Request.Cookies(),"https://music.163.com/weapi/djradio/"+t, params, encSecKey)
	if resErr != nil {
		log.Println(resErr)
	}
	c.String(200, res)
}



// 订阅电台列表
func DjSublist(c *gin.Context){
	limit :=c.DefaultQuery("limit","30")
	offset :=c.DefaultQuery("offset","0")
	params, encSecKey, encErr := utils.Encrypt(fmt.Sprintf(`{"limit":"%s","offset":"%s","total":ture,"csrf_token":""}`,
		limit,offset))
	if encErr != nil {
		log.Println(encErr)
	}
	res, resErr := utils.DoPostRequest(c.Request.Cookies(),"https://music.163.com/weapi/djradio/get/subed", params, encSecKey)
	if resErr != nil {
		log.Println(resErr)
	}
	c.String(200, res)
}

