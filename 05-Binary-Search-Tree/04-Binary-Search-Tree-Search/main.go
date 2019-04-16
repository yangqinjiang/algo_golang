package main

import (
	core2 "algo_golang/04-Heap/core"
	"algo_golang/05-Binary-Search-Tree/core"
)

func main() {
	bst := core.NewBST()
	bst.Insert(1, 1)
	bst.Insert(2, 1)
	bst.Insert(3, 1)
	core2.Assert(bst.IsEmpty() != true)
	core2.Assert(bst.Size() == 3)
	core2.Assert(bst.Contain(1))
	core2.Assert(bst.Contain(2))
	core2.Assert(bst.Contain(3))
	core2.Assert(false == bst.Contain(4))

	//search
	_, err := bst.Search(3)
	core2.Assert(nil == err)

	_, err = bst.Search(4)
	core2.Assert(nil != err)

}
