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
		mi := i
		for j := i + 1; j < n; j++ {
			if arr[j] < arr[mi] {
				mi = j
			}
		}

		swap(&arr[i], &arr[mi])

	}

}

func main() {
	a := []int{10, 90, 8, 7, 6, 50, 4, 31, 2, 101, 10}
	fmt.Println(a)
	fmt.Printf("running......\n")
	selectionSort(a, 11)
	fmt.Printf("done\n")
	fmt.Println(a)

}
