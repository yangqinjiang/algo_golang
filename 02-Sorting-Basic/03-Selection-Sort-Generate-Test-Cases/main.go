package main

import "fmt"

/**
交换数组或切片的两个元素
*/
func swap(a, b *int) {
	*a, *b = *b, *a
}

/**
选择排序
*/
func selectionSort(arr []int, n int) {

	for i := 0; i < n; i++ {
		//寻找[i, n)区间里的最小值
		minIndex := i
		for j := i + 1; j < n; j++ {
			if arr[j] < arr[minIndex] {
				minIndex = j
			}
		}

		swap(&arr[i], &arr[minIndex])

	}

}

func main() {
	fmt.Println("测试排序算法辅助函数")
	// 测试排序算法辅助函数
	N := 20000
	st := new(SortTestHelper)
	arr := st.generateRandomArray(N, 0, 100000)
	fmt.Println(arr)
	selectionSort(arr, N)
	// 打印arr数组的所有内容

	st.printArray(arr, N)

}
