package user

import (
	"fmt"
	"github.com/WithLin/ncma/utils"
	"github.com/gin-gonic/gin"
	"log"
)

func UserSubCount(c *gin.Context) {
	params, encSecKey, encErr := utils.Encrypt(fmt.Sprintf(`{"csrf_token":""}`))
	if encErr != nil {
		log.Println(encErr)
	}
	res, resErr := utils.DoPostRequest(c.Request.Cookies(),"https://music.163.com/weapi/subcount", params, encSecKey)
	if resErr != nil {
		log.Println(resErr)
	}
	c.String(200,res)
}