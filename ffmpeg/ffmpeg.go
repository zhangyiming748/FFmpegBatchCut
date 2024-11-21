package ffmpeg

import (
	"FFmpegBatchCut/util"
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"regexp"
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
	if !IsValidate(timestamps) {
		return fmt.Errorf("给定的时间戳文件:%v格式非法\n", timestamps)
	}
	timestamps = formatTimestamps(timestamps)
	fname := fp
	folder := strings.Split(fname, ".")[0]
	_ = os.Mkdir(folder, 0777)
	if !strings.HasSuffix(fname, "mp4") {
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
	log.Printf("时间戳%v\n", timestamps)
	for i := 0; i < length-1; i++ {
		var index string
		if i < 10 {
			index = fmt.Sprintf("%02d", i+1)
		} else {
			index = fmt.Sprintf("%02d", i+1)
		}
		mp4 := strings.Join([]string{index, "mp4"}, ".")
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
	var last string
	if length < 10 {
		last = fmt.Sprintf("%02d", length)
	} else {
		last = fmt.Sprintf("%02d", length)
	}
	mp4 := strings.Join([]string{last, "mp4"}, ".")
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
func IsValidate(timestamps []string) bool {
	var s []int
	for index, v := range timestamps {
		i, err := strconv.Atoi(v)
		if err != nil {
			log.Printf("可能包含非数字字符:%v\n", v)
			return false
		}
		if isNineDigitNumber(v) {
			//fmt.Printf("%s 是九位纯数字\n", v)
		} else {
			log.Printf("第%d行的数字有问题:%s不是九位纯数字\n", index+2, v)
			return false
		}
		s = append(s, i)
	}
	for i := 0; i < len(s)-1; i++ {
		if s[i] > s[i+1] {
			log.Printf("第%v行的数字有问题:%s\n", i+2, timestamps[i+1])
			return false
		}
	}
	return true
}
func isNineDigitNumber(str string) bool {
	match, _ := regexp.MatchString("^[0-9]{9}$", str)
	return match
}
