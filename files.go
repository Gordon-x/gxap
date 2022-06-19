package gxap

import (
	"os"
)

func PathExist(f string) bool {
	_, err := os.Stat(f)
	return err == nil || os.IsExist(err)
}

// CreatePath 创建目录
func CreatePath(dir string) error {
	if PathExist(dir) {
		return nil
	}

	return os.Mkdir(dir, 0755)
}
