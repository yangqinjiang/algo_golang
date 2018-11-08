package core

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
