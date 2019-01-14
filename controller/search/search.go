package search

import (
	"fmt"
	"github.com/WithLin/ncma/utils"
	"github.com/gin-gonic/gin"
	"log"
)

// 热门搜索
func SearchHot(c *gin.Context){

	params, encSecKey, encErr := utils.Encrypt(fmt.Sprintf(`{"type":1111,"csrf_token":""}`))
	if encErr != nil {
		log.Println(encErr)
	}
	res, resErr := utils.DoPostRequest(c.Request.Cookies(),"https://music.163.com/weapi/search/hot", params, encSecKey)
	if resErr != nil {
		log.Println(resErr)
	}
	c.String(200, res)

}


// 多类型搜索
func SearchMultiMatch(c *gin.Context){

	ty :=c.DefaultQuery("type","1")
	s :=c.DefaultQuery("keywords","")
	params, encSecKey, encErr := utils.Encrypt(fmt.Sprintf(`{"type":%s,"keywords":"%s","csrf_token":""}`,
		ty,s))
	if encErr != nil {
		log.Println(encErr)
	}
	res, resErr := utils.DoPostRequest(c.Request.Cookies(),"https://music.163.com/weapi/search/suggest/multimatch", params, encSecKey)
	if resErr != nil {
		log.Println(resErr)
	}
	c.String(200, res)

}



//搜索建议
func SearchSuggest(c *gin.Context){

	s :=c.DefaultQuery("keywords","")
	params, encSecKey, encErr := utils.Encrypt(fmt.Sprintf(`{"keywords":"%s","csrf_token":""}`,s))
	if encErr != nil {
		log.Println(encErr)
	}
	res, resErr := utils.DoPostRequest(c.Request.Cookies(),"https://music.163.com/weapi/search/suggest/web", params, encSecKey)
	if resErr != nil {
		log.Println(resErr)
	}
	c.String(200, res)

}


//搜索
func Search(c *gin.Context){

	limit :=c.DefaultQuery("limit","30")
	offset :=c.DefaultQuery("offset","0")
	ty :=c.DefaultQuery("type","1")
	params, encSecKey, encErr := utils.Encrypt(fmt.Sprintf(`{"type":%s,"limit":%s,"offset":%s,"csrf_token":""}`,
		ty,limit,offset))
	if encErr != nil {
		log.Println(encErr)
	}
	res, resErr := utils.DoPostRequest(c.Request.Cookies(),"https://music.163.com/weapi/search/get", params, encSecKey)
	if resErr != nil {
		log.Println(resErr)
	}
	c.String(200, res)

}


