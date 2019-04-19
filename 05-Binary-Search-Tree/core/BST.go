package core

import (
	"algo_golang/the3rd/queue"
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
func NewNode2(n *node) *node {
	return &node{key: n.key, value: n.value, left: n.left, right: n.right}
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

//二分搜索树的前序遍历
func (bst *BST) PreOrder() {
	bst.preOrder(bst.root)
}

//二分搜索树的中序遍历
func (bst *BST) InOrder() {
	bst.inOrder(bst.root)
}

//二分搜索树的中序遍历
func (bst *BST) PostOrder() {
	bst.postOrder(bst.root)
}

//二分搜索树的层序遍历
func (bst *BST) LevelOrder() {
	//使用队列来操作
	q := queue.New()
	q.Add(bst.root)
	for 0 != q.Length() { //不为空

		n, ok := q.Remove().(*node) //,转换成 *node, 通过断言实现类型转换

		if !ok {
			fmt.Println("node 转换失败", ok, n)
			continue
		}
		fmt.Print(n.key, " ")

		//先判断左子树, 再判断右子树
		//有左子树,则入队列
		if nil != n.left {
			q.Add(n.left)
		}
		//有右子树,则入队列
		if nil != n.right {
			q.Add(n.right)
		}
	}
}

//寻找二分搜索树最小的键值
func (bst *BST) Minimum() string {
	if 0 == bst.count{
		return ""
	}
	minNode := bst.minimun(bst.root)
	return minNode.key
}
//寻找二分搜索树最大的键值
func (bst *BST) Maximum() (string) {
	if 0 == bst.count{
		return ""
	}
	maxNode := bst.maximun(bst.root)
	return maxNode.key
}
//从二分搜索树中删除最小值所在节点
func (bst *BST)RemoveMin()  {
	if(bst.root != nil){
		bst.root = bst.removeMin(bst.root)
	}
}
//从二分搜索树中删除最大值所在节点
func (bst *BST)RemoveMax()  {
	if(bst.root != nil){
		bst.root = bst.removeMax(bst.root)
	}
}
// 从二分搜索树中删除键值为key的节点
func (bst *BST)Remove(key string)  {
	bst.root = bst.remove(bst.root,key)
}
// 释放以node为根的二分搜索树的所有节点
// 采用后续遍历的递归算法
func (bst *BST) Destroy() {
	bst.destroy(bst.root)
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

func (bst *BST) preOrder(node *node) {
	if nil != node {
		fmt.Print(node.key, " ")
		bst.preOrder(node.left)
		bst.preOrder(node.right)
	}
}

func (bst *BST) inOrder(node *node) {
	if nil != node {
		bst.inOrder(node.left)
		fmt.Print(node.key, " ")
		bst.inOrder(node.right)
	}
}

func (bst *BST) postOrder(node *node) {
	if nil != node {
		bst.postOrder(node.left)
		bst.postOrder(node.right)
		fmt.Print(node.key, " ")
	}
}
// 返回以node为根的二分搜索树的最小键值所在的节点
func (bst *BST)minimun(n *node) *node {
	if n.left == nil{
		return n
	}
	return bst.minimun(n.left)
}
// 返回以node为根的二分搜索树的最大键值所在的节点
func (bst *BST)maximun(n *node) *node {
	if n.right == nil{
		return n
	}
	return bst.maximun(n.right)
}
// 删除掉以node为根的二分搜索树中的最小节点
// 返回删除节点后新的二分搜索树的根
func (bst *BST)removeMin(n *node) *node {
	if nil == n.left{
		bst.count --
		return n.right
	}
	n.left = bst.removeMin(n.left)
	return n
}
// 删除掉以node为根的二分搜索树中的最大节点
// 返回删除节点后新的二分搜索树的根
func (bst *BST)removeMax(n *node) *node {
	if nil == n.right{
		bst.count --
		return n.left
	}
	n.right = bst.removeMax(n.right)
	return n
}
// 删除掉以node为根的二分搜索树中键值为key的节点, 递归算法
// 返回删除节点后新的二分搜索树的根
func (bst *BST)remove(node *node,key string) (*node) {
	if nil == node{
		return nil
	}
	if key < node.key{ //删左子树的元素
		node.left = bst.remove(node.left,key)
		return  node
	}else if key > node.key{//删右子树的元素
		node.right = bst.remove(node.right,key)
		return node
	}else{// key == node.key // 找到此节点
		//待删除节点左子树为空的情况
		if nil == node.left{
			bst.count --
			return node.right
		}

		//待删除节点右子树为空的情况
		if nil == node.right{
			bst.count --
			return node.left
		}

		//待删除节点左右子树都不为空的情况

		//找到 比 待删除节点 [大] 的 [最小节点],即待删除节点右子树的最小节点
		// 用这个节点顶替待删除节点的位置
		mini := bst.minimun(node.right)
		successor := NewNode2(mini)
		bst.count ++

		successor.right = bst.removeMin(node.right)
		successor.left = node.left
		// delete node
		bst.count --
		return successor
	}
}
func (bst *BST) destroy(node *node) {
	if nil != node {
		bst.destroy(node.left)
		bst.destroy(node.right)
		node = nil //删除本节点
		bst.count--
	}
}

// 构造函数, 默认构造一棵空二分搜索树
func NewBST() *BST {
	return &BST{root: nil, count: 0}
}
