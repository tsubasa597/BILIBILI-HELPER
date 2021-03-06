package main

import (
	"bili/task"
)

var (
	rs     *task.Response = &task.Response{}
	status task.Status    = task.Status{}
)

func init() {
	status.UserCheck(rs)
	status.DailyLiveCheckin(rs)
	status.DailySliver2Coin(rs)
	status.DailyVideo(rs)
}

func main() {
	// fmt.Println(status)
}
