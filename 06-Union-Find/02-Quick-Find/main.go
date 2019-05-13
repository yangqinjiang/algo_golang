package main

import (
	"algo_golang/06-Union-Find/core"
	"fmt"
)

func main() {
	fmt.Println("Quick-find")
	//使用10000的数据规模
	n := 10000*1

	//虽然isConnected只需要O(1)的时间,但由于union操作需要O(n)的时间
	//总体测试过程的算法复杂度是O(n^2)的
	ufth := core.UnionFindTestHelper{}
	err := ufth.TestUF1(n)
	if err != nil{
		fmt.Println(err.Error())
	}
}
