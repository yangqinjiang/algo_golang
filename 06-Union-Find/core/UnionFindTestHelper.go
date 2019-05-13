package core

import (
	"fmt"
	"math/rand"
	"time"
)

var seed  int64 = 100
type UnionFindTestHelper struct {

}
// 测试第一版本的并查集, 测试元素个数为n
func (this *UnionFindTestHelper)TestUF1(n int)   {
	uf := NewUnionFind(n)

	startTime := time.Now()
	r := rand.New(rand.NewSource(seed))
	//进行n次操作,每次随机选择两个元素进行合并操作
	for i:=0;i<n ;i++  {
		a := r.Int()%n
		b := r.Int()%n
		uf.UnionElements(a,b)  // n^2
	}
	//再进行n次操作,每次随机选择两个元素,, 查询他们是否同属一个集合
	for i:=0;i<n ;i++  {
		a := r.Int()%n
		b := r.Int()%n
		uf.IsConnected(a,b)
	}
	elapsed := time.Since(startTime)


	// 打印输出对这2n个操作的耗时
	fmt.Printf("UF1 2*n=%d,took %f%s",2*n, elapsed.Seconds(), "s \n")

}
// 测试第二版本的并查集, 测试元素个数为n, 测试逻辑和之前是完全一样的
// 思考一下: 这样的冗余代码如何消除?
// 由于这个课程不是设计模式课程, 在这里就不过多引入相关的问题讲解了。留作给大家的思考题:)
func (this *UnionFindTestHelper)TestUF2(n int)   {
	uf := NewUnionFind2(n)

	startTime := time.Now()
	r := rand.New(rand.NewSource(seed))
	//进行n次操作,每次随机选择两个元素进行合并操作
	for i:=0;i<n ;i++  {
		a := r.Int()%n
		b := r.Int()%n
		uf.UnionElements(a,b)
	}
	//再进行n次操作,每次随机选择两个元素,, 查询他们是否同属一个集合
	for i:=0;i<n ;i++  {
		a := r.Int()%n
		b := r.Int()%n
		uf.IsConnected(a,b)

	}
	elapsed := time.Since(startTime)


	// 打印输出对这2n个操作的耗时
	fmt.Printf("UF2 2*n=%d,took %f%s",2*n, elapsed.Seconds(), "s \n")

}
// 测试第三版本的并查集, 测试元素个数为n
func (this *UnionFindTestHelper)TestUF3(n int)   {
	uf := NewUnionFind3(n)

	startTime := time.Now()
	r := rand.New(rand.NewSource(seed))
	//进行n次操作,每次随机选择两个元素进行合并操作
	for i:=0;i<n ;i++  {
		a := r.Int()%n
		b := r.Int()%n
		uf.UnionElements(a,b)
	}
	//再进行n次操作,每次随机选择两个元素,, 查询他们是否同属一个集合
	for i:=0;i<n ;i++  {
		a := r.Int()%n
		b := r.Int()%n
		uf.IsConnected(a,b)

	}
	elapsed := time.Since(startTime)


	// 打印输出对这2n个操作的耗时
	fmt.Printf("UF3 2*n=%d,took %f%s",2*n, elapsed.Seconds(), "s \n")

}