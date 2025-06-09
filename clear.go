/*
 * @Author: nijineko
 * @Date: 2025-06-09 11:20:48
 * @LastEditTime: 2025-06-09 11:58:28
 * @LastEditors: nijineko
 * @Description: clear log files
 * @FilePath: \noa-timer\clear.go
 */
package noatimer

import (
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/noa-log/noa"
	"github.com/noa-log/noa-timer/tools/file"
)

/**
 * @description: Clear log files
 * @param {noa.LogConfig} Log noa log instance
 * @param {time.Duration} ExpireTime expiration time for log files
 */
func Clear(Log *noa.LogConfig, ExpireTime time.Duration) {
	CurrentTime := time.Now().Unix()

	// get log directory file paths
	FilePaths, err := file.GetDirectoryFilePaths(Log.Writer.FolderPath)
	if err != nil {
		Log.Error(DEFAULT_LOG_SOURCE, err)
		return
	}

	for _, FilePath := range FilePaths {
		FileExt := filepath.Ext(FilePath)
		FileNameNoExt := strings.TrimSuffix(filepath.Base(FilePath), FileExt)

		// parse the timestamp from the filename based on the log file's time format
		FileTime, err := time.ParseInLocation(Log.Writer.TimeFormat, FileNameNoExt, time.Local)
		if err != nil {
			Log.Error(DEFAULT_LOG_SOURCE, err)
			continue
		}

		// check if expired
		if CurrentTime-FileTime.Unix() > int64(ExpireTime.Seconds()) {
			// delete expired file
			if err := os.Remove(FilePath); err != nil {
				Log.Error(DEFAULT_LOG_SOURCE, err)
				continue
			}
		}
	}
}
