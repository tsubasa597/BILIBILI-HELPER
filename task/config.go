package task

import (
	_ "embed"
	"log"

	nested "github.com/antonfisher/nested-logrus-formatter"
	"github.com/sirupsen/logrus"
	"gopkg.in/yaml.v3"
)

var (
	//go:embed conf.yaml
	yamlFile []byte

	loger *logrus.Logger = newLogFormat()
	conf  *config        = newConf()
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
	Cookie    Cookie
	UserAgent string `yaml:"userAgent"`
	Status    Status
}

// TaskStatus 任务信息
type Status struct {
	IsVideoWatch   bool `yaml:"isVideoWatch"`
	IsVideoShare   bool `yaml:"isVideoShare"`
	IsLiveCheckin  bool `yaml:"isLiveCheckin"`
	IsSliver2Coins bool `yaml:"isSliver2Coins"`
}

func newConf() *config {
	conf := &config{}
	yaml.Unmarshal(yamlFile, conf)
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
