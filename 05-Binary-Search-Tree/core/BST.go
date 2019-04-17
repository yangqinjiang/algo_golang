package core

import (
	"errors"
	"fmt"
)

// 二分搜索树中的节点为私有的结构体, 外界不需要了解二分搜索树节点的具体实现
type node struct {
	key   string
	value int
	left  *node
	right *node
}

func NewNode(key string, value int) *node {
	return &node{key: key, value: value, left: nil, right: nil}
}

type BST struct {
	root  *node //根节点
	count int   //节点个数
}

// 返回二分搜索树的节点个数
func (bst *BST) Size() int {
	return bst.count
}

// 返回二分搜索树是否为空
func (bst *BST) IsEmpty() bool {
	return bst.count == 0
}

// 向二分搜索树中插入一个新的(key, value)数据对
func (bst *BST) Insert(key string, value int) {
	bst.root = bst.insert(bst.root, key, value)
}

//查看二分搜索树是否存在 键key
func (bst *BST) Contain(key string) bool {
	return bst.contain(bst.root, key)
}

// 在二分搜索树中搜索键key所对应的值。如果这个值不存在, 则返回NULL
func (bst *BST) Search(key string) (int, error) {
	return bst.search(bst.root, key)
}

///----------------private-------------------
func (bst *BST) insert(node *node, key string, value int) *node {
	// 向以node为根的二分搜索树中, 插入节点(key, value), 使用递归算法
	// 返回插入新节点后的二分搜索树的根
	if node == nil {
		bst.count++
		return NewNode(key, value)
	}
	if key == node.key {
		node.value = value
	} else if key < node.key {
		node.left = bst.insert(node.left, key, value)
	} else { // key > node.key
		node.right = bst.insert(node.right, key, value)
	}
	return node
}

// 查看以node为根的二分搜索树中是否包含键值为key的节点, 使用递归算法
func (bst *BST) contain(node *node, key string) bool {

	if nil == node {
		return false
	}
	if key == node.key {
		return true
	} else if key < node.key {
		return bst.contain(node.left, key)
	} else { //key > node.key
		return bst.contain(node.right, key)
	}

}

// 在以node为根的二分搜索树中查找key所对应的value, 递归算法
// 若value不存在, 则返回error
func (bst *BST) search(node *node, key string) (int, error) {
	if nil == node {
		return 0, errors.New("404")
	}
	if key == node.key {
		return node.value, nil
	} else if key < node.key {
		return bst.search(node.left, key)
	} else { // key > node.key
		return bst.search(node.right, key)
	}
}

//二分搜索树的前序遍历
func (bst *BST) PreOrder()  {
	bst.preOrder(bst.root)
}
func (bst *BST) preOrder(node *node)  {
	if nil != node{
		fmt.Print(node.key," ")
		bst.preOrder(node.left)
		bst.preOrder(node.right)
	}
}
//二分搜索树的中序遍历
func (bst *BST) InOrder()  {
	bst.inOrder(bst.root)
}
func (bst *BST) inOrder(node *node)  {
	if nil != node{
		bst.inOrder(node.left)
		fmt.Print(node.key," ")
		bst.inOrder(node.right)
	}
}

//二分搜索树的中序遍历
func (bst *BST) PostOrder()  {
	bst.postOrder(bst.root)
}
func (bst *BST) postOrder(node *node)  {
	if nil != node{
		bst.postOrder(node.left)
		bst.postOrder(node.right)
		fmt.Print(node.key," ")
	}
}

// 释放以node为根的二分搜索树的所有节点
// 采用后续遍历的递归算法
func (bst *BST)Destroy()  {
	bst.destroy(bst.root)
}
func (bst *BST)destroy(node *node)  {
	if nil != node{
		bst.destroy(node.left)
		bst.destroy(node.right)
		node = nil //删除本节点
		bst.count --
	}
}
// 构造函数, 默认构造一棵空二分搜索树
func NewBST() *BST {
	return &BST{root: nil, count: 0}
}
