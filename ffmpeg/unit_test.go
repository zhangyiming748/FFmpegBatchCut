package ffmpeg

import (
	"FFmpegBatchCut/util"
	"log"
	"os/exec"
	"path/filepath"
	"strings"
	"testing"
)
func FastClean(root string) {
	files := util.GetFiles(root)
	for _, file := range files {
		ext := filepath.Ext(file)
		var newFile string
		if ext == ".mp4" {
			newFile = strings.TrimSuffix(file, ext) + "fast" + ext
		} else {
			newFile = strings.Replace(file, ext, ".mp4", 1)
		}

		cmd := exec.Command("ffmpeg", "-i", file, "-vcodec", "h264_nvenc", "-acodec", "aac", newFile)
		if out, err := cmd.CombinedOutput(); err != nil {
			log.Fatalf("cmd.Run() failed with %s\t out %s\n", err, string(out))
		} else {
			log.Printf("Successfully converted %s\t out %s\n", newFile, string(out))
		}
	}
}

// go test -v -timeout 10h -run TestFastClean
func TestFastClean(t *testing.T) {
	root := "D:\\pikpak\\另类精品系列の精液狂射10斤\\试运行"
	FastClean(root)
}