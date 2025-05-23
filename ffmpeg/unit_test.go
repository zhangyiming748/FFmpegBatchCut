package ffmpeg

import (
	"FFmpegBatchCut/util"
	"log"
	"testing"
)

func init() {
	util.SetLog("BitchCut.log")
	log.SetFlags(2 | 16)
}

// go test -v -timeout 10h -run TestFastClean
func TestFastClean(t *testing.T) {
	root := "D:\\pikpak"
	FastClean(root)
}

func TestZero(t *testing.T) {
	ret := util.FormatSecondToHMS(0)
	t.Log(ret)
	//00:00:00.000
}
