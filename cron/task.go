package cron

import (
	"sync"
	"time"

	"github.com/tsubasa597/BILIBILI-HELPER/info"
)

// Tasker 任务接口
type Tasker interface {
	Run(chan<- interface{}, *sync.WaitGroup)
	Next(time.Time) time.Time
	Info() info.Task
}
