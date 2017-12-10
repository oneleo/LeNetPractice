package mygzip

import (
	"archive/tar"
	"compress/gzip"
	"fmt"
	"io"
	"os"
	"path/filepath"
)

// GzDecompress 函數會將指定的 .gz 壓縮檔解壓縮至特定位置。
// tar 格式：多個檔案打包但不壓縮。gz 格式：只能將單一個檔案進行壓縮。
// Reference：https://studygolang.com/articles/7481
func GzDecompress(src, dstDir string) (err error) {

	// 根據作業系統調整正反鈄線。
	//src = filepath.Clean(src)
	// 根據作業系統調整路徑的正反鈄線，以及將目錄路徑轉換成絕對路徑。
	src, err = filepath.Abs(src)
	if err != nil {
		fmt.Println("Error while finding absolute path", src, "-", err)
		return err
	}

	// 開啟已存在的 .gz 檔。
	srcFile, err := os.Open(src)
	if err != nil {
		return err
	}
	// 在 function 結束前關閉已開啟的檔案。
	defer srcFile.Close()

	// 以 gzip 格式解析已開啟的檔案。
	gzReader, err := gzip.NewReader(srcFile)
	if err != nil {
		return err
	}
	// 在 function 結束前關閉已開啟 gzip reader。
	defer gzReader.Close()

	// 將目前路徑，加上已解析後的檔名，作為解壓縮的目的地。
	//dstFile := dstDir + "\\" + gzReader.Name
	dstFile := filepath.Join(dstDir, gzReader.Name)

	// 根據作業系統調整鈄線正反。
	//dstFile = filepath.Clean(dstFile)
	// 根據作業系統調整路徑的正反鈄線，以及將目錄路徑轉換成絕對路徑。
	dstFile, err = filepath.Abs(dstFile)
	if err != nil {
		fmt.Println("Error while finding absolute path", dstFile, "-", err)
		return err
	}

	// 建立一個新檔作為將上述解壓縮後的檔案存取位置。
	file, err := os.Create(dstFile)
	if err != nil {
		fmt.Println("Error while creating", dstFile, "-", err)
		return err
	}
	// 在 function 結束前關閉已開啟檔案。
	defer file.Close()

	// 將解壓縮後的檔案寫入至開啟的新檔。
	if _, err = io.Copy(file, gzReader); err != nil {
		return err
	}

	// 若無錯誤，回傳 nil。
	return nil
}

// TarGzDecompress 函數會將指定的 .tar.gz 壓縮檔解壓縮至特定位置。
// tar 格式：多個檔案打包但不壓縮。gz 格式：只能將單一個檔案進行壓縮。
// Reference：https://studygolang.com/articles/7481
func TarGzDecompress(src, dstDir string) (err error) {

	// 根據作業系統調整路徑的正反鈄線，以及將目錄路徑轉換成絕對路徑。
	src, err = filepath.Abs(src)
	if err != nil {
		fmt.Println("Error while finding absolute path", src, "-", err)
		return err
	}

	// 開啟已存在的 .tar.gz 檔。
	srcFile, err := os.Open(src)
	if err != nil {
		return err
	}
	// 在 function 結束前關閉已開啟的檔案。
	defer srcFile.Close()

	// 以 gzip 格式解析已開啟的檔案。
	gzReader, err := gzip.NewReader(srcFile)
	if err != nil {
		return err
	}
	// 在 function 結束前關閉已開啟 gzip reader。
	defer gzReader.Close()

	// 以 tar 格式解析已開啟的檔案。
	tarReader := tar.NewReader(gzReader)
	// 因 .tar 檔解開後會包含多個檔案及資料夾，故要一個一個將檔寫至目的地。
	for {
		// 取得 .tar 檔內的下一個檔名。
		tr, err := tarReader.Next()
		if err != nil {
			// 若已經沒有檔名就離開迴圈。
			if err == io.EOF {
				break
				// 若發生錯誤則回傳錯誤訊息。
			} else {
				return err
			}
		}
		// 將目前路徑，加上已解析後的檔名，作為解壓縮的目的地。
		//dstFile := dstDir + "\\" + tr.Name
		dstFile := filepath.Join(dstDir, tr.Name)

		// 根據作業系統調整路徑的正反鈄線，以及將目錄路徑轉換成絕對路徑。
		dstFile, err = filepath.Abs(dstFile)
		if err != nil {
			fmt.Println("Error while finding absolute path", dstFile, "-", err)
			return err
		}

		// 取得目前檔名的資訊。
		info := tr.FileInfo()
		// 若目前指到的檔名是資料夾的話，則建立此資料夾。
		// Reference：http://blog.ralch.com/tutorial/golang-working-with-tar-and-gzip/
		if info.IsDir() {
			if err = os.MkdirAll(dstFile, info.Mode()); err != nil {
				return err
			}
			// 資料夾建立完成，重頭開始迴圈。
			continue
		} else {
			// 若目前指到的檔名是檔案的話，則新建一檔案。
			file, err := os.Create(dstFile)
			if err != nil {
				fmt.Println("Error while creating", dstFile, "-", err)
				return err
			}
			// 在 function 結束前關閉已開啟檔案。
			defer file.Close()
			// 將解壓縮後的檔案寫入至開啟的新檔。
			if _, err = io.Copy(file, tarReader); err != nil {
				return err
			}
		}
	}
	// 若無錯誤，回傳 nil。
	return nil
}
