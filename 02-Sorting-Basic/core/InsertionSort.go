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
