package main

import (
	"algo_golang/02-Sorting-Basic/core"
	core2 "algo_golang/03-Sorting-Advance/core"
	"fmt"
)

func main() {

	fmt.Println("running Quick Sort 3 ways")

	// 比较Merge Sort和双路快速排序和三路快排三种排序算法的性能效率
	// 对于包含有大量重复数据的数组, 三路快排有巨大的优势
	// 对于一般性的随机数组和近乎有序的数组, 三路快排的效率虽然不是最优的, 但是是在非常可以接受的范围里
	// 因此, 在一些语言中, 三路快排是默认的语言库函数中使用的排序算法。比如Java:)

	n := 1000000
	ms := new(core2.MergeSort)
	qs_advance := new(core2.QuickSortAdvance)
	qs_2_ways := new(core2.QuickSortTwoWays)
	qs_3_ways := new(core2.QuickSortThreeWays)
	// 测试1 一般性测试
	fmt.Printf("Test for random array,size = %d , random range [0,%d]\n", n, n)

	st := new(core.SortTestHelper)
	arr1 := st.GenerateRandomArray(n, 0, n)
	arr3 := st.CopyIntArray(arr1, n)
	arr4 := st.CopyIntArray(arr1, n)
	arr5 := st.CopyIntArray(arr1, n)

	st.TestSort("Merge Sort MergeSortBUAdvance", ms.MergeSortBUAdvance, arr1, n)

	st.TestSort("Quick Sort advance", qs_advance.QuickSort, arr3, n)
	st.TestSort("Quick Sort 2 ways", qs_2_ways.QuickSort, arr4, n)
	st.TestSort("Quick Sort 3 ways", qs_3_ways.QuickSort, arr5, n)

	fmt.Println("-------------")
	// 测试2 测试近乎有序的数组
	// 加入了随机选择标定点的步骤后, 我们的快速排序可以轻松处理近乎有序的数组
	// 但是对于近乎有序的数组, 其效率比优化后的归并排序要低, 但完全再容忍范围里
	// 思考一下为什么对于近乎有序的数组, 快排的性能比优化后的归并排序低? :)
	swapTimes := 100
	fmt.Printf("Test for nearly ordered array, size = %d , swap time =  [0,%d]\n", n, swapTimes)
	arr1 = st.GenerateNearlyOrderedArray(n, swapTimes)
	arr3 = st.CopyIntArray(arr1, n)
	arr4 = st.CopyIntArray(arr1, n)
	arr5 = st.CopyIntArray(arr1, n)

	st.TestSort("Merge Sort MergeSortBUAdvance", ms.MergeSortBUAdvance, arr1, n)

	st.TestSort("Quick Sort advance", qs_advance.QuickSort, arr3, n)
	st.TestSort("Quick Sort 2 ways", qs_2_ways.QuickSort, arr4, n)
	st.TestSort("Quick Sort 3 ways", qs_3_ways.QuickSort, arr5, n)

	fmt.Println("-------------")
	// 测试3 测试存在包含大量相同元素的数组
	// 但此时, 对于含有大量相同元素的数组, 我们的快速排序算法再次退化成了O(n^2)级别的算法
	// 思考一下为什么在这种情况下, 快排退化成了O(n^2)的算法? :)
	swapTimes = 10
	fmt.Printf("Test for random array, size = %d ,, random range [0,%d]\n", n, swapTimes)
	arr1 = st.GenerateRandomArray(n, 0, swapTimes)
	arr3 = st.CopyIntArray(arr1, n)
	arr4 = st.CopyIntArray(arr1, n)
	arr5 = st.CopyIntArray(arr1, n)

	st.TestSort("Merge Sort MergeSortBUAdvance", ms.MergeSortBUAdvance, arr1, n)

	st.TestSort("Quick Sort advance", qs_advance.QuickSort, arr3, n)
	st.TestSort("Quick Sort 2 ways", qs_2_ways.QuickSort, arr4, n)
	st.TestSort("Quick Sort 3 ways", qs_3_ways.QuickSort, arr5, n)

}
