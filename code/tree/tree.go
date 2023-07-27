package tree

import "fmt"

// 二叉树
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// MaxDepth 104.获取二叉树的最大深度
// 深度优先（后续遍历）
func MaxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	// 递归的计算左右子树的最大深度
	leftMax := MaxDepth(root.Left)
	rightMax := MaxDepth(root.Right)

	return max(leftMax, rightMax) + 1
}

// max 辅助判断int大小的函数
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

var (
	res   int // 最大深度
	depth int // 遍历到节点的深度
)

// MaxDepth2 广度优先遍历（BFS）
func MaxDepth2(root *TreeNode) int {
	traverse(root)
	return res
}

func traverse(root *TreeNode) {
	if root == nil {
		return
	}
	// 前序位置
	depth++
	if root.Left == nil && root.Right == nil {
		// 到达叶子节点
		res = max(res, depth)
	}
	traverse(root.Left)
	traverse(root.Right)
	// 后续位置
	depth--
}

// Traverse1 打印树的每一个节点所在层数
// 前序遍历
func Traverse1(root *TreeNode, level int) {
	if root == nil {
		return
	}
	// 前序位置
	fmt.Printf("节点 %d 在第 %d 层\n", root.Val, level)
	Traverse1(root.Left, level+1)
	Traverse1(root.Right, level+1)
}

// ChildrenTreeCount 打印树的左右子树各有多少节点
// 后续遍历
func ChildrenTreeCount(root *TreeNode) int {
	if root == nil {
		return 0
	}
	leftCount := ChildrenTreeCount(root.Left)
	rightCount := ChildrenTreeCount(root.Right)
	// 后续位置
	fmt.Printf("节点 %d 的左子树有 %d 个节点，右子树有 %d 个节点\n",
		root.Val, leftCount, rightCount)
	return leftCount + rightCount + 1
}

// DiameterOfBinaryTree 543.二叉树的直径
// 【直径】长度 = 左右子树深度之和
func DiameterOfBinaryTree(root *TreeNode) int {
	var maxLen int
	var dfs func(*TreeNode) int
	dfs = func(node *TreeNode) int {
		// 叶子节点后面的nil为-1，到叶子节点刚好为0
		if node == nil {
			return -1
		}
		leftLen := dfs(node.Left) + 1
		rightLen := dfs(node.Right) + 1
		// 更新全局最大长度
		maxLen = max(maxLen, leftLen+rightLen)
		// 返回当前子树最大链长
		return max(leftLen, rightLen)
	}
	dfs(root)
	return maxLen
}

// DiameterOfBinaryTree2 通过遍历查出二叉树直径
// 通过前序遍历
// 缺陷：前序位置无法获取子树信息，所以只能让每个节点调用 getMaxDepth 函数去算子树的深度。
func DiameterOfBinaryTree2(root *TreeNode) int {
	// 记录最大直径
	maxDepth := 0
	var (
		traverse    func(*TreeNode)
		getMaxDepth func(*TreeNode) int
	)
	// 1.遍历二叉树
	traverse = func(node *TreeNode) {
		if node == nil {
			return
		}
		leftMax := getMaxDepth(node)
		rightMax := getMaxDepth(node)
		maxDepth = max(maxDepth, leftMax+rightMax)
		traverse(node.Left)
		traverse(node.Right)
	}
	// 2.计算二叉树的最大深度
	getMaxDepth = func(node *TreeNode) int {
		if node == nil {
			return -1
		}
		leftMax := getMaxDepth(node.Left)
		rightMax := getMaxDepth(node.Right)
		return 1 + max(leftMax, rightMax)
	}
	// 对每个节点计算直径，求出最大的
	traverse(root)
	return maxDepth
}
