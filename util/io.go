// 提供文件和目录操作的工具函数
package util

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	"github.com/h2non/filetype"
)

// ReadByLine 按行读取文件内容
// fp: 文件路径
// 返回: 字符串切片，每个元素为文件的一行
func ReadByLine(fp string) []string {
	lines := []string{}
	fi, err := os.Open(fp)
	if err != nil {
		fmt.Printf("Error: %s\n", err)
		log.Println("按行读文件出错")
		return []string{}
	}
	defer fi.Close()

	br := bufio.NewReader(fi)
	for {
		a, _, c := br.ReadLine()
		if c == io.EOF {
			break
		}
		lines = append(lines, string(a))
	}
	return lines
}

// 按行写文件
func WriteByLine(fp string, s []string) {
	file, err := os.OpenFile(fp, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0777)
	if err != nil {
		return
	}
	defer file.Close()
	writer := bufio.NewWriter(file)
	for _, v := range s {
		writer.WriteString(v)
		writer.WriteString("\n")
	}
	writer.Flush()
}

/*
获取当前文件夹和全部子文件夹下视频文件
*/

func GetAllFiles(root string) (files []string) {
	patterns := []string{"webm", "m4v", "mp4", "mov", "avi", "wmv", "ts", "rmvb", "wma", "avi", "flv", "rmvb", "mpg", "f4v"}
	for _, pattern := range patterns {
		files = append(files, getFilesByExtension(root, pattern)...)
	}
	return files
}

/*
获取当前文件夹下视频文件
*/

func GetFiles(root string) (files []string) {
	files = append(files, getFilesByHead(root)...)
	return files
}

/*
获取当前文件夹和全部子文件夹下指定扩展名的全部文件
*/
func getFilesByHead(root string) []string {
	var files []string
	defer func() {
		if err := recover(); err != nil {
			log.Println("获取文件出错")
			os.Exit(-1)
		}
	}()
	filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if !info.IsDir() {
			// Open a file descriptor
			file, _ := os.Open(path)
			// We only have to pass the file header = first 261 bytes
			head := make([]byte, 261)
			file.Read(head)
			if filetype.IsVideo(head) {
				fmt.Printf("File: %v is a video\n", path)
				files = append(files, path)
			}

		}
		return nil
	})
	return files
}

/*
获取当前文件夹和全部子文件夹下指定扩展名的全部文件
*/
func getFilesByExtension(root, extension string) []string {
	var files []string
	defer func() {
		if err := recover(); err != nil {
			log.Println("获取文件出错")
			os.Exit(-1)
		}
	}()
	filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if !info.IsDir() && strings.HasSuffix(info.Name(), extension) {
			files = append(files, path)
		}
		return nil
	})
	return files
}
func IsExist(fp string) bool {
	// 使用 os.Stat 函数获取文件信息
	if f, err := os.Stat(fp); err == nil {
		log.Println("Path exists")
		if f.IsDir() {
			log.Println("Path is a directory")
		} else {
			log.Println("Path is a file")
		}
		return true
	} else if os.IsNotExist(err) {
		log.Println("Path does not exist")
		return false
	} else {
		log.Println("Error occurred:", err)
		return false
	}
}

// IsVideo 检查文件是否为视频文件
// fname: 文件路径
// 返回: 布尔值，表示是否为视频文件
func IsVideo(fname string) bool {
	// Open a file descriptor
	file, _ := os.Open(fname)
    defer file.Close()
	// We only have to pass the file header = first 261 bytes
	head := make([]byte, 261)
	file.Read(head)
	return filetype.IsVideo(head)
}
func GetAllVideoFilesButMp4(dir string) ([]string, error) {
	var files []string
	err := filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		// 判断是否是文件，如果是文件则将其绝对路径添加到files切片中
		if !info.IsDir() {
			if IsVideo(path) {
				if filter(path) {
					files = append(files, path)
				}
			}

		}
		return nil
	})
	return files, err
}

// GetAllVideoFilesInDir 获取指定目录下的所有视频文件
// dir: 目录路径
// 返回: 视频文件路径列表和可能的错误
func GetAllVideoFilesInDir(dir string) ([]string, error) {
	var files []string
	err := filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		// 判断是否是文件，如果是文件则将其绝对路径添加到files切片中
		if !info.IsDir() {
			if IsVideo(path) {
				files = append(files, path)
			}

		}
		return nil
	})
	return files, err
}

/*
视频但非mp4
*/
func filter(fp string) bool {
	file, _ := os.Open(fp)

	// We only have to pass the file header = first 261 bytes
	head := make([]byte, 261)
	file.Read(head)

	if filetype.IsVideo(head) {
		if strings.HasSuffix(fp, "mp4") {
			return false
		} else {
			return true
		}
	} else {
		return false
	}
}
func IfFileExists(path string) bool {
	_, err := os.Stat(path)
	if err == nil {
		return true
	}
	if os.IsNotExist(err) {
		return false
	}
	return false
}

func ConvMp4(videos []string) {
	for _, video := range videos {
		before := filepath.Ext(video)
		after := strings.Replace(video, before, ".mp4", 1)
		cmd := exec.Command("ffmpeg", "-i", video, "-c:v", "libx265", "-c:a", "aac", "-map_metadata", "-1", after)
		if hostname, _ := os.Hostname(); hostname == "DESKTOP-VGFTVD8" {
			cmd = exec.Command("ffmpeg", "-hwaccel", "cuda", "-i", video, "-c:v", "h264_nvenc", "-c:a", "aac", "-ac", "1", "-preset", "medium", "-cq", "20", "-map_metadata", "-1", after)
		}
		if err := Exec(cmd); err == nil {
			os.Remove(video)
		}
	}
}
