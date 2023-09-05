package main

// node 树节点
type node struct {
	data        int
	left, right *node
}

// getBlackLen 获取两个黑色节点中间的距离
func getBlackLen(root *node) int {
	var getLength func(*node) int
	var getParent func(root *node) *node
	// 1. 先获取两个黑色节点的公共祖先节点
	// 2. 获取的同时得到黑色节点到该公共祖先节点的距离
	// 3. 两个距离相加

	// getLength 节点到1的距离
	getLength = func(node *node) int {
		// 如果遇到黑色叶子节点
		if node.data == 1 {
			return 0
		}
		if node.data == 0 && node.left == nil && node.right == nil {
			return -1
		}
		leftLen := getLength(node.left) + 1
		rightLen := getLength(node.right) + 1
		return max(leftLen, rightLen)
	}

	// 获取两个值为1的黑色节点的公共节点
	getParent = func(root *node) *node {
		// 如果找到1
		if root == nil || root.data == 1 {
			return root
		}
		left := getParent(root.left)
		right := getParent(root.right)
		// 如果左右为1，则当前节点为最近祖先节点
		if left == nil {
			return right
		}
		if right == nil {
			return left
		}
		return root
	}
	parentNode := getParent(root)
	return getLength(parentNode.left) + getLength(parentNode.right) + 2
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
