package main

import (
	"fmt"
	"math/rand"
	"time"
)

type SortTestHelper struct {
	c int
}

// 生成有n个元素的随机数组,每个元素的随机范围为[rangeL, rangeR]
func (e SortTestHelper) generateRandomArray(n, rangeL, rangeR int) []int {
	fmt.Println("call generateRandomArray")
	if rangeL > rangeR {
		panic("随机范围错误")
	}

	arr := make([]int, n)
	//随机数种子
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < n; i++ {
		//随机生成
		arr[i] = r.Int()%(rangeR-rangeL+1) + rangeL
	}
	return arr
}

func (e SortTestHelper) printArray(arr []int, n int) {
	fmt.Println("call printArray")
	// 打印arr数组的所有内容
	fmt.Println(arr)
}
