package main

import (
	"algo_golang/06-Union-Find/core"
	"fmt"
)
// 对比UF1, UF2, UF3和UF4的时间性能
func main() {
	fmt.Println("06-Path-Compression")
	//使用10000的数据规模
	n := 10000*1000

	//虽然isConnected只需要O(1)的时间,但由于union操作需要O(n)的时间
	//总体测试过程的算法复杂度是O(n^2)的
	ufth := core.UnionFindTestHelper{}
	// 100万数据对于UF1来说太慢了, 不再测试
	//ufth.TestUF1(n)  // 100w+, 不要运行它
	// 对于UF2来说, 其时间性能是O(n*h)的, h为并查集表达的树的最大高度
	// 这里严格来讲, h和logn没有关系, 不过大家可以简单这么理解
	// 我们后续内容会对h进行优化, 总体而言, 这个h是远小于n的
	// 所以我们实现的UF2测试结果远远好于UF1, n越大越明显:)
	// 100万数据对于UF2来说也是很慢的, 不再测试
	//ufth.TestUF2(n)
	// 对于UF3来说, 其时间性能依然是O(n*h)的, h为并查集表达的树的最大高度
	// 但由于UF3能更高概率的保证树的平衡, 所以性能更优
	ufth.TestUF3(n)
	// UF4虽然相对UF3进行有了优化, 但优化的地方出现的情况较少,
	// 所以性能更优表现的不明显, 甚至在一些数据下性能会更差
	ufth.TestUF4(n)
	//路径压缩
	ufth.TestUF5(n)
	/**
	参考结果:
	06-Path-Compression
	UF3 Size 2*n=20000000,took 6.323000s
	UF4 Rank 2*n=20000000,took 6.529000s
	UF5 Path-Compression 2*n=20000000,took 4.906000s
	 */
}
