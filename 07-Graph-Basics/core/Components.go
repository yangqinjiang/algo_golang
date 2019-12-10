package core

//求无权图的联通分量
type Components struct {
	G       Graph  //图的引用
	visited []bool //记录dfs的过程中,节点是否被访问
	ccount  int    //记录联通分量个数
	id      []int  //每个节点所对应的联通分量标记
}

//图的深度优先遍历
func (this *Components) dfs(v int) {
	this.visited[v] = true
	this.id[v] = this.ccount
	for _, i := range this.G.Adj(v) {
		if !this.visited[i] {
			this.dfs(i)
		}
	}
}

//返回图的联通分量个数
func (this *Components) Count() int {
	return this.ccount
}

//查询点v和点w是否联通
func (this *Components) IsConnected(v, w int) bool {
	if !(v >= 0 && v < this.G.V()) {
		panic("v param error")
	}
	if !(w >= 0 && w < this.G.V()) {
		panic("w param error")
	}
	return this.id[v] == this.id[w]
}

//构造函数,求出无权图的联通分量
func NewComponents(graph Graph) *Components {
	//算法初始化
	_visited := make([]bool, graph.V())
	_id := make([]int, graph.V())
	//初始化数组的值
	for i := 0; i < graph.V(); i++ {
		_visited[i] = false
		_id[i] = -1
	}
	c := &Components{
		G:       graph,
		visited: _visited,
		id:      _id,
		ccount:  0,
	}
	//求图的联通分量
	for i := 0; i < c.G.V(); i++ {
		if !c.visited[i] {
			c.dfs(i)
			c.ccount++
		}
	}
	return c
}
