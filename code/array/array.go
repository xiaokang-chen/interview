package array

import (
	"interview/code/list"
	"math"
)

// RemoveDuplicates 26.删除有序数组中的重复项
// 时刻保证[0,slow]
func RemoveDuplicates(nums []int) []int {
	if len(nums) == 0 {
		return []int{}
	}
	slow, fast := 0, 0
	for fast < len(nums) {
		if nums[fast] != nums[slow] {
			slow++
			// 维护[0, slow]都是去重后的
			nums[slow] = nums[fast]
		}
		fast++
	}
	// 数组长度为索引+1
	return nums[:slow+1]
}

// 83.删除排序链表中的重复元素
// 和数组的写法完全一样
func DeleteDuplicates(head *list.ListNode) *list.ListNode {
	if head == nil {
		return nil
	}
	slow, fast := head, head
	for fast != nil {
		if fast.Val != slow.Val {
			slow.Next = fast
			slow = slow.Next
		}
		fast = fast.Next
	}
	// 末尾相同的处理
	slow.Next = nil
	return head
}

// RemoveElement 27.移除元素
// 每次调整后保持[0, slow-1]是不包含val值的
func RemoveElement(nums []int, val int) int {
	slow, fast := 0, 0
	for fast < len(nums) {
		if nums[fast] != val {
			nums[slow] = nums[fast]
			slow += 1
		}
		fast++
	}
	return slow
}

// MoveZeroes 283.移动零
// 将所有的0移动到数组后面
// 我们将所有 0 移到最后，其实就相当于移除 nums 中的所有 0，然后再把后面的元素都赋值为 0 即可
func MoveZeroes(nums []int) {
	// p为移除0后的数组长度，需要往后再插入0
	p := RemoveElement(nums, 0)
	for p < len(nums) {
		nums[p] = 0
		p++
	}
}

// MoveZeroes2 可以参考快速排序
func MoveZeroes2(nums []int) {
	if len(nums) == 0 {
		return
	}
	i, j := 0, 0
	// 维持[0, j-1]是非0数组
	for i < len(nums) {
		if nums[i] != 0 {
			nums[i], nums[j] = nums[j], nums[i]
			j++
		}
		i++
	}
}

// TwoSum 167.两数之和II
func TwoSum(numbers []int, target int) []int {
	left, right := 0, len(numbers)-1
	for left < right {
		sum := numbers[left] + numbers[right]
		if sum == target {
			return []int{left + 1, right + 1}
		} else if sum > target {
			right--
		} else if sum < target {
			left++
		}
	}
	return []int{-1, -1}
}

// ReverseString 反转字符串
func ReverseString(s []byte) {
	left, right := 0, len(s)-1
	for left < right {
		s[left], s[right] = s[right], s[left]
		left++
		right--
	}
}

// IsPalindrome 判断是否是回文串
func IsPalindrome(s string) bool {
	left, right := 0, len(s)-1
	for left < right {
		if s[left] != s[right] {
			return false
		}
		left--
		right++
	}
	return true
}

// LongestPalindrome 5.最长回文子串
// 找回文串的难点在于，回文串的的长度可能是奇数也可能是偶数，
// 解决该问题的核心是从中心向两端扩散的双指针技巧
func LongestPalindrome(s string) string {
	res := ""
	for i := 0; i < len(s); i++ {
		// 以s[i]为中心的最长回文子串
		s1 := getLongestPalindrome(s, i, i)
		// 以s[i]和s[i+1]为中心的最长回文子串
		s2 := getLongestPalindrome(s, i, i+1)
		if len(res) < len(s1) {
			res = s1
		}
		if len(res) < len(s2) {
			res = s2
		}
	}
	return res
}

// getLongestPalindrome 返回以s[l]和s[r]为中心的最长回文子串
// 1. 如果输入相同的 l 和 r，就相当于寻找长度为奇数的回文串
// 2. 如果输入相邻的 l 和 r，则相当于寻找长度为偶数的回文串
func getLongestPalindrome(s string, l int, r int) string {
	// 从中心向外展开
	for l >= 0 && r < len(s) && s[l] == s[r] {
		l--
		r++
	}
	// 边界条件问题：因为最后一次相同后，l和r又更新了，此时[l+1, r-1]是目标回文
	return s[l+1 : r]
}

