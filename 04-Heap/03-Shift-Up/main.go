package main

import (
	"algo_golang/04-Heap/core"
	"fmt"
	"math/rand"
	"time"
)

func main() {

	maxheap := core.NewMaxHeap(100)
	fmt.Println("maxheap=",maxheap)
	fmt.Println("IsEmpty? ",maxheap.IsEmpty())
	fmt.Println("Size =",maxheap.Size())

	//插入元素
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < 50; i++ {
		//随机生成
		maxheap.Insert(r.Int()% 100)
	}

	maxheap.TestPrint()
}
