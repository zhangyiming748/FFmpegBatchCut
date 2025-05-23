# FFmpegBatchCut

基于 FFmpeg 的视频批量切割工具，支持从 [LosslessCut](https://github.com/mifi/lossless-cut) 项目文件中读取时间点进行视频分割。
A batch video cutting tool based on FFmpeg, supporting video segmentation using timestamps from LosslessCut project files.

## 功能特点 | Features

- ✅ 支持读取 [LosslessCut](https://github.com/mifi/lossless-cut) 的 proj.llc 项目文件
  - Support reading proj.llc project files from [LosslessCut](https://github.com/mifi/lossless-cut)
- ✅ 自动按序号生成分割后的视频文件
  - Automatically generate segmented video files with sequential numbering
- ✅ 支持 NVIDIA GPU 加速（在特定设备上）
  - Support NVIDIA GPU acceleration (on specific devices)
- ✅ 保持视频质量（使用 libx265/h264_nvenc 编码）
  - Maintain video quality (using libx265/h264_nvenc encoding)
- ✅ 完整的错误处理和日志记录
  - Complete error handling and logging

## 系统要求 | Requirements

- FFmpeg 命令行工具
- Go 1.23.3 或更高版本
- （可选）NVIDIA GPU 及相关驱动

## 安装 | Installation

```bash
git clone https://github.com/zhangyiming748/FFmpegBatchCut.git
cd FFmpegBatchCut
go mod download
```

## 使用方法 | Usage

1. 准备项目文件 | Prepare project file

创建 proj.llc 文件，格式如下 | Create a proj.llc file in the following format:
```json
{
    cutSegments: [
        {
            start: 0,
            end: 90.5,
            name: 'segment1'
        },
        {
            start: 90.5,
            end: 180.3,
            name: 'segment2'
        }
    ]
}
```

2. 文件放置 | File placement
- 将 proj.llc 文件和需要切割的视频文件放在同一目录下
- Place the proj.llc file and the video file in the same directory

3. 运行程序 | Run the program
```bash
go run main.go
```

## 输出说明 | Output

- 程序将在视频所在目录生成编号的片段文件（01.mp4, 02.mp4, ...）
- The program will generate numbered segment files (01.mp4, 02.mp4, ...) in the video directory
- 所有操作日志将记录在 BitchCut.log 文件中
- All operation logs will be recorded in BitchCut.log file

## 编码设置 | Encoding Settings

- NVIDIA GPU 设备 | NVIDIA GPU devices:
  - 视频编码：h264_nvenc
  - 音频编码：aac
  - 预设：slow
  - CQ：18

- 其他设备 | Other devices:
  - 视频编码：libx265
  - 音频编码：aac
  - 标签：hvc1

## 注意事项 | Notes

1. 建议先用小视频文件测试
   - Recommend testing with small video files first
2. 确保输出文件名符合预期
   - Ensure output filenames meet expectations
3. 验证切割点的准确性
   - Verify the accuracy of cutting points
4. 检查输出视频的质量
   - Check the quality of output videos

## 许可证 | License

[License Name] - 查看 LICENSE 文件了解更多信息
[License Name] - See LICENSE file for more information

        