package user

import (
	"fmt"
	"github.com/WithLin/ncma/utils"
	"github.com/gin-gonic/gin"
	"log"
)

func UserDetail(c *gin.Context) {
	id := c.Query("uid")
	params, encSecKey, encErr := utils.Encrypt(fmt.Sprintf(`{"id":"%s","csrf_token":""}`, id))
	if encErr != nil {
		log.Println(encErr)
	}
	res, resErr := utils.DoPostRequest(c.Request.Cookies(),"https://music.163.com/weapi/v1/user/detail/"+id, params, encSecKey)
	if resErr != nil {
		log.Println(resErr)
	}
	c.String(200,res)
}