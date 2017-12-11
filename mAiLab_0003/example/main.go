package main

import (
	"image"
	"strconv"

	"github.com/oneleo/LeNetPractice/mAiLab_0003/mygzip"
	"github.com/oneleo/LeNetPractice/mAiLab_0003/myhttp"
	"github.com/oneleo/LeNetPractice/mAiLab_0003/mymnist"
)

func main() {

	// 定義欲下載 MNIST 的四個檔名。
	fileNames := []string{"train-images-idx3-ubyte.gz", "train-labels-idx1-ubyte.gz", "t10k-images-idx3-ubyte.gz", "t10k-labels-idx1-ubyte.gz"}

	// 定義放置解答檔的目錄。
	dstDir := "..\\answer"

	// 若欲放置解答檔的目錄不存在則建立之。
	mygzip.CreateFolder(dstDir)

	// 解答 1、下載以下四個檔案。
	for _, fn := range fileNames {
		// 下載 MNIST 的四個檔案。
		_ = myhttp.DownloadFromURL("http://yann.lecun.com/exdb/mnist/"+fn, dstDir)

		// 將下載的 MNIST gz 檔解壓縮。
		_ = mygzip.GzDecompress(dstDir+"\\"+fn, dstDir)
	}

	// 解析指定的 *images.idx?-ubyte 檔，將所有影像儲存至 imgs slice。
	imgs, _ := mymnist.ReadMnistImages(dstDir + "\\train-images.idx3-ubyte")

	// 解答 2、輸出 train-images.idx3-ubyte 檔案中的第一個圖，大小為 28x28。
	// 並將解答寫入至「2. First images.txt」文字檔內。
	_ = mymnist.WriteImgsAvgToTxt(dstDir+"\\2. First images.txt", imgs[0:1])

	// 解答 3、輸出 train-images.idx3-ubyte 檔案中前十個圖的平均圖，採無條件捨去，大小為 28x28。
	// 並將解答寫入至「3. Images 1 - 10 average.txt」文字檔內。
	_ = mymnist.WriteImgsAvgToTxt(dstDir+"\\3. Images 1 - 10 average.txt", imgs[0:10])

	// 解答 4、輸出 train-labels.idx1-ubyte 檔案中前十個 labels 的平均，精確度取至小數點以下兩位，採無條件捨去。
	// 並將解答寫入至「4. Labels 1 - 10 average.txt」文字檔內。
	lbls, _ := mymnist.ReadMnistLabels(dstDir + "\\train-labels.idx1-ubyte")
	_ = mymnist.WriteLblsAvgToTxt(dstDir+"\\4. Labels 1 - 10 average.txt", lbls[0:10])

	// 因 mymnist.ImgAddZero 輸出為單一元素，故先建立只有單一元素的 slice。
	imgAddZero := make([]image.Gray, 1)
	// 將第一張 28x28 圖片補零成 32x32 的圖片存至 tmp 變數。
	tmp, _ := mymnist.ImgAddZero(imgs[0], 32, 32)
	// 補零成 32x32 的圖片位址存成 mymnist.WriteImgsAvgToTxt 可以接受的指標格式至 imgAddZero slice。
	imgAddZero[0] = *tmp

	// 解答 5、輸出 train-images.idx3-ubyte 檔案中的第一個圖，大小為 32x32。原圖置中，多出來的地方補 0。
	// 並將解答寫入至「5. First resized image.txt」文字檔內。
	_ = mymnist.WriteImgsAvgToTxt(dstDir+"\\5. First resized image.txt", imgAddZero[0:1])

	// 解答 6、做基本題 2 時，將圖檔存成 BMP 格式。
	// 並將解答寫入至「6. First image.bmp」影像檔內。
	_ = mymnist.WriteBMP(dstDir+"\\6. First image.bmp", &imgs[0])

	// 補充 1、輸出 train-images.idx3-ubyte 檔案中的前十個圖，大小為 28x28，將圖檔存成 BMP 格式。
	for i := 0; i < 10; i++ {
		_ = mymnist.WriteBMP(dstDir+"\\train-images_"+strconv.Itoa(i)+".bmp", &imgs[i])
	}

	// 補充 2、輸出 train-images.idx3-ubyte 檔案中的前十個圖，大小為 32x32。原圖置中，多出來的地方補 0，將圖檔存成 BMP 格式。
	for i := 0; i < 10; i++ {
		imgAddZero, _ := mymnist.ImgAddZero(imgs[i], 32, 32)
		_ = mymnist.WriteBMP(dstDir+"\\train-images-resized_"+strconv.Itoa(i)+".bmp", imgAddZero)
	}

}
