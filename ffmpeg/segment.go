package ffmpeg
import (
	"FFmpegBatchCut/util"
	"fmt"
	"log"
	"os"
	"os/exec"
)

func CutBySegments(mp4 string, segments []util.Segment)  {
	for i, segment := range segments {
		// 将 i+1 转换为两位数的字符串，不足两位前面补0
		index := fmt.Sprintf("%02d", i+1)
		// 构造输出文件名，格式为 "01.mp4"
		start:=util.FormatSecondToHMS(segment.Start)
		end:=util.FormatSecondToHMS(segment.End)
		// 调用 CutBySegment 函数进行切割
		if err:= CutBySegment(index,mp4,start,end); err!= nil {
			log.Printf("Cut File: %s By Segment error: %v",mp4, err)
		}
	}
}

func CutBySegment(index,mp4 ,start,end string) error {
	cmd := exec.Command("ffmpeg")
	if hostname, _ := os.Hostname(); hostname == "DESKTOP-VGFTVD8" {
		cmd.Args = append(cmd.Args, "-hwaccel", "cuda")
		cmd.Args = append(cmd.Args, "-i", mp4)
		cmd.Args = append(cmd.Args, "-ss", start)
		cmd.Args = append(cmd.Args, "-to", end)
		cmd.Args = append(cmd.Args, "-c:v", "h264_nvenc")
		cmd.Args = append(cmd.Args, "-c:a", "aac")
		cmd.Args = append(cmd.Args, "-preset", "slow")
		cmd.Args = append(cmd.Args, "-cq", "18")
	} else {
		cmd.Args = append(cmd.Args, "-i", mp4)
		cmd.Args = append(cmd.Args, "-ss", start)
		cmd.Args = append(cmd.Args, "-to", end)
		cmd.Args = append(cmd.Args, "-c:v", "libx265")
		cmd.Args = append(cmd.Args, "-tag:v", "hvc1")
		cmd.Args = append(cmd.Args, "-c:a", "aac")
	}
	cmd.Args = append(cmd.Args, "-map_metadata", "-1")
	cmd.Args = append(cmd.Args, "-vsync", "0") // 添加这行
	cmd.Args = append(cmd.Args, "-copyts")     // 添加这行
	cmd.Args = append(cmd.Args, mp4)
	err := util.Exec(cmd)
	if err != nil {
		return err
	}
	return nil
}