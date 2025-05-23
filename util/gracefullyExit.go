package util

import (
	"log"
	"os"
	"os/signal"
	"syscall"
)

func ExitAfterRun(final func()) {
	c := make(chan os.Signal, 1)
	// 监听信号
	signal.Notify(c, syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)
	go func() {
		for s := range c {
			switch s { // 终端控制进程结束(终端连接断开)|用户发送INTR字符(Ctrl+C)触发|结束程序
			case syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM:
				log.Println("接收到退出信号:", s)
				final()
			default:
				log.Println("其他信号:", s)
			}
		}
	}()
}

func Exit() {
	log.Fatalln("这是一次不完整的退出,根据上一条日志文件恢复数据")
}
