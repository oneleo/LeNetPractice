package mymnist

import (
	"bufio"
	"encoding/binary"
	"errors"
	"fmt"
	"image"
	"io"
	"os"
	"path/filepath"

	"golang.org/x/image/bmp"
)

// MNIST 訓練資料 train-images-idx3-ubyte 格式
// TRAINING SET IMAGE FILE (train-images-idx3-ubyte):
// [offset] [type]          [value]          [description]
// 0000     32 bit integer  0x00000803(2051) magic number
// 0004     32 bit integer  60000            number of images
// 0008     32 bit integer  28               number of rows
// 0012     32 bit integer  28               number of columns
// 0016     unsigned byte   ??               pixel
// 0017     unsigned byte   ??               pixel
// ........
// xxxx     unsigned byte   ??               pixel
// Pixels are organized row-wise. Pixel values are 0 to 255. 0 means background (white), 255 means foreground (black).

// MNIST 測試資料 t10k-images-idx3-ubyte 格式
// TEST SET IMAGE FILE (t10k-images-idx3-ubyte):
// [offset] [type]          [value]          [description]
// 0000     32 bit integer  0x00000803(2051) magic number
// 0004     32 bit integer  10000            number of images
// 0008     32 bit integer  28               number of rows
// 0012     32 bit integer  28               number of columns
// 0016     unsigned byte   ??               pixel
// 0017     unsigned byte   ??               pixel
// ........
// xxxx     unsigned byte   ??               pixel
// Pixels are organized row-wise. Pixel values are 0 to 255. 0 means background (white), 255 means foreground (black).

// ReadMnistImages 是將 MNIST 的 *images.idx?-ubyte 封裝檔
// 解開並轉換成各別的灰階圖格式。
// Reference：https://github.com/petar/GoMNIST/blob/master/mnist.go
// Reference：https://github.com/fumin/neural-network-mnist/blob/master/mnist/mnist.go
// Reference：https://github.com/kortschak/mnist/blob/master/mnist.go
func ReadMnistImages(src string) (imgs []image.Gray, err error) {

	// 根據作業系統調整路徑的正反鈄線，以及將目錄路徑轉換成絕對路徑。
	src, err = filepath.Abs(src)
	if err != nil {
		fmt.Println("Error while finding absolute path", src, "-", err)
		return nil, err
	}

	// 開啟已存在的 *images.idx?-ubyte 檔。
	srcFile, err := os.Open(src)
	if err != nil {
		return nil, err
	}
	// 在 function 結束前關閉已開啟的檔案。
	defer srcFile.Close()

	var (
		// 用來儲存 *images.idx?-ubyte 檔第一個檔頭 magic number。
		mgc uint32
		// 用來儲存 *images.idx?-ubyte 檔第二個檔頭 number of images。
		num uint32
		// 用來儲存 *images.idx?-ubyte 檔第三個檔頭 number of rows。
		row uint32
		// 用來儲存 *images.idx?-ubyte 檔第四個檔頭 number of columns。
		col uint32
	)

	// 剪下前 32bits 的資料（因為 mgc 是 32bits），再按 Big-Endian 順序寫入至 mgc。
	// 位元組順序（Byte Order）請參考：
	// Reference：https://zh.wikipedia.org/wiki/%E5%AD%97%E8%8A%82%E5%BA%8F
	// Reference：http://lihaoquan.me/2016/11/5/golang-byteorder.html
	if err = binary.Read(srcFile, binary.BigEndian, &mgc); err != nil {
		return nil, err
	}

	// 剪下前 32bits 的資料（因為 num 是 32bits），再按 Big-Endian 順序寫入至 num。
	if err = binary.Read(srcFile, binary.BigEndian, &num); err != nil {
		return nil, err
	}

	// 剪下前 32bits 的資料（因為 row 是 32bits），再按 Big-Endian 順序寫入至 row。
	if err = binary.Read(srcFile, binary.BigEndian, &row); err != nil {
		return nil, err
	}

	// 剪下前 32bits 的資料（因為 col 是 32bits），再按 Big-Endian 順序寫入至 col。
	if err = binary.Read(srcFile, binary.BigEndian, &col); err != nil {
		return nil, err
	}

	// 將 imgs 變數宣告用來儲存灰階影像的 slice，個數為 num 個。
	imgs = make([]image.Gray, num)

	for i := 0; i < int(num); i++ {
		// 建立用來暫存解析影像後的像素 pix 資訊 slice，共有 row*col 個元素。
		rawImages := make([]byte, row*col)

		// io.ReadFull 會將 srcFile 內容將 rawImages[i] 填滿
		// 若讀取的 byte 數 != len(rawImages[i])，則回傳 err 錯誤。
		_, err := io.ReadFull(srcFile, rawImages)
		if err != nil {
			return nil, err
		}
		// 建立一個大小為 row*col 的空灰階影像變數。
		imgs[i] = *image.NewGray(image.Rect(0, 0, int(col), int(row)))
		// 將 rawImages 儲存的像素 pix 資訊寫入至 imgs[i] 的像素 pix 內。
		imgs[i].Pix = rawImages
	}
	// 回傳所有的影像，及回傳 nil（無）錯誤。
	return imgs, nil
}

