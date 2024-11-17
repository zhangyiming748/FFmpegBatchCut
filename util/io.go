package util

import (
	"bufio"
	"fmt"
	"github.com/h2non/filetype"
	"io"
	"log"
	"os"
	"path/filepath"
	"strings"
)

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
	return

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

func IsVideo(fname string) bool {
	// Open a file descriptor
	file, _ := os.Open(fname)

	// We only have to pass the file header = first 261 bytes
	head := make([]byte, 261)
	file.Read(head)
	return filetype.IsVideo(head)
}
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