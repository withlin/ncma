package album

import (
	"fmt"
	"github.com/WithLin/ncma/utils"
	"github.com/gin-gonic/gin"
	"log"
)

func Album(c *gin.Context) {
	id := c.Param("id")
	params, encSecKey, encErr := utils.Encrypt(fmt.Sprintf(`{"id":"%s","csrf_token":""}`, id))
	if encErr != nil {
		log.Println(encErr)
	}
	res, resErr := utils.DoRequest("http://music.163.com/weapi/v1/album/"+id, params, encSecKey)
	if resErr != nil {
		log.Println(resErr)
	}
	c.String(200, `{"code":200,"album":%s}`,res)
}