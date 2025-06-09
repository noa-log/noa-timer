/*
 * @Author: nijineko
 * @Date: 2025-06-09 10:49:41
 * @LastEditTime: 2025-06-09 10:53:34
 * @LastEditors: nijineko
 * @Description: file path utilities
 * @FilePath: \noa-timer\tools\file\path.go
 */
package file

import (
	"os"
	"path/filepath"
)

/**
 * @description: get all file paths in a directory
 * @param {string} Path directory path
 * @return {[]string} file paths in the directory
 * @return {error} error
 */
func GetDirectoryFilePaths(Path string) ([]string, error) {
	var FileList []string
	Files, err := os.ReadDir(Path)
	if err != nil {
		return FileList, err
	}
	for _, File := range Files {
		if !File.IsDir() {
			FileList = append(FileList, filepath.Join(Path, File.Name()))
		}
	}
	return FileList, err
}
