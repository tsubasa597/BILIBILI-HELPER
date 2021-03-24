package task

import (
	_ "embed"
	"io/ioutil"
	"log"

	nested "github.com/antonfisher/nested-logrus-formatter"
	"github.com/sirupsen/logrus"
	"gopkg.in/yaml.v3"
)

// Cookie 用于登录的必要参数
type Cookie struct {
	UserID   string `yaml:"userId"`
	SessData string `yaml:"sessData"`
	BiliJct  string `yaml:"biliJct"`
}

// GetVerify 将参数进行组合
func (cookie Cookie) GetVerify() string {
	// TODO 修改组合方式
	return "bili_jct=" + cookie.BiliJct + ";SESSDATA=" + cookie.SessData + ";DedeUserID=" + cookie.UserID + ";"
}

// config 配置
type config struct {
	Cookie
}

func NewConfig(path string) *config {
	conf := &config{}
	yaml.Unmarshal(loadYamlFile(path), conf)
	if conf == (&config{}) {
		log.Fatalln("无法读取数据")
	}
	return conf
}

func newLogFormat() *logrus.Logger {
	log := logrus.New()
	log.SetFormatter(&nested.Formatter{
		HideKeys:    true,
		FieldsOrder: []string{"component", "category"},
	})
	return log
}

func loadYamlFile(path string) []byte {
	file, err := ioutil.ReadFile(path)
	if err != nil {
		log.Fatalln("无法读取文件", err)
	}
	return file
}
