# FFmpegBatchCut

~~给定时间 批量截取视频 比如给 000130300 表示 00:01:30.300 自动生成 ffmpeg -ss 00:00:00 -to 00:01:30.300 并且调用 exec.command.Run 执行~~
可以省略开始时间 000000000
# 更新
直接使用[LosslessCut](https://github.com/mifi/lossless-cut)提供的proj.llc文件读取时间点