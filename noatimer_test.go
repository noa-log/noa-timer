/*
 * @Author: nijineko
 * @Date: 2025-06-09 11:37:32
 * @LastEditTime: 2025-06-09 12:35:14
 * @LastEditors: nijineko
 * @Description: noa timer test package
 * @FilePath: \noa-timer\noatimer_test.go
 */
package noatimer

import (
	"testing"
	"time"

	"github.com/jasonlvhit/gocron"
	"github.com/noa-log/noa"
)

func TestDefaultTask(t *testing.T) {
	Log := noa.NewLog()
	Log.Writer.FolderPath = "./logs"
	Log.Writer.TimeFormat = "2006-01-02"

	Log.Info("Test", "This is a test log entry")

	// Start the default task scheduler
	StartDefaultTask(Log)

	// Wait for a while to let the tasks run
	select {}
}

func TestTask(t *testing.T) {
	Log := noa.NewLog()
	Log.Writer.FolderPath = "./logs"
	Log.Writer.TimeFormat = "2006-01-02 15-04-05"

	go func() {
		Task := gocron.NewScheduler()

		Task.Every(1).Minute().Do(Compress, Log)
		Task.Every(65).Second().Do(Clear, Log, time.Second)

		<-Task.Start()
	}()

	for range 70 {
		Log.Info("Test", "This is a test log entry")
		time.Sleep(time.Second)
	}
}
