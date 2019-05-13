package core

import (
	"fmt"
	"math/rand"
	"time"
)

type UnionFindTestHelper struct {

}
// 测试第一版本的并查集, 测试元素个数为n
func (this *UnionFindTestHelper)TestUF1(n int) error  {
	uf := NewUnionFind(n)

	startTime := time.Now()
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
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
		ic := uf.IsConnected(a,b)
		if !ic{
			//return  errors.New(fmt.Sprintf("%d,%d 不属于一个集合",a,b))
		}
	}
	elapsed := time.Since(startTime)


	// 打印输出对这2n个操作的耗时
	fmt.Printf("took %f%s", elapsed.Seconds(), "s \n")
	return nil
}
