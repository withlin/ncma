package recommend

import (
	"fmt"
	"github.com/WithLin/ncma/utils"
	"github.com/gin-gonic/gin"
	"log"
)

// 推荐节目
func ProgramRecommend(c *gin.Context) {

	cateId :=c.Query("type")
	limit :=c.DefaultQuery("limit","10")
	offset :=c.DefaultQuery("limit","0")
	params, encSecKey, encErr := utils.Encrypt(fmt.Sprintf(`{"cateId":"%s","limit":%s,"offset":%s,"csrf_token":""}`,
		cateId,limit,offset))
	if encErr != nil {
		log.Println(encErr)
	}
	res, resErr := utils.DoPostRequest(c.Request.Cookies(),"https://music.163.com/weapi/program/recommend/v1", params, encSecKey)
	if resErr != nil {
		log.Println(resErr)
	}
	c.String(200, res)
}


// 推荐节目
func RecommendResource(c *gin.Context) {

	params, encSecKey, encErr := utils.Encrypt(fmt.Sprintf(`{"csrf_token":""}`))
	if encErr != nil {
		log.Println(encErr)
	}
	res, resErr := utils.DoPostRequest(c.Request.Cookies(),"https://music.163.com/weapi/v1/discovery/recommend/resource", params, encSecKey)
	if resErr != nil {
		log.Println(resErr)
	}
	c.String(200, res)
}


// 每日推荐歌曲
func RecommendSongs(c *gin.Context){

	limit :=c.DefaultQuery("limit","20")
	offset :=c.DefaultQuery("limit","0")

	params, encSecKey, encErr := utils.Encrypt(fmt.Sprintf(`{"limit":%s,"offset":%s,"total":true,"csrf_token":""}`,
		limit,offset))
	if encErr != nil {
		log.Println(encErr)
	}
	res, resErr := utils.DoPostRequest(c.Request.Cookies(),"https://music.163.com/weapi/v1/discovery/recommend/songs", params, encSecKey)
	if resErr != nil {
		log.Println(resErr)
	}
	c.String(200, res)

}
