package core

import "algo_golang/04-Heap/core"

type UnionFind4 struct {
	// 我们的第二版Union-Find, 使用一个数组构建一棵指向父节点的树
	// parent[i]表示第i个元素所指向的父节点
	parent []int
	count  int   //数据个数
	rank   []int // rank[i]表示以i为根的集合所表示的树的层数
}

// 构造函数
func NewUnionFind4(n int) *UnionFind4 {
	t := UnionFind4{count: n, parent: make([]int, n, n), rank: make([]int, n, n)}
	// 初始化, 每一个parent[i]指向自己, 表示每一个元素自己自成一个集合
	for i := 0; i < n; i++ {
		t.parent[i] = i
		t.rank[i] = 1
	}
	return &t
}

// 查找过程, 查找元素p所对应的集合编号
// O(h)复杂度, h为树的高度
func (this *UnionFind4) Find(p int) int {
	core.Assert(p >= 0 && p < this.count)
	// 不断去查询自己的父亲节点, 直到到达根节点
	// 根节点的特点: parent[p] == p
	for p != this.parent[p] {
		p = this.parent[p]
	}
	return p
}

// 查看元素p和元素q是否所属一个集合
// O(h)复杂度, h为树的高度
func (this *UnionFind4) IsConnected(p, q int) bool {
	return this.Find(p) == this.Find(q)
}

// 合并元素p和元素q所属的集合
// O(h)复杂度, h为树的高度
func (this *UnionFind4) UnionElements(p, q int) {
	pRoot := this.Find(p)
	qRoot := this.Find(q)

	if pRoot == qRoot {
		return
	}

	// 根据两个元素所在树的元素个数不同判断合并方向
	// 将元素个数少的集合合并到元素个数多的集合上
	if this.rank[pRoot] < this.rank[qRoot] {
		this.parent[pRoot] = qRoot
	} else if this.rank[qRoot] < this.rank[pRoot] {
		this.parent[qRoot] = pRoot

	} else { // rank[pRoot] == rank[qRoot]
		this.parent[pRoot] = qRoot
		this.rank[qRoot] += 1 // 此时, 我维护rank的值
	}

}
