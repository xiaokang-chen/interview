package main

// 题目为leetcode第752题.打开转盘锁
// 该题是典型的路径搜索问题，从 0000 到 target 的距离
// 每个中间态（可以看为树上的一个节点）都有8个子节点：
// 因为每一个步到下一步，都可以对4个拨轮进行2种不同的操作（+1或-1）

// OpenLockWithDFS 通过广度优先遍历开锁
// 起始字符串为 0000
// deadends：死亡字符串的数组
// target：目标字符串
// return：起始字符串到目标字符串最小步数
func OpenLockWithDFS(deadends []string, target string) int {
	// 死亡字符串map
	deads := make(map[string]bool)
	for _, item := range deadends {
		deads[item] = true
	}
	// 已遍历字符串map
	visited := make(map[string]bool)
	// 队列
	q := []string{"0000"}
	// 最小步长
	step := 0

	for len(q) > 0 {
		n := len(q)
		for i := 0; i < n; i++ {
			cur := q[0]
			q = q[1:]
			// 判断字符串是否为死亡字符串
			if deads[cur] {
				continue
			}
			// 如果已经到达最终节点，则返回步长
			if cur == target {
				return step
			}
			// 将子元素添加到队列
			for j := 0; j < 4; j++ {
				up := PlusOne(cur, j)
				if !visited[up] {
					q = append(q, up)
					visited[up] = true
				}
				down := MinusOne(cur, j)
				if !visited[down] {
					q = append(q, down)
					visited[down] = true
				}
			}
		}
		// 层数+1
		step++
	}

	return -1
}

// PlusOne +1
func PlusOne(s string, i int) string {
	c := []byte(s)
	if c[i] == '9' {
		c[i] = '0'
	} else {
		c[i] += 1
	}
	return string(c)
}

// MinusOne -1
func MinusOne(s string, i int) string {
	c := []byte(s)
	if c[i] == '0' {
		c[i] = '9'
	} else {
		c[i] -= 1
	}
	return string(c)
}

// OpenLockWithDoubleDFS 使用双向的广度优先遍历搜索算法解锁
func OpenLockWithDoubleDFS(deadends []string, target string) int {
	// 死亡字符串map
	deads := make(map[string]bool)
	for _, item := range deadends {
		deads[item] = true
	}
	// 已遍历字符串map
	visited := make(map[string]bool)
	// 队列申请两个，因为可以无序遍历，所以使用map
	m1 := make(map[string]bool)
	m2 := make(map[string]bool)
	m1["0000"] = true
	m2[target] = true
	// 最小步长
	step := 0
	for len(m1) > 0 && len(m2) > 0 {
		// hash遍历过程中不能修改，用temp存储扩散结果
		temp := make(map[string]bool)

		// 将m1中所有节点向周围扩散
		for cur := range m1 {
			if deads[cur] {
				continue
			}
			if m2[cur] {
				return step
			}
			visited[cur] = true

			for j := 0; j < 4; j++ {
				up := PlusOne(cur, j)
				if !visited[up] {
					temp[up] = true
				}
				down := MinusOne(cur, j)
				if !visited[down] {
					temp[down] = true
				}
			}
		}
		step++
		// 交换m1和m2，两个队列交替扩展
		m1 = m2
		m2 = temp
	}

	return -1
}
