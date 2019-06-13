package main

import (
	core2 "algo_golang/04-Heap/core"
	"algo_golang/05-Binary-Search-Tree/core"
	"algo_golang/05-Binary-Search-Tree/core/SequenceST"
	"fmt"
	"time"
)

func main() {
	sst := SequenceST.NewSequenceST()
	sst.Insert("1",1)
	sst.Insert("2",2)
	sst.Insert("3",3)
	core2.Assert(sst.Contain("1"))
	core2.Assert(sst.Contain("2"))
	core2.Assert(sst.Contain("3"))
	core2.Assert(!sst.Contain("4"))


	bst := core.NewBST()
	bst.Insert("1", 1)
	bst.Insert("2", 1)
	bst.Insert("3", 1)
	core2.Assert(bst.IsEmpty() != true)
	core2.Assert(bst.Size() == 3)
	core2.Assert(bst.Contain("1"))
	core2.Assert(bst.Contain("2"))
	core2.Assert(bst.Contain("3"))
	core2.Assert(false == bst.Contain("4"))

	_,err0 := bst.Search("1")
	core2.Assert(err0 == nil)
	_,err0 = bst.Search("4")
	core2.Assert(err0 != nil)

	//search
	_, err := bst.Search("3")
	core2.Assert(nil == err)

	_, err = bst.Search("4")
	core2.Assert(nil != err)



	filepath := "D:/yangqinjiang/code/golang/src/algo_golang/05-Binary-Search-Tree/04-Binary-Search-Tree-Search/bible.txt"
	fmt.Println("读取文件...",filepath)
	startTime0 := time.Now()
	words,err := core.ReadFile(filepath)
	fmt.Println("单词数量",len(words),err)
	if nil != err{
		panic(err)
	}
	elapsed0 := time.Since(startTime0)
	fmt.Printf("%s: took %f%s", "读取文件", elapsed0.Seconds(), "s \n")
	// 统计圣经中所有词的词频
	// 注: 这个词频统计法相对简陋, 没有考虑很多文本处理中的特殊问题
	// 在这里只做性能测试用
	startTime := time.Now()
	fmt.Println("测试 BST...")
	bst_words := core.NewBST()
	for _,iter := range words {

		if res,err1 := bst_words.Search(iter);err1 != nil{
			//fmt.Println(iter,err1)
			bst_words.Insert(iter,1)
		}else{
			//fmt.Println(iter,res)
			bst_words.Insert(iter,res + 1) // 已存在,则 + 1
		}
	}
	// 输出圣经中god一词出现的频率
	if bst_words.Contain("god"){
		count,_ :=bst_words.Search("god")
		fmt.Println("'god' :",count)
	}else{
		fmt.Println("No word 'god' in ",filepath)
	}
	//被测试的函数指针
	elapsed := time.Since(startTime)
	fmt.Printf("%s: took %f%s", "递归二分搜索树", elapsed.Seconds(), "s \n")


	// 统计圣经中所有词的词频
	// 注: 这个词频统计法相对简陋, 没有考虑很多文本处理中的特殊问题
	// 在这里只做性能测试用
	fmt.Println("测试 顺序查找表 SST...")
	startTime1 := time.Now()
	sst_words := SequenceST.NewSequenceST()
	for _,iter := range words {

		if res,err1 := sst_words.Search(iter);err1 != nil{
			//fmt.Println(iter,err1)
			sst_words.Insert(iter,1)
		}else{
			//fmt.Println(iter,res)
			sst_words.Insert(iter,res + 1) // 已存在,则 + 1
		}
	}
	// 输出圣经中god一词出现的频率
	if sst_words.Contain("god"){
		count,_ :=sst_words.Search("god")
		fmt.Println("'god' :",count)
	}else{
		fmt.Println("No word 'god' in ",filepath)
	}
	//被测试的函数指针
	elapsed1 := time.Since(startTime1)
	fmt.Printf("%s: took %f%s", "测试顺序查找表 SST", elapsed1.Seconds(), "s \n")

}