package algorithm

// BubbleSort 冒泡排序
// 时间复杂度：O(N^2)
// 空间复杂度：O(1)
func BubbleSort(arr []int) []int {
	for i := 0; i < len(arr); i++ {
		for j := 0; j < len(arr)-1-i; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}
	return arr
}

// BubbleSort2 改进的冒泡
// 假如初始数组为[7,6,5,2,8,9]，一轮排序后变为[6,5,2,7,8,9]，
// 在第一轮就可以把最大的三个数确定了，即7,8,9（因为7在于8和9比
// 较的过程中指针j没移动），所以第二轮可以直接将待排序数组缩减为[6,5,2]
func BubbleSort2(arr []int) []int {
	var i = len(arr) - 1
	for i > 0 {
		// 记录待排序序列的尾端
		pos := 0
		for j := 0; j < i; j++ {
			// 每次遇到前者大于后者，则进行交换，并且从最后一次
			// 交换的位置开始，之后都是从小到大有序的（因为在此之前已经判断过了）
			if arr[j] > arr[j+1] {
				pos = j
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
		i = pos
	}
	return arr
}

// SelectionSort 选择排序
// 时间复杂度：O(N^2)
// 空间复杂度：O(1)
func SelectionSort(arr []int) []int {
	// 这里边界条件需要注意
	// 因为比较过程中，最后一个数字一定是最大的，所以只需要len(arr) - 1次
	for i := 0; i < len(arr)-1; i++ {
		var minIndex = i
		for j := i; j < len(arr); j++ {
			if arr[j] < arr[minIndex] {
				minIndex = j
			}
		}
		// minIndex不等于i，代表需要交换位置
		if minIndex != i {
			arr[i], arr[minIndex] = arr[minIndex], arr[i]
		}
	}
	return arr
}

// InsertSort 插入排序
// 时间复杂度：O(N^2)
// 空间复杂度：O(1)
func InsertSort(arr []int) []int {
	// 这里边界条件需要注意
	// 因为比较过程是两个一起比较，所以外层从1开始
	for i := 1; i < len(arr); i++ {
		key := arr[i]
		j := i - 1
		// 内层一直循环，直到key找到自己的位置（即前方都小于自己，后面都大于自己）
		for j >= 0 && arr[j] > key {
			// 大元素向后移位置，指针向前移动
			arr[j+1] = arr[j]
			j--
		}
		// 最后的边界条件，需要注意，j+1的位置是“始终空出来”的，而j是始终用来比较的元素
		arr[j+1] = key
	}
	return arr
}

// InsertSort2 采用二分查找的插入排序
// 在元素key插入时利用二分查找
func InsertSort2(arr []int) []int {
	for i := 1; i < len(arr); i++ {
		left, right := 0, i-1
		key := arr[i]
		for left <= right {
			mid := left + (right-left)/2
			if arr[mid] == key {
				left = mid + 1
			} else if arr[mid] < key {
				left = mid + 1
			} else if arr[mid] > key {
				right = mid - 1
			}
		}
		// 最后left之后需要后移
		for j := i - 1; j >= left; j-- {
			arr[j+1] = arr[j]
		}
		arr[left] = key
	}
	return arr
}

// ShellSort 希尔排序
func ShellSort(arr []int) []int {
	d := len(arr)
	for d > 1 {
		d = d / 2
		for i := 0; i < d; i++ {

			// 进行插入排序
			for j := i + d; j < len(arr); j += d {
				key := arr[j]
				k := j - d
				for k >= 0 && arr[k] > key {
					// 后移元素
					arr[k], arr[k+d] = arr[k+d], arr[k]
					k -= d
				}
				// 将key插入找到的位置
				arr[k+d] = key
			}
		}

	}
	return arr
}