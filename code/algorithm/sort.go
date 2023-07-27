package algorithm

// BubbleSort 1.冒泡排序
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

// SelectionSort 2.选择排序
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

// InsertSort 3.插入排序
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

// ShellSort 4.希尔排序
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

// MergeSort 归并排序
// 分治法
// 分：利用递归的方法将复杂问题转化为相同类型的子问题
// 治：针对这些子问题的共性统一处理
// 诀窍：找到临界条件，在此条件上将递归初始值具体化
func MergeSort(arr []int) []int {
	if len(arr) < 2 {
		return arr
	}
	middle := len(arr) / 2
	leftArr := arr[:middle]
	rightArr := arr[middle:]
	return merge(MergeSort(leftArr), MergeSort(rightArr))
}

// merge 合并
func merge(left []int, right []int) []int {
	var result []int
	for len(left) != 0 && len(right) != 0 {
		if left[0] < right[0] {
			result = append(result, left[0])
			left = left[1:]
		} else {
			result = append(result, right[0])
			right = right[1:]
		}
	}
	// 边界：剩余元素一起append
	if len(left) != 0 {
		result = append(result, left...)
	}
	if len(right) != 0 {
		result = append(result, right...)
	}
	return result
}

// QuickSort 6.快速排序
// 1. 从数列中找出一个“基准”元素
// 2. 重新排列数列，所有比基准小的放在基准前面，所有比基准大的放在后面
// 3. 递归地把小于基准的子数列和大于基准的子数列排序
func QuickSort(arr []int, left int, right int) []int {
	if left < right {
		pivot := partition(arr, left, right)
		QuickSort(arr, left, pivot-1)
		QuickSort(arr, pivot+1, right)
	}
	return arr
}

// partition 分区
func partition(arr []int, left int, right int) int {
	pivot := arr[left]
	for left < right {
		// 右侧指针往左侧移动，直到找到小于pivot的值
		for left < right && pivot < arr[right] {
			right--
		}
		arr[left] = arr[right]
		// 左侧指针往右侧移动，直到找到大于pivot的值
		for left < right && pivot >= arr[left] {
			left++
		}
		arr[right] = arr[left]
	}
	// 临界条件：left=right
	arr[left] = pivot
	return left
}

// HeapSort 堆排序
//
// 堆从逻辑结构上就是一个完全二叉树
//
// 1. 首先将待排序数组构造成一个大根堆，此时，数组最大的值就在根节点
// 2. 将根节点的数和末尾的数交换，此时，末尾的数为最大值，剩余待排序数组长度为n-1
// 3. 将n-1的待排序数组再调整为大根堆，如此反复执行，最后会得到有序数组
func HeapSort(arr []int) []int {
	// 建堆
	for i := len(arr) / 2; i >= 0; i++ {
		adjustHeap(arr, i, len(arr))
	}
	// 调整（构造堆）
	for j := len(arr) - 1; j > 0; j-- {
		// 将头和尾换位置，最大值放最后
		arr[0], arr[j] = arr[j], arr[0]
		adjustHeap(arr, 0, j)
	}
	return arr
}

// adjustHeap 构造堆
// 1. 从最后一颗子树开始，从后往前调整
// 2. 每次调整，从上往下调整
// 3. 调整为大根堆
func adjustHeap(arr []int, i int, len int) {
	return
}
