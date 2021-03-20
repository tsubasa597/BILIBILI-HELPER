package task

import (
	"bili/config"
)

// DailyInfo 任务信息
type Daily struct {
	conf         config.Config
	tasks        []Task
	params       []string
	Level        float64
	NextLevelExp float64
	IsLogin      bool
	Slivers      float64
	Coins        float64
	logInfo      chan []interface{}
	done         chan int
}

// New 启动日常任务
func New() (status *Daily) {
	status = &Daily{
		conf:    *config.Init(),
		tasks:   []Task{},
		logInfo: make(chan []interface{}, 4),
		done:    make(chan int),
	}
	go status.readLog()
	status.UserCheck()
	if status.IsLogin {
		if status.conf.Status.IsLiveCheckin {
			status.tasks = append(status.tasks, Task(status.DailyLiveCheckin))
		}
		if status.conf.Status.IsSliver2Coins {
			status.tasks = append(status.tasks, Task(status.DailySliver2Coin))
		}
		if status.conf.Status.IsVideoWatch {
			status.params = []string{"BV1NT4y137Jc"}
			status.tasks = append(status.tasks, Task(status.DailyVideo))
		}
		if status.conf.Status.IsVideoShare {
			status.params = []string{"BV1NT4y137Jc"}
			status.tasks = append(status.tasks, Task(status.DailyVideoShare))
		}
	}
	return status
}

func (status *Daily) readLog() {
Log:
	for {
		select {
		case info := <-status.logInfo:
			switch info[0].(string) {
			case "Info":
				config.Log.Info(info[1:])
			case "Warn":
				config.Log.Warnln(info[1:])
			case "Error":
				config.Log.Errorln(info[1:])
			case "Fatal":
				config.Log.Fatal(info[1:])
			}
		case <-status.done:
			break Log
		}
	}
}
