package task

// Info 任务信息
type Info interface {
}

// Status 任务信息
type Status struct {
	Level                int
	NextLevelExp         int
	IsLogin              bool
	IsVideoWatch         bool
	IsLiveCheckin        bool
	IsSliver2CoinsStatus bool
	Slivers              int
	Coins                int
}
