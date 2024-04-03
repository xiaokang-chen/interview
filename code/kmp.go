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
	j := -1
	next := getNext(pattern)
	for i := 0; i < len(text); i++ {
		for j > -1 && text[i] != pattern[j+1] {
			j = next[j] // 向前回溯
		}
		if text[i] == pattern[j+1] { // 匹配则j往后
			j++
		}
		if j == len(pattern)-1 { // 全匹配则返回匹配的起点索引
			return i - len(pattern) + 1
		}
	}
	return -1
}

// getNext 获取最长相同前后缀
// next保存的是最长前缀（最后一个字母）的下标
func getNext(pattern string) []int {
	next := make([]int, len(pattern))
	next[0] = -1
	j := -1
	for i := 1; i < len(pattern); i++ {
		for j > -1 && pattern[i] != pattern[j+1] { // 不同则j向前回退
			j = next[j]
		}
		if pattern[i] == pattern[j+1] { // 相同就继续
			j++
		}
		next[i] = j // 将前缀长度赋给next[j]
	}
	return next
}
