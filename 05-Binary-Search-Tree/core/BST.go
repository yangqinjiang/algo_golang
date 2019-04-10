package core
// 二分搜索树中的节点为私有的结构体, 外界不需要了解二分搜索树节点的具体实现
type node struct {
	key int
	value int
	left *node
	right *node
}

func NewNode(key int,value int) *node  {
	return &node{key:key,value:value,left:nil,right:nil}
}
type BST struct {
	root *node //根节点
	count int //节点个数
}
// 返回二分搜索树的节点个数
func (bst *BST)Size() int  {
	return bst.count
}
// 返回二分搜索树是否为空
func (bst *BST) IsEmpty() bool  {
	return bst.count == 0
}
func (bst *BST)Insert(key int,value int)  {
	bst.insert(bst.root,key,value)
}
func (bst *BST) insert(node *node,key int,value int) *node  {
	// 向以node为根的二分搜索树中, 插入节点(key, value), 使用递归算法
	// 返回插入新节点后的二分搜索树的根
	if node == nil{
		bst.count ++
		return NewNode(key,value)
	}
	if key == node.key{
		node.value = value
	}else if key < node.key{
		node.left = bst.insert(node.left,key,value)
	}else{ // key > node.key
		node.right = bst.insert(node.right,key,value)
	}
	return node
}
// 构造函数, 默认构造一棵空二分搜索树
func NewBST() *BST  {
	return &BST{root:nil,count:0}
}