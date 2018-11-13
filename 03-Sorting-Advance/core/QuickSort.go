package core

type QuickSort struct {
	
}
/**
交换数组或切片的两个元素
*/
func (s *QuickSort) swap(a, b *int) {
	*a, *b = *b, *a
}
// 对arr[l...r]部分进行partition操作
// 返回p, 使得arr[l...p-1] < arr[p] ; arr[p+1...r] > arr[p]
func (e *QuickSort)__partition(arr []int, l,r int) int  {

	v:=arr[l]
	j:=l// arr[l+1...j] < v ; arr[j+1...i) > v
	for i:=l+1;i<=r;i++{
		if arr[i] < v{
			j++
			e.swap(&arr[j],&arr[i])
		}
	}
	e.swap(&arr[l],&arr[j])

	return j
}

// 对arr[l...r]部分进行快速排序
func (e *QuickSort)__quickSort(arr []int, l,r int)  {
	if l >= r{
		return 
	}
	p := e.__partition(arr,l,r)
	e.__quickSort(arr,l,p-1)
	e.__quickSort(arr,p+1,r)
}
func (e *QuickSort)QuickSort(arr []int, n int)  {
	e.__quickSort(arr,0,n-1)
}
