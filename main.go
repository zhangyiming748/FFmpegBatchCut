package main

import (
	"FFmpegBatchCut/ffmpeg"
	"FFmpegBatchCut/util"
	"fmt"
	"log"
	"os"
	"runtime"
)

func init() {
	util.SetLog("BitchCut.log")
	log.SetFlags(2 | 16)
	util.ExitAfterRun(util.Exit)
	if runtime.GOOS != "windows" {
		log.Println("极其不建议在Windows下运行")
	}
}

func main() {
	root := "/mnt/f/整理/AV/综艺"
	// videos, _ := util.GetAllVideoFilesButMp4(root)
	// util.ConvMp4(videos)
	folders, _ := util.GetFoldersWithLLCFiles(root)
	if len(folders) == 0 {
		log.Fatalln("没有找到任何符合条件的文件")
	}
	for _, folder := range folders {
		defer func ()  {
			if err:=recover(); err!= nil {
				log.Printf("get panic : %v\n",err)
			}
		}()
		fmt.Printf("符合筛选条件的目录:%v\n", folders)
		llcFile, has := util.FindProjLLCFile(folder)
		if !has {
			log.Println("未找到文件")
			continue
		}
		log.Printf("找到的工程文件:%v\n", llcFile)
		videos, _ := util.GetAllVideoFilesInDir(folder)
		if len(videos) > 1 {
			log.Printf("跳过包含多个视频,可能是分割后的文件夹%v\n", folder)
			continue
		}
		mp4 := videos[0]
		log.Printf("找到的视频文件:%v\n", mp4)
		timestamps := util.UseProjLLCFile(llcFile)
		for i, v := range timestamps {
			log.Printf("第%d个时间戳:%s\n", i, v)
		}
		log.Printf("目录%v\t文件%v\n", folder, mp4)
		err := ffmpeg.CutOne(mp4, timestamps)
		if err != nil {
			log.Fatal(err)
		} else {
			if err := os.RemoveAll(mp4); err != nil {
				log.Printf("删除%v失败\t%v\n", mp4, err)
			}
			log.Printf("分割文件结束,删除%v成功\n", mp4)
		}
	}
}
