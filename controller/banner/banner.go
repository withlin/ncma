package banner

import (
	"fmt"
	"github.com/WithLin/ncma/utils"
	"github.com/gin-gonic/gin"
	"log"
)

// 轮播图
func Banner(c *gin.Context){
	params, encSecKey, encErr := utils.Encrypt(fmt.Sprintf(`{"clientType":"pc","csrf_token":""}`))
	if encErr != nil {
		log.Println(encErr)
	}
	res, resErr := utils.DoPostRequest(c.Request.Cookies(),"https://music.163.com/api/v2/banner/get", params, encSecKey)
	if resErr != nil {
		log.Println(resErr)
	}
	c.String(200,res)

}