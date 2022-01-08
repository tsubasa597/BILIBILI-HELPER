// Package info 监听返回的消息内容
package info

import "time"

// Info 监听返回值的基础结构体
type Info struct {
	Name string
	Time int64
}

const (
	PauseDuration, TwoDay time.Duration = time.Minute * 30, time.Hour * 24 * 2
)