// WriteBMP 函數是將讀入的 img 灰階影像變數，編碼為 bmp 格式後，儲存成名為 dstFile 檔。
func WriteBMP(dstFile string, img *image.Gray) (err error) {

	// 根據作業系統調整路徑的正反鈄線，以及將目錄路徑轉換成絕對路徑。
	dstFile, err = filepath.Abs(dstFile)
	if err != nil {
		fmt.Println("Error while finding absolute path", dstFile, "-", err)
		return err
	}

	// 建立一個新檔作為將上述灰階影像變數存取成目標檔。
	file, err := os.Create(dstFile)
	if err != nil {
		fmt.Println("Error while creating", dstFile, "-", err)
		return err
	}
	// 在 function 結束前關閉已開啟檔案。
	defer file.Close()

	// 將 img 影像變數儲存成目標 file 檔，若正確執行則回傳 nil（無）錯誤。
	return bmp.Encode(file, img)
}

// WriteImgsAvgToTxt 函數將讀入的灰階影像變數每一個像素 pix 累加平均後，將每一筆像素 pix 資訊寫入至目的檔。
func WriteImgsAvgToTxt(dstFile string, imgs []image.Gray) (err error) {

	// 讀取第一張灰階影像的列數（高度）。
	row := imgs[0].Bounds().Dy()
	// 讀取第一張灰階影像的行數（寬度）。
	col := imgs[0].Bounds().Dx()
	// 建立用來暫存解析影像後的像素 pix 資訊 slice，共有 row*col 個元素。
	newPixs := make([]uint64, row*col)

	// 將 newPixs 變數元素清空，但 Golang 預設已在創建時清空。
	// for i := 0; i < row; i++ {
	// 	for j := 0; j < col; j++ {
	// 		newImg.Pix[col*i+j] = uint8(0)
	// 	}
	// }

	// 輸入的每一張影像都會進行累加。
	for h := 0; h < len(imgs); h++ {
		// 若某一張影像的列數（高度）或行數（寬度）和第一張影像不同，則回傳錯誤。
		if row != imgs[h].Bounds().Dy() || col != imgs[h].Bounds().Dx() {
			fmt.Println("Row or column number from importing images are not same!")
			return errors.New("Row or column number from importing images are not same!")
		}
		// 每一張影像的列數（高度）。
		for i := 0; i < row; i++ {
			// 每一張影像的行數（寬度）。
			for j := 0; j < col; j++ {
				// 將每一個像素的 Pix 個別進行累加，並儲存至 newPixs slice 內。
				newPixs[col*i+j] += uint64(imgs[h].Pix[col*i+j])
			}
		}
	}

	// 根據作業系統調整路徑的正反鈄線，以及將目錄路徑轉換成絕對路徑。
	dstFile, err = filepath.Abs(dstFile)
	if err != nil {
		fmt.Println("Error while finding absolute path", dstFile, "-", err)
		return err
	}
	// 建立一個新檔作為將像素 Pix 資訊儲存的目的檔位置。
	file, err := os.Create(dstFile)
	if err != nil {
		fmt.Println("Error while creating", dstFile, "-", err)
		return err
	}
	// 在 function 結束前關閉已開啟檔案。
	defer file.Close()

	// 建立一個寫檔緩衝。
	fileWriter := bufio.NewWriter(file)

	// 每一張影像的列數（高度）。
	for i := 0; i < row; i++ {
		// 每一張影像的行數（寬度）。
		for j := 0; j < col; j++ {
			// 將每一個像素 Pix 除以影像總數取得平均值。
			tmp := uint8(float64(newPixs[col*i+j]) / float64(len(imgs)))
			// 將 uint8 轉成 16 進制格式（%02X）後寫入緩衝內。
			fmt.Fprintf(fileWriter, "%02X ", tmp)
		}
		// 將下一行分隔符號寫入緩衝內。
		fmt.Fprintln(fileWriter, "")
	}

	// 寫入檔案。
	fileWriter.Flush()
	// 同步檔案。
	file.Sync()

	// 回傳 nil（無）錯誤。
	return nil
}

