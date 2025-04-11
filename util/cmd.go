package util

import (
	"bufio"
	"fmt"
	"log"
	"os/exec"
)

func Exec(cmd *exec.Cmd) error {
	log.Printf("当前运行的命令是:%s\n", cmd.String())

	// 合并标准输出和错误输出
	pipe, err := cmd.StderrPipe()
	if err != nil {
		return fmt.Errorf("创建管道失败: %v", err)
	}
	cmd.Stdout = cmd.Stderr

	if err := cmd.Start(); err != nil {
		return fmt.Errorf("启动命令失败: %v", err)
	}

	scanner := bufio.NewScanner(pipe)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}

	return cmd.Wait()
}
