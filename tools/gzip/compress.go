/*
 * @Author: nijineko
 * @Date: 2025-06-09 10:57:55
 * @LastEditTime: 2025-06-09 12:27:16
 * @LastEditors: nijineko
 * @Description: gzip compress utilities
 * @FilePath: \noa-timer\tools\gzip\compress.go
 */
package gzip

import (
	"compress/gzip"
	"io"
	"os"
	"path/filepath"
	"strings"

	"github.com/noa-log/noa-timer/tools/file"
)

/**
 * @description: Compress a file by its path
 * @param {string} FilePath
 * @return {*}
 */
func CompressFileByPath(FilePath string) error {
	File, err := os.Open(FilePath)
	if err != nil {
		return err
	}
	defer File.Close()

	// get file name and extension
	FileExt := filepath.Ext(FilePath)
	FileNameNoExt := strings.TrimSuffix(filepath.Base(FilePath), FileExt)

	// create gzip file
	DestFilePath := filepath.Join(filepath.Dir(FilePath), FileNameNoExt+".gz")
	return CompressFile(DestFilePath, File)
}

/**
 * @description: Compress a file and save it to a destination path
 * @param {string} DestFilePath destination file path
 * @param {*os.File} SourceFile source file to compress
 * @return {error} error if any
 */
func CompressFile(DestFilePath string, SourceFile *os.File) error {
	// create folder if it doesn't exist
	if err := file.Mkdir(filepath.Dir(DestFilePath)); err != nil {
		return err
	}

	GZipFile, err := os.Create(DestFilePath)
	if err != nil {
		return err
	}
	defer GZipFile.Close()

	// create gzip writer
	Writer := gzip.NewWriter(GZipFile)
	defer Writer.Close()

	// set gzip header
	Writer.Name = filepath.Base(SourceFile.Name())
	Writer.Comment = "Compressed by noa-timer"
	
	if _, err := SourceFile.Seek(0, io.SeekStart); err != nil {
		return err
	}
	if _, err = io.Copy(Writer, SourceFile); err != nil {
		return err
	}

	return nil
}
