package core

import (
	"algo_golang/02-Sorting-Basic/core"
	"math/rand"
	"time"
)

/**
三路快速排序法
*/
type QuickSortThreeWays struct {
	myrand *rand.Rand //随机数生成器
}

/**
交换数组或切片的两个元素
*/
func (e *QuickSortThreeWays) swap(a, b *int) {
	*a, *b = *b, *a
}

// 对arr[l...r]部分进行partition操作
// 返回p, 使得arr[l...p-1] < arr[p] ; arr[p+1...r] > arr[p]
func (e *QuickSortThreeWays) __partition(arr []int, l, r int) (int, int) {
	// 随机在arr[l...r]的范围中, 选择一个数值作为标定点pivot
	e.swap(&arr[l], &arr[e.myrand.Int()%(r-l+1)+l])

	v := arr[l]

	lt := l     // arr[l+1...lt] < v
	gt := r + 1 // arr[gt...r] > v
	i := l + 1  // arr[lt+1...i) == v
	for i < gt {

		if arr[i] < v {
			e.swap(&arr[i], &arr[lt+1])
			i++
			lt++
		} else if arr[i] > v {
			e.swap(&arr[i], &arr[gt-1])
			gt--
		} else { // arr[i] == v
			i++
		}

	}
	e.swap(&arr[l], &arr[lt])
	return lt, gt

}

// 对arr[l...r]部分进行快速排序
func (e *QuickSortThreeWays) __quickSort(arr []int, l, r int) {

	// 对于小规模数组, 使用插入排序进行优化
	if l >= r {
		iss := new(core.InsertionSort)
		iss.InsertionSortLR(arr, l, r)
		return
	}
	// 调用双路快速排序的partition
	lt, gt := e.__partition(arr, l, r)
	e.__quickSort(arr, l, lt-1)
	e.__quickSort(arr, gt, r)
}
func (e *QuickSortThreeWays) QuickSort(arr []int, n int) {
	e.myrand = rand.New(rand.NewSource(time.Now().UnixNano()))
	e.__quickSort(arr, 0, n-1)
}
