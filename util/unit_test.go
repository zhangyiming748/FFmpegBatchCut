package util

import (
	"fmt"
	"os"
	"testing"
)

/*
windows下测试通过
*/
func TestGetFolders(t *testing.T) {
	dir := "F:\\media" // 替换为你要检查的目录
	file, err := GetAllVideoFilesInDir(dir)
	if err != nil {
		return
	} else {
		t.Log(file)
	}
}

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
	timestamps := secondToHMS(seconds)
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
