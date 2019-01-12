package video

import (
	"fmt"
	"github.com/WithLin/ncma/utils"
	"github.com/gin-gonic/gin"
	"log"
	"strconv"
)

// 视频详情
func VedioDetail(c *gin.Context) {
	id := c.Query("uid")
	params, encSecKey, encErr := utils.Encrypt(fmt.Sprintf(`{"id":"%s","csrf_token":""}`, id))
	if encErr != nil {
		log.Println(encErr)
	}
	res, resErr := utils.DoPostRequest(c.Request.Cookies(),"https://music.163.com/weapi/cloudvideo/v1/video/detail", params, encSecKey)
	if resErr != nil {
		log.Println(resErr)
	}
	c.String(200,res)
}


// 收藏与取消收藏视频
func VedioSub(c *gin.Context) {
	t := c.Query("t")
	if t=="1" {
		t = "sub"
	}else {
		t = "unsub"
	}
	id := c.Query("uid")
	params, encSecKey, encErr := utils.Encrypt(fmt.Sprintf(`{"id":"%s","t":"%s","csrf_token":""}`, id,t))
	if encErr != nil {
		log.Println(encErr)
	}
	res, resErr := utils.DoPostRequest(c.Request.Cookies(),"https://music.163.com/weapi/cloudvideo/video/"+t, params, encSecKey)
	if resErr != nil {
		log.Println(resErr)
	}
	c.String(200,res)
}


// 视频链接
func VedioUrl(c *gin.Context) {
	id := c.Query("id")
	res := c.DefaultQuery("res","1080")
	p, _ := strconv.Atoi(res)
	params, encSecKey, encErr := utils.Encrypt(fmt.Sprintf(`{"ids":"[%s]","resolution":%d,"csrf_token":""}`, id,p))
	if encErr != nil {
		log.Println(encErr)
	}
	res, resErr := utils.DoPostRequest(c.Request.Cookies(),"https://music.163.com/weapi/cloudvideo/playurl", params, encSecKey)
	if resErr != nil {
		log.Println(resErr)
	}
	c.String(200,res)
}