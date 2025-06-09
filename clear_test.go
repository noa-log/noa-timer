/*
 * @Author: nijineko
 * @Date: 2025-06-09 11:20:48
 * @LastEditTime: 2025-06-09 12:06:40
 * @LastEditors: nijineko
 * @Description: clear log files
 * @FilePath: \noa-timer\clear_test.go
 */
package noatimer

import (
	"testing"
	"time"

	"github.com/noa-log/noa"
)

func TestClear(t *testing.T) {
	Log := noa.NewLog()
	Log.Writer.TimeFormat = "2006-01-02 15-04-05"

	for range 61 {
		Log.Info("Test", "This is a test log entry")
		time.Sleep(time.Second)
	}

	Clear(Log, time.Second)
}
