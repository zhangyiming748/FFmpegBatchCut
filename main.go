package main

import (
	"FFmpegBatchCut/ffmpeg"
	"FFmpegBatchCut/util"
	"fmt"
	"log"
	"os"
	"path/filepath"
)

func init() {
	util.SetLog("BitchCut.log")
	util.ExitAfterRun(util.Exit)
}
func main() {
	root := "D:\\迅雷下载"
	folders, _ := util.GetFoldersWithTimestamps(root)
	if len(folders) == 0 {
		log.Fatalln("没有找到任何符合条件的文件")
	}
	for _, folder := range folders {
		fmt.Printf("符合筛选条件的目录:%v\n", folders)
		timestampsFile := filepath.Join(folder, "timestamps.txt")
		if !util.IfFileExists(timestampsFile) {
			continue
		}
		videos, _ := util.GetAllVideoFilesInDir(folder)
		if len(videos) > 1 {
			log.Printf("跳过包含多个视频,可能是分割后的文件夹%v\n", folder)
			continue
		}
		mp4 := videos[0]
		timestamps := util.ReadByLine(timestampsFile)
		timestamps = removeEmptyStrings(timestamps)
		log.Printf("目录%v\t文件%v\n", folder, mp4)
		err := ffmpeg.CutOne(mp4, timestamps)
		if err != nil {
			log.Fatal(err)
		} else {
			if err := os.RemoveAll(timestampsFile); err != nil {
				log.Printf("删除%v失败\t%v\n", timestampsFile, err)
			} else {
				if err := os.RemoveAll(mp4); err != nil {
					log.Printf("删除%v失败\t%v\n", mp4, err)
				}
				log.Printf("分割文件结束,删除%v和%v成功\n", timestampsFile, mp4)
			}
		}
	}
}
func removeEmptyStrings(input []string) []string {
	var result []string
	for _, str := range input {
		if str != "" {
			result = append(result, str)
		}
	}
	return result
}
