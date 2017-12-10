package myhttp

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strings"
)

// DownloadFromURL 函數會將指定的網址檔案下載下來。
// Reference：https://github.com/thbar/golang-playground/blob/master/download-files.go
func DownloadFromURL(url, dstDir string) (err error) {

	// 將字串依指定符號（/）分割成數個字串。
	fileTokens := strings.Split(url, "/")
	// 最後的分割字串為下載檔名。
	fileName := fileTokens[len(fileTokens)-1]
	//dst := dstDir + "\\" + fileName
	dst := filepath.Join(dstDir, fileName)

	// 根據作業系統調整路徑的正反鈄線，以及將目錄路徑轉換成絕對路徑。
	dst, err = filepath.Abs(dst)
	if err != nil {
		fmt.Println("Error while finding absolute path", dst, "-", err)
		return err
	}

	// 新建檔案並確認此檔名無重覆。
	out, err := os.Create(dst)
	if err != nil {
		fmt.Println("Error while creating", dst, "-", err)
		return err
	}
	// 在 function 結束前關閉已開啟的檔案。
	defer out.Close()

	// 發出 Http Get 請求並確認是否成功。
	response, err := http.Get(url)
	if err != nil {
		fmt.Println("Error while downloading", url, "-", err)
		return err
	}
	// 在 function 結束前關閉已開啟 Http 請求。
	defer response.Body.Close()

	// 將檔案下載下來。
	_, err = io.Copy(out, response.Body)
	if err != nil {
		fmt.Println("Error while downloading", url, "-", err)
		return err
	}

	// 回傳 nil（無）錯誤。
	return nil
}
