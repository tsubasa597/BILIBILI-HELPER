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
	status.UserCheckInfo(jrs)
	status.LiveCheckinInfo(jrs)
	status.Sliver2CoinsInfo(jrs)
	status.VideoWatchInfo(jrs)
}

func main() {
	fmt.Println(status)
}
