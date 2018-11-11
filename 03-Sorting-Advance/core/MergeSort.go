package core

import (
	"algo_golang/02-Sorting-Basic/core"
)

type MergeSort struct {
}

func (e *MergeSort) MergeSort(arr []int, n int) {

	e.__mergeSort(arr, 0, n-1)
}

// 递归使用归并排序,对arr[l...r]的范围进行排序
func (e *MergeSort) __mergeSort(arr []int, l, r int) {
	//边界
	if l >= r {
		return
	}
	mid := (l + r) / 2 //注意int的范围
	e.__mergeSort(arr, l, mid)
	e.__mergeSort(arr, mid+1, r) //这里是 加一
	e.__merge(arr, l, mid, r)
}

// 将arr[l...mid]和arr[mid+1...r]两部分进行归并
func (e *MergeSort) __merge(arr []int, l, mid, r int) {
	//申请另一片数组空间
	aux := make([]int, r-l+1) //这里是 加一

	for i := l; i <= r; i++ {
		aux[i-l] = arr[i]
	}
	// 初始化，i指向左半部分的起始索引位置l；j指向右半部分起始索引位置mid+1
	i := l
	j := mid + 1 //这里是 加一
	for k := l; k <= r; k++ {
		if i > mid { // 如果左半部分元素已经全部处理完毕
			arr[k] = aux[j-l]
			j++
		} else if j > r { // 如果右半部分元素已经全部处理完毕
			arr[k] = aux[i-l]
			i++
		} else if aux[i-l] < aux[j-l] { // 左半部分所指元素 < 右半部分所指元素
			arr[k] = aux[i-l]
			i++
		} else { // 左半部分所指元素 >= 右半部分所指元素
			arr[k] = aux[j-l]
			j++
		}
	}

}

//优化的合并排序

func (e *MergeSort)MergeSort2(arr []int, n int)  {
	e.__mergeSort2(arr, 0, n-1)
}
// 使用优化的归并排序算法, 对arr[l...r]的范围进行排序
func (e *MergeSort) __mergeSort2(arr []int, l,r int) {

	//优化2:对于小规模数组,使用插入排序
	if r-l <=15{
		iss := new(core.InsertionSort)
		iss.InsertionSortLR(arr,l,r)
		return
	}
	mid := (l + r) / 2 //注意int的范围
	e.__mergeSort2(arr, l, mid)
	e.__mergeSort2(arr, mid+1, r) //这里是 加一

	// 优化1: 对于arr[mid] <= arr[mid+1]的情况,不进行merge
	// 对于近乎有序的数组非常有效,但是对于一般情况,有一定的性能损失
	if arr[mid] > arr[mid+1]{
		e.__merge(arr, l, mid, r)
	}

}

// 使用自底向上的归并排序算法
func (e *MergeSort)MergeSortBU(arr []int, n int)  {

	// Merge Sort Bottom Up 无优化版本
	for sz := 1; sz<n;sz += sz{
		//注意边界
		for i := 0;i< n-sz ; i+= sz+sz{
			// 对 arr[i...i+sz-1] 和 arr[i+sz...i+2*sz-1] 进行归并
			//注意边界
			if i+sz+sz-1 > n-1{
				e.__merge(arr,i , i+sz -1 ,n-1);
			}else{
				e.__merge(arr,i , i+sz -1 ,i+sz+sz-1);
			}
		}
	}

	//优化版本 见 MergeSortBUAdvance 函数

}


// 使用自底向上的归并排序算法
//优化版本
//https://github.com/liuyubobobo/Play-with-Algorithms/blob/master/03-Sorting-Advance/Course%20Code%20(C%2B%2B)/04-Merge-Sort-Bottom-Up/main.cpp
// Merge Sort BU 也是一个O(nlogn)复杂度的算法，虽然只使用两重for循环
// 所以，Merge Sort BU也可以在1秒之内轻松处理100万数量级的数据
// 注意：不要轻易根据循环层数来判断算法的复杂度，Merge Sort BU就是一个反例
// 关于这部分陷阱，推荐看我的《玩转算法面试》课程，第二章：《面试中的复杂度分析》：）
func (e *MergeSort)MergeSortBUAdvance(arr []int, n int)  {

	step := 16 //小数组 的容量大小
	step_1 := step -1
	iss := new(core.InsertionSort)

	// 对于小数组, 使用插入排序优化
	for i:=0 ; i<n ;i+=step{
		//注意边界
		if i+step_1 > n-1{
			iss.InsertionSortLR(arr,i,n-1)
		}else{
			iss.InsertionSortLR(arr,i,i+step_1)
		}

	}

	// Merge Sort Bottom Up 无优化版本
	for sz := step; sz<n;sz += sz{
		//注意边界
		for i := 0;i< n-sz ; i+= sz+sz{
			// 对于arr[mid] <= arr[mid+1]的情况,不进行merge
			if arr[i+sz-1] > arr[i+sz]{
				if i+sz+sz-1 > n-1{
					e.__merge(arr,i , i+sz -1 ,n-1);
				}else{
					e.__merge(arr,i , i+sz -1 ,i+sz+sz-1);
				}
			}
		}
	}


}