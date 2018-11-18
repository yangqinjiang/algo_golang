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

//Heapify
// 构造函数, 通过一个给定数组创建一个最大堆
// 该构造堆的过程, 时间复杂度为O(n)
func NewMaxHeapByArray(arr []int,n int) *MaxHeap  {

	// 索引从1开始
	o := &MaxHeap{data:make([]int,n+1,n+1),
		count:0,
		capacity:n}

		//更新堆元素
		for i:=0;i<n;i++{
			o.data[i+1] = arr[i]
		}
		//更新计数器
		o.count=n

		//从不是子节点开始,进行shiftDown
		for i:= o.count/2 ;i>=1;i--{
			o.shiftDown(i)
		}


	return o
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

func (e *MaxHeap)shiftDown(k int)  {

	for 2*k <= e.count{
		j := 2*k//在此轮循环中,data[k]和data[j]交换位置
		if j+1 <= e.count && e.data[j+1] > e.data[j]{
			j++
		}
		// data[j] 是 data[2*k]和data[2*k+1]中的最大值
		if e.data[k] >= e.data[j] {
			break
		}

		e.swap(&e.data[k],&e.data[j])
		k=j
	}
}

/**
交换数组或切片的两个元素
*/
func (e *MaxHeap) swap(a, b *int) {
	*a, *b = *b, *a
}

// 向最大堆中插入一个新的元素 item
func (e *MaxHeap)Insert(item int)  {

	//边界
	assertions.ShouldBeLessThanOrEqualTo(e.count + 1, e.capacity)


	e.data[e.count + 1] = item
	e.count ++
	e.shiftUp(e.count)
}
// 从最大堆中取出堆顶元素, 即堆中所存储的最大数据
func (e *MaxHeap)ExtractMax() int  {
	assertions.ShouldBeTrue(e.count > 0)
	ret := e.data[1] //读取第一个,是最大值

	//交换最后和第一个元素,使得不是最大堆
	e.swap(&e.data[1],&e.data[e.count])
	e.count --
	//进行 shiftDown
	e.shiftDown(1)

	//返回第一个
	return  ret

}

func (e *MaxHeap)GetMax() int  {
	assertions.ShouldBeTrue(e.count > 0)
	return e.data[1]
}

// 以树状打印整个堆结构
//我"抄"不出来啊~~>
func (e *MaxHeap)TestPrint()  {
	fmt.Println(e.data)
}




