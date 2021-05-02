package main

import (
	"fmt"

	"github.com/tsubasa597/BILIBILI-HELPER/daily"
	"github.com/tsubasa597/BILIBILI-HELPER/global"
)

func main() {
	c, err := global.NewConfig("./config.yaml")
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(daily.Run(daily.New(*c)))
}
