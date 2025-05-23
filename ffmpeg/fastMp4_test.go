package ffmpeg

import (
	"FFmpegBatchCut/util"
	"testing"
)

// go test -v -timeout 30h -run TestAnyVideoToMP4
func TestAnyVideoToMP4(t *testing.T) {
	// 测试用例：将指定文件夹中的所有视频文件转换为 MP4 格式
	// 输入：指定文件夹路径
	// 输出：转换后的 MP4 文件路径
	// 预期结果：所有视频文件都被成功转换为 MP4 格式
	root := "D:\\pikpak\\角色扮演成熟女神\\My Pack"
	files, err := util.GetAllVideoButMP4FilesInRootFolder(root)
	if err != nil {
		t.Fatal(err)
	}
	for _, file := range files {
		if err := AnyVideoToMP4(file); err != nil {
			t.Fatal(err)
		}
	}
}
