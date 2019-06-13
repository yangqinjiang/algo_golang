package core

import "algo_golang/04-Heap/core"

type UnionFind struct {
	id []int //我们的第一版Union-Find本质就是一个数组
	count int //数据个数
}
// 构造函数
func NewUnionFind(n int) *UnionFind {
	t := UnionFind{count:n,id:make([]int,n)}
	for i:=0;i<n;i++  {
		t.id[i] =i
	}
	return &t
}
//查找过程, 查找元素p所对应的集合编号
func (this *UnionFind)Find(p int) int{
	core.Assert(p >= 0 && p< this.count)
	return this.id[p]
}

//查看元素p和元素q是否所属于一个集合
// O(1)复杂度
func (this *UnionFind)IsConnected(p,q int) bool  {
	return this.Find(p) == this.Find(q)
}
//合并元素p和元素q所属的集合
//O(n)复杂度
func (this *UnionFind)UnionElements(p,q int)  {
	pID := this.Find(p)
	qID := this.Find(q)
	if pID == qID{
		return
	}
	//合并过程需要遍历一遍所有元素,将两个元素的所属集合编号合并
	for i:=0;i<this.count;i++ {
		if this.id[i] == pID{
			this.id[i] = qID
		}
	}
}