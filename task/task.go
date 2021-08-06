package task

import "github.com/tsubasa597/BILIBILI-HELPER/api"

type Tasker interface {
	Run() string
	Init(api.API)
}
