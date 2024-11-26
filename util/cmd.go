package util

import (
	"fmt"
	"log"
	"os/exec"
)

func Exec(cmd *exec.Cmd) error {
	log.Printf("当前运行的命令是:%s\n", cmd.String())
	output, err := cmd.CombinedOutput()
	if err != nil {
		log.Printf("命令%v遇到了问题%v\n", cmd.String(), err)
		return err
	} else {
		fmt.Printf("当前命令输出:%v\n", string(output))
		return nil
	}
}
