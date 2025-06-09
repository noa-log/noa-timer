# Noa Timer
Noa Log 的定时任务模块，支持自动化执行压缩日志，清理过期日志等功能。

## 安装
```bash
go get -u github.com/noa-log/noa-timer
```

## 快速开始
```go
package main

import (
    "github.com/noa-log/noa"
    noatimer "github.com/noa-log/noa-timer"
)

func main() {
    // 创建一个新的日志实例
    logger := noa.NewLog()

    // 注册默认定时任务，每天00:30压缩日志，00:35清理过期7天日志
    go noatimer.StartDefaultTask(logger)

    logger.Info("Test", "Starting Noa Timer")

    // ... 执行其他业务逻辑操作
}
```

## 自定义定时任务
你也可以自定义定时任务，以下是一个简单的示例：
```go
package main

import (
    "time"
    "github.com/noa-log/noa"
    noatimer "github.com/noa-log/noa-timer"
    "github.com/jasonlvhit/gocron"
)

func logTask(logger *noa.LogConfig) {
	task := gocron.NewScheduler()

	task.Every(1).Day().At("00:30").Do(noatimer.Compress, logger) // 每天00:30执行日志压缩
	task.Every(1).Day().At("00:35").Do(noatimer.Clear, logger, time.Hour*24*7) // 每天00:35清理过期7天的日志

	<-task.Start()
}

func main() {
    // 创建一个新的日志实例
    logger := noa.NewLog()

    // 注册自定义定时任务
    go logTask(logger)

    logger.Info("Test", "Starting Noa Timer with Custom Task")

    // ... 执行其他业务逻辑操作
}
```

## 许可
本项目基于[Apache License 2.0](https://www.apache.org/licenses/LICENSE-2.0)协议开源。使用时请遵守协议的条款。