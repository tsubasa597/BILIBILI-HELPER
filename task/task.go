package task

import (
	"sync"
	"time"
)

// TODO: 参数传递方式

// DailyInfo 任务信息
type Info struct {
	Level        float64
	NextLevelExp float64
	Slivers      float64
	Coins        float64
	Tasks        []TaskFunc

	*API
	params  map[string]string
	isLogin bool
	logInfo chan []interface{}
	done    chan int
}

// Default 启动日常任务
func Default() (status *Info) {
	status = &Info{
		logInfo: make(chan []interface{}, 4),
		done:    make(chan int),
		params:  make(map[string]string),
		API:     newApi(*NewConfig("./conf.yaml")),
	}

	go status.readLog()

	status.UserCheck()
	if status.isLogin {
		status.Tasks = append(status.Tasks, TaskFunc(status.DailyLiveCheckin))
		status.Tasks = append(status.Tasks, TaskFunc(status.DailySliver2Coin))
		status.params["bvid"] = "BV1NT4y137Jc"
		status.Tasks = append(status.Tasks, TaskFunc(status.DailyVideo))
		status.params["bvid"] = "BV1NT4y137Jc"
		status.Tasks = append(status.Tasks, TaskFunc(status.DailyVideoShare))
	}
	return
}

func (status *Info) readLog() {
	loger := newLogFormat()

Log:
	for {
		select {
		case info := <-status.logInfo:
			switch info[0].(string) {
			case "Info":
				loger.Info(info[1:])
			case "Warn":
				loger.Warnln(info[1:])
			case "Error":
				loger.Errorln(info[1:])
			case "Fatal":
				loger.Fatal(info[1:])
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
type TaskFunc func(v map[string]string)

// Run task 类型的函数调用
func (t TaskFunc) Run(wg *sync.WaitGroup, v map[string]string) {
	defer wg.Done()
	t(v)
}

// Start 启动任务
func Start(task *Info) {
	var wg sync.WaitGroup

	for _, i := range task.Tasks {
		// 防止请求过快出错
		time.Sleep(time.Second)
		wg.Add(1)
		go i.Run(&wg, task.params)
	}

	wg.Wait()
	task.done <- 1
}

// UserCheck 用户检查
func (info *Info) UserCheck() {
	info.isLogin = info.userCheck(info.logInfo, nil)
}

// DailyVideo 观看视频
func (info *Info) DailyVideo(param map[string]string) {
	info.watchVideo(info.logInfo, info.params)
}

// DailyVideoShare 分享视频
func (info *Info) DailyVideoShare(param map[string]string) {
	info.shareVideo(info.logInfo, info.params)
}

// DailySliver2Coin 银瓜子换硬币信息
func (info *Info) DailySliver2Coin(param map[string]string) {
	info.sliver2Coins(info.logInfo, nil)
}

// DailyLiveCheckin 直播签到信息
func (info *Info) DailyLiveCheckin(param map[string]string) {
	info.checkLive(info.logInfo, nil)
}
