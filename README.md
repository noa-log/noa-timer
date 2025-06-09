# Noa Timer
The timer task module of Noa Log, supporting automated actions such as compressing logs and clearing expired logs.

## Installation
```bash
go get -u github.com/noa-log/noa-timer
```

## Quick Start
```go
package main

import (
    "github.com/noa-log/noa"
    noatimer "github.com/noa-log/noa-timer"
)

func main() {
    // Create a new log instance
    logger := noa.NewLog()

    // Register default timer tasks: compress logs at 00:30 and clear logs older than 7 days at 00:35 every day
    go noatimer.StartDefaultTask(logger)

    logger.Info("Test", "Starting Noa Timer")

    // ... Execute other business logic
}
```

## Custom timer tasks
You can also define your own timer tasks. Here is a simple example:
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

	task.Every(1).Day().At("00:30").Do(noatimer.Compress, logger) // Compress logs daily at 00:30
	task.Every(1).Day().At("00:35").Do(noatimer.Clear, logger, time.Hour*24*7) // Clear logs older than 7 days daily at 00:35

	<-task.Start()
}

func main() {
    // Create a new log instance
    logger := noa.NewLog()

    // Register custom timer tasks
    go logTask(logger)

    logger.Info("Test", "Starting Noa Timer with Custom Task")

    // ... Execute other business logic
}
```

## License
This project is open-sourced under the [Apache License 2.0](https://www.apache.org/licenses/LICENSE-2.0). Please comply with the terms when using it.