package ffmpeg

import (
	"fmt"
	"log"
	"os/exec"
	"path/filepath"
	"strings"
)

func AnyVideoToMP4(file string) error {
	out := strings.Replace(file, filepath.Ext(file), ".mp4", 1)
	cmd := exec.Command("ffmpeg", "-i", file, "-c:v", "h264_nvenc", out)
	log.Printf("执行命令:%v\n", cmd.String())
	o, err := cmd.CombinedOutput()
	if err != nil {
		return fmt.Errorf("执行命令%v失败:%v\n%v", cmd.String(), err, string(o))
	}
	fmt.Println(string(o))
	return nil
}
