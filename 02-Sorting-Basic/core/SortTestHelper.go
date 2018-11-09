package core

import (
	"fmt"
	"math/rand"
	"time"
)

/**
排序的helper
*/
type SortTestHelper struct {
}

// 生成有n个元素的随机数组,每个元素的随机范围为[rangeL, rangeR]
func (e *SortTestHelper) GenerateRandomArray(n, rangeL, rangeR int) []int {
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

func (e *SortTestHelper) GenerateNearlyOrderedArray(n, swapTimes int) []int {
	fmt.Println("call GenerateNearlyOrderedArray")
	arr := make([]int, n)
	for i := 0; i < n; i++ {
		//生成
		arr[i] = i
	}

	//随机数种子
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	//交换其中一些数据
	for j := 0; j < swapTimes; j++ {
		posx := r.Int() % n
		posy := r.Int() % n
		//交换
		arr[posx], arr[posy] = arr[posy], arr[posx]
	}
	return arr
}

func (e SortTestHelper) PrintArray(arr []int, n int) {
	fmt.Println("call printArray")
	// 打印arr数组的所有内容
	fmt.Println(arr)
}

// 判断arr数组是否有序
func (e SortTestHelper) IsSorted(arr []int, n int) bool {

	for i := 0; i < n-1; i++ {
		if arr[i] > arr[i+1] {
			return false
		}
	}
	return true
}

// 测试sort排序算法排序arr数组所得到结果的正确性和算法运行时间
//TODO:your_sort 应该使用接口形式
func (e SortTestHelper) TestSort(sortName string, your_sort func(arr []int, n int), arr []int, n int) {
	startTime := time.Now()
	//被测试的函数指针
	your_sort(arr, n)
	elapsed := time.Since(startTime)

	if !e.IsSorted(arr, n) {
		panic(sortName + "排序不正确")
	}

	fmt.Printf("%s: took %f%s", sortName, elapsed.Seconds(), "s \n")
}

//拷贝整型数组a中的所有元素到一个新的数组, 并返回新的数组
func (e SortTestHelper) CopyIntArray(arr []int, n int) []int {
	// 预分配足够多的元素切片
	copyData := make([]int, n)
	copy(copyData, arr)
	return copyData
}
