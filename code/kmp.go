package main

// bfSearch 暴力解法
func bfSearch(text, pattern string) int {
	n := len(text)
	m := len(pattern)
	for i := 0; i <= n-m; i++ {
		var j int
		for j = 0; j < m; j++ {
			if text[i+j] != pattern[j] {
				break
			}
		}
		if j == m {
			return i
		}
	}
	return -1
}

// 暴力解法会导致模式串每次都需要回溯到首个位置重新进行比较
// kmp算法解决的是让模式串“尽可能不去回溯”，这就需要建立一个
// 记录字符串前缀的next数组
// ===============================================
// kmp算法本身也是一个dp问题：
// KMP 算法永不回退txt的指针i，不走回头路（不会重复扫描txt），而是借助dp数组中储存的信息把pat移到正确的位置继续匹配
func KmpSearch(text, pattern string) int {
	// n := len(text)
	// m := len(pattern)
	// for i := 0; i < n; i++ {
	// 	j = dp[j][string(text[i])]
	// 	if j == m {
	// 		return i - m + 1
	// 	}
	// }
	return -1
}
