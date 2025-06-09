/*
 * @Author: nijineko
 * @Date: 2025-06-09 11:01:28
 * @LastEditTime: 2025-06-09 11:01:29
 * @LastEditors: nijineko
 * @Description: directory utilities
 * @FilePath: \noa-timer\tools\file\directory.go
 */
package file

import "os"

/**
 * @description: Create directory
 * @param {string} Path directory path
 * @return {error} error
 */
func Mkdir(Path string) error {
	if _, err := os.Stat(Path); os.IsNotExist(err) {
		if err := os.MkdirAll(Path, os.ModePerm); err != nil {
			return err
		}
	}

	return nil
}
