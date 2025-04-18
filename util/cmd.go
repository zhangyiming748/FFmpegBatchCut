package util

import (
	"bufio"
	"fmt"
	"log"
	"os/exec"
)

func Exec(cmd *exec.Cmd) error {
	log.Printf("当前运行的命令是:%s\n", cmd.String())
	
	stdout, err := cmd.StdoutPipe()
	if err != nil {
		return fmt.Errorf("创建stdout管道失败: %v", err)
	}
	stderr, err := cmd.StderrPipe()
	if err != nil {
		return fmt.Errorf("创建stderr管道失败: %v", err)
	}

	if err := cmd.Start(); err != nil {
		return fmt.Errorf("启动命令失败: %v", err)
	}

	// 处理输出
	go func() {
		scanner := bufio.NewScanner(stdout)
		for scanner.Scan() {
			fmt.Printf("进度: %s\n", scanner.Text())
		}
	}()

	go func() {
		scanner := bufio.NewScanner(stderr)
		for scanner.Scan() {
			fmt.Printf("错误: %s\n", scanner.Text())
		}
	}()

	return cmd.Wait()
}
