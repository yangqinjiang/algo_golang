package main
import (
	"algo_golang/07-Graph-Basics/core"
	"fmt"
)

func main() {
	G_txt := "07-Graph-Basics/06-Finding-a-Path/testG.txt"
	rg := core.ReadGraph{}
	g1 := core.NewSparseGraph(7,false)
	err := rg.Read(g1, G_txt)
	if err != nil{
		fmt.Println(err)
		return
	}

	fmt.Println("test G1 in Sparse Graph:")
	g1.Show()
	fmt.Println()
	path := core.NewPath(g1,0)
	fmt.Println("Path from 0 to 6 : ")
	path.ShowPath(6) //output: 0 -> 5 -> 3 -> 4 -> 6
}
