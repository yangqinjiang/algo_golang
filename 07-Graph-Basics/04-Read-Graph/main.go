package main

import (
	"algo_golang/07-Graph-Basics/core"
	"fmt"
)

func main()  {
	fmt.Println("\n\n==========使用两种图的存储方式读取testG1.txt文件==========\n\n")
	// 使用两种图的存储方式读取testG1.txt文件
	rg := core.ReadGraph{}
	g1 := core.NewSparseGraph(13,false)
	err := rg.Read(g1, "07-Graph-Basics/04-Read-Graph/testG1.txt")
	if err != nil{
		fmt.Println(err)
		return
	}
	fmt.Println("test G1 in Sparse Graph:")
	g1.Show()


	g2 := core.NewDenseGraph(13,false)
	err = rg.Read(g2, "07-Graph-Basics/04-Read-Graph/testG1.txt")
	if err != nil{
		fmt.Println(err)
		return
	}
	fmt.Println("test G1 in Dense Graph:")
	g2.Show()


	fmt.Println("\n\n==========使用两种图的存储方式读取testG2.txt文件==========\n\n")
	// 使用两种图的存储方式读取testG2.txt文件
	rg2 := core.ReadGraph{}
	g3 := core.NewSparseGraph(6,false)
	err  = rg2.Read(g3, "07-Graph-Basics/04-Read-Graph/testG2.txt")
	if err != nil{
		fmt.Println(err)
		return
	}
	fmt.Println("test G2 in Sparse Graph:")
	g3.Show()


	g4 := core.NewDenseGraph(6,false)
	err = rg2.Read(g4, "07-Graph-Basics/04-Read-Graph/testG2.txt")
	if err != nil{
		fmt.Println(err)
		return
	}
	fmt.Println("test G2 in Dense Graph:")
	g4.Show()
}
