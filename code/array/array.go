package array

import "interview/code/list"

// RemoveDuplicates 26.删除有序数组中的重复项
// 时刻保证[0,slow]
func RemoveDuplicates(nums []int) int {
	if len(nums) == 0 {
		return 0
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
	return slow + 1
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
