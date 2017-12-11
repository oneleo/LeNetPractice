# LeNet 實作團

## mAiLab_0002：Random Number

題目：[mAiLab_0002：Random Number](http://hemingwang.blogspot.tw/2017/04/mailab0002random-number.html)

實作團目錄：[LeNet 實作團（目錄）](http://hemingwang.blogspot.tw/2017/04/lenet.html)

### 1、建議的環境變數（Windows）

* 進到【控制台】→【系統及安全性】→【系統】→【進階系統設定】→【進階】標籤→【環境變數(N)...】→在【%USERNAME% 的使用者變數(U)】新增環境變數。

``` bat
> set GOROOT="C:\Go"
> set GOPATH="%USERPROFILE%\go;%USERPROFILE%\work"
> set GOBIN="%USERPROFILE%\go\bin"
> set PATH="%PATH%;%GOROOT%\bin;%GOBIN%;%USERPROFILE%\work\bin"
```

### 1、建議的環境變數（Linux，macOS）

``` bash
$ vim "$HOME/.bash_profile"
```

    export GOROOT="/usr/local/go"
    export GOPATH="$HOME/go:$HOME/work"
    export GOBIN="$HOME/go/bin"
    export PATH="$PATH:$GOROOT/bin:$GOBIN:$HOME/work/bin”

    source "$HOME/.profile"

``` bash
$ exec /bin/bash
```
### 2、執行與結果

``` bat
> cd %USERPROFILE%\work\src\github.com\oneleo\LeNetPractice\mAiLab_0002\example
> go run .\main.go
```

	1、產生五個亂數，並將其輸出。
	[0.21642098778892313 0.2868893195840724 0.028758584555295806 0.22360007613247682 0.5789217058500571]

	2、產生 N 個介於 -1 與 1 之間的亂數，計算其平均值與標準差並輸出，每個亂數的值則不用輸出。
	-0.043197877217242164 0.52955115264394 

	4、自己寫一個亂數產生器。
	-0.020147616917228117 0.5565119909196391

### 3、使用 go test 內建指令計算處理時間

#### 3-1、第 1 個函數（產生五個亂數，並將其輸出）處理時間。

* 【200000】：函數執行 200,000 次。
* 【8895 ns/op】：函數執行平均花費 10,782 奈秒。
* 【1.928s】：本次測試總共花費 2.294 秒。

``` bat
> cd %USERPROFILE%\work\src\github.com\oneleo\LeNetPractice\mAiLab_0002\myrandom
> go test -bench=Benchmark_MultiRdmFlt64
```

	goos: windows
	goarch: amd64
	pkg: github.com/LeNetPractice/mAiLab_0002/myrandom
	Benchmark_MultiRdmFlt64-4         200000             10782 ns/op
	PASS
	ok      github.com/LeNetPractice/mAiLab_0002/myrandom   2.294s

#### 3-2 第 2 個函數（產生 N 個介於 -1 與 1 之間的亂數，計算其平均值與標準差並輸出，每個亂數的值則不用輸出）處理時間。

* 函數執行 100,000 次。
* 函數執行平均花費 21,674 奈秒。
* 本次測試總共花費 2.426 秒。

``` bat
> cd %USERPROFILE%\work\src\github.com\oneleo\LeNetPractice\mAiLab_0002\myrandom
> go test -bench=Benchmark_RdmFlt64
```

	goos: windows
	goarch: amd64
	pkg: github.com/LeNetPractice/mAiLab_0002/myrandom
	Benchmark_RdmFlt64-4      100000             21674 ns/op
	PASS
	ok      github.com/LeNetPractice/mAiLab_0002/myrandom   2.426s

#### 3-3 第 4 個函數（自己寫一個亂數產生器）處理時間。

* 函數執行 100,000 次。
* 函數執行平均花費 22,344 奈秒。
* 本次測試總共花費 2.495 秒。

``` bat
> cd %USERPROFILE%\work\src\github.com\oneleo\LeNetPractice\mAiLab_0002\myrandom
> go test -bench=Benchmark_MyRdmFlt64
```

	goos: windows
	goarch: amd64
	pkg: github.com/LeNetPractice/mAiLab_0002/myrandom
	Benchmark_MyRdmFlt64-4            100000             22344 ns/op
	PASS
	ok      github.com/LeNetPractice/mAiLab_0002/myrandom   2.495s

### 4、結論

* 自己從 Go 的 rand 原始碼中截取亂數程式，雖然已去掉用不到的部份，但似乎不會比直接使用 Go 的 rand 還來得快。