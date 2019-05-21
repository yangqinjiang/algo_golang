package main

import (
	"algo_golang/07-Graph-Basics/core"
	"fmt"
	"math/rand"
)

func main()  {
	N := 20
	M := 100
	fmt.Println("--------SparseGraph-------------")
	g1 := core.NewSparseGraph(N,false)
	for i:=0;i<M;i++{
		a := rand.Int()%N
		b := rand.Int()%N
		//fmt.Println(a,b)
		g1.AddEdge(a,b)
	}
	// O(V^2)

	for v:=0;v<N;v++{
		fmt.Print(" : ")
		adj := core.NewAdjIteratorSparseGraph(g1,v)
		for w:= adj.Begin(); !adj.End() ; w= adj.Next()  {
			fmt.Print(w," ")
		}
		fmt.Println()
	}
	fmt.Println("---------DenseGraph------------")
	//r := rand.New(rand.NewSource(time.Now().UnixNano()))
	// Dense Graph
	g2 := core.NewDenseGraph(N,false)
	for i:=0;i<M;i++{
		a := rand.Int()%N
		b := rand.Int()%N
		//fmt.Println(a,b)
		g2.AddEdge(a,b)
	}
	// O(V^2)

	for v:=0;v<N;v++{
		fmt.Print(" : ")
		adj := core.NewAdjIterator(g2,v)
		for w:= adj.Begin(); !adj.End() ; w= adj.Next()  {
			fmt.Print(w," ")
		}
		fmt.Println()
	}
}
