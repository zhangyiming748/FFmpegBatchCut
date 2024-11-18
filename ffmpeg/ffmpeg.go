package ffmpeg

import (
	"FFmpegBatchCut/util"
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"strconv"
	"strings"
)

var (
	OperatingSystem string
	Architecture    string
)

func init() {
	OperatingSystem = runtime.GOOS
	Architecture = runtime.GOARCH
}

/*
输入文件名和时间点切片
*/
func CutOne(fp string, timestamps []string) (err error) {
	defer func() {
		log.Println("运行完成")
	}()
	if timestamps[0] != "000000000" {
		newElement := "000000000"
		newSlice := append([]string{newElement}, timestamps...)
		timestamps = newSlice
	}
	timestamps = formatTimestamps(timestamps)
	fname := fp
	folder := strings.Split(fname, ".")[0]
	os.Mkdir(folder, 0777)
	if !strings.HasSuffix(fname, "mp4") && !strings.HasSuffix(fname, "mkv") {
		log.Printf("开始转换%v为mp4标准格式\n", fname)
		mp4 := strings.Replace(fname, filepath.Ext(fname), ".mp4", -1)
		log.Printf("命令原文%v\n", exec.Command("ffmpeg", "-hwaccel", "cuda", "-i", fname, "-c:v", "h264_nvenc", "-c:a", "libopus", "-ac", "1", "-preset", "medium", "-cq", "20", mp4).String())
		cmd := exec.Command("ffmpeg", "-hwaccel", "cuda", "-i", fname, "-c:v", "h264_nvenc", "-c:a", "libopus", "-ac", "1", "-preset", "medium", "-cq", "20", mp4)
		// ffmpeg -hwaccel cuda -i -c:v h264_nvenc -preset medium -cq 20
		if OperatingSystem == "darwin" && Architecture == "amd64" {
			cmd = exec.Command("ffmpeg", "-i", fname, "-c:v", "libx265", "-tag:v", "hevc", "-c:a", "libopus", "-ac", "1", mp4)
		}
		err = util.Exec(cmd)
		if err != nil {
			return err
		}
		return
	}
	length := len(timestamps)
	for i := 0; i < length-1; i++ {
		mp4 := strings.Join([]string{strconv.Itoa(i + 1), "mp4"}, ".")
		mp4 = strings.Join([]string{folder, mp4}, string(os.PathSeparator))
		log.Printf("命令原文:%s\n", exec.Command("ffmpeg", "-hwaccel", "cuda", "-i", fname, "-ss", timestamps[i], "-to", timestamps[i+1], "-c:v", "h264_nvenc", "-c:a", "libopus", "-ac", "1", "-preset", "medium", "-cq", "20", mp4).String())
		cmd := exec.Command("ffmpeg", "-hwaccel", "cuda", "-i", fname, "-ss", timestamps[i], "-to", timestamps[i+1], "-c:v", "h264_nvenc", "-c:a", "libopus", "-ac", "1", "-preset", "medium", "-cq", "20", mp4)
		if OperatingSystem == "darwin" && Architecture == "amd64" {
			cmd = exec.Command("ffmpeg", "-i", fname, "-ss", timestamps[i], "-to", timestamps[i+1], "-c:v", "libx265", "-tag:v", "hevc", "-c:a", "libopus", "-ac", "1", mp4)
		}
		err = util.Exec(cmd)
		if err != nil {
			return err
		}
	}
	mp4 := strings.Join([]string{strconv.Itoa(length), "mp4"}, ".")
	mp4 = strings.Join([]string{folder, mp4}, string(os.PathSeparator))
	log.Printf("命令原文:%s\n", exec.Command("ffmpeg", "-hwaccel", "cuda", "-i", fname, "-ss", timestamps[length-1], "-c:v", "h264_nvenc", "-c:a", "libopus", "-ac", "1", "-preset", "medium", "-cq", "20", mp4).String())
	cmd := exec.Command("ffmpeg", "-hwaccel", "cuda", "-i", fname, "-ss", timestamps[length-1], "-c:v", "h264_nvenc", "-c:a", "libopus", "-ac", "1", "-preset", "medium", "-cq", "20", mp4)
	if OperatingSystem == "darwin" && Architecture == "amd64" {
		cmd = exec.Command("ffmpeg", "-i", fname, "-ss", timestamps[length-1], "-c:v", "libx265", "-c:a", "libopus", "-tag:v", "hevc", "-ac", "1", mp4)
	}
	err = util.Exec(cmd)
	if err != nil {
		return err
	}
	return nil
}

func formatTimestamps(timestamps []string) []string {
	var formatted []string
	for _, ts := range timestamps {
		// 将字符串分割为小时、分钟、秒和毫秒
		hours := ts[0:2]
		minutes := ts[2:4]
		seconds := ts[4:6]
		milliseconds := ts[6:9]

		// 格式化为所需的格式
		formattedTimestamp := fmt.Sprintf("%s:%s:%s.%s", hours, minutes, seconds, milliseconds)
		formatted = append(formatted, formattedTimestamp)
	}
	return formatted
}
