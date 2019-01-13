package follow

import (
	"fmt"
	"github.com/WithLin/ncma/utils"
	"github.com/gin-gonic/gin"
	"log"
)

func Follow(c *gin.Context) {

    t :=c.Query("t")
	if t == "1"{
		t = "follow"
	}else {
		t = "delfollow"
	}
    id :=c.Query("id")
	params, encSecKey, encErr := utils.Encrypt(fmt.Sprintf(`{"t":"%s","id":"%s","csrf_token":""}`,t,id))
	if encErr != nil {
		log.Println(encErr)
	}
	res, resErr := utils.DoPostRequest(c.Request.Cookies(),"https://music.163.com/weapi/user/"+t+"/"+id, params, encSecKey)
	if resErr != nil {
		log.Println(resErr)
	}
	c.String(200, res)
}