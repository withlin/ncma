package signin

import (
	"fmt"
	"github.com/WithLin/ncma/utils"
	"github.com/gin-gonic/gin"
	"log"
)


// 签到

/*
    0为安卓端签到 3点经验, 1为网页签到,2点经验
    签到成功 {'android': {'point': 3, 'code': 200}, 'web': {'point': 2, 'code': 200}}
    重复签到 {'android': {'code': -2, 'msg': '重复签到'}, 'web': {'code': -2, 'msg': '重复签到'}}
    未登录 {'android': {'code': 301}, 'web': {'code': 301}}
*/
func DailySignin(c *gin.Context){


	ty :=c.DefaultQuery("type","0")
	params, encSecKey, encErr := utils.Encrypt(fmt.Sprintf(`{"type":%s`,ty))
	if encErr != nil {
		log.Println(encErr)
	}

	res, resErr := utils.DoPostRequest(c.Request.Cookies(),"https://music.163.com/weapi/point/dailyTask", params, encSecKey)
	if resErr != nil {
		log.Println(resErr)
	}
	c.String(200, res)

}