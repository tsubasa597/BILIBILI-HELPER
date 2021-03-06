package config

import (
	"bili/login"
	"fmt"
	"io/ioutil"

	"gopkg.in/yaml.v3"
)

// Config 配置
type Config struct {
	Cookie    login.Cookie
	UserAgent string `yaml:"userAgent"`
}

var (
	Conf *Config = &Config{}
)

func init() {
	yamlFile, err := ioutil.ReadFile("./conf.yaml")
	if err != nil {
		fmt.Println(err.Error())
	}
	yaml.Unmarshal(yamlFile, Conf)
}
