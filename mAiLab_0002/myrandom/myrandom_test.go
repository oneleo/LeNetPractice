// How to use:
//
// 1. Testing
// (1) > cd "%GOPATH%\src\github.com\LeNetPractice\mAiLab_0002\myrandom"
// or (1) $ cd "$GOPATH/src/github.com/LeNetPractice/mAiLab_0002/myrandom"
// (2) $> go test -v
//
// 2. Benchmark:
// (1) > cd "%GOPATH%\src\github.com\LeNetPractice\mAiLab_0002\myrandom"
// or (1) $ cd "$GOPATH/src/github.com/LeNetPractice/mAiLab_0002/myrandom"
// (2) $> go test -bench={Mathod Name} -v
// such as: $> go test -bench=Benchmark_MultiRdmFlt64 -v
//
// 3. Test all:
// $> go test -bench=. -v

package myrandom

import (
	"testing"
)

// Test_MultiRdmFlt64 是測試 MultiRdmFlt64 函數所產生出來的亂數平均值及標準差要在合理範圍內。
func Test_MultiRdmFlt64(t *testing.T) {
	// 定義測試集 Struct。
	var tests = []struct {
		min float64
		max float64
		num uint8
	}{
		// 測試 1。
		{-1.0, 1.0, 100},
		// 測試 2。
		{-1000.0, -100.0, 100},
		// 測試 3。
		{100.0, 1000.0, 100},
	}
	for _, test := range tests {
		rands := MultiRdmFlt64(test.min, test.max, test.num)
		for _, r := range rands {
			// 如果產生出的 num 筆亂數不在 [min, max) 內，測試失敗。
			if r < test.min || r >= test.max {
				t.Errorf("Error range: %g, should in [%g, %g).", r, test.min, test.max)
			}
		}
	}
}

// Test_RdmFlt64 是測試 RdmFlt64 函數所產生出來的亂數平均值及標準差要在合理範圍內。
func Test_RdmFlt64(t *testing.T) {
	// 定義測試集 Struct。
	var tests = []struct {
		min float64
		max float64
		num uint8
	}{
		// 測試 1。
		{-1.0, 1.0, 100},
		// 測試 2。
		{-1000.0, -100.0, 100},
		// 測試 3。
		{100.0, 1000.0, 100},
	}
	for _, test := range tests {
		mean, sigma := RdmFlt64(test.min, test.max, test.num)
		// 如果產生出的 mean 值不在 [min, max) 內，測試失敗。
		if mean < test.min || mean >= test.max {
			t.Errorf("Error range: mean = %g, should in [%g, %g).", mean, test.min, test.max)
		}
		// 如果產生出的 sigma 值小於 0，測試失敗。
		if sigma < 0 {
			t.Errorf("Error range: sigma = %g, should be >= 0.", sigma)
		}
	}
}

// Test_MyRdmFlt64 是測試 MyRdmFlt64 函數所產生出來的亂數平均值及標準差要在合理範圍內。
func Test_MyRdmFlt64(t *testing.T) {
	// 定義測試集 Struct。
	var tests = []struct {
		min float64
		max float64
		num uint8
	}{
		// 測試 1。
		{-1.0, 1.0, 100},
		// 測試 2。
		{-1000.0, -100.0, 100},
		// 測試 3。
		{100.0, 1000.0, 100},
	}
	for _, test := range tests {
		mean, sigma := MyRdmFlt64(test.min, test.max, test.num)
		// 如果產生出的 mean 值不在 [min, max) 內，測試失敗。
		if mean < test.min || mean >= test.max {
			t.Errorf("Error range: mean = %g, should in [%g, %g).", mean, test.min, test.max)
		}
		// 如果產生出的 sigma 值小於 0，測試失敗。
		if sigma < 0 {
			t.Errorf("Error range: sigma = %g, should be >= 0.", sigma)
		}
	}
}

// Benchmark_MultiRdmFlt64 為計算 MultiRdmFlt64 函數的效能評估，計算執行 b.N 次所花費的時間。
func Benchmark_MultiRdmFlt64(b *testing.B) {
	// b.N 執行次數會依系統進行自動調整。
	for i := 0; i < b.N; i++ {
		MultiRdmFlt64(0, 1, 5)
	}
}

// Benchmark_RdmFlt64 為計算 RdmFlt64 函數的效能評估，計算執行 b.N 次所花費的時間。
func Benchmark_RdmFlt64(b *testing.B) {
	for i := 0; i < b.N; i++ {
		RdmFlt64(-1, 1, 100)
	}
}

// Benchmark_MyRdmFlt64 為計算 MyRdmFlt64 函數的效能評估，計算執行 b.N 次所花費的時間。
func Benchmark_MyRdmFlt64(b *testing.B) {
	for i := 0; i < b.N; i++ {
		MyRdmFlt64(-1, 1, 100)
	}
}
