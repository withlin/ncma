package user

import (
	"fmt"
	"github.com/WithLin/ncma/utils"
	"github.com/gin-gonic/gin"
	"log"
	"strconv"
)

//用户详情
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



// 收藏计数
func UserSubCount(c *gin.Context) {
	params, encSecKey, encErr := utils.Encrypt(fmt.Sprintf(`{"csrf_token":""}`))
	if encErr != nil {
		log.Println(encErr)
	}
	res, resErr := utils.DoPostRequest(c.Request.Cookies(),"https://music.163.com/weapi/subcount", params, encSecKey)
	if resErr != nil {
		log.Println(resErr)
	}
	c.String(200,res)
}



// 用户创建的电台
func UserAudio(c *gin.Context){
	uid :=c.Query("uid")
	params, encSecKey, encErr := utils.Encrypt(fmt.Sprintf(`{"uid":"%s","csrf_token":""}`,uid))
	if encErr != nil {
		log.Println(encErr)
	}
	res, resErr := utils.DoPostRequest(c.Request.Cookies(),"https://music.163.com/weapi/djradio/get/byuser", params, encSecKey)
	if resErr != nil {
		log.Println(resErr)
	}
	c.String(200,res)
}



//云盘
func UserCloud(c *gin.Context){
	limit := c.DefaultQuery("limit","30")
	offset := c.DefaultQuery("offset","0")
	params, encSecKey, encErr := utils.Encrypt(fmt.Sprintf(`{"limit":"%s","offset":"%s","csrf_token":""}`,limit,offset))
	if encErr != nil {
		log.Println(encErr)
	}
	res, resErr := utils.DoPostRequest(c.Request.Cookies(),"https://music.163.com/weapi/djradio/get/byuser", params, encSecKey)
	if resErr != nil {
		log.Println(resErr)
	}
	c.String(200,res)
}



//云盘数据详情(暂时不要使用)
func UserCloudSearch(c *gin.Context){
	byids := c.Query("byids")
	id := c.Query("id")
	params, encSecKey, encErr := utils.Encrypt(fmt.Sprintf(`{"byids":"%s","id":"%s","csrf_token":""}`,byids,id))
	if encErr != nil {
		log.Println(encErr)
	}
	res, resErr := utils.DoPostRequest(c.Request.Cookies(),"https://music.163.com/weapi/djradio/get/byuser", params, encSecKey)
	if resErr != nil {
		log.Println(resErr)
	}
	c.String(200,res)

}



//用户电台节目
func UserDj(c *gin.Context){
	uid :=c.Query("uid")
	limit := c.DefaultQuery("limit","30")
	offset := c.DefaultQuery("offset","0")
	params, encSecKey, encErr := utils.Encrypt(fmt.Sprintf(`{"uid":"%s","limit":"%s","offset":"%s","csrf_token":""}`,limit,offset))
	if encErr != nil {
		log.Println(encErr)
	}
	res, resErr := utils.DoPostRequest(c.Request.Cookies(),"https://music.163.com/weapi/dj/program/"+uid, params, encSecKey)
	if resErr != nil {
		log.Println(resErr)
	}
	c.String(200,res)
}



//用户动态
func UserEvent(c *gin.Context){
	uid :=c.Query("uid")
	params, encSecKey, encErr := utils.Encrypt(fmt.Sprintf(`{"uid":"%s","time":-1,"getcounts":true,"csrf_token":""}`,uid))
	if encErr != nil {
		log.Println(encErr)
	}
	res, resErr := utils.DoPostRequest(c.Request.Cookies(),"https://music.163.com/weapi/event/get/"+uid, params, encSecKey)
	if resErr != nil {
		log.Println(resErr)
	}
	c.String(200,res)
}



// 关注TA的人(粉丝)
func UserFolloweds(c *gin.Context){
	uid :=c.Query("uid")
	limit := c.DefaultQuery("limit","30")
	offset := c.DefaultQuery("offset","0")
	params, encSecKey, encErr := utils.Encrypt(fmt.Sprintf(`{"uid":"%s","limit":"%s","offset":"%s","csrf_token":""}`,limit,offset))
	if encErr != nil {
		log.Println(encErr)
	}
	res, resErr := utils.DoPostRequest(c.Request.Cookies(),"https://music.163.com/weapi/user/getfolloweds"+uid, params, encSecKey)
	if resErr != nil {
		log.Println(resErr)
	}
	c.String(200,res)

}



// TA关注的人(关注)
func UserFollows(c *gin.Context){

	uid :=c.Query("uid")
	limit := c.DefaultQuery("limit","30")
	offset := c.DefaultQuery("offset","0")
	params, encSecKey, encErr := utils.Encrypt(fmt.Sprintf(`{"uid":"%s","limit":"%s","offset":"%s","order":true,"csrf_token":""}`,limit,offset))
	if encErr != nil {
		log.Println(encErr)
	}
	res, resErr := utils.DoPostRequest(c.Request.Cookies(),"https://music.163.com/weapi/user/getfollows/"+uid, params, encSecKey)
	if resErr != nil {
		log.Println(resErr)
	}
	c.String(200,res)
}



//获取用户歌单
func UserPlayList(c *gin.Context){
	uid := c.Query("uid")
	limit := c.DefaultQuery("limit","30")
	offset := c.DefaultQuery("offset","0")

	params, encSecKey, encErr := utils.Encrypt(fmt.Sprintf(`{"uid":"%s","limit":"%s","offset":"%s","csrf_token":""}`,
		uid,limit,offset))
	if encErr != nil {
		log.Println(encErr)
	}
	res, resErr := utils.DoPostRequest(c.Request.Cookies(),"https://music.163.com/weapi/user/playlist", params, encSecKey)
	if resErr != nil {
		log.Println(resErr)
	}
	c.String(200,res)


}


// 听歌排行
func UserRecord(c *gin.Context){
	uid :=c.Query("uid")
	ty := c.DefaultQuery("type","1")
	result, _ := strconv.Atoi(ty)
	params, encSecKey, encErr := utils.Encrypt(fmt.Sprintf(`{"uid":"%s","type":%d,"csrf_token":""}`,uid,result))
	if encErr != nil {
		log.Println(encErr)
	}
	res, resErr := utils.DoPostRequest(c.Request.Cookies(),"https://music.163.com/weapi/v1/play/record", params, encSecKey)
	if resErr != nil {
		log.Println(resErr)
	}
	c.String(200,res)
}



//更新用户信息
func UserUpdate(c *gin.Context){
	birthday :=c.Query("birthday")
	birthday =fmt.Sprintf(`"birthday":"%s"`,birthday)
	city :=c.Query("city")
	city =fmt.Sprintf(`"city":"%s"`,city)
	gender :=c.Query("gender")
	gender =fmt.Sprintf(`"gender":"%s"`,gender)
	nickname :=c.Query("nickname")
	nickname =fmt.Sprintf(`"nickname":"%s"`,nickname)
	signature :=c.Query("signature")
	signature =fmt.Sprintf(`"signature":"%s"`,signature)
	params, encSecKey, encErr := utils.Encrypt(fmt.Sprintf(`{"avatarImgId":"0",%s,%s,%s,%s,%s,"csrf_token":""}`,
		birthday,city,gender,nickname,signature))
	if encErr != nil {
		log.Println(encErr)
	}
	res, resErr := utils.DoPostRequest(c.Request.Cookies(),"https://music.163.com/weapi/user/profile/update", params, encSecKey)
	if resErr != nil {
		log.Println(resErr)
	}
	c.String(200,res)
}
