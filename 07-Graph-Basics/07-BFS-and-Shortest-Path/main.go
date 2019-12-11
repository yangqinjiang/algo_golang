package main
import (
	"algo_golang/07-Graph-Basics/core"
	"fmt"
)

func main() {
	G_txt := "07-Graph-Basics/07-BFS-and-Shortest-Path/testG.txt"

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
	//比较使用深度优先遍历和广度优先遍历,获得路径的不同
	//广度优先遍历获得的是无权图的最短路径
	dfs := core.NewPath(g1,0)
	fmt.Println("DFS: ")
	dfs.ShowPath(6) //output: 0 -> 5 -> 3 -> 4 -> 6

	bfs := core.NewShortestPath(g1,0)
	fmt.Println("BFS: ")
	bfs.ShowPath(6)//output:0 -> 6
	fmt.Println()

	//----------------------------
	G1_txt := "07-Graph-Basics/07-BFS-and-Shortest-Path/testG1.txt"
	rg2 := core.ReadGraph{}
	g2 := core.NewSparseGraph(13,false)
	err2 := rg2.Read(g2, G1_txt)
	if err2 != nil{
		fmt.Println(err2)
		return
	}
	fmt.Println("test G1 in Sparse Graph:")
	g2.Show()
	fmt.Println()
	//比较使用深度优先遍历和广度优先遍历,获得路径的不同
	//广度优先遍历获得的是无权图的最短路径
	dfs2 := core.NewPath(g2,0)
	fmt.Println("DFS 2: ")
	dfs2.ShowPath(3) //output:0 -> 5 -> 4 -> 3

	bfs2 := core.NewShortestPath(g2,0)
	fmt.Println("BFS 2: ")
	bfs2.ShowPath(3)//output:0 -> 5 -> 3
	fmt.Println()
}
