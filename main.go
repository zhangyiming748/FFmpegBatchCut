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
}
func main() {
	root := "D:\\迅雷下载"
	folders, _ := util.GetFoldersWithTimestamps(root)
	for _, folder := range folders {
		fmt.Printf("符合筛选条件的目录:%v\n", folders)
		timestampsFile := filepath.Join(folder, "timestamps.txt")
		if !util.IfFileExists(timestampsFile) {
			continue
		}
		videos, _ := util.GetAllVideoFilesInDir(folder)
		if len(videos) > 1 {
			continue
		}
		mp4 := videos[0]
		timestamps := util.ReadByLine(timestampsFile)
		timestamps = removeEmptyStrings(timestamps)
		log.Printf("目录%v\n文件%v\n时间戳%v\n", folder, mp4, timestamps)
		err := ffmpeg.CutOne(mp4, timestamps)
		if err != nil {
			log.Fatal(err)
		} else {
			if err := os.RemoveAll(timestampsFile); err != nil {
				log.Printf("删除%v失败\t%v\n", timestamps, err)
			} else {
				if err := os.RemoveAll(mp4); err != nil {
					log.Printf("删除%v失败\n", mp4)
				}
				log.Printf("分割文件结束,删除%v和%v失败\n", timestampsFile, mp4)
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
