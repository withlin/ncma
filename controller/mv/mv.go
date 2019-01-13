package mv

import (
	"fmt"
	"github.com/WithLin/ncma/utils"
	"github.com/gin-gonic/gin"
	"log"
)

// MV详情
func MvDetail(c *gin.Context) {
	id :=c.Query("mvid")
	params, encSecKey, encErr := utils.Encrypt(fmt.Sprintf(`{"mvid":"%s","csrf_token":""}`,id))
	if encErr != nil {
		log.Println(encErr)
	}
	res, resErr := utils.DoPostRequest(c.Request.Cookies(),"https://music.163.com/weapi/mv/detail", params, encSecKey)
	if resErr != nil {
		log.Println(resErr)
	}
	c.String(200, res)
}


// 最新MV
func MvFist(c *gin.Context) {

	limit :=c.DefaultQuery("limit","30")

	params, encSecKey, encErr := utils.Encrypt(fmt.Sprintf(`{"limit":%s,"total":true,"csrf_token":""}`,limit))
	if encErr != nil {
		log.Println(encErr)
	}
	res, resErr := utils.DoPostRequest(c.Request.Cookies(),"https://music.163.com/weapi/mv/first", params, encSecKey)
	if resErr != nil {
		log.Println(resErr)
	}
	c.String(200, res)
}



// 收藏与取消收藏MV
func MvSub(c *gin.Context) {

    t :=c.Query("t")
	if t=="1" {
		t="sub"
	}else{
		t="unsub"
	}
	mvid :=c.Query("mvid")

	params, encSecKey, encErr := utils.Encrypt(fmt.Sprintf(`{"t":"%s","mvid":"%s","mvids":"[%s]","csrf_token":""}`,
		t,mvid,mvid))
	if encErr != nil {
		log.Println(encErr)
	}
	res, resErr := utils.DoPostRequest(c.Request.Cookies(),"https://music.163.com/weapi/mv/"+t, params, encSecKey)
	if resErr != nil {
		log.Println(resErr)
	}
	c.String(200, res)
}


// // 已收藏MV列表
func MvList(c *gin.Context) {


	limit :=c.DefaultQuery("limit","30")
	offset :=c.DefaultQuery("offset","0")

	params, encSecKey, encErr := utils.Encrypt(fmt.Sprintf(`{"limit":"%s","offset":"%s","total":true,"csrf_token":""}`,
		limit,offset))
	if encErr != nil {
		log.Println(encErr)
	}
	res, resErr := utils.DoPostRequest(c.Request.Cookies(),"https://music.163.com/weapi/cloudvideo/allvideo/sublist", params, encSecKey)
	if resErr != nil {
		log.Println(resErr)
	}
	c.String(200, res)
}


// mv链接
func MvUrl(c *gin.Context) {


	id :=c.Query("id")
	r :=c.DefaultQuery("res","1080")

	params, encSecKey, encErr := utils.Encrypt(fmt.Sprintf(`{"id":"%s","r":%s,"csrf_token":""}`,
		id,r))
	if encErr != nil {
		log.Println(encErr)
	}
	res, resErr := utils.DoPostRequest(c.Request.Cookies(),"https://music.163.com/weapi/song/enhance/play/mv/url", params, encSecKey)
	if resErr != nil {
		log.Println(resErr)
	}
	c.String(200, res)
}