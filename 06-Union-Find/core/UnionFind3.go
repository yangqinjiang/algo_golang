package core

import "algo_golang/04-Heap/core"
/**
UnionFind2比UnionFind还慢
 */
type UnionFind3 struct {
	// 我们的第二版Union-Find, 使用一个数组构建一棵指向父节点的树
	// parent[i]表示第i个元素所指向的父节点
	parent []int
	count  int   //数据个数
	sz []int // sz[i]表示以i为根的集合中元素个数
}
// 构造函数
func NewUnionFind3(n int) *UnionFind3 {
	t := UnionFind3{count:n,parent:make([]int,n,n),sz:make([]int,n,n)}
	// 初始化, 每一个parent[i]指向自己, 表示每一个元素自己自成一个集合
	for i:=0;i<n;i++  {
		t.parent[i] =i
		t.sz[i] = 1
	}
	return &t
}
// 查找过程, 查找元素p所对应的集合编号
// O(h)复杂度, h为树的高度
func (this *UnionFind3)Find(p int) int{
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
func (this *UnionFind3)IsConnected(p,q int) bool  {
	return this.Find(p) == this.Find(q)
}
// 合并元素p和元素q所属的集合
// O(h)复杂度, h为树的高度
func (this *UnionFind3)UnionElements(p,q int)  {
	pRoot := this.Find(p)
	qRoot := this.Find(q)
	if pRoot == qRoot{
		return
	}

	//根据两个元素所在树的元素个数不同判断合并方向
	//将元素个数少的集合合并到元素个数多的集合上
	if this.sz[pRoot] < this.sz[qRoot]{
		this.parent[pRoot] = qRoot
		this.sz[qRoot] += this.sz[pRoot]
	}else{
		this.parent[qRoot] = pRoot
		this.sz[pRoot] += this.sz[qRoot]
	}

}