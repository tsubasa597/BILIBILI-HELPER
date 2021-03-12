package config

import (
	"bili/verify"
	"io/ioutil"
	"log"
	"os"

	"gopkg.in/yaml.v3"
)

// Config 配置
type Config struct {
	Cookie    verify.Cookie
	UserAgent string `yaml:"userAgent"`
	Status    TaskStatus
}

// TaskStatus 任务信息
type TaskStatus struct {
	IsVideoWatch   bool `yaml:"isVideoWatch"`
	IsVideoShare   bool `yaml:"isVideoShare"`
	IsLiveCheckin  bool `yaml:"isLiveCheckin"`
	IsSliver2Coins bool `yaml:"isSliver2Coins"`
}

var (
	// Conf 公共配置
	Conf *Config = &Config{}
)

func init() {
	pwd, _ := os.Getwd()
	yamlFile, err := ioutil.ReadFile(pwd + "/conf.yaml")
	if err != nil {
		log.Fatalln(err.Error())
	}
	yaml.Unmarshal(yamlFile, Conf)
	if Conf == (&Config{}) {
		log.Fatalln("无法读取数据")
	}
}
