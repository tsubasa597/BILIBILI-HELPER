package global

import (
	"fmt"
	"io/ioutil"

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
	return fmt.Sprintf("bili_jct=%s;SESSDATA=%s;DedeUserID=%s;", cookie.BiliJct, cookie.SessData, cookie.UserID)
}

func NewConfig(path string) (*Cookie, error) {
	file, err := loadYamlFile(path)
	if err != nil {
		return nil, err
	}

	c := &Cookie{}
	err = yaml.Unmarshal(file, c)
	if err != nil {
		return nil, err
	}

	return c, nil
}

func loadYamlFile(path string) ([]byte, error) {
	file, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}
	return file, nil
}
