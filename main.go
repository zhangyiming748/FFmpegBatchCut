// 程序入口点，用于批量处理视频切割任务
package main

import (
	"FFmpegBatchCut/ffmpeg"
	"FFmpegBatchCut/util"
	"fmt"
	"log"
	"os"
	"runtime"
	"time"
)

func init() {
	// 初始化日志文件和配置
	util.SetLog("BitchCut.log")
	// 设置日志标志：包含文件名和行号
	log.SetFlags(2 | 16)
	// 注册程序退出时的清理函数
	util.ExitAfterRun(util.Exit)
	// 检查操作系统类型，提供运行环境建议
	if runtime.GOOS != "windows" {
		log.Println("极其不建议在Windows下运行")
	}
}

// waitUntil7AM 等待直到早上7点
func waitUntil7AM() {
	for {
		now := time.Now()
		hour := now.Hour()

		// 如果当前是早上7点，返回
		if hour == 7 {
			return
		}

		// 不是早上7点，等待30分钟
		time.Sleep(30 * time.Minute)
	}
}
func main() {
	// 在开始处理之前等待到早上7点
	// waitUntil7AM()
	// 指定要处理的根目录
	root := "F:\\原始视频\\pantyhose"

	// 获取包含LLC文件的所有文件夹
	folders, _ := util.GetFoldersWithLLCFiles(root)
	if len(folders) == 0 {
		log.Fatalln("没有找到任何符合条件的文件")
	}

	// 遍历每个文件夹进行处理
	for _, folder := range folders {
		defer func() {
			if err := recover(); err != nil {
				log.Printf("get panic : %v\n", err)
			}
		}()
		for _, folder := range folders {
			fmt.Printf("符合筛选条件的目录:%v\n", folder)
		}
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
		segments, err := util.ParseSegments(llcFile)
		if err != nil {
			log.Printf("解析%v失败:%v\n", llcFile, err)
			continue
		}
		log.Printf("目录%v\t文件%v\n", folder, mp4)
		if err = ffmpeg.CutBySegments(mp4, segments); err != nil {
			log.Printf("%v\n", err)
			continue
		} else {
			if err := os.RemoveAll(mp4); err != nil {
				log.Printf("删除%v失败\t%v\n", mp4, err)
			}
			if err := os.RemoveAll(llcFile); err != nil {
				log.Printf("删除%v失败\t%v\n", llcFile, err)
			}
			log.Printf("分割文件结束,删除%v成功\n", mp4)
		}
	}
}
