package check

import (
	"fmt"
	"github.com/WithLin/ncma/utils"
	"github.com/gin-gonic/gin"
	"log"
)


// 歌曲可用性
func CheckMusic(c *gin.Context){

	id :=c.Query("id")
	br :=c.DefaultQuery("br","999000")
	params, encSecKey, encErr := utils.Encrypt(fmt.Sprintf(`{"ids":%s,"br":%s,"csrf_token":""}`,
		id,br))
	if encErr != nil {
		log.Println(encErr)
	}
	_, resErr := utils.DoPostRequest(c.Request.Cookies(),"https://music.163.com/weapi/song/enhance/player/url", params, encSecKey)
	playable := false
	if resErr != nil {
		playable = true
	}
	if playable {

			c.String(200,`{"success"": true, "message"": "ok""}`)
	}else {
		c.String(200,`{"success"": false, "message"": "亲爱的,暂无版权"}`)
	}


}