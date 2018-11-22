package main

import (
	"algo_golang/04-Heap/core"
	"fmt"
	"math/rand"
	"time"
)

func main() {

	n := 100
	maxheap := core.NewMaxHeap(n)
	fmt.Println("maxheap=",maxheap)
	fmt.Println("IsEmpty? ",maxheap.IsEmpty())
	fmt.Println("Size =",maxheap.Size())


	//插入元素
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < n; i++ {
		//随机生成
		maxheap.Insert(r.Int()% 100)
	}
	fmt.Println("堆里的数据顺序")
	maxheap.TestPrint()

	// 将maxheap中的数据逐渐使用extractMax取出来
	// 取出来的顺序应该是按照从大到小的顺序取出来的
	fmt.Println("取出来的顺序")
	arr := make([]int,n,n)
	for i:= 0 ;i<n;i++{
		arr[i] = maxheap.ExtractMax()
		fmt.Print(arr[i]," ")
	}
	fmt.Println()
	// 确保arr数组是从大到小排列的
	for i:=1 ;i<n;i++{
		core.Assert(arr[i-1] >= arr[i])
	}

}
