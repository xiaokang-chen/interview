package main

import (
	"fmt"
	"sync"
)

// 数组-线性访问
func traverse(arr []int) {
	for i := 0; i < len(arr); i++ {
		// 迭代访问arr
		fmt.Println("item: ", arr[i])
	}
}

// 二叉树
type TreeNode1 struct {
	Val   int
	Left  *TreeNode1
	Right *TreeNode1
}

// 二叉树-非线性递归
func traverse1(root *TreeNode1) {
	if root != nil {
		traverse1(root.Left)
		fmt.Println(root.Val)
		traverse1(root.Right)
	}
}

// N叉树
type TreeNode2 struct {
	Val      int
	Children []*TreeNode2
}

// N叉数遍历
func traverse2(root *TreeNode2) {
	fmt.Println(root.Val)
	for _, child := range root.Children {
		traverse2(child)
	}
}

// ==================B+树=======================
// B树称为“多路平衡查找树”，B+树是其变种。B树中所有结点的孩子个数的最大值称为B树的阶，通常用M表示。一般从查找效率考虑，通常要求M>=3。一棵M阶B树，有如下特性：
// 1. 若根节点不是叶子结点，则至少有两棵树
// 2. 每一个节点最多M棵子树，最多有M-1个关键字
// 3. 除根节点外，其他的每个分支至少有ceil(M/2)个子树，至少含有ceil(M/2)-1个关键字
// BPItem B+树数据节点（叶子节点）
type BPItem struct {
	Key int
	Val int
}

// BPNode B+树索引节点（非叶子节点）
type BPNode struct {
	MaxKey int       // 存储子树的最大关键字
	Nodes  []*BPNode // 节点的子树
	Items  []*BPItem // 叶子节点的数据记录
	Pre    *BPNode   // 叶子节点中指向前一叶子节点，实现叶子节点链表
	Next   *BPNode   // 叶子节点中指向下一叶子节点，实现叶子节点链表
}

// BPTree B+树
type BPTree struct {
	mutex sync.Mutex
	root  *BPNode // B+树的根节点
	width int     // B+树的阶
	halfw int     // [M/2]
}

// ==================B+树 END=======================
