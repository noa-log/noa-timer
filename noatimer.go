/*
 * @Author: nijineko
 * @Date: 2025-06-09 11:14:13
 * @LastEditTime: 2025-06-09 11:38:45
 * @LastEditors: nijineko
 * @Description: noa timer package
 * @FilePath: \noa-timer\noatimer.go
 */
package noatimer

import (
	"time"

	"github.com/jasonlvhit/gocron"
	"github.com/noa-log/noa"
)

const (
	DEFAULT_LOG_SOURCE = "Noatimer" // default log source
)

/**
 * @description: Start the default task scheduler
 * @param {noa.LogConfig} Log noa log instance
 */
func StartDefaultTask(Log *noa.LogConfig) {
	// Create a new scheduler
	Task := gocron.NewScheduler()

	Task.Every(1).Day().At("00:30").Do(Compress, Log)              // Compress log files every day at 00:30
	Task.Every(1).Day().At("00:35").Do(Clear, Log, time.Hour*24*7) // Clear log files older than 7 days every day at 00:35

	<-Task.Start()
}
