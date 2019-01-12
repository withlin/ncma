package utils

import (
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"net/http/cookiejar"
	"net/url"
	"strings"
	"time"
)

var userAgentList = [19]string {
	"Mozilla/5.0 (Macintosh; Intel Mac OS X 10_12_5) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/59.0.3071.115 Safari/537.36",
	"Mozilla/5.0 (iPhone; CPU iPhone OS 9_1 like Mac OS X) AppleWebKit/601.1.46 (KHTML, like Gecko) Version/9.0 Mobile/13B143 Safari/601.1",
	"Mozilla/5.0 (iPhone; CPU iPhone OS 9_1 like Mac OS X) AppleWebKit/601.1.46 (KHTML, like Gecko) Version/9.0 Mobile/13B143 Safari/601.1",
	"Mozilla/5.0 (Linux; Android 5.0; SM-G900P Build/LRX21T) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/59.0.3071.115 Mobile Safari/537.36",
	"Mozilla/5.0 (Linux; Android 6.0; Nexus 5 Build/MRA58N) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/59.0.3071.115 Mobile Safari/537.36",
	"Mozilla/5.0 (Linux; Android 5.1.1; Nexus 6 Build/LYZ28E) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/59.0.3071.115 Mobile Safari/537.36",
	"Mozilla/5.0 (iPhone; CPU iPhone OS 10_3_2 like Mac OS X) AppleWebKit/603.2.4 (KHTML, like Gecko) Mobile/14F89;GameHelper",
	"Mozilla/5.0 (Macintosh; Intel Mac OS X 10_12_5) AppleWebKit/603.2.4 (KHTML, like Gecko) Version/10.1.1 Safari/603.2.4",
	"Mozilla/5.0 (iPhone; CPU iPhone OS 10_0 like Mac OS X) AppleWebKit/602.1.38 (KHTML, like Gecko) Version/10.0 Mobile/14A300 Safari/602.1",
	"Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/51.0.2704.103 Safari/537.36",
	"Mozilla/5.0 (Macintosh; Intel Mac OS X 10.12; rv:46.0) Gecko/20100101 Firefox/46.0",
	"Mozilla/5.0 (Windows NT 10.0; Win64; x64; rv:46.0) Gecko/20100101 Firefox/46.0",
	"Mozilla/4.0 (compatible; MSIE 7.0; Windows NT 6.0)",
	"Mozilla/4.0 (compatible; MSIE 8.0; Windows NT 6.0; Trident/4.0)",
	"Mozilla/5.0 (compatible; MSIE 9.0; Windows NT 6.1; Trident/5.0)",
	"Mozilla/5.0 (compatible; MSIE 10.0; Windows NT 6.2; Win64; x64; Trident/6.0)",
	"Mozilla/5.0 (Windows NT 6.3; Win64, x64; Trident/7.0; rv:11.0) like Gecko",
	"Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/42.0.2311.135 Safari/537.36 Edge/13.10586",
	"Mozilla/5.0 (iPad; CPU OS 10_0 like Mac OS X) AppleWebKit/602.1.38 (KHTML, like Gecko) Version/10.0 Mobile/14A300 Safari/602.1",
}

func DoPostRequest(cookies []*http.Cookie,address, params, encSecKey string)(string ,error){
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	result := url.Values{}
	result.Set("params", params)
	result.Set("encSecKey", encSecKey)
	body := strings.NewReader(result.Encode())
	j, err := cookiejar.New(nil)
	if err!=nil {
		return "",nil
	}
	client := &http.Client{Jar: j}
	request, err:= http.NewRequest("POST", address, body)
	if err!=nil {
		return "",nil
	}
	reqUrl, err := url.Parse(address)
	if err!=nil {
		return "",nil
	}
	j.SetCookies(reqUrl,cookies)
	log.Println(request.Cookies())
	request.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	request.Header.Set("Referer", "http://music.163.com")
	request.Header.Set("Accept", "*/*")
	request.Header.Set("Accept-Language","zh-CN,zh;q=0.8,gl;q=0.6,zh-TW;q=0.4")
	request.Header.Set("Cookie","appver=2.0.2")
	request.Header.Set("User-Agent", userAgentList[r.Intn(19)])
	res , err := client.Do(request)
	defer res.Body.Close()
	if err != nil {
		return  "" ,err
	}
	resBody,err := ioutil.ReadAll(res.Body)
	if err !=nil{
		return "",err
	}
	return string(resBody), nil
}


func DoGetRequest(cookies []*http.Cookie,address string) (string,error){
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	j, err := cookiejar.New(nil)
	if err!=nil {
		return "",nil
	}
	client := &http.Client{Jar: j}
	reqUrl, err := url.Parse(address)
	if err!=nil {
		return "",nil
	}
	j.SetCookies(reqUrl,cookies)
	request,err :=http.NewRequest("GET",address,nil)
	request.Header.Set("Referer", "http://music.163.com")
	request.Header.Set("User-Agent", userAgentList[r.Intn(19)])
	res, err := client.Do(request)
	defer res.Body.Close()
	// 错误处理
	if err != nil {
		return "", err
	}
	rs, err := ioutil.ReadAll(res.Body)
	return string(rs), err
}
