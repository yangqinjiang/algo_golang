package core

type Graph interface{
 	V() int
	E() int
	AddEdge(v,w int)
	HasEdge(v,w int) bool
 	Show()
 	Adj(v int) []int
}
