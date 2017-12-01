package myrandom

import (
	"math"
	"math/rand"
	"time"

	// 手動截取 Golang math/rand 原始檔中最關鍵產生 float64 亂數的部份，作為自製的亂數產生函數。
	"github.com/LeNetPractice/mAiLab_0002/rand_fromgo"
)

// region function

// MultiRdmFlt64 會回傳一個亂數 Slice，Slice 元素個數為 num 個 [min, max) 範圍內的亂數。
func MultiRdmFlt64(min, max float64, num uint8) (out []float64) {
	// 以目前時間作為亂度種子。
	rand.Seed(time.Now().UnixNano() + 1234)
	// 建立一個具有 num 個元素的 Slice。
	out = make([]float64, num)
	for n := uint8(0); n < num; n++ {
		// 產生一個 [min, max) 範圍內的亂數。
		out[n] = (max-min)*rand.Float64() + min
	}
	return out
}

// RdmFlt64 會回傳計算 num 個亂數的 mean 值以及 sigma 值。
func RdmFlt64(min, max float64, num uint8) (mean, sigma float64) {
	// 以目前時間作為亂度種子。
	rand.Seed(time.Now().UnixNano() + 2468)
	// 定義一個具有 num 個元素的 Slice。
	rnd := make([]float64, num)
	// 定義一個加總用變數。
	sum := float64(0.0)
	for n := uint8(0); n < num; n++ {
		// 產生一個 [min, max) 範圍內的亂數。
		rnd[n] = (max-min)*rand.Float64() + min
		sum += rnd[n]
	}
	// 計算平均值。
	mean = sum / float64(num)

	sigma = 0.0
	for n := uint8(0); n < num; n++ {
		// 計算差平方，並加總。
		sigma += math.Pow(rnd[n]-mean, 2)
	}
	// 計算標準差。
	sigma = math.Sqrt(sigma / float64(num-1))

	return mean, sigma
}

// MyRdmFlt64 會回傳計算 num 個亂數的 mean 值以及 sigma 值。
func MyRdmFlt64(min, max float64, num uint8) (mean, sigma float64) {
	// 以目前時間作為亂度種子。
	rand_fromgo.Seed(time.Now().UnixNano() + 1357)
	// 定義一個具有 num 個元素的 Slice。
	rnd := make([]float64, num)
	// 定義一個加總用變數。
	sum := float64(0.0)
	for n := uint8(0); n < num; n++ {
		// 產生一個 [min, max) 範圍內的亂數，並加總。
		rnd[n] = (max-min)*rand_fromgo.Float64() + min
		sum += rnd[n]
	}
	// 計算平均值。
	mean = sum / float64(num)

	sigma = 0.0
	for n := uint8(0); n < num; n++ {
		// 計算差平方，並加總。
		sigma += math.Pow(rnd[n]-mean, 2)
	}
	// 計算標準差。
	sigma = math.Sqrt(sigma / float64(num-1))

	return mean, sigma
}

// endregion function
