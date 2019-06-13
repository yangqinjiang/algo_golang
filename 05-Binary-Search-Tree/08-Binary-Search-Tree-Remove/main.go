package main

import (
	"algo_golang/05-Binary-Search-Tree/core"
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

func main() {
	bst := core.NewBST()
	//取n个取值范围在[0,...m)的随机数放进二分搜索树中

	N:=10
	M:=100
	order := []string{}
	for i:=0;i<N;i++{
		key := rand.Int()%M
		//为了后续测试方便,这里value值取和key值一样
		value := key
		fmt.Print(key," ")
		key_str := strconv.Itoa(key)
		bst.Insert(key_str,value)
		order = append(order,key_str)
	}
	fmt.Println("OK")
	//测试二分搜索树的size()
	fmt.Println("Size=",bst.Size())

	fmt.Println("Before Shuffle:",order)
	//打乱order数组的顺序
	Shuffle(order,N);
	fmt.Println("After Shuffle:",order)

	// 乱序删除[0...n)范围里的所有元素
	for i:=0;i<N ;i++  {
		key := order[i]
		if bst.Contain(key){
			bst.Remove(key)
			fmt.Println("After remove,key= ",order[i]," ,size=",bst.Size())
		}
	}
	//最终整个二分搜索树应该为空
	if bst.Size() != 0{
		panic("最终整个二分搜索树应该为空")
	}else{
		fmt.Println("ok")
	}

}
func Shuffle(arr []string,n int)  {
	rand.Seed(time.Now().Unix()) //随机种子
	for i:= n-1; i>=0 ;i--  {
		x:=rand.Int() %(i+1)
		arr[i],arr[x] = arr[x],arr[i]
	}
}
