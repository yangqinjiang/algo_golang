package core

type InsertionSort struct {
}

/**
交换数组或切片的两个元素
*/
func (s *InsertionSort) swap(a, b *int) {
	*a, *b = *b, *a
}
func (e *InsertionSort) InsertionSort(arr []int, n int) {

	for i := 1; i < n; i++ {
		//寻找元素arr[i]合适的插入元素

		//写法一
		//		for j:=i;j>0;j--{
		//			if arr[j] < arr[j-1]{
		//				e.swap(&arr[j],&arr[j-1])
		//			}else{
		//				break
		//			}
		//		}

		//写法二
		//注意边界
		for j := i; j > 0 && arr[j] < arr[j-1]; j-- {
			e.swap(&arr[j], &arr[j-1])
		}
	}

}

//更好的插入排序
func (e *InsertionSort) InsertionSortAdvance(arr []int, n int) {
	for i := 1; i < n; i++ {
		//TODO:写法三
		//注意边界
		e := arr[i]
		j := 0 // j保存元素e应该插入的位置
		for j = i; j > 0 && arr[j-1] > e; j-- {
			arr[j] = arr[j-1] //前一元素,往后移动
		}
		arr[j] = e //不满足条件,则是当前比较的元素
	}
}

// 对arr[l...r]范围的数组进行插入排序
func (e *InsertionSort)InsertionSortLR(arr []int,l,r int)  {
	for i := l+1; i <= r; i++ {
		//TODO:写法三
		//注意边界
		e := arr[i]
		j := 0 // j保存元素e应该插入的位置
		for j = i; j > l && arr[j-1] > e; j-- {
			arr[j] = arr[j-1] //前一元素,往后移动
		}
		arr[j] = e //不满足条件,则是当前比较的元素
	}
}
