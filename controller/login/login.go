package login

import (
	"bufio"
	"crypto/md5"
	"fmt"
	"github.com/WithLin/ncma/utils"
	"github.com/gin-gonic/gin"
	"log"
	"strings"
)


//手机登陆
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



//Email登陆
func LoginEmail(c *gin.Context) {
	email := c.Query("email")
	password := c.Query("password")
	params, encSecKey, encErr := utils.Encrypt(fmt.Sprintf(`{"phone":"%s","password":"%s","rememberLogin":"true"}`, email,fmt.Sprintf("%x", md5.Sum([]byte(password)))))
	if encErr != nil {
		log.Println(encErr)
	}
	res, resErr := utils.DoPostRequest(c.Request.Cookies(),"https://music.163.com/weapi/login", params, encSecKey)
	if resErr != nil {
		log.Println(resErr)
	}
	c.String(200, res)
}



//刷新登陆
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



//登陆状态
func LoginStatus(c *gin.Context) {
	res, resErr := utils.DoGetRequest(c.Request.Cookies(),"https://music.163.com")
	if resErr != nil {

	}
	scanner := bufio.NewScanner(strings.NewReader(res))
	i :=1
	profile :=""
	bindings :=""
	for scanner.Scan() {
		i++
		if i ==27 {
			profile = string(scanner.Bytes()[11:len(scanner.Bytes())-2])
			fmt.Println(profile)
		}
		if i ==296 {
			bindings = string(scanner.Bytes()[11:len(scanner.Bytes())-1])
			fmt.Println(bindings)
			break
		}
	}
	var userId,nickname,avatarUrl,birthday,userType,djStatus string
	p :=strings.Split(profile,",")
	for i :=0;i <len(p) ;i ++  {
		result :=strings.Split(p[i],":")
		switch result[0] {
		case "userId":
			userId=result[1]
		case "nickname":
			nickname=result[1]
		case "avatarUrl":
			avatarUrl=result[1]+":"+result[2]
		case "birthday":
			birthday=result[1]
		case "userType":
			userType=result[1]
		case "djStatus":
			djStatus=result[1]
		}
	}
	result := fmt.Sprintf(`{"userId":%s,"nickname":%s,"avatarUrl":%s,"birthday":%s,"userType":%s,"djStatus":%s}`,
		userId,nickname,avatarUrl,birthday,userType,djStatus)
	c.String(200,fmt.Sprintf(`{"code": 200, "profile":%s, "bindings": %s}`,result,bindings))
}



//退出登陆
func LogOut(c *gin.Context) {
	params, encSecKey, encErr := utils.Encrypt(fmt.Sprintf(`{"csrf_token":""}`))
	if encErr != nil {
		log.Println(encErr)
	}
	res, resErr := utils.DoPostRequest(c.Request.Cookies(),"https://music.163.com/weapi/logout", params, encSecKey)
	if resErr != nil {
		log.Println(resErr)
	}
	c.String(200, res)
}