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

	//测试二分搜索树的前序遍历 preOrder
	fmt.Println("PreOrder...")
	bst.PreOrder()
	fmt.Println()

	//测试二分搜索树的中序遍历 inOrder
	fmt.Println("InOrder...")
	bst.InOrder()
	fmt.Println()

	//测试二分搜索树的后序遍历 postOrder
	fmt.Println("postOrder...")
	bst.PostOrder()
	fmt.Println()

	bst.Destroy()//TODO:删除失败
	//测试二分搜索树的size()
	fmt.Println("Size=",bst.Size())
	fmt.Println("postOrder...")
	bst.PostOrder()
	fmt.Println()
}
