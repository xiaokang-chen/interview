package main

import "fmt"

// 数组-线性访问
func traverse(arr []int) {
	for i := 0; i < len(arr); i++ {
		// 迭代访问arr
		fmt.Println("item: ", arr[i])
	}
}

// 二叉树
type TreeNode1 struct {
	val   int
	left  *TreeNode1
	right *TreeNode1
}

// 二叉树-非线性递归
func traverse1(root *TreeNode1) {
	if root != nil {
		traverse1(root.left)
		fmt.Println(root.val)
		traverse1(root.right)
	}
}

// N叉树
type TreeNode2 struct {
	val      int
	children []*TreeNode2
}

// N叉数遍历
func traverse2(root *TreeNode2) {
	fmt.Println(root.val)
	for _, child := range root.children {
		traverse2(child)
	}
}
