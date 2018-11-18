package core


type MaxHeap struct {
	data []int
	count int
}

// 构造函数, 构造一个空堆, 可容纳capacity个元素

//在Go语言中没有构造函数的概念,
// 对象的创建通常交由一个全局的创建函数来完成,
// 以NewXXX来命令,表示"构造函数":
func NewMaxHeap(capacity int) *MaxHeap {
	//在创建slice时第一个参数用于
	// 确定初始化该slice的大小该slice中的值为零值，
	// 第三个参数用于确定该slice的长度
	//例如: MaxHeap, make([]int,2,5)  => &{[0 0] 0}
	//这里全部设置为 0
	return &MaxHeap{data:make([]int,capacity+1,capacity+1),count:0}
}
// 返回堆中的元素个数
func (e *MaxHeap)Size() int  {
	return e.count
}
// 返回一个布尔值, 表示堆中是否为空
func (e *MaxHeap)IsEmpty() bool {
	return e.count == 0
}




