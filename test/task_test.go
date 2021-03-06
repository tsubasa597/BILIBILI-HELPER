package bilitest

import (
	"bili/task"
	"testing"
)

var (
	testTask task.Response
	testInfo task.Status
)

func TestLiveCheckin(t *testing.T) {
	//testTask.LiveCheckin()
	//t.Log(testTask)
}

func TestVideoWatch(t *testing.T) {
	//testTask.VideoWatch("BV1NT4y137Jc")
	//t.Log(testTask)
}

func TestUserCheck(t *testing.T) {
	//testTask.UserCheck()
	//t.Log(testTask)
}

func TestSliverCoins(t *testing.T) {
	//testTask.Sliver2CoinsStatus()
	//t.Log(testTask)
	//testTask.Sliver2Coins()
	//t.Log(testTask)
}

func TestLiveCheckinInfo(t *testing.T) {
	// testInfo.LiveCheckinInfo(&testTask)
}

func TestSliverCoinsInfo(t *testing.T) {
	// testInfo.Sliver2CoinsInfo(&testTask)
}

func TestUserCheckInfo(t *testing.T) {
	testInfo.UserCheck(&testTask)
}