// ImgAddZero 函數是將輸入的 src 圖檔放置在 maxRow 列、maxCol 行的空圖檔中央
// 以擴充 src 圖檔的列及行。
func ImgAddZero(src image.Gray, maxRow, maxCol int) (newImg *image.Gray, err error) {

	// 建立一個新的灰階影像變數，像素為 maxCol*maxRow。
	newImg = image.NewGray(image.Rect(0, 0, maxCol, maxRow))

	// 將 newImg 變數元素清空，但 Golang 預設已在創建時清空。
	// for i := 0; i < maxRow; i++ {
	// 	for j := 0; j < maxCol; j++ {
	// 		newImg.Pix[maxCol*i+j] = uint8(0)
	// 	}
	// }
	// 讀取輸入的第一張灰階影像的列數（高度）。
	srcRow := src.Bounds().Dy()
	// 讀取輸入的第一張灰階影像的行數（寬度）。
	srcCol := src.Bounds().Dx()
	// 取得將原始圖檔置中的列數（高度）偏移值。
	rowOffset := int((maxRow - srcRow) / 2)
	// 取得將原始圖檔置中的行數（寬度）偏移值。
	colOffset := int((maxCol - srcCol) / 2)
	// 若輸入的原始影像比欲擴增補零的影像來的大，則出現錯誤。
	if rowOffset < 0 || colOffset < 0 {
		fmt.Println("maxRow < srcRow or maxCol < srcCol")
		return nil, errors.New("Error: maxRow < srcRow or maxCol < srcCol")
	}

	// 將原始影像像素 pix 值加入偏移值後寫入新的灰階影像。
	for i := 0; i < srcRow; i++ {
		for j := 0; j < srcCol; j++ {
			newImg.Pix[maxCol*(i+rowOffset)+(j+colOffset)] = src.Pix[srcCol*i+j]
		}
	}

	// 回傳新的灰階影像，並回傳 nil（無）錯誤。
	return newImg, nil
}

// MNIST 訓練資料欄位名稱 train-labels-idx1-ubyte 格式
// TRAINING SET LABEL FILE (train-labels-idx1-ubyte):
// [offset] [type]          [value]          [description]
// 0000     32 bit integer  0x00000801(2049) magic number (MSB first)
// 0004     32 bit integer  60000            number of items
// 0008     unsigned byte   ??               label
// 0009     unsigned byte   ??               label
// ........
// xxxx     unsigned byte   ??               label
// The labels values are 0 to 9.

// MNIST 測試資料欄位名稱 t10k-labels-idx1-ubyte 格式
// TEST SET LABEL FILE (t10k-labels-idx1-ubyte):
// [offset] [type]          [value]          [description]
// 0000     32 bit integer  0x00000801(2049) magic number (MSB first)
// 0004     32 bit integer  10000            number of items
// 0008     unsigned byte   ??               label
// 0009     unsigned byte   ??               label
// ........
// xxxx     unsigned byte   ??               label
// The labels values are 0 to 9.

