package core

type Graph interface{
 	V() int
	AddEdge(v,w int)
}
