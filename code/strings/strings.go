package strings

import "strings"

// LengthOfLongestSubstring 无重复字符的最长子串
// 利用滑动窗口
func LengthOfLongestSubstring(s string) int {
	window := make(map[byte]int)
	left, right := 0, 0
	maxLen := 0
	for right < len(s) {
		c := s[right]
		window[c]++
		right++
		// 滑动窗口
		for window[c] > 1 {
			d := s[left]
			window[d]--
			left++
		}
		maxLen = max(maxLen, right-left)
	}
	return maxLen
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// ReverseWords 反转字符串中的单词
func ReverseWords(s string) string {
	// 按照空格切割单词
	t := strings.Fields(s)
	// 遍历前半段
	for i := 0; i < len(t)/2; i++ {
		t[i], t[len(t)-1-i] = t[len(t)-1-i], t[i]
	}
	return strings.Join(t, " ")
}
