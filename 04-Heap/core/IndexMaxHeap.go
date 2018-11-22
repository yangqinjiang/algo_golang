package core

import (
	"fmt"
	"github.com/smartystreets/assertions"
)

type IndexMaxHeap struct {
	data []int // 最大索引堆中的数据
	indexes []int// 最大索引堆中的索引
	count int //数组元素个数
	capacity int //容量
}

// 返回堆中的元素个数
func (e *IndexMaxHeap)Size() int  {
	return e.count
}
// 返回一个布尔值, 表示堆中是否为空
func (e *IndexMaxHeap)IsEmpty() bool {
	return e.count == 0
}

// 索引堆中, 数据之间的比较根据data的大小进行比较, 但实际操作的是索引
func (e *IndexMaxHeap)shiftUp(k int)  {

	//如果父节点小于子节点,则swap
	for k>1 && e.data[e.indexes[k/2]] < e.data[e.indexes[k]]{
		e.swap(&e.indexes[k/2],&e.indexes[k])//交换索引
		k /= 2
	}
}
// 索引堆中, 数据之间的比较根据data的大小进行比较, 但实际操作的是索引
func (e *IndexMaxHeap)shiftDown(k int)  {

	for 2*k <= e.count{
		j := 2*k//在此轮循环中,data[k]和data[j]交换位置
		if j+1 <= e.count && e.data[e.indexes[j+1]] > e.data[e.indexes[j]]{
			//左右子节点最大的一个
			j++
		}
		// data[j] 是 data[2*k]和data[2*k+1]中的最大值
		if e.data[e.indexes[k]] >= e.data[e.indexes[j]] {
			break
		}

		e.swap(&e.indexes[k],&e.indexes[j])
		k=j
	}
}

/**
交换数组或切片的两个元素
*/
func (e *IndexMaxHeap) swap(a, b *int) {
	*a, *b = *b, *a
}

// 向最大堆中插入一个新的元素 item
func (e *IndexMaxHeap)Insert(i,item int)  {

	//边界
	assertions.ShouldBeLessThanOrEqualTo(e.count + 1, e.capacity)
	assertions.ShouldBeTrue(i+1>=1 && i+1 <= e.capacity)

	i+=1
	e.data[i] =  item
	e.indexes[e.count+1] = i
	e.count ++

	e.shiftUp(e.count)
}
// 从最大索引堆中取出堆顶元素, 即索引堆中所存储的最大数据
func (e *IndexMaxHeap)ExtractMax() int  {
	assertions.ShouldBeTrue(e.count > 0)
	ret := e.data[e.indexes[1]] //读取第一个,是最大值

	//交换最后和第一个元素,使得不是最大堆
	e.swap(&e.indexes[1],&e.indexes[e.count])
	e.count --
	//进行 shiftDown
	e.shiftDown(1)

	//返回第一个
	return  ret

}
// 从最大索引堆中取出堆顶元素的索引
func (e *IndexMaxHeap)ExtractMaxIndex()  int{
	assertions.ShouldBeTrue(e.count > 0)
	ret := e.indexes[1] //读取第一个,是最大值

	//交换最后和第一个元素,使得不是最大堆
	e.swap(&e.indexes[1],&e.indexes[e.count])
	e.count --
	//进行 shiftDown
	e.shiftDown(1)

	//返回第一个
	return  ret
}
// 获取最大索引堆中的堆顶元素
func (e *IndexMaxHeap)GetMax() int  {
	assertions.ShouldBeTrue(e.count > 0)
	return e.data[e.indexes[1]]
}
// 获取最大索引堆中的堆顶元素的索引
func (e *IndexMaxHeap)GetMaxIndex() int {
	assertions.ShouldBeTrue(e.count > 0)
	return e.indexes[1] -1
}
// 获取最大索引堆中索引为i的元素
func (e *IndexMaxHeap)GetItem(i int) int  {
	assertions.ShouldBeTrue(i+1 >= 1 && i+1 <= e.capacity)
	return e.data[i+1]
}
// 将最大索引堆中索引为i的元素修改为newItem
func (e *IndexMaxHeap)Change(i,newItem int)  {

	i+=1
	e.data[i]= newItem
	// 找到indexes[j] = i, j表示data[i]在堆中的位置
	// 之后shiftUp(j), 再shiftDown(j)
	for j:=1;j<=e.count;j++{
		if e.indexes[j] == i{
			e.shiftUp(j)
			e.shiftDown(j)
			return
		}
	}
}
// 测试索引堆中的索引数组index
// 注意:这个测试在向堆中插入元素以后, 不进行extract操作有效
func (e *IndexMaxHeap)TestIndexes() bool  {
	return false
}

// 以树状打印整个堆结构
//我"抄"不出来啊~~>
func (e *IndexMaxHeap)TestPrint()  {
	fmt.Println(e.data)
}



// 构造函数, 构造一个空堆, 可容纳capacity个元素

//在Go语言中没有构造函数的概念,
// 对象的创建通常交由一个全局的创建函数来完成,
// 以NewXXX来命令,表示"构造函数":
func NewIndexMaxHeap(capacity int) *IndexMaxHeap {
	//在创建slice时第一个参数用于
	// 确定初始化该slice的大小该slice中的值为零值，
	// 第三个参数用于确定该slice的长度
	//例如: MaxHeap, make([]int,2,5)  => &{[0 0] 0}
	//这里全部设置为 0
	return &IndexMaxHeap{data:make([]int,capacity+1,capacity+1),
		indexes:make([]int,capacity+1,capacity+1),
		count:0,
		capacity:capacity}
}

func HeapSortUsingIndexMaxHeap(arr []int,n int)  {
	maxheap := NewIndexMaxHeap(n)

	//创建堆
	for i := 0; i < n; i++ {
		maxheap.Insert(i,arr[i])
	}


	//从堆中依次取出元素
	for i:=n-1;i>=0;i--{
		arr[i] = maxheap.ExtractMax()
	}
}


