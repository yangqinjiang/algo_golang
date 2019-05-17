package core

import "algo_golang/04-Heap/core"
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

	if this.hasEdge(v,w){
		return
	}
	this.g[v][w] = true
	if !this.directed{
		this.g[w][v] = true
	}
	this.m ++
}
// 验证图中是否有从v到w的边
func (this *DenseGraph)hasEdge(v,w int) bool {
	core.Assert( v >= 0 && v < this.n)
	core.Assert( w >= 0 && w < this.n)
	return this.g[v][w]
}
