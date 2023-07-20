package util

import (
	"fmt"
	"os"
)

func CheckFileExist(filePath string) bool {
	if _, err := os.Stat(filePath); err == nil {
		return true
	} else {
		return false
	}
}

func CheckDir(dirPath string) error {
	_, err := os.Stat(dirPath)
	if os.IsNotExist(err) {
		// 目录不存在，新建目录
		err := os.MkdirAll(dirPath, os.ModePerm)
		if err != nil {
			return fmt.Errorf("创建目录失败: %v", err)
		}
	}
	return nil
}
