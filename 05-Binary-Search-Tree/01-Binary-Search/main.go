package main

import (
	core2 "algo_golang/04-Heap/core"
	"algo_golang/05-Binary-Search-Tree/core"
	"fmt"
	"time"
)

// 比较非递归和递归写法的二分查找的效率
// 非递归算法在性能上有微弱优势
func main() {
	fmt.Println("二分查找法")
	n := 10000000
	a := make([]int,n)
	for i:=0 ;i< n;i++{
		a[i]=  i
	}
	startTime := time.Now()

	// 测试非递归二分查找法
	// 对于我们的待查找数组[0...N)
	// 对[0...N)区间的数值使用二分查找，最终结果应该就是数字本身
	// 对[N...2*N)区间的数值使用二分查找，因为这些数字不在arr中，结果为-1
	for i:=0;i<2*n ;i++  {
		v := core.BinarySearch(a,n,i)
		//fmt.Println("v=",v,"i=",i)//打印会消耗很大
		if i < n {
			core2.Assert(v == i)
		}else{
			core2.Assert( v == -1)
		}
	}
	//被测试的函数指针
	elapsed := time.Since(startTime)
	fmt.Printf("%s: took %f%s", "非递归二分查找法", elapsed.Seconds(), "s \n")

	startTime2 := time.Now()

	// 测试非递归二分查找法
	// 对于我们的待查找数组[0...N)
	// 对[0...N)区间的数值使用二分查找，最终结果应该就是数字本身
	// 对[N...2*N)区间的数值使用二分查找，因为这些数字不在arr中，结果为-1
	for i:=0;i<2*n ;i++  {
		v := core.BinarySearch2(a,n,i)
		//fmt.Println("v=",v,"i=",i)//打印会消耗很大
		if i < n {
			core2.Assert(v == i)
		}else{
			core2.Assert( v == -1)
		}
	}
	//被测试的函数指针
	elapsed2 := time.Since(startTime2)
	fmt.Printf("%s: took %f%s", "递归二分查找法", elapsed2.Seconds(), "s \n")
}
