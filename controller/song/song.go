package song

import (
	"fmt"
	"github.com/WithLin/ncma/utils"
	"github.com/gin-gonic/gin"
	"log"
	"strings"
)

// 相似歌手
func SongDetail(c *gin.Context){

	ids :=strings.Split(c.Query("ids"),",")
	re :=""
	for id :=0;id < len(ids) ;id++  {
		re +=fmt.Sprintf(`{"id":"%s"},`,id)
	}
	re =re[0:len(re)-1]

	params, encSecKey, encErr := utils.Encrypt(fmt.Sprintf(`{"ids":[%s],"c":"%s","csrf_token":""}`,
		strings.Join(ids,","),re))
	if encErr != nil {
		log.Println(encErr)
	}
	res, resErr := utils.DoPostRequest(c.Request.Cookies(),"https://music.163.com/weapi/v3/song/detail", params, encSecKey)
	if resErr != nil {
		log.Println(resErr)
	}
	c.String(200, res)

}



// 相似歌手
func SongUrl(c *gin.Context){

	//if(!('MUSIC_U' in query.cookie)) query.cookie._ntes_nuid = crypto.randomBytes(16).toString("hex")
	id :=c.Query("id")
	br :=c.DefaultQuery("br","999000")

	params, encSecKey, encErr := utils.Encrypt(fmt.Sprintf(`{"ids":[%s],"br":"%s","csrf_token":""}`,
		id,br))
	if encErr != nil {
		log.Println(encErr)
	}
	res, resErr := utils.DoPostRequest(c.Request.Cookies(),"https://music.163.com/weapi/song/enhance/player/url", params, encSecKey)
	if resErr != nil {
		log.Println(resErr)
	}
	c.String(200, res)

}