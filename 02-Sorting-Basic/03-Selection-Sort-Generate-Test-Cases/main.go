package main

import (
	"algo_golang/02-Sorting-Basic/core"
	"fmt"
)

func main() {
	fmt.Println("测试排序算法辅助函数")
	// 测试排序算法辅助函数
	N := 20000
	st := new(core.SortTestHelper)
	arr := st.GenerateRandomArray(N, 0, 100000)
	fmt.Println(arr)
	ss := new(core.SelectionSort)
	ss.SelectionSort(arr, N)
	// 打印arr数组的所有内容

	st.PrintArray(arr, N)

}
