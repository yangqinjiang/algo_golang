package core

import "algo_golang/04-Heap/core"
/**
UnionFind2比UnionFind还慢
 */
type UnionFind2 struct {
	// 我们的第二版Union-Find, 使用一个数组构建一棵指向父节点的树
	// parent[i]表示第i个元素所指向的父节点
	parent []int
	count  int   //数据个数
}
// 构造函数
func NewUnionFind2(n int) *UnionFind2 {
	t := UnionFind2{count:n,parent:make([]int,n,n)}
	// 初始化, 每一个parent[i]指向自己, 表示每一个元素自己自成一个集合
	for i:=0;i<n;i++  {
		t.parent[i] =i
	}
	return &t
}
// 查找过程, 查找元素p所对应的集合编号
// O(h)复杂度, h为树的高度
func (this *UnionFind2)Find(p int) int{
	core.Assert(p >= 0 && p< this.count)
	// 不断去查询自己的父亲节点, 直到到达根节点
	// 根节点的特点: parent[p] == p
	for p != this.parent[p]{
		p = this.parent[p]
	}
	return p
}

// 查看元素p和元素q是否所属一个集合
// O(h)复杂度, h为树的高度
func (this *UnionFind2)IsConnected(p,q int) bool  {
	return this.Find(p) == this.Find(q)
}
// 合并元素p和元素q所属的集合
// O(h)复杂度, h为树的高度
func (this *UnionFind2)UnionElements(p,q int)  {
	pRoot := this.Find(p)
	qRoot := this.Find(q)
	if pRoot == qRoot{
		return
	}
	this.parent[pRoot] = qRoot
}