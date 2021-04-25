package main

import (
	"fmt"

	"github.com/tsubasa597/BILIBILI-HELPER/global"
	"github.com/tsubasa597/BILIBILI-HELPER/task"
)

func main() {
	c, err := global.NewConfig("./config.yaml")
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(task.Run(task.New(*c)))
}
