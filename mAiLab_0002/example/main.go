package main

import (
	"fmt"

	"github.com/LeNetPractice/mAiLab_0002/myrandom"
)

func main() {
	ran := myrandom.MultiRdmFlt64(0, 1, 5)
	m1, s1 := myrandom.RdmFlt64(-1, 1, 100)
	m2, s2 := myrandom.MyRdmFlt64(-1, 1, 100)

	fmt.Println("1、產生五個亂數，並將其輸出。")
	fmt.Println(ran, "\n")

	fmt.Println("2、產生 N 個介於 -1 與 1 之間的亂數，計算其平均值與標準差並輸出，每個亂數的值則不用輸出。")
	fmt.Println(m1, s1, "\n")

	fmt.Println("4、自己寫一個亂數產生器。")
	fmt.Println(m2, s2, "\n")
}
