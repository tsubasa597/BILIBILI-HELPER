package cron

import (
	"sync"
	"time"

	"github.com/tsubasa597/BILIBILI-HELPER/api/user"
	"go.uber.org/zap"
)

// User 用户信息
type User struct {
	UID      int64
	Name     string
	Face     string
	Sign     string
	timeCell time.Duration
	log      *zap.Logger
	mutex    *sync.RWMutex
}

var (
	_ Tasker = (*User)(nil)
)

// NewUser 初始化
func NewUser(uid int64, timeCell time.Duration, log *zap.Logger) *User {
	if log == nil {
		log = zap.NewExample()
	}

	if info, err := user.GetInfo(uid); err == nil {
		return &User{
			UID:      uid,
			Name:     info.Data.Name,
			Face:     info.Data.Face,
			Sign:     info.Data.Sign,
			timeCell: timeCell,
			log:      log,
			mutex:    &sync.RWMutex{},
		}
	}

	return &User{
		UID:      uid,
		timeCell: timeCell,
		log:      log,
		mutex:    &sync.RWMutex{},
	}
}

// Run 监听用户信息变更
func (u *User) Run() interface{} {
	var us user.Info

	info, err := user.GetInfo(u.UID)
	if err != nil {
		u.log.Error(err.Error())
		return nil
	}

	u.mutex.Lock()

	if u.Name != info.Data.Name {
		u.Name = info.Data.Name
		us.Name = info.Data.Name
	}

	if u.Face != info.Data.Face {
		u.Face = info.Data.Face
		us.Face = info.Data.Face
	}

	if u.Sign != info.Data.Sign {
		u.Sign = info.Data.Sign
		us.Sign = info.Data.Sign
	}

	u.mutex.Unlock()

	return us
}

// Next 下次运行时间
func (user User) Next(t time.Time) time.Time {
	return t.Add(user.timeCell)
}

// Info 监听用户的部分信息
func (user User) Info() Info {
	user.mutex.RLock()
	defer user.mutex.RLocker().Unlock()

	return Info{
		ID:       user.UID,
		Name:     user.Name,
		TimeCell: user.timeCell,
	}
}
