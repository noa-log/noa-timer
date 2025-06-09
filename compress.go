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
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/noa-log/noa"
	"github.com/noa-log/noa-timer/tools/file"
	"github.com/noa-log/noa-timer/tools/gzip"
)

/**
 * @description: Compress log files
 * @param {noa.LogConfig} Log noa log instance
 */
func Compress(Log *noa.LogConfig) {
	// get current log file name with time format
	FileNameTime := time.Now().Format(Log.Writer.TimeFormat)

	// get log directory file paths
	FilePaths, err := file.GetDirectoryFilePaths(Log.Writer.FolderPath)
	if err != nil {
		Log.Error(DEFAULT_LOG_SOURCE, err)
		return
	}

	for _, FilePath := range FilePaths {
		FileExt := filepath.Ext(FilePath)
		FileName := strings.TrimSuffix(filepath.Base(FilePath), FileExt)

		// skip current log file
		if FileName == FileNameTime {
			continue
		}

		// skip already compressed files
		if FileExt == ".gz" {
			continue
		}

		// compress file
		if err := gzip.CompressFileByPath(FilePath); err != nil {
			Log.Error(DEFAULT_LOG_SOURCE, err)
			continue
		}

		// delete original file after compression
		if err := os.Remove(FilePath); err != nil {
			Log.Error(DEFAULT_LOG_SOURCE, err)
			continue
		}
	}
}
