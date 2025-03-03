package ffmpeg

import (
	"FFmpegBatchCut/util"
	"fmt"
	"log"
	"os"
	"os/exec"
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
		return fmt.Errorf("给定的时间戳文件:%v格式非法", timestamps)
		//error strings should not end with punctuation or newlines (ST1005)
	}
	timestamps = formatTimestamps(timestamps)
	fname := fp
	folder := strings.Split(fname, ".")[0]
	_ = os.Mkdir(folder, 0777)
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
		cmd := exec.Command("ffmpeg")
		if hostname, _ := os.Hostname(); hostname == "DESKTOP-VGFTVD8" {
			cmd.Args = append(cmd.Args, "-hwaccel", "cuda")
			cmd.Args = append(cmd.Args, "-i", fname)
			cmd.Args = append(cmd.Args, "-ss", timestamps[i])
			cmd.Args = append(cmd.Args, "-to", timestamps[i+1])
			cmd.Args = append(cmd.Args, "-c:v", "h264_nvenc")
			cmd.Args = append(cmd.Args, "-c:a", "libmp3lame")
			cmd.Args = append(cmd.Args, "-preset", "slow")
			cmd.Args = append(cmd.Args, "-cq", "18")
			cmd.Args = append(cmd.Args, "-map_metadata", "-1")
			cmd.Args = append(cmd.Args, mp4)
		} else {
			cmd.Args = append(cmd.Args, "-i", fname)
			cmd.Args = append(cmd.Args, "-ss", timestamps[i])
			cmd.Args = append(cmd.Args, "-to", timestamps[i+1])
			cmd.Args = append(cmd.Args, "-c:v", "libx265")
			cmd.Args = append(cmd.Args, "-tag:v", "hvc1")
			cmd.Args = append(cmd.Args, "-c:a", "libmp3lame")
			cmd.Args = append(cmd.Args, "-map_metadata", "-1")
			cmd.Args = append(cmd.Args, mp4)

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
	cmd := exec.Command("ffmpeg", "-i", fname, "-ss", timestamps[length-1], "-c:v", "libx265", "-c:a", "libopus", "-tag:v", "hevc", "-ac", "1", mp4)
	if hostname, _ := os.Hostname(); hostname == "DESKTOP-VGFTVD8" {
		cmd = exec.Command("ffmpeg", "-hwaccel", "cuda", "-i", fname, "-ss", timestamps[length-1], "-c:v", "h264_nvenc", "-c:a", "libopus", "-ac", "1", "-preset", "medium", "-cq", "20", mp4)
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
