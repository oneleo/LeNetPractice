package mygzip

import (
	"fmt"
	"os"
	"path/filepath"
)

// CreateFolder 函數會使用遞迴的方式來建立目錄。
func CreateFolder(path string) (err error) {
	// 根據作業系統調整正反鈄線，以及將目錄路徑轉換成絕對路徑。
	path, err = filepath.Abs(path)
	if err != nil {
		fmt.Println("Error while finding absolute path", path, "-", err)
		return err
	}

	// 以遞迴的方式建立目錄。
	os.MkdirAll(path, os.ModePerm)
	return nil
}
