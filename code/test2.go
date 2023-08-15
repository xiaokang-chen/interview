package main

// TreeNode 树节点
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// FindTreeMax 找出二叉树中最大值
func FindTreeMax(root *TreeNode) int {
	maxValue := 0
	if root != nil {
		return 0
	}
	var dfs func(node *TreeNode)
	dfs = func(node *TreeNode) {
		if node == nil {
			return
		}
		dfs(node.Left)
		dfs(node.Right)
		if node.Val > maxValue {
			maxValue = node.Val
		}
	}
	dfs(root)
	return maxValue
}

// // 1
// // [1,2,3,4]
// // [6,7,8,9]

// // 2
// // true true
// // true false

// // 3
// // nil
// // "b"
// // "c"

// // 4
// // [1,4,2,3]
// func MaxProfit(arr []int) int {
// 	if len(arr) < 1 {
// 		return 0
// 	}
// 	left, right := 0, 0
// 	maxProfit := 0
// 	for right < len(arr) {
// 		// 向右移动，直到遇到下跌点
// 		for arr[left] < arr[right] && right < len(arr) {
// 			right++
// 		}
// 		if left != right {
// 			maxProfit += arr[right-1] - arr[left]
// 			left = right
// 		} else {
// 			right++
// 		}
// 	}
// 	return maxProfit
// }

// // 5
// // [5,1,2,3,4]
// // [4,5,1,2,3]
// func CycleMoveNum(arr []int) int {
// 	k := 0
// 	for i := 1; i < len(arr); i++ {
// 		if arr[i] < arr[i-1] {
// 			k = i
// 			break
// 		}
// 	}
// 	return k
// }
