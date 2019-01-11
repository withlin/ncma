package config

import (
	"fmt"
	"io/ioutil"
	"os"

	yaml "gopkg.in/yaml.v2"
)

var jsonData map[string]interface{}

// ESConfi
type conf struct {
	GoConf struct {
		APIPrefix   string `yaml:"apiprefix"`
		Port        int    `yaml:"port"`
		TokenMaxAge int    `yaml:"tokenmaxage"`
		Env         string `yaml:"env"`
		LogDir      string `yaml:"logdir"`
		TokenSecret  string `yaml:"tokensecret"`
	} `yaml:"go"`
}

//Conf 配置文件
var Conf conf

func initYaml() {
	bytes, err := ioutil.ReadFile("./config.yml")
	if err != nil {
		fmt.Println("ReadFile: ", err.Error())
		os.Exit(-1)
	}
	err = yaml.Unmarshal(bytes, &Conf)
	if err != nil {
		fmt.Println("Yml File Unmarshal :", err.Error())
	}
}


//InitAll 初始化全部的数据
func init() {
	initYaml()
}