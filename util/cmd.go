package util

import (
	"log"
	"os/exec"
)

func Exec(cmd *exec.Cmd) error {
	log.Printf("当前运行的命令是:%s\n", cmd.String())

	if output, err := cmd.CombinedOutput(); err != nil {
		log.Printf("命令执行失败:%s\n", err.Error())
		return err
	} else {
		log.Printf("命令执行成功:%v\n", string(output))
	}
	return nil
}
