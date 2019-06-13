package main

import (
	"algo_golang/02-Sorting-Basic/core"
	core2 "algo_golang/03-Sorting-Advance/core"
	"fmt"
)

func main() {

	fmt.Println("running mergeSortBU")

	// Merge Sort是我们学习的第一个O(nlogn)复杂度的算法
	// 可以在1秒之内轻松处理100万数量级的数据
	// 注意：不要轻易尝试使用SelectionSort, InsertionSort或者BubbleSort处理100万级的数据
	// 否则，你就见识了O(n^2)的算法和O(nlogn)算法的本质差异：）

	n := 50000
	iss := new(core.InsertionSort)
	ms := new(core2.MergeSort)
	// 测试1 一般性测试
	fmt.Printf("Test for random array,size = %d , random range [0,%d]\n", n, n)

	st := new(core.SortTestHelper)
	arr1 := st.GenerateRandomArray(n, 0, n)
	arr2 := st.CopyIntArray(arr1, n)
	arr3 := st.CopyIntArray(arr1, n)
	arr4 := st.CopyIntArray(arr1, n)
	arr5 := st.CopyIntArray(arr1, n)

	st.TestSort("Inserion Sort", iss.InsertionSortAdvance, arr1, n)
	st.TestSort("Merge Sort", ms.MergeSort, arr2, n)
	st.TestSort("Merge Sort 2", ms.MergeSort2, arr3, n)
	st.TestSort("Merge Sort mergeSortBU", ms.MergeSortBU, arr4, n)
	st.TestSort("Merge Sort MergeSortBUAdvance", ms.MergeSortBUAdvance, arr5, n)
	//

	fmt.Println("----")
	// 测试2 测试近乎有序的数组
	// 对于近乎有序的数组, 数组越有序, InsertionSort的时间性能越趋近于O(n)
	// 所以可以尝试, 当swapTimes比较大时, MergeSort更快
	// 但是当swapTimes小到一定程度, InsertionSort变得比MergeSort快

	swapTimes := 10
	fmt.Printf("Test for nearly ordered array, size = %d , , swap time = %d\n", n, n)
	arr1 = st.GenerateNearlyOrderedArray(n, swapTimes)
	arr2 = st.CopyIntArray(arr1, n)
	arr3 = st.CopyIntArray(arr1, n)
	arr4 = st.CopyIntArray(arr1, n)
	arr5 = st.CopyIntArray(arr1, n)

	st.TestSort("Inserion Sort", iss.InsertionSortAdvance, arr1, n)
	st.TestSort("Merge Sort", ms.MergeSort, arr2, n)
	st.TestSort("Merge Sort 2", ms.MergeSort2, arr3, n)
	st.TestSort("Merge Sort mergeSortBU", ms.MergeSortBU, arr4, n)
	st.TestSort("Merge Sort MergeSortBUAdvance", ms.MergeSortBUAdvance, arr5, n)

	fmt.Println("----Merge Sort 1000万数量级-----")

	n = 10000000 //Merge Sort BU也可以在1秒之内轻松处理1000万数量级的数据
	arr2 = st.GenerateRandomArray(n, 0, n)
	arr3 = st.CopyIntArray(arr1, n)
	arr4 = st.CopyIntArray(arr1, n)
	arr5 = st.CopyIntArray(arr1, n)

	st.TestSort("Merge Sort", ms.MergeSort, arr2, n)
	st.TestSort("Merge Sort 2", ms.MergeSort2, arr3, n)
	st.TestSort("Merge Sort mergeSortBU", ms.MergeSortBU, arr4, n)
	st.TestSort("Merge Sort MergeSortBUAdvance", ms.MergeSortBUAdvance, arr5, n)
}
