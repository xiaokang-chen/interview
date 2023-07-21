package algorithm

// BinarySearch 二分查找
func BinarySearch(nums []int, target int) int {
	left, right := 0, len(nums)-1
	for left <= right {
		// [ left + (right-left)/2 ] 和 [ (left + right)/2 ] 结果相同
		// 但是前者有效防止了left和right太大，直接相加导致的溢出问题
		mid := left + (right-left)/2
		if nums[mid] == target {
			return mid
		} else if nums[mid] < target {
			// 右侧寻找
			left = mid + 1
		} else if nums[mid] > target {
			right = mid - 1
		}
	}
	return -1
}
