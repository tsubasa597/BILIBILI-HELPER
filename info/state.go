package info

// State 运行状态
type State int32

const (
	StateStop State = iota
	// StatePause 暂停状态
	StatePause
	// StateRuning 正在运行状态
	StateRuning
)
