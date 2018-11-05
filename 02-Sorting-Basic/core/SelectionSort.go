package core

/**
选择排序
*/
type SelectionSort struct {
}

/**
交换数组或切片的两个元素
*/
func (s *SelectionSort) swap(a, b *int) {
	*a, *b = *b, *a
}

/**
选择排序
*/
func (s *SelectionSort) SelectionSort(arr []int, n int) {

	for i := 0; i < n; i++ {
		//寻找[i, n)区间里的最小值
		minIndex := i
		for j := i + 1; j < n; j++ {
			if arr[j] < arr[minIndex] {
				minIndex = j
			}
		}

		s.swap(&arr[i], &arr[minIndex])

	}

}
