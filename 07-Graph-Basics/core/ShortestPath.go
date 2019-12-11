package core
import (
	"container/list"
	"fmt"
)
type ShortestPath struct {
	g       Graph  //图的引用
	s       int    //起始点
	visited []bool //记录dfs的过程中节点是否被访问
	from    []int  //记录路径,from[i]表示查找的路径上i的上一个节点
	ord     []int  //记录路径中节点的次序,ord[i]表示i节点在路径中的次序
}

func NewShortestPath(graph Graph,s int) *ShortestPath  {
	//算法初始化
	if !(s >= 0 && s < graph.V()){
		panic("param error")
	}
	_visited := make([]bool,graph.V())
	_from := make([]int,graph.V())
	_ord := make([]int,graph.V())
	for i:=0;i<graph.V();i++{
		_visited[i] = false
		_from[i] = -1
		_ord[i] = -1
	}
	sp :=  &ShortestPath{s:s,visited:_visited,from:_from,ord:_ord,g:graph}
	//无向图最短路径算法,从s开始广度优先遍历整张图
	//golang队列的使用
	//https://blog.wolfogre.com/posts/slice-queue-vs-list-queue/
	queue := list.New()
	queue.PushBack(s)//入栈
	sp.visited[s] = true
	sp.ord[s] = 0
	for 0 != queue.Len(){
		v := queue.Remove(queue.Front()).(int)
		for _,i := range sp.g.Adj(v){
			if !sp.visited[i]{
				queue.PushBack(i)//入栈
				sp.visited[i] = true
				sp.from[i] = v
				sp.ord[i] = sp.ord[v] + 1
			}
		}
	}
	return sp
}
//查询从s点到w点是否有路径
func (this *ShortestPath)HasPath(w int) bool{
	if !(w >= 0 && w < this.g.V()){
		panic("assert fail")
	}
	return  this.visited[w]
}
// 查询从s点到w点的路径, 存放在res中
func (this *ShortestPath)Path(w int) []int  {
	if !this.HasPath(w){
		panic("not path to w")
	}
	//利用堆栈数据结构
	stack := list.New()
	//通过from数组逆向查找从s到w的路径,存放到堆栈中
	p := w
	for p!= -1{
		stack.PushBack(p)//入栈
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
func (this *ShortestPath)ShowPath(w int){
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
// 查看从s点到w点的最短路径长度
// 若从s到w不可达，返回-1
func (this *ShortestPath)Length(w int) int  {
	if !(w >= 0 && w < this.g.V()){
		panic("assert fail")
	}
	return this.ord[w]
}