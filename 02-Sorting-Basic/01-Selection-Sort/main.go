package main

import (
	"algo_golang/02-Sorting-Basic/core"
	"fmt"
)

func main() {
	a := []int{10, 90, 8, 7, 6, 50, 4, 31, 2, 101, 10}
	fmt.Println(a)
	fmt.Printf("running......\n")
	ss := new(core.SelectionSort)
	ss.SelectionSort(a, len(a))
	fmt.Printf("done\n")
	fmt.Println(a)

}
