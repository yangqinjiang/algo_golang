package core

import (
	"fmt"
	"github.com/smartystreets/assertions"
)

type MaxHeap struct {
	data []int //用数组保存堆数据
	count int //数组元素个数
	capacity int //容量
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
	return &MaxHeap{data:make([]int,capacity+1,capacity+1),
	count:0,
	capacity:capacity}
}
// 返回堆中的元素个数
func (e *MaxHeap)Size() int  {
	return e.count
}
// 返回一个布尔值, 表示堆中是否为空
func (e *MaxHeap)IsEmpty() bool {
	return e.count == 0
}

func (e *MaxHeap)shiftUp(k int)  {

	//如果父节点小于子节点,则swap
	for k>1 && e.data[k/2] < e.data[k]{
		e.swap(&e.data[k/2],&e.data[k])
		k /= 2
	}
}

/**
交换数组或切片的两个元素
*/
func (e *MaxHeap) swap(a, b *int) {
	*a, *b = *b, *a
}

func (e *MaxHeap)Insert(item int)  {

	//边界
	assertions.ShouldBeLessThanOrEqualTo(e.count + 1, e.capacity)


	e.data[e.count + 1] = item
	e.count ++
	e.shiftUp(e.count)
}

// 以树状打印整个堆结构
//我"抄"不出来啊~~>
func (e *MaxHeap)TestPrint()  {
	fmt.Println(e.data)
}




