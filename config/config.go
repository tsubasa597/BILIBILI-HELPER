package config

import (
	_ "embed"
	"log"

	"gopkg.in/yaml.v3"
)

// Config 配置
type Config struct {
	Cookie    Cookie
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

//go:embed conf.yaml
var yamlFile []byte

func Init() *Config {
	var conf *Config = &Config{}
	yaml.Unmarshal(yamlFile, conf)
	if conf == (&Config{}) {
		log.Fatalln("无法读取数据")
	}
	return conf
}