// ReadMnistLabels 函數是將 MNIST 的 *labels.idx?-ubyte 封裝檔
// 解開並轉換成各別的 label 名稱。
func ReadMnistLabels(src string) (lbls []byte, err error) {

	// 根據作業系統調整路徑的正反鈄線，以及將目錄路徑轉換成絕對路徑。
	src, err = filepath.Abs(src)
	if err != nil {
		fmt.Println("Error while finding absolute path", src, "-", err)
		return nil, err
	}

	// 開啟已存在的 *labels.idx?-ubyte 檔。
	srcFile, err := os.Open(src)
	if err != nil {
		return nil, err
	}
	// 在 function 結束前關閉已開啟的檔案。
	defer srcFile.Close()

	var (
		// 用來儲存 *labels.idx?-ubyte 檔第一個檔頭 magic number。
		mgc uint32
		// 用來儲存 *labels.idx?-ubyte 檔第二個檔頭 number of items。
		num uint32
	)

	// 剪下前 32bits 的資料（因為 mgc 是 32bits），再按 Big-Endian 順序寫入至 mgc。
	if err = binary.Read(srcFile, binary.BigEndian, &mgc); err != nil {
		return nil, err
	}

	// 剪下前 32bits 的資料（因為 num 是 32bits），再按 Big-Endian 順序寫入至 num。
	if err = binary.Read(srcFile, binary.BigEndian, &num); err != nil {
		return nil, err
	}

	// 建立一用來儲存所有 lebel 的 slice。
	lbls = make([]byte, num)

	for i := 0; i < int(num); i++ {
		// 宣告一個 uint8 變數 l，用來暫存取得的 label 數值。
		var l byte
		// 用 io.ReadFull 要用 lbls[i:i+1] 來指定單一元素的 slice，
		// 若用 lbls[i] 則會出錯，因為就不是 slice 而是 byte 變數了。
		//_, err := io.ReadFull(srcFile, lbls[i:i+1])
		if err := binary.Read(srcFile, binary.BigEndian, &l); err != nil {
			return nil, err
		}
		lbls[i] = l
	}
	// 回傳所有的 label 及回傳 nil（無）錯誤。
	return lbls, nil
}

// WriteLblsAvgToTxt 函數將讀入的 label 數值累加平均後，再寫入至目的檔。
func WriteLblsAvgToTxt(dstFile string, lbls []byte) (err error) {
	var newlbl byte

	// 將輸入的 label 值累加。
	for h := 0; h < len(lbls); h++ {
		newlbl += lbls[h]
	}

	// 根據作業系統調整路徑的正反鈄線，以及將目錄路徑轉換成絕對路徑。
	dstFile, err = filepath.Abs(dstFile)
	if err != nil {
		fmt.Println("Error while finding absolute path", dstFile, "-", err)
		return err
	}

	// 開啟已存在的 *labels.idx?-ubyte 檔。
	file, err := os.Create(dstFile)
	if err != nil {
		fmt.Println("Error while creating", dstFile, "-", err)
		return err
	}
	// 在 function 結束前關閉已開啟的檔案。
	defer file.Close()
	// 建立一個寫檔緩衝。
	fileWriter := bufio.NewWriter(file)
	// 將 label 的值進行平均。
	tmp := float64(newlbl) / float64(len(lbls))
	// 將 label 平均值取小數點至第二位後（%.2f）寫入緩衝。
	fmt.Fprintf(fileWriter, "%05.2f", tmp)
	// 將下一行分隔符號寫入緩衝內。
	fmt.Fprintln(fileWriter, "")
	// 寫入檔案。
	fileWriter.Flush()
	// 同步檔案。
	file.Sync()
	// 回傳 nil（無）錯誤。
	return nil
}

//http://www.jianshu.com/p/dad73b68e7eb
//func ReadFull(r Reader,buf []byte) (n int,err error)```
//函數文檔：
//>ReadFull 精確地從 r 中將 len(buf) 個字節讀取到 buf 中。它返回覆制的字節數，如果讀取的字節較少，還會返回一個錯誤。若沒有讀取到字節，錯誤就只是 EOF。如果一個 EOF 發生在讀取了一些但不是所有的字節後，ReadFull 就會返回 ErrUnexpectedEOF。對於返回值，當且僅當 err == nil 時，才有 n == len(buf)。
