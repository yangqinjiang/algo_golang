package core

import "algo_golang/04-Heap/core"
// 稀疏图 - 邻接表
type SparseGraph struct {
	n int //节点数
	m int //边数
	directed bool //是否为有向图
	g [][]int// 图的具体数据
}

func NewSparseGraph(n int,directed bool) *SparseGraph  {
	core.Assert( n >= 0)
	// g初始化为n个空的vector, 表示每一个g[i]都为空, 即没有任和边
	// m:0  // 初始化没有任何边
	return &SparseGraph{n:n,m:0,directed:directed,g: make([][]int,n)}
}
// 返回节点个数
func (this *SparseGraph)V() int  {
	return  this.n
}
// 返回边的个数
func (this *SparseGraph)E() int  {
	return  this.m
}
// 向图中添加一个边
func (this *SparseGraph)AddEdge(v,w int)  {
	core.Assert( v >= 0 && v < this.n)
	core.Assert( w >= 0 && w < this.n)

	this.g[v] = append(this.g[v], w)
	if v != w && !this.directed{
		this.g[w] = append(this.g[w], v)
	}

	this.m ++
}
// 验证图中是否有从v到w的边
func (this *SparseGraph)hasEdge(v,w int) bool {
	core.Assert( v >= 0 && v < this.n)
	core.Assert( w >= 0 && w < this.n)
	for i:=0 ; i<len(this.g[v]);i++{
		if this.g[v][i] == w{
			return  true
		}
	}
	return false
}
// 邻边迭代器, 传入一个图和一个顶点,
// 迭代在这个图中和这个顶点向连的所有顶点
type AdjIteratorSparseGraph struct {
	g *SparseGraph// 图G的引用
	v ,index int
}
//邻边迭代器
func NewAdjIteratorSparseGraph(g *SparseGraph,v int) *AdjIteratorSparseGraph{
	core.Assert( v >= 0 && v <g.n)
	return &AdjIteratorSparseGraph{v:v,index:-1,g:g}
}
// 返回图G中与顶点v相连接的第一个顶点
func (this *AdjIteratorSparseGraph)Begin() int  {

	this.index = 0
	if len(this.g.g[this.v]) > 0{
		return this.g.g[this.v][this.index]
	}
	// 若没有顶点和v相连接, 则返回-1
	return -1
}
// 返回图G中与顶点v相连接的下一个顶点
func (this *AdjIteratorSparseGraph)Next() int  {
	this.index ++
	if this.index < len(this.g.g[this.v]){
		return  this.g.g[this.v][this.index]
	}

	//若没有顶点和v相连接,则返回-1
	return -1
}
// 查看是否已经迭代完了图G中与顶点v相连接的所有顶点
func (this *AdjIteratorSparseGraph)End()  bool{
	return this.index >= len(this.g.g[this.v])
}
