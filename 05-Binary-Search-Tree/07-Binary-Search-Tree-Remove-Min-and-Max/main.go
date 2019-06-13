package main

import (
	"algo_golang/05-Binary-Search-Tree/core"
	"fmt"
	"math/rand"
	"strconv"
)

func main() {
	bst := core.NewBST()
	//取n个取值范围在[0,...m)的随机数放进二分搜索树中

	N:=10
	M:=100
	for i:=0;i<N;i++{
		key := rand.Int()%M
		//为了后续测试方便,这里value值取和key值一样
		value := key
		fmt.Print(key," ")
		bst.Insert(strconv.Itoa(key),value)
	}
	fmt.Println("OK")
	//测试二分搜索树的size()
	fmt.Println("Size=",bst.Size())

	mini := bst.Minimum()
	fmt.Println("minium=",mini)
	max := bst.Maximum()
	fmt.Println("maxium=",max )

	fmt.Println("测试 removeMin....")
	// 测试 removeMin
	// 输出的元素应该是从小到大排列的
	for 0 != bst.Size() {
		fmt.Print("min:",bst.Minimum())
		bst.RemoveMin()
		fmt.Println(" After removeMin,size=",bst.Size())
	}

	//---------------------
	for i:=0;i<N;i++{
		key := rand.Int()%M
		//为了后续测试方便,这里value值取和key值一样
		value := key
		fmt.Print(key," ")
		bst.Insert(strconv.Itoa(key),value)
	}
	fmt.Println("\n测试 removeMax....")
	// 测试 removeMax
	// 输出的元素应该是从大到小排列的
	for 0 != bst.Size() {
		fmt.Print("max:",bst.Maximum())
		bst.RemoveMax()
		fmt.Println(" After removeMax,size=",bst.Size())
	}
}
