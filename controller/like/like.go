package like

import (
	"fmt"
	"github.com/WithLin/ncma/utils"
	"github.com/gin-gonic/gin"
	"log"
)


// 红心与取消红心歌曲
func Like(c *gin.Context) {

	like :=c.Query("like")
	if like == "" {
		like="true"
	}else{
		like = "false"
	}
	trackId :=c.Query("id")
	time := c.DefaultQuery("time","25")
	alg :=c.DefaultQuery("alg","itembased")

	params, encSecKey, encErr := utils.Encrypt(fmt.Sprintf(`{"like":%s,"trackId":"%s","time":%s,"alg":"%s","csrf_token":""}`,
		like,trackId,time,alg))
	if encErr != nil {
		log.Println(encErr)
	}
	res, resErr := utils.DoPostRequest(c.Request.Cookies(),"https://music.163.com/weapi/radio/like?alg="+
		alg+"&trackId="+trackId+"&like="+like+"&time="+time, params, encSecKey)
	if resErr != nil {
		log.Println(resErr)
	}
	c.String(200, res)
}



// 喜欢的歌曲(无序)
func LikeList(c *gin.Context) {

	uid :=c.Query("uid")

	params, encSecKey, encErr := utils.Encrypt(fmt.Sprintf(`{"uid":"%s","csrf_token":""}`,
		uid))
	if encErr != nil {
		log.Println(encErr)
	}
	res, resErr := utils.DoPostRequest(c.Request.Cookies(),"https://music.163.com/weapi/song/like/get", params, encSecKey)
	if resErr != nil {
		log.Println(resErr)
	}
	c.String(200, res)
}