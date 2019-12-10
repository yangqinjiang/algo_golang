package core

import (
	"algo_golang/04-Heap/core"
	"fmt"
)
// 稠密图 - 邻接矩阵
type DenseGraph struct {
	n int //节点数
	m int //边数
	directed bool //是否为有向图
	g [][]bool// 图的具体数据

}

func NewDenseGraph(n int,directed bool) *DenseGraph  {
	core.Assert( n >= 0)

	// m:0  // 初始化没有任何边
	// g初始化为n*n的布尔矩阵, 每一个g[i][j]均为false, 表示没有任和边

	_g := make([][]bool,n)
	for i:=0;i<n ;i++  {
		_g[i] = make([]bool,n)//false为boolean型变量的默认值
	}
	return &DenseGraph{n:n,m:0,directed:directed,g:_g }
}
// 返回节点个数
func (this *DenseGraph)V() int  {
	return  this.n
}
// 返回边的个数
func (this *DenseGraph)E() int  {
	return  this.m
}
// 向图中添加一个边
func (this *DenseGraph)AddEdge(v,w int)  {
	core.Assert( v >= 0 && v < this.n)
	core.Assert( w >= 0 && w < this.n)

	if this.HasEdge(v,w){
		return
	}
	this.g[v][w] = true
	if !this.directed{
		this.g[w][v] = true
	}
	this.m ++
}
// 验证图中是否有从v到w的边
func (this *DenseGraph)HasEdge(v,w int) bool {
	core.Assert( v >= 0 && v < this.n)
	core.Assert( w >= 0 && w < this.n)
	return this.g[v][w]
}

//显示图的信息
func (this *DenseGraph)Show()  {
	for i:=0;i<this.n ;i++  {
		for j:=0;j<this.n ;j++  {
			fmt.Print(this.g[i][j],"\t")
		}
		fmt.Println();
	}
}
//返回图中一个顶点的所有邻边
func (this *DenseGraph)Adj(v int) []int   {
	if !(v >=0 && v < this.n){
		panic("v error")
	}
	var adjV []int
	for i:=0;i<this.n ; i++ {
		if this.g[v][i]{
			adjV = append(adjV, i)
		}
	}
	return adjV
}

// 邻边迭代器, 传入一个图和一个顶点,
// 迭代在这个图中和这个顶点向连的所有顶点
type AdjIterator struct {
	g *DenseGraph// 图G的引用
	v ,index int
}
//邻边迭代器
func NewAdjIterator(g *DenseGraph,v int) *AdjIterator{
	core.Assert( v >= 0 && v <g.n)
	return &AdjIterator{v:v,index:-1,g:g}
}
//返回图G中与顶点v相连接的第一个顶点
func (this *AdjIterator)Begin() int  {
	//索引从-1开始,因为每次遍历都需要调用一次Next()
	this.index = -1
	return this.Next()
}
//返回图G中顶点V相连接的下一个顶点
func (this *AdjIterator)Next() int  {
	//从当前index开始向后搜索,直到找到一个g[v][index]为true
	for this.index +=1;this.index < this.g.V();this.index++{
		if this.g.g[this.v][this.index]{
			return this.index
		}
	}
	//若没有顶点和v相连接,则返回-1
	return -1
}
// 查看是否已经迭代完了图G中与顶点v相连接的所有顶点
func (this *AdjIterator)End()  bool{
	return this.index >= this.g.V()
}
