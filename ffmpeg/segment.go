// 视频切割相关功能的实现
package ffmpeg

import (
	"FFmpegBatchCut/util"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
)

// CutBySegments 根据给定的片段列表切割视频文件
// mp4: 输入视频文件路径
// segments: 切割片段列表
func CutBySegments(mp4 string, segments []util.Segment) error {
	for i, segment := range segments {
		// 将 i+1 转换为两位数的字符串，不足两位前面补0
		index := fmt.Sprintf("%02d", i+1)
		// 构造输出文件名，格式为 "01.mp4"
		start := util.FormatSecondToHMS(segment.Start)
		end := util.FormatSecondToHMS(segment.End)
		// 调用 CutBySegment 函数进行切割
		if err := CutBySegment(index, mp4, start, end); err != nil {
			return fmt.Errorf("Cut File: %s By Segment error: %v", mp4, err)
		}
	}
	return nil
}

// CutBySegment 执行单个视频片段的切割
// index: 输出文件的序号（两位数字）
// mp4: 输入视频文件路径
// start: 开始时间点
// end: 结束时间点
func CutBySegment(index, mp4, start, end string) error {
	out := filepath.Join(filepath.Dir(mp4), index+".mp4")
	cmd := exec.Command("ffmpeg")
	if hostname, _ := os.Hostname(); hostname == "DESKTOP-VGFTVD8" {
		cmd.Args = append(cmd.Args, "-hwaccel", "cuda")
		cmd.Args = append(cmd.Args, "-i", mp4)
		if start != "00:00:00.000" {
			cmd.Args = append(cmd.Args, "-ss", start)
		}
		if end != "00:00:00.000" {
			cmd.Args = append(cmd.Args, "-to", end)
		}
		cmd.Args = append(cmd.Args, "-c:v", "h264_nvenc")
		cmd.Args = append(cmd.Args, "-preset", "slow")
		cmd.Args = append(cmd.Args, "-cq", "18")
	} else {
		cmd.Args = append(cmd.Args, "-i", mp4)
		if start != "00:00:00.000" {
			cmd.Args = append(cmd.Args, "-ss", start)
		}
		if end != "00:00:00.000" {
			cmd.Args = append(cmd.Args, "-to", end)
		}
		cmd.Args = append(cmd.Args, "-c:v", "libx265")
		cmd.Args = append(cmd.Args, "-tag:v", "hvc1")
	}
	cmd.Args = append(cmd.Args, "-c:a", "aac")
	cmd.Args = append(cmd.Args, "-map_metadata", "-1")
	cmd.Args = append(cmd.Args, "-vsync", "0") // 添加这行
	cmd.Args = append(cmd.Args, "-copyts")     // 添加这行
	cmd.Args = append(cmd.Args, out)
	err := util.Exec(cmd)
	if err != nil {
		return err
	}
	return nil
}