// // MaxProfit 121.买卖股票的最佳时机（仅能买卖一次）
// func MaxProfit(prices []int) int {
// 	n := len(prices)
// 	dp0, dp1 := 0, math.MinInt
// 	for i := 0; i < n; i++ {
// 		// 今天卖了股票：昨天买的股票，获取收益
// 		dp0 = max(dp0, dp1+prices[i])
// 		// 今天买了股票：因为只能买一次，之前肯定没买过，所以利润肯定是-今天价格
// 		dp1 = max(dp1, -prices[i])
// 	}
// 	return dp0
// }

// MaxProfit 121.买卖股票的最佳时机（仅能买卖一次）
func MaxProfit(prices []int) int {
	buy, profit := math.MaxInt32, 0
	for _, item := range prices {
		buy = min(buy, item)
		profit = max(profit, item-buy)
	}
	return profit
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// // MaxProfit2 122.买卖股票的最佳时机II（可以多次买卖）
// func MaxProfit2(prices []int) int {
// 	n := len(prices)
// 	dp0, dp1 := 0, math.MinInt32
// 	for i := 0; i < n; i++ {
// 		temp := dp0
// 		dp0 = max(dp0, dp1+prices[i])
// 		dp1 = max(dp1, temp-prices[i])
// 	}
// 	return dp0
// }

// MaxProfit2 122.买卖股票的最佳时机II（可以多次买卖）
// 7,1,5,3,6,4
func MaxProfit2(prices []int) int {
	buy, profit := prices[0], 0
	for i := 1; i < len(prices); i++ {
		if prices[i] < prices[i-1] {
			profit += prices[i-1] - buy
			buy = prices[i]
		}
	}
	if prices[len(prices)-1] > buy {
		profit += prices[len(prices)-1] - buy
	}
	return profit
}

// Rotate 48.旋转图像
func Rotate(matrix [][]int) {
	n := len(matrix)
	// 先对矩阵每一行进行反转，只需要反转一半数量
	for i := 0; i < n/2; i++ {
		matrix[i], matrix[n-1-i] = matrix[n-1-i], matrix[i]
	}
	// 再沿对角线镜像对称二维矩阵
	for i := 0; i < n; i++ {
		for j := i; j < n; j++ {
			matrix[i][j], matrix[j][i] = matrix[j][i], matrix[i][j]
		}
	}
}

// Rotate2 旋转图像
func Rotate2(matrix [][]int) {
	n := len(matrix)
	// 先沿对角线反转元素
	for i := 0; i < n; i++ {
		for j := i; j < n; j++ {
			matrix[i][j], matrix[j][i] = matrix[j][i], matrix[i][j]
		}
	}
	// 再反转每一行
	for _, row := range matrix {
		m := len(row)
		for i := 0; i < m/2; i++ {
			row[i], row[m-1-i] = row[m-1-i], row[i]
		}
	}
}

// SpiralOrder 螺旋矩阵
func SpiralOrder(matrix [][]int) []int {
	// m行n列
	m, n := len(matrix), len(matrix[0])
	top, bottom := 0, m-1
	left, right := 0, n-1
	res := make([]int, 0, m*n)
	for len(res) < m*n {
		// 1. 在上侧：从左向右
		// 确保上下空间
		if top <= right {
			for i := left; i <= right; i++ {
				res = append(res, matrix[top][i])
			}
			// 上边界下移
			top++
		}
		// 2. 在右侧：从上向下
		// 确保左右空间
		if left <= right {
			for i := top; i <= bottom; i++ {
				res = append(res, matrix[i][right])
			}
			// 右边界左移
			right--
		}
		// 3. 在下侧：从右向左
		// 确保上下空间
		if top <= right {
			for i := right; i >= left; i-- {
				res = append(res, matrix[bottom][i])
			}
			// 下边界上移
			bottom--
		}
		// 4. 在左侧：从下向上
		// 确保左右空间
		if left <= right {
			for i := bottom; i >= top; i-- {
				res = append(res, matrix[i][left])
			}
			// 左边界右移
			left++
		}
	}
	return res
}
