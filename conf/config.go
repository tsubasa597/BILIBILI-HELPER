package conf

import (
	_ "embed"
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

var (
	Log *logrus.Logger
)

func init() {
	Log = logrus.New()
	Log.SetFormatter(&nested.Formatter{
		HideKeys:    true,
		FieldsOrder: []string{"component", "category"},
	})
}
