package main

import (
	"bili/config"
	"bili/task"
	"fmt"
)

var (
	status task.Status = task.Status{}
)

func main() {
	fmt.Println(config.Conf)
}
