package main

import (
	core2 "algo_golang/02-Sorting-Basic/core"
	"algo_golang/04-Heap/core"
)

// 将所有的元素依次添加到最小堆中, 再在将所有元素从堆中依次取出来, 完成排序
// 使用这样的一个最小堆排序, 来检验我们的最小堆的正确性
// 思考：使用最小堆可不可以编写如第6小节所介绍的优化的快速排序算法？
func heapSortUsingMinHeap(arr []int,n int)  {
	minHeap := core.NewMinHeap(n)
	for i:=0;i<n ; i++ {
		minHeap.Insert(arr[i])
	}
	for i:=0;i<n ;i++  {
		arr[i] = minHeap.ExtractMin()
	}
}
func main()  {
	n := 10000000 // 100 w
	sth := core2.SortTestHelper{}
	arr := sth.GenerateRandomArray(n,0,n)
	sth.TestSort("Heap Sort Using Min-Heap", heapSortUsingMinHeap, arr, n)
}
