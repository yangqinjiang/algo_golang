package core

import (
	"algo_golang/02-Sorting-Basic/core"
	"math/rand"
	"time"
)

/**
双路快速排序法
// 调用双路快速排序的partition
*/
type QuickSortTwoWays struct {
	myrand *rand.Rand //随机数生成器
}

/**
交换数组或切片的两个元素
*/
func (e *QuickSortTwoWays) swap(a, b *int) {
	*a, *b = *b, *a
}

// 对arr[l...r]部分进行partition操作
// 返回p, 使得arr[l...p-1] < arr[p] ; arr[p+1...r] > arr[p]
func (e *QuickSortTwoWays) __partition(arr []int, l, r int) int {
	// 随机在arr[l...r]的范围中, 选择一个数值作为标定点pivot
	e.swap(&arr[l], &arr[e.myrand.Int()%(r-l+1)+l])

	v := arr[l]
	// arr[l+1...i) <= v; arr(j...r] >= v
	i := l + 1
	j := r
	for {
		// 注意这里的边界, arr[i] < v, 不能是arr[i] <= v
		// 思考一下为什么?
		for {
			if i <= r && arr[i] < v {
				i++
			} else {
				break
			}
		}

		// 注意这里的边界, arr[j] > v, 不能是arr[j] >= v
		// 思考一下为什么?

		for {
			if j >= l+1 && arr[j] > v {
				j--
			} else {
				break
			}
		}

		// 对于上面的两个边界的设定, 有的同学在课程的问答区有很好的回答:)
		// 大家可以参考: http://coding.imooc.com/learn/questiondetail/4920.html

		if i > j {
			break
		}

		e.swap(&arr[i], &arr[j])
		i++
		j--
	}
	e.swap(&arr[l], &arr[j])

	return j
}

// 对arr[l...r]部分进行快速排序
func (e *QuickSortTwoWays) __quickSort(arr []int, l, r int) {

	// 对于小规模数组, 使用插入排序进行优化
	if l >= r {
		iss := new(core.InsertionSort)
		iss.InsertionSortLR(arr, l, r)
		return
	}
	// 调用双路快速排序的partition
	p := e.__partition(arr, l, r)
	e.__quickSort(arr, l, p-1)
	e.__quickSort(arr, p+1, r)
}
func (e *QuickSortTwoWays) QuickSort(arr []int, n int) {
	e.myrand = rand.New(rand.NewSource(time.Now().UnixNano()))
	e.__quickSort(arr, 0, n-1)
}
