package util

import (
	"fmt"
	"os"
	"testing"
)

func TestReadLLC(t *testing.T) {
	llcFile, has := FindProjLLCFile("D:\\迅雷下载")
	if !has {
		t.Log("未找到文件")
		return
	}
	seconds, err := extractStartsFromTextFile(llcFile)
	if err != nil {
		t.Log(err)
	}
	for _, v := range seconds {
		t.Log(v)
	}
	timestamps := SecondToHMS(seconds)
	for _, v := range timestamps {
		t.Log(v)
	}
}

func TestWindowsName(t *testing.T) {
	hostname, err := os.Hostname()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(hostname)
}
