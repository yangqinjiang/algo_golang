package main

import (
	"algo_golang/04-Heap/core"
	"fmt"
)

func main() {

	maxheap := core.NewMaxHeap(10)
	fmt.Println("maxheap=",maxheap)
	fmt.Println("IsEmpty? ",maxheap.IsEmpty())
	fmt.Println("Size =",maxheap.Size())
}
