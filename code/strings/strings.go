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

// 验证回文串
func ValidPalindrome(s string) bool {
	i, j := 0, len(s)-1
	for i < j {
		if !isalnums(s[i]) {
			i++
			continue
		}
		if !isalnums(s[j]) {
			j--
			continue
		}
		if s[i] != s[j] {
			return false
		}
		i++
		j--
	}
	return true
}

func isalnums(char byte) bool {
	return (char >= 'A' && char <= 'Z') ||
		(char >= 'a' && char <= 'z') ||
		(char >= '0' && char <= '9')
}

// IsSubsequence 判断子序列
// 判断s是否为t的子序列
func IsSubsequence(s string, t string) bool {
	i, j := 0, 0
	for i < len(s) && j < len(t) {
		if s[i] == t[j] {
			i++
		}
		j++
	}
	return i == len(s)
}

// // IsSubsequence2 进阶判断子序列
// // 有大量的s
// func IsSubsequence2(s string, t string) bool {
// 	n, m := len(s), len(t)
// 	f := make([][26]int, m)
// 	for i := 0; i < 26; i++ {
// 		f[m][i] = m
// 	}
// 	for i := m - 1; i >= 0; i-- {
// 		for j := 0; j < 26; j++ {
// 			if t[i] == byte(j+'a') {
// 				f[i][j] = i
// 			} else {
// 				f[i][j] = f[i+1][j]
// 			}
// 		}
// 	}
// 	add := 0
// 	for i:= 0; i < n; i++ {

// 	}
// }
