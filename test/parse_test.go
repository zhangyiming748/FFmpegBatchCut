package main

import (
    "bufio"
    "os"
    "strings"
    "strconv"
    "regexp"
    "fmt"
    "testing"
)
type Segment struct {
    Start float64
    End   float64
    Name  string
}
func parseSegments(filename string) ([]Segment, error) {
    file, err := os.Open(filename)
    if err != nil {
        return nil, err
    }
    defer file.Close()

    var segments []Segment
    var currentSegment *Segment
    
    // 用于提取数值的正则表达式
    numRegex := regexp.MustCompile(`[-]?\d+\.?\d*`)
    
    scanner := bufio.NewScanner(file)
    for scanner.Scan() {
        line := strings.TrimSpace(scanner.Text())
        
        // 开始新的段落
        if strings.Contains(line, "{") {
            currentSegment = &Segment{}
            continue
        }
        
        // 结束当前段落
        if strings.Contains(line, "},") {
            if currentSegment != nil {
                segments = append(segments, *currentSegment)
                currentSegment = nil
            }
            continue
        }
        
        // 解析字段
        if currentSegment != nil {
            if strings.Contains(line, "start:") {
                if num := numRegex.FindString(line); num != "" {
                    currentSegment.Start, _ = strconv.ParseFloat(num, 64)
                }
            } else if strings.Contains(line, "end:") {
                if num := numRegex.FindString(line); num != "" {
                    currentSegment.End, _ = strconv.ParseFloat(num, 64)
                }
            } else if strings.Contains(line, "name:") {
                currentSegment.Name = strings.Trim(strings.Split(line, ":")[1], " ',")
            }
        }
    }
    
    // 处理最后一个段落（如果有的话）
    if currentSegment != nil {
        segments = append(segments, *currentSegment)
    }
    
    return segments, scanner.Err()
}

func TestParse(t *testing.T) {
    segments, err := parseSegments("exam.json")
    if err != nil {
        fmt.Println("Error:", err)
        return
    }
    
    // 打印解析结果
    for i, seg := range segments {
        fmt.Printf("Segment %d:\n", i+1)
        fmt.Printf("  Start: %f\n", seg.Start)
        fmt.Printf("  End: %f\n", seg.End)
        fmt.Printf("  Name: %s\n", seg.Name)
    }
}