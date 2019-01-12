package login

import (
	"bufio"
	"fmt"
	"github.com/WithLin/ncma/utils"
	"github.com/gin-gonic/gin"
	"strings"
)


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