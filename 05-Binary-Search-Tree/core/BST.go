package core
// 二分搜索树中的节点为私有的结构体, 外界不需要了解二分搜索树节点的具体实现
type node struct {
	key string
	value int
	left *node
	right *node
}

func NewNode(key string,value int) *node  {
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
// 构造函数, 默认构造一棵空二分搜索树
func NewBST() *BST  {
	return &BST{root:nil,count:0}
}