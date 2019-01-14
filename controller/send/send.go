package send

import (
	"fmt"
	"github.com/WithLin/ncma/utils"
	"github.com/gin-gonic/gin"
	"log"
)

// 私信歌单
func SendPlayList(c *gin.Context){

	id :=c.Query("playlist")
	msg :=c.Query("msg")
	userIds :=c.Query("user_ids")
	params, encSecKey, encErr := utils.Encrypt(fmt.Sprintf(`{"id":"%s","type":"playlist","msg":"%s","userIds":"%s","csrf_token":""}`,
		id,msg,userIds))
	if encErr != nil {
		log.Println(encErr)
	}
	res, resErr := utils.DoPostRequest(c.Request.Cookies(),"https://music.163.com/weapi/msg/private/send", params, encSecKey)
	if resErr != nil {
		log.Println(resErr)
	}
	c.String(200, res)

}

// 私信
func SendText(c *gin.Context){

	id :=c.Query("id")
	msg :=c.Query("msg")
	userIds :=c.Query("user_ids")
	params, encSecKey, encErr := utils.Encrypt(fmt.Sprintf(`{"id":"%s","type":"text","msg":"%s","userIds":"%s","csrf_token":""}`,
		id,msg,userIds))
	if encErr != nil {
		log.Println(encErr)
	}
	res, resErr := utils.DoPostRequest(c.Request.Cookies(),"https://music.163.com/weapi/discovery/simiArtist", params, encSecKey)
	if resErr != nil {
		log.Println(resErr)
	}
	c.String(200, res)

}

