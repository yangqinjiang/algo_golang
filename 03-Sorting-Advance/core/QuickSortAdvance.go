package core

import (
	"math/rand"
	"time"
	"algo_golang/02-Sorting-Basic/core"
)

type QuickSort2Way struct {
	myrand *rand.Rand  //随机数生成器
}
/**
交换数组或切片的两个元素
*/
func (e *QuickSort2Way) swap(a, b *int) {
	*a, *b = *b, *a
}
// 对arr[l...r]部分进行partition操作
// 返回p, 使得arr[l...p-1] < arr[p] ; arr[p+1...r] > arr[p]
func (e *QuickSort2Way)__partition(arr []int, l,r int) int  {
	// 随机在arr[l...r]的范围中, 选择一个数值作为标定点pivot
	e.swap( &arr[l] , &arr[e.myrand.Int()%(r-l+1)+l] );

	v:=arr[l]
	j:=l// arr[l+1...j] < v ; arr[j+1...i) > v
	for i:=l+1;i<=r;i++{
		if arr[i] < v{
			j++
			e.swap(&arr[j],&arr[i])
		}
	}
	e.swap(&arr[l],&arr[j])

	return j
}

// 对arr[l...r]部分进行快速排序
func (e *QuickSort2Way)__quickSort(arr []int, l,r int)  {

	// 对于小规模数组, 使用插入排序进行优化
	if l >= r{
		iss := new(core.InsertionSort)
		iss.InsertionSortLR(arr,l,r)
		return
	}
	p := e.__partition(arr,l,r)
	e.__quickSort(arr,l,p-1)
	e.__quickSort(arr,p+1,r)
}
func (e *QuickSort2Way)QuickSort(arr []int, n int)  {
	e.myrand = rand.New(rand.NewSource(time.Now().UnixNano()))
	e.__quickSort(arr,0,n-1)
}
