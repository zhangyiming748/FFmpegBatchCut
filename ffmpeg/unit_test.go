package ffmpeg

import (
	"FFmpegBatchCut/util"
	"log"
	"testing"
)

func init() {
	util.SetLog("BitchCut.log")
	log.SetFlags(2|16)
}


// go test -v -timeout 10h -run TestFastClean
func TestFastClean(t *testing.T) {
	root := "D:\\pikpak\\另类精品系列の精液狂射10斤\\试运行"
	FastClean(root)
}