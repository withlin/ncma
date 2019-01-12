package login

import (
	"fmt"
	"github.com/WithLin/ncma/utils"
	"github.com/gin-gonic/gin"
	"log"
)

func LoginRefresh(c *gin.Context) {
	params, encSecKey, encErr := utils.Encrypt(fmt.Sprintf(`{"csrf_token":""}`))
	if encErr != nil {
		log.Println(encErr)
	}
	res, resErr := utils.DoPostRequest(c.Request.Cookies(),"https://music.163.com/weapi/login/token/refresh", params, encSecKey)
	if resErr != nil {
		log.Println(resErr)
	}
	c.String(200, res)
}