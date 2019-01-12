package comment

import (
	"fmt"
	"github.com/WithLin/ncma/utils"
	"github.com/gin-gonic/gin"
	"log"
)

// 专辑评论
func  CommentAlbum(c *gin.Context){

	rid :=c.Query("uid")
	limit := c.DefaultQuery("limit","20")
	offset := c.DefaultQuery("offset","0")
	params, encSecKey, encErr := utils.Encrypt(fmt.Sprintf(`{"rid":"%s","limit":"%s","offset":"%s","csrf_token":""}`,rid,limit,offset))
	if encErr != nil {
		log.Println(encErr)
	}
	res, resErr := utils.DoPostRequest(c.Request.Cookies(),"https://music.163.com/weapi/v1/resource/comments/R_AL_3_"+rid, params, encSecKey)
	if resErr != nil {
		log.Println(resErr)
	}
	c.String(200,res)


}


func  CommentDj(c *gin.Context){

	rid :=c.Query("uid")
	limit := c.DefaultQuery("limit","20")
	offset := c.DefaultQuery("offset","0")
	params, encSecKey, encErr := utils.Encrypt(fmt.Sprintf(`{"rid":"%s","limit":"%s","offset":"%s","csrf_token":""}`,rid,limit,offset))
	if encErr != nil {
		log.Println(encErr)
	}
	res, resErr := utils.DoPostRequest(c.Request.Cookies(),"https://music.163.com/weapi/v1/resource/comments/A_DJ_1_"+rid, params, encSecKey)
	if resErr != nil {
		log.Println(resErr)
	}
	c.String(200,res)


}



// 热门评论
func  CommentHot(c *gin.Context){

	ty :=c.Query("type")
	switch ty {
	case "0":
		ty = "R_SO_4_"
	case "1":
		ty = "R_MV_5_"
	case "2":
		ty = "A_PL_0_"
	case "3":
		ty = "R_AL_3_"
	case "4":
		ty = "A_DJ_1_"
	case "5":
		ty = "R_VI_62_"

	}

	rid :=c.Query("id")
	limit := c.DefaultQuery("limit","20")
	offset := c.DefaultQuery("offset","0")
	params, encSecKey, encErr := utils.Encrypt(fmt.Sprintf(`{"type":"%s","rid":"%s","limit":"%s","offset":"%s","csrf_token":""}`,
		ty,rid,limit,offset))
	if encErr != nil {
		log.Println(encErr)
	}
	res, resErr := utils.DoPostRequest(c.Request.Cookies(),"https://music.163.com/weapi/v1/resource/hotcomments/"+ty+rid, params, encSecKey)
	if resErr != nil {
		log.Println(resErr)
	}
	c.String(200,res)


}



func  CommentLike(c *gin.Context){

	t :=c.Query("t")
	if t== "1" {
		t = "like"
	}else {
		t = "unlike"
	}
	ty :=c.Query("type")
	switch ty {
	case "0":
		ty = "R_SO_4_"
	case "1":
		ty = "R_MV_5_"
	case "2":
		ty = "A_PL_0_"
	case "3":
		ty = "R_AL_3_"
	case "4":
		ty = "A_DJ_1_"
	case "5":
		ty = "R_VI_62_"

	}

	id :=c.Query("id")
	cid :=c.Query("cid")
	params, encSecKey, encErr := utils.Encrypt(fmt.Sprintf(`{"threadId":"%s","commentId":"%s","csrf_token":""}`,
		ty+id,cid))
	if encErr != nil {
		log.Println(encErr)
	}
	res, resErr := utils.DoPostRequest(c.Request.Cookies(),"https://music.163.com/weapi/v1/comment/"+t, params, encSecKey)
	if resErr != nil {
		log.Println(resErr)
	}
	c.String(200,res)
}




