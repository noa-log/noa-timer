/*
 * @Author: nijineko
 * @Date: 2025-06-09 10:40:33
 * @LastEditTime: 2025-06-09 11:21:50
 * @LastEditors: nijineko
 * @Description: compress log
 * @FilePath: \noa-timer\compress.go
 */
package noatimer

import (
	"testing"
	"time"

	"github.com/noa-log/noa"
)

func TestCompress(t *testing.T) {
	Log := noa.NewLog()
	Log.Writer.TimeFormat = "2006-01-02 15-04-05"

	for range 61 {
		Log.Info("Test", "This is a test log entry")
		time.Sleep(time.Second)
	}

	Compress(Log)
}
