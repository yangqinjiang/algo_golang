package main

import (
	"algo_golang/02-Sorting-Basic/core"
	"fmt"
)

func main() {
	fmt.Println("running...")
	N := 100000
	st := new(core.SortTestHelper)
	arr1 := st.GenerateRandomArray(N, 0, 100000)

	arr2 := st.CopyIntArray(arr1, N)

	//两个数组不是引用,是copy
	fmt.Printf("arr1 point=%p \n", &arr1)
	fmt.Printf("arr2 point=%p \n", arr2)

	//测试两种不同的排序算法
	ss := new(core.SelectionSort)
	//1
	st.TestSort("selection Sort", ss.SelectionSort, arr1, N)
	//2
	iso := new(core.InsertionSort)
	st.TestSort("Insertion Sort", iso.InsertionSort, arr2, N)

}
