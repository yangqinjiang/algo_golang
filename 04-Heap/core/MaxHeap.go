package core

import (
	"fmt"
	"github.com/smartystreets/assertions"
)

type MaxHeap struct {
	data     []int //用数组保存堆数据
	count    int   //数组元素个数
	capacity int   //容量
}

// 返回堆中的元素个数
func (e *MaxHeap) Size() int {
	return e.count
}

// 返回一个布尔值, 表示堆中是否为空
func (e *MaxHeap) IsEmpty() bool {
	return e.count == 0
}

func (e *MaxHeap) shiftUp(k int) {

	//如果父节点小于子节点,则swap
	for k > 1 && e.data[k/2] < e.data[k] {
		e.swap(&e.data[k/2], &e.data[k])
		k /= 2
	}
}

func (e *MaxHeap) shiftDown(k int) {

	for 2*k <= e.count {
		j := 2 * k //在此轮循环中,data[k]和data[j]交换位置
		if j+1 <= e.count && e.data[j+1] > e.data[j] {
			//左右子节点最大的一个
			j++
		}
		// data[j] 是 data[2*k]和data[2*k+1]中的最大值
		if e.data[k] >= e.data[j] {
			break
		}

		e.swap(&e.data[k], &e.data[j])
		k = j
	}
}

// 原始的shiftDown过程
func (e *MaxHeap) shiftDown2(n, k int) {

	for 2*k+1 < n {
		j := 2*k + 1 //在此轮循环中,data[k]和data[j]交换位置
		if j+1 < n && e.data[j+1] > e.data[j] {
			//左右子节点最大的一个
			j++
		}
		// data[j] 是 data[2*k]和data[2*k+1]中的最大值
		if e.data[k] >= e.data[j] {
			break
		}

		e.swap(&e.data[k], &e.data[j])
		k = j
	}
}

// 优化的shiftDown过程, 使用赋值的方式取代不断的swap,
// 该优化思想和我们之前对插入排序进行优化的思路是一致的
func (e *MaxHeap) shiftDown3(n, k int) {

	v := e.data[k]

	for 2*k+1 < n {
		j := 2*k + 1 //在此轮循环中,data[k]和data[j]交换位置
		if j+1 < n && e.data[j+1] > e.data[j] {
			//左右子节点最大的一个
			j++
		}
		// data[j] 是 data[2*k]和data[2*k+1]中的最大值
		if v >= e.data[j] {
			break
		}
		e.data[k] = e.data[j]
		k = j
	}

	e.data[k] = v
}

/**
交换数组或切片的两个元素
*/
func (e *MaxHeap) swap(a, b *int) {
	*a, *b = *b, *a
}

// 向最大堆中插入一个新的元素 item
func (e *MaxHeap) Insert(item int) {

	//边界
	assertions.ShouldBeLessThanOrEqualTo(e.count+1, e.capacity)

	e.data[e.count+1] = item
	e.count++
	e.shiftUp(e.count)
}

// 从最大堆中取出堆顶元素, 即堆中所存储的最大数据
func (e *MaxHeap) ExtractMax() int {
	assertions.ShouldBeTrue(e.count > 0)
	ret := e.data[1] //读取第一个,是最大值

	//交换最后和第一个元素,使得不是最大堆
	e.swap(&e.data[1], &e.data[e.count])
	e.count--
	//进行 shiftDown
	e.shiftDown(1)

	//返回第一个
	return ret

}

func (e *MaxHeap) GetMax() int {
	assertions.ShouldBeTrue(e.count > 0)
	return e.data[1]
}

// 以树状打印整个堆结构
//我"抄"不出来啊~~>
func (e *MaxHeap) TestPrint() {
	fmt.Println(e.data)
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
	return &MaxHeap{data: make([]int, capacity+1, capacity+1),
		count:    0,
		capacity: capacity}
}

//Heapify
// 构造函数, 通过一个给定数组创建一个最大堆
// 该构造堆的过程, 时间复杂度为O(n)
func NewMaxHeapByArray(arr []int, n int) *MaxHeap {

	// 索引从1开始
	o := &MaxHeap{data: make([]int, n+1, n+1),
		count:    0,
		capacity: n}

	//更新堆元素
	for i := 0; i < n; i++ {
		o.data[i+1] = arr[i]
	}
	//更新计数器
	o.count = n

	//从不是子节点开始,进行shiftDown
	for i := o.count / 2; i >= 1; i-- {
		o.shiftDown(i)
	}

	return o
}

// heapSort1, 将所有的元素依次添加到堆中, 在将所有元素从堆中依次取出来, 即完成了排序
// 无论是创建堆的过程, 还是从堆中依次取出元素的过程, 时间复杂度均为O(nlogn)
// 整个堆排序的整体时间复杂度为O(nlogn)
func HeapSort1(arr []int, n int) {
	maxheap := NewMaxHeap(n)

	//创建堆
	for i := 0; i < n; i++ {
		maxheap.Insert(arr[i])
	}

	//从堆中依次取出元素
	for i := n - 1; i >= 0; i-- {
		arr[i] = maxheap.ExtractMax()
	}
}

// heapSort2, 借助我们的heapify过程创建堆
// 此时, 创建堆的过程时间复杂度为O(n), 将所有元素依次从堆中取出来, 实践复杂度为O(nlogn)
// 堆排序的总体时间复杂度依然是O(nlogn), 但是比上述heapSort1性能更优, 因为创建堆的性能更优
func HeapSort2(arr []int, n int) {

	maxheap := NewMaxHeapByArray(arr, n)

	//从堆中依次取出元素
	for i := n - 1; i >= 0; i-- {
		arr[i] = maxheap.ExtractMax()
	}
}

/**
不使用一个额外的最大堆,直接在原数组上进行原地的堆排序
*/
func HeapSort3(arr []int, n int) {
	// 注意，此时我们的堆是从0开始索引的
	// 从(最后一个元素的索引-1)/2开始
	// 最后一个元素的索引 = n-1

	// 索引从1开始
	o := &MaxHeap{data: make([]int, n, n),
		count:    0,
		capacity: n}

	//更新堆元素
	o.data = arr

	//更新计数器
	o.count = n

	//从不是子节点开始,进行shiftDown
	for i := (n - 1 - 1) / 2; i >= 0; i-- {
		o.shiftDown2(n, i)
	}

	for i := n - 1; i > 0; i-- {
		//交换最大值到最后的位置
		o.swap(&o.data[0], &o.data[i])
		o.shiftDown2(i, 0)
	}

}

/**
不使用一个额外的最大堆,直接在原数组上进行原地的堆排序
*/
func HeapSort4(arr []int, n int) {
	// 注意，此时我们的堆是从0开始索引的
	// 从(最后一个元素的索引-1)/2开始
	// 最后一个元素的索引 = n-1

	// 索引从1开始
	o := &MaxHeap{data: make([]int, n, n),
		count:    0,
		capacity: n}

	//更新堆元素
	o.data = arr

	//更新计数器
	o.count = n

	//从不是子节点开始,进行shiftDown
	for i := (n - 1 - 1) / 2; i >= 0; i-- {
		o.shiftDown3(n, i)
	}

	for i := n - 1; i > 0; i-- {
		//交换最大值到最后的位置
		o.swap(&o.data[0], &o.data[i])
		o.shiftDown3(i, 0)
	}

}
