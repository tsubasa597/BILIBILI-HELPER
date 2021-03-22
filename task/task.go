package task

import (
	"sync"
	"time"

	"github.com/tsubasa597/BILIBILI-HELPER/conf"
)

// DailyInfo 任务信息
type Info struct {
	Level        float64
	NextLevelExp float64
	Slivers      float64
	Coins        float64

	params  []string
	conf    conf.Config
	tasks   []Task
	isLogin bool
	logInfo chan []interface{}
	done    chan int
}

// New 启动日常任务
func New() (status *Info) {
	status = &Info{
		conf:    *conf.Init(),
		tasks:   []Task{},
		logInfo: make(chan []interface{}, 4),
		done:    make(chan int),
	}
	go status.readLog()
	status.UserCheck()
	if status.isLogin {
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

func (status *Info) readLog() {
Log:
	for {
		select {
		case info := <-status.logInfo:
			switch info[0].(string) {
			case "Info":
				conf.Log.Info(info[1:])
			case "Warn":
				conf.Log.Warnln(info[1:])
			case "Error":
				conf.Log.Errorln(info[1:])
			case "Fatal":
				conf.Log.Fatal(info[1:])
			}
		case <-status.done:
			break Log
		}
	}
}

// Tasker 任务
type Tasker interface {
	Run()
}

// Task task 类型的函数
type Task func(v ...string)

// Run task 类型的函数调用
func (t Task) Run(wg *sync.WaitGroup, v ...string) {
	defer wg.Done()
	t(v...)
}

// Start 启动任务
func Start() {
	var wg sync.WaitGroup
	task := New()
	for _, i := range task.tasks {
		// 防止请求过快出错
		time.Sleep(time.Second)
		wg.Add(1)
		go i.Run(&wg, task.params...)
	}

	wg.Wait()
	task.done <- 1
}

// UserCheck 用户检查
func (info *Info) UserCheck() {
	info.isLogin = userCheck(info.logInfo)
}

// DailyVideo 观看视频
func (info *Info) DailyVideo(param ...string) {
	watchVideo(info.logInfo, info.params...)
}

// DailyVideoShare 分享视频
func (info *Info) DailyVideoShare(param ...string) {
	shareVideo(info.logInfo, info.params...)
}

// DailySliver2Coin 银瓜子换硬币信息
func (info *Info) DailySliver2Coin(param ...string) {
	sliver2Coins(info.logInfo, info.params...)
}

// DailyLiveCheckin 直播签到信息
func (info *Info) DailyLiveCheckin(param ...string) {
	info.logInfo <- checkLive(info.params...)
}
