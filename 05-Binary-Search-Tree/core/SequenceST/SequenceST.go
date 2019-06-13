package SequenceST

import "errors"

// 顺序查找表中的节点为私有的结构体, 外界不需要了解顺序查找表中节点的具体实现
// 我们的顺序查找表, 内部本质是一个链表
type node struct {
	key   string
	value int
	next  *node
}

func NewNode(key string, value int) *node {
	return &node{key: key, value: value, next: nil}
}

type SequenceST struct {
	head  *node //表头
	count int   //顺序查找表中的节点个数
}

// 返回顺序查找表中的节点个数
func (bst *SequenceST) Size() int {
	return bst.count
}

// 返回顺序查找表是否为空
func (bst *SequenceST) IsEmpty() bool {
	return bst.count == 0
}

// 向顺序查找表中插入一个新的(key, value)数据对
func (bst *SequenceST) Insert(key string, value int) {
	// 查找一下整个顺序表，肯是否存在同样大小的key
	node := bst.head
	for nil != node  {
		//若在顺序表中找到了同样大小key的节点
		//则当前节点不需要插入,将该key所对应的值更新为value,后返回
		if key == node.key{
			node.value = value
			return
		}
		node = node.next//找下一个
	}
	//若顺序表中没有同样大小的key,则创建新节点,将新节点直接插在表头
	newNode := NewNode(key,value)
	newNode.next = bst.head
	bst.head = newNode
	bst.count ++
}

// 查看顺序查找表中是否包含键值为key的节点
func (bst *SequenceST) Contain(key string) bool {
	node := bst.head
	for nil != node{
		if key == node.key{
			return  true
		}
		//找下一个节点
		node = node.next
	}
	return false
}

// 在顺序查找表中查找key所对应的value, 若value不存在, 则返回NULL
func (bst *SequenceST) Search(key string) (int, error) {
	node := bst.head
	for nil != node{
		if key == node.key{
			return node.value, nil
		}
		node = node.next//找下一个
	}
	return 0, errors.New("404")
}


// 构造函数, 默认构造一棵空二分搜索树
func NewSequenceST() *SequenceST {
	return &SequenceST{head: nil, count: 0}
}
