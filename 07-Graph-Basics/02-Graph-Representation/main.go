package main

import (
	"algo_golang/07-Graph-Basics/core"
	"fmt"
)

func main()  {
	dg := core.NewDenseGraph(10,true)
	fmt.Println(dg)

	sg := core.NewSparseGraph(10,true)
	fmt.Println(sg)
}
