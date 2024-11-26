package util

import (
	"fmt"
	"math"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

func UseProjLLCFile(llcFile string) []string {
	seconds, _ := extractStartsFromTextFile(llcFile)
	timestamps := secondToHMS(seconds)
	return timestamps
}

// 搜索目标文件夹是否包含后缀为proj.llc的文件
func FindProjLLCFile(folderPath string) (string, bool) {
	var projLLCFile string
	err := filepath.Walk(folderPath, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() && strings.HasSuffix(path, "proj.llc") {
			projLLCFile = path
			return filepath.SkipDir
		}
		return nil
	})
	if err != nil {
		fmt.Printf("遍历文件夹时出错: %v\n", err)
		return "", false
	}
	if projLLCFile != "" {
		return projLLCFile, true
	}
	return "", false
}

// 提取start后边的秒数
func extractStartsFromTextFile(filePath string) ([]float64, error) {
	data, err := os.ReadFile(filePath)
	if err != nil {
		return nil, err
	}
	lines := strings.Split(string(data), "\n")
	var startValues []float64
	for _, line := range lines {
		line = strings.TrimSpace(line)
		if strings.HasPrefix(line, "start:") {
			line = strings.Replace(line, ",", "", 1)
			parts := strings.Split(line, ":")
			if len(parts) > 1 {
				valueStr := strings.TrimSpace(parts[1])
				value, err := strconv.ParseFloat(valueStr, 6)
				if err != nil {
					return nil, err
				}
				startValues = append(startValues, value)
			}
		}
	}
	return startValues, nil
}

// 秒转换为时间
func formatSecondToHMS(seconds float64) string {
	hours := int(seconds / 3600)
	seconds -= float64(hours * 3600)
	minutes := int(seconds / 60)
	seconds -= float64(minutes * 60)
	milliseconds := int(math.Round(seconds * 1000))
	//hh
	//fmt.Printf("hh=%02d\n", hours)
	//mm
	//fmt.Printf("mm=%02d\n", minutes)
	//ss
	//fmt.Printf("ss=%02d\n", int(seconds))
	//ms
	//fmt.Printf("ms=%03d\n", milliseconds)
	times := fmt.Sprintf("%02d:%02d:%02d.%03d", hours, minutes, int(seconds), milliseconds)
	times = times[:12]
	//fmt.Println(times)
	times = strings.Replace(times, ":", "", -1)
	times = strings.Replace(times, ".", "", -1)
	return times
}
func secondToHMS(currentTime []float64) []string {
	var timestamps []string
	for _, second := range currentTime {
		timestamps = append(timestamps, formatSecondToHMS(second))
	}
	return timestamps
}
