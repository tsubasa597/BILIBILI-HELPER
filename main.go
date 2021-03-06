package main

import (
	"bili/task"
	"fmt"
)

var (
	jrs    *task.JSONResponse = &task.JSONResponse{}
	status task.Status        = task.Status{}
)

func init() {
	status.UserCheck(jrs)
	status.DailyLiveCheckin(jrs)
	status.DailySliver2Coin(jrs)
	status.DailyVideo(jrs)
}

func main() {
	fmt.Println(status)
}
