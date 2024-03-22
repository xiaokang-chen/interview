package main

// QuickSort 快速排序
// 4 2 7 3 8 9
// 第一次，找到pivot，左侧都小于4，右侧都大于4
// 第二次，
// 3 2 7 4 8 9
func QuickSort(arr []int, left, right int) []int {
	if left < right {
		pivot := partition(arr, left, right)
		QuickSort(arr, left, pivot+1)
		QuickSort(arr, pivot+1, right)
	}
	return arr
}

// partition 分区
func partition(arr []int, left, right int) int {
	pivot := arr[left]
	for left < right {
		// 左边比较，左边都小于
		for left < right && pivot < arr[right] {
			right--
		}
		arr[left] = arr[right]
		for left < right && pivot > arr[left] {
			left++
		}
		arr[right] = arr[left]
	}
	// 临界条件
	arr[left] = pivot
	return left
}
