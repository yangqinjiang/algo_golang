package main

import (
	"algo_golang/02-Sorting-Basic/core"
	"fmt"
)

func main() {
	//测试1 一般测试
	fmt.Println("测试1 一般测试 running...")
	N := 10000
	st := new(core.SortTestHelper)
	arr1 := st.GenerateRandomArray(N, 0, N)

	arr2 := st.CopyIntArray(arr1, N)

	//两个数组不是引用,是copymain.go
	fmt.Printf("arr1 point=%p    |arr2 point=%p \n", arr1, arr2)

	//测试两种不同的排序算法
	ss := new(core.SelectionSort)
	//1
	st.TestSort("selection Sort:", ss.SelectionSort, arr1, N)
	//2
	iso := new(core.InsertionSort)
	st.TestSort("Insertion Sort Advance:", iso.InsertionSortAdvance, arr2, N)

	//// 测试2 有序性更强的测试
	fmt.Println("\n\n测试2 有序性更强的测试 running...")

	arr1 = st.GenerateRandomArray(N, 0, 3)

	arr2 = st.CopyIntArray(arr1, N)

	//两个数组不是引用,是copymain.go
	fmt.Printf("arr1 point=%p    |arr2 point=%p \n", arr1, arr2)

	//测试两种不同的排序算法

	//1
	st.TestSort("selection Sort:", ss.SelectionSort, arr1, N)
	//2

	st.TestSort("Insertion Sort Advance:", iso.InsertionSortAdvance, arr2, N)

	// 测试3 测试近乎有序的数组
	fmt.Println("\n\n测试3 测试近乎有序的数组 running...")

	swapTimes := 100
	arr1 = st.GenerateNearlyOrderedArray(N, swapTimes)

	arr2 = st.CopyIntArray(arr1, N)

	//两个数组不是引用,是copymain.go
	fmt.Printf("arr1 point=%p    |arr2 point=%p \n", arr1, arr2)

	//测试两种不同的排序算法
	//1
	st.TestSort("selection Sort:", ss.SelectionSort, arr1, N)
	//2

	st.TestSort("Insertion Sort Advance:", iso.InsertionSortAdvance, arr2, N)

}
