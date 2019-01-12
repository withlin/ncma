package login

import (
	"crypto/md5"
	"fmt"
	"github.com/WithLin/ncma/utils"
	"github.com/gin-gonic/gin"
	"log"
)

func LoginCellPhone(c *gin.Context) {
	phone := c.Query("phone")
	password := c.Query("password")
	params, encSecKey, encErr := utils.Encrypt(fmt.Sprintf(`{"phone":"%s","password":"%s","rememberLogin":"true"}`, phone,fmt.Sprintf("%x", md5.Sum([]byte(password)))))
	if encErr != nil {
		log.Println(encErr)
	}

	res, resErr := utils.DoPostRequest(c.Request.Cookies(),"https://music.163.com/weapi/login/cellphone", params, encSecKey)
	if resErr != nil {
		log.Println(resErr)
	}
	c.String(200, res)
}