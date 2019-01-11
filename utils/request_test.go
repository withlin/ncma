package utils

import (
	"log"
	"testing"
)

func TestDoRequest(t *testing.T){
	param :=`{"id":"72029984","csrf_token":""}`
	params, encSecKey, encErr := Encrypt(param)
	if encErr != nil {
		log.Println(encErr)
	}
	res, resErr := DoRequest("http://music.163.com/weapi/v1/album/72029984", params, encSecKey)
	if resErr != nil {
		log.Println(resErr)
	}
	log.Println(res)
}