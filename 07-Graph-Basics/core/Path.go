package core

import (
	"container/list"
	"fmt"
)
//计算路径
type Path struct {
	g Graph //图的引用
	s int //起始点
	visited []bool //记录dfs的过程中,节点是否被访问
	from []int //记录路径,from[i]表示查找的路径上i的上一个节点
}
//图的深度优先遍历
func (this *Path)dfs(v int){
	this.visited[v] = true
	for _,i := range this.g.Adj(v) {
		if !this.visited[i]{
			//记录路径
			this.from[i] = v
			this.dfs(i)
		}
	}
}
//构造函数,寻路算法,寻找图graph从s点到其他点的路径
func NewPath(graph Graph,s int) *Path{
	//算法初始化
	if !(s >= 0 && s < graph.V()){
		panic("assert fail")
	}
	_visited := make([]bool,graph.V())
	_from := make([]int,graph.V())
	for i:=0;i<graph.V() ;i++  {
		_visited[i] = false
		_from[i] =-1
	}
	p := &Path{g:graph,visited:_visited,from:_from,s:s}
	//寻路算法
	p.dfs(s)

	return p
}
//查询从s点到w点是否有路径
func (this *Path)HasPath(w int) bool{
	if !(w >= 0 && w < this.g.V()){
		panic("assert fail")
	}
	return  this.visited[w]
}
func (this *Path)Path(w int) []int  {
	if !this.HasPath(w){
		panic("not path to w")
	}
	//利用堆栈数据结构
	stack := list.New()
	//通过from数组逆向查找从s到w的路径,存放到堆栈中
	p := w
	for p!= -1{
		stack.PushBack(p)
		p = this.from[p]
	}

	//从栈中依次取出元素,获得顺序的从s到w的路径
	var res []int
	for 0 != stack.Len(){
		//拿出来,用临时变量保存一下
		e := stack.Back()
		//删除stack的e
		stack.Remove(e)
		res = append(res, e.Value.(int))
	}
	return res
}
func (this *Path)ShowPath(w int){
	p := this.Path(w)
	for i:=0;i< len(p);i++  {
		fmt.Print(p[i])
		if i== len(p) - 1{
			fmt.Println()
		}else{
			fmt.Print(" -> ")
		}
	}
}