// 歌曲评论
func  CommentMusic(c *gin.Context){

	rid :=c.Query("id")
	limit := c.DefaultQuery("limit","20")
	offset := c.DefaultQuery("offset","0")
	params, encSecKey, encErr := utils.Encrypt(fmt.Sprintf(`{"rid":"%s","limit":"%s","offset":"%s","csrf_token":""}`,
		rid,limit,offset))
	if encErr != nil {
		log.Println(encErr)
	}
	res, resErr := utils.DoPostRequest(c.Request.Cookies(),"https://music.163.com/weapi/v1/resource/comments/R_SO_4_"+rid, params, encSecKey)
	if resErr != nil {
		log.Println(resErr)
	}
	c.String(200,res)


}




// MV评论
func  CommentMV(c *gin.Context){


	rid :=c.Query("id")
	limit := c.DefaultQuery("limit","20")
	offset := c.DefaultQuery("offset","0")
	params, encSecKey, encErr := utils.Encrypt(fmt.Sprintf(`{"rid":"%s","limit":"%s","offset":"%s","csrf_token":""}`,
		rid,limit,offset))
	if encErr != nil {
		log.Println(encErr)
	}
	res, resErr := utils.DoPostRequest(c.Request.Cookies(),"https://music.163.com/weapi/v1/resource/comments/R_MV_5_"+rid, params, encSecKey)
	if resErr != nil {
		log.Println(resErr)
	}
	c.String(200,res)


}




// 歌单评论
func  CommentPlayList(c *gin.Context){

	rid :=c.Query("id")
	limit := c.DefaultQuery("limit","20")
	offset := c.DefaultQuery("offset","0")
	params, encSecKey, encErr := utils.Encrypt(fmt.Sprintf(`{"rid":"%s","limit":"%s","offset":"%s","csrf_token":""}`,rid,limit,offset))
	if encErr != nil {
		log.Println(encErr)
	}
	res, resErr := utils.DoPostRequest(c.Request.Cookies(),"https://music.163.com/weapi/v1/resource/comments/A_PL_0_"+rid, params, encSecKey)
	if resErr != nil {
		log.Println(resErr)
	}
	c.String(200,res)


}




// 视频评论
func  CommentVidio(c *gin.Context){

	rid :=c.Query("id")
	limit := c.DefaultQuery("limit","20")
	offset := c.DefaultQuery("offset","0")
	params, encSecKey, encErr := utils.Encrypt(fmt.Sprintf(`{"rid":"%s","limit":"%s","offset":"%s","csrf_token":""}`,rid,limit,offset))
	if encErr != nil {
		log.Println(encErr)
	}
	res, resErr := utils.DoPostRequest(c.Request.Cookies(),"https://music.163.com/weapi/v1/resource/comments/R_VI_62_"+rid, params, encSecKey)
	if resErr != nil {
		log.Println(resErr)
	}
	c.String(200,res)


}


// 发送与删除评论
func  Comment(c *gin.Context){


	content :=c.Query("content")
	commentId :=c.Query("commentId")
	ty :=c.Query("type")
	switch ty {
	case "0":
		ty = "R_SO_4_"
	case "1":
		ty = "R_MV_5_"
	case "2":
		ty = "A_PL_0_"
	case "3":
		ty = "R_AL_3_"
	case "4":
		ty = "A_DJ_1_"
	case "5":
		ty = "R_VI_62_"

	}
	id :=c.Query("id")
	t :=c.Query("t")
	params, encSecKey :="",""
	if t== "1" {
		t = "add"
		params,encSecKey,_=utils.Encrypt(fmt.Sprintf(`{"threadId":"%s","content":"%s","csrf_token":""}`,
			ty+id,content))

	}else {
		t = "delete"
		params,encSecKey,_=utils.Encrypt(fmt.Sprintf(`{"threadId":"%s","commentId":"%s","csrf_token":""}`,
			ty+id,commentId))
	}

	res, resErr := utils.DoPostRequest(c.Request.Cookies(),"https://music.163.com/weapi/resource/comments/"+t, params, encSecKey)
	if resErr != nil {
		log.Println(resErr)
	}
	c.String(200,res)


}
