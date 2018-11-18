package main

import (
	"algo_golang/04-Heap/core"
	"fmt"
	core2 "algo_golang/03-Sorting-Advance/core"
	core3 "algo_golang/02-Sorting-Basic/core"
)

// heapSort1, 将所有的元素依次添加到堆中, 在将所有元素从堆中依次取出来, 即完成了排序
// 无论是创建堆的过程, 还是从堆中依次取出元素的过程, 时间复杂度均为O(nlogn)
// 整个堆排序的整体时间复杂度为O(nlogn)
func heapSort1(arr []int,n int)  {
	maxheap := core.NewMaxHeap(n)

	//创建堆
	for i := 0; i < n; i++ {
		maxheap.Insert(arr[i])
	}

	//从堆中依次取出元素
	for i:=n-1;i>=0;i--{
		arr[i] = maxheap.ExtractMax()
	}
}
// heapSort2, 借助我们的heapify过程创建堆
// 此时, 创建堆的过程时间复杂度为O(n), 将所有元素依次从堆中取出来, 实践复杂度为O(nlogn)
// 堆排序的总体时间复杂度依然是O(nlogn), 但是比上述heapSort1性能更优, 因为创建堆的性能更优
func heapSort2(arr []int,n int)  {


	maxheap := core.NewMaxHeapByArray(arr,n)

	//从堆中依次取出元素
	for i:=n-1;i>=0;i--{
		arr[i] = maxheap.ExtractMax()
	}
}


func main() {

	n := 1000000
	ms := new(core2.MergeSort)
	qs_advance := new(core2.QuickSortAdvance)
	qs_2_ways := new(core2.QuickSortTwoWays)
	qs_3_ways := new(core2.QuickSortThreeWays)
	// 测试1 一般性测试
	fmt.Printf("Test for random array,size = %d , random range [0,%d]\n", n, n)

	st := new(core3.SortTestHelper)
	arr1 := st.GenerateRandomArray(n, 0, n)
	arr2 := st.CopyIntArray(arr1, n)
	arr3 := st.CopyIntArray(arr1, n)
	arr4 := st.CopyIntArray(arr1, n)
	arr5 := st.CopyIntArray(arr1, n)
	arr6 := st.CopyIntArray(arr1, n)

	st.TestSort("Merge Sort MergeSortBUAdvance", ms.MergeSortBUAdvance, arr1, n)
	st.TestSort("Quick Sort advance", qs_advance.QuickSort, arr2, n)
	st.TestSort("Quick Sort 2 ways", qs_2_ways.QuickSort, arr3, n)
	st.TestSort("Quick Sort 3 ways", qs_3_ways.QuickSort, arr4, n)
	st.TestSort("Heap Sort 1", heapSort1, arr5, n)
	st.TestSort("Heap Sort 2 Heapify", heapSort2, arr6, n)

	fmt.Println("-------------")




	// 测试2 测试近乎有序的数组
	swapTimes:=100
	fmt.Printf("Test for nearly ordered array, size = %d , swap time = %d \n", n, swapTimes)

	arr1 = st.GenerateNearlyOrderedArray(n, swapTimes)
	arr2 = st.CopyIntArray(arr1, n)
	arr3 = st.CopyIntArray(arr1, n)
	arr4 = st.CopyIntArray(arr1, n)
	arr5 = st.CopyIntArray(arr1, n)
	arr6 = st.CopyIntArray(arr1, n)

	st.TestSort("Merge Sort MergeSortBUAdvance", ms.MergeSortBUAdvance, arr1, n)
	st.TestSort("Quick Sort advance", qs_advance.QuickSort, arr2, n)
	st.TestSort("Quick Sort 2 ways", qs_2_ways.QuickSort, arr3, n)
	st.TestSort("Quick Sort 3 ways", qs_3_ways.QuickSort, arr4, n)
	st.TestSort("Heap Sort 1", heapSort1, arr5, n)
	st.TestSort("Heap Sort 2 Heapify", heapSort2, arr6, n)

	fmt.Println("-------------")

	// 测试3 测试存在包含大量相同元素的数组
	fmt.Printf("Test for random array, size = %d , random range [0,10] \n", n)

	arr1 = st.GenerateRandomArray(n, 0,10)
	//arr2 = st.CopyIntArray(arr1, n)
	arr3 = st.CopyIntArray(arr1, n)
	arr4 = st.CopyIntArray(arr1, n)
	arr5 = st.CopyIntArray(arr1, n)
	arr6 = st.CopyIntArray(arr1, n)

	st.TestSort("Merge Sort MergeSortBUAdvance", ms.MergeSortBUAdvance, arr1, n)
	// 这种情况下, 普通的QuickSort退化为O(n^2)的算法, 不做测试
	//st.TestSort("Quick Sort advance", qs_advance.QuickSort, arr2, n)
	st.TestSort("Quick Sort 2 ways", qs_2_ways.QuickSort, arr3, n)
	st.TestSort("Quick Sort 3 ways", qs_3_ways.QuickSort, arr4, n)
	st.TestSort("Heap Sort 1", heapSort1, arr5, n)
	st.TestSort("Heap Sort 2 Heapify", heapSort2, arr6, n)

	fmt.Println("-------------")
}
