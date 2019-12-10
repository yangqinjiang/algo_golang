package main

import (
	"algo_golang/07-Graph-Basics/core"
	"fmt"
)

func main()  {
	G1_txt := "07-Graph-Basics/05-DFS-and-Components/testG1.txt"
	G2_txt := "07-Graph-Basics/05-DFS-and-Components/testG2.txt"
	fmt.Println("\n\n==========使用两种图的存储方式读取testG1.txt文件==========\n\n")
	// 使用两种图的存储方式读取testG1.txt文件
	rg := core.ReadGraph{}
	g1 := core.NewSparseGraph(13,false)
	err := rg.Read(g1, G1_txt)
	if err != nil{
		fmt.Println(err)
		return
	}
	components1 := core.NewComponents(g1)
	fmt.Printf("Test G1.txt ,SparseGraph Component Count= %d \n",components1.Count())
	if 3 != components1.Count(){
		panic("G1.txt component count error")
	}
	fmt.Println("test G1 in Sparse Graph:")
	g1.Show()


	g2 := core.NewDenseGraph(13,false)
	err = rg.Read(g2, G1_txt)
	if err != nil{
		fmt.Println(err)
		return
	}
	components2 := core.NewComponents(g2)
	fmt.Printf("Test G1.txt ,DenseGraph Component Count= %d \n",components2.Count())
	if 3 != components2.Count(){
		panic("G1.txt component count error")
	}
	fmt.Println("test G1 in Dense Graph:")
	g2.Show()


	fmt.Println("\n\n==========使用两种图的存储方式读取testG2.txt文件==========\n\n")
	// 使用两种图的存储方式读取testG2.txt文件
	rg2 := core.ReadGraph{}
	g3 := core.NewSparseGraph(6,false)
	err  = rg2.Read(g3, G2_txt)
	if err != nil{
		fmt.Println(err)
		return
	}
	components3 := core.NewComponents(g3)
	fmt.Printf("Test G2.txt ,SparseGraph Component Count= %d \n",components3.Count())
	if 1 != components3.Count(){
		panic("G2.txt component count error")
	}
	fmt.Println("test G2 in Sparse Graph:")
	g3.Show()


	g4 := core.NewDenseGraph(6,false)
	err = rg2.Read(g4, G2_txt)
	if err != nil{
		fmt.Println(err)
		return
	}
	components4 := core.NewComponents(g4)
	fmt.Printf("Test G2.txt ,DenseGraph Component Count= %d \n",components4.Count())
	if 1 != components4.Count(){
		panic("G2.txt component count error")
	}
	fmt.Println("test G2 in Dense Graph:")
	g4.Show()
}
