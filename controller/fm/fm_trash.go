package fm

import (
	"fmt"
	"github.com/WithLin/ncma/utils"
	"github.com/gin-gonic/gin"
	"log"
)

func FmTrash(c *gin.Context) {
	id := c.Query("id")
	time := c.DefaultQuery("time","25")
	params, encSecKey, encErr := utils.Encrypt(fmt.Sprintf(`{"songId":"%s","time":%s,"csrf_token":""}`,
		id,time))
	if encErr != nil {
		log.Println(encErr)
	}
	res, resErr := utils.DoPostRequest(c.Request.Cookies(),"https://music.163.com/weapi/radio/trash/add?alg=RT&songId"+id+"&time="+time, params, encSecKey)
	if resErr != nil {
		log.Println(resErr)
	}
	c.String(200, res)
}