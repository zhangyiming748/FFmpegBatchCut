package util

import (
	"testing"
)

/*
windows下测试通过
*/
func TestGetFolders(t *testing.T) {
	dir := "F:\\media\\STAR-994-C" // 替换为你要检查的目录
	file, err := GetAllVideoFilesInDir(dir)
	if err != nil {
		return
	} else {
		t.Log(file)
	}
}
