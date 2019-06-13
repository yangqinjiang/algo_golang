package core

// 二分查找法,在有序数组arr中,查找target
// 如果找到target,返回相应的索引index
// 如果没有找到target,返回-1
func BinarySearch(arr []int,n,target int) int  {
	// 在arr[l...r]之中查找target
	l ,r := 0, n-1
	for l <= r {
		//mid := (l+r)/2
		// 防止极端情况下的整形溢出，使用下面的逻辑求出mid
		mid := l + (r - l) / 2
		if arr[mid] == target{
			return mid
		}
		if arr[mid] > target{
			r = mid -1
		}else{
			l = mid + 1
		}
	}
	return -1
}

func BinarySearch2(arr []int,n,target int) int {
	return _BinarySearch2(arr,0,n-1,target)
}
// 用递归的方式写二分查找法
func _BinarySearch2(arr []int,l,r,target int) int  {
	if l > r{
		return  -1
	}
	//mid := (l+r)/2
	// 防止极端情况下的整形溢出，使用下面的逻辑求出mid
	mid := l + (r - l) / 2

	if arr[mid] == target{
		return mid
	}else if arr[mid] > target{
		return _BinarySearch2(arr,l,mid -1 ,target)
	}else{
		return _BinarySearch2(arr,mid+1,r,target)
	}
}

