package tree

import (
	"fmt"
)

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

// MinDepth 二叉树的最小深度
func MinDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	// 初始化一个队列
	q := []*TreeNode{root}
	depth := 1
	for len(q) > 0 {
		for i := 0; i < len(q); i++ {
			cur := q[0]
			q = q[1:]
			// 判断是否到达终点
			if cur.Left == nil && cur.Right == nil {
				return depth
			}
			if cur.Left != nil {
				q = append(q, cur.Left)
			}
			if cur.Right != nil {
				q = append(q, cur.Right)
			}
		}
		depth++
	}
	return depth
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

// WidthOfBinaryTree 二叉树最大宽度
// BFS 广度优先遍历
// 一个编号为index的节点的左子节点编号为[2*index]，右子节点编号为[2*index+1]
func WidthOfBinaryTree(root *TreeNode) int {
	if root == nil {
		return 0
	}
	// 使用双端队列
	type Pair struct {
		node  *TreeNode
		index int
	}
	res := 1
	q := []Pair{{root, 1}}
	for q != nil {
		width := q[len(q)-1].index - q[0].index + 1
		res = max(res, width)
		temp := q
		q = nil
		// 遍历每一层
		for _, p := range temp {
			if p.node.Left != nil {
				q = append(q, Pair{p.node.Left, p.index * 2})
			}
			if p.node.Right != nil {
				q = append(q, Pair{p.node.Right, p.index*2 + 1})
			}
		}
	}
	return res
}

// WidthOfBinaryTree2 深度优先遍历
// DFS 深度优先遍历
// 一层一层的遍历，用两个变量left和right去分别达到每一层的最左边的节点和最右边的节点
// 然后每一层计算right-left+1，进而更新全局maxWidth
func WidthOfBinaryTree2(root *TreeNode) int {
	levelMin := map[int]int{}
	// 每层的宽度
	var dfs func(*TreeNode, int, int) int
	dfs = func(node *TreeNode, depth, index int) int {
		if node == nil {
			return 0
		}
		if _, ok := levelMin[depth]; !ok {
			// 每一层最先访问的节点是最左边的节点
			levelMin[depth] = index
		}
		return max(index-levelMin[depth]+1, max(dfs(node.Left, depth+1, index*2), dfs(node.Right, depth+1, index*2+1)))
	}
	return dfs(root, 1, 1)
}

// 树的广度优先遍历（层序遍历）
// 需要借助队列实现
func BFS(node *TreeNode) []int {
	res := []int{}
	if node == nil {
		return res
	}
	queue := []*TreeNode{node}
	for len(queue) > 0 {
		cur := queue[0]
		res = append(res, cur.Val)
		queue = queue[1:]
		if cur.Left != nil {
			queue = append(queue, cur.Left)
		}
		if cur.Right != nil {
			queue = append(queue, cur.Right)
		}
	}
	return res
}

// 广度-递归
func BFSWithRecursion(node *TreeNode) []int {
	res := []int{}
	if node == nil {
		return res
	}
	slice := []*TreeNode{node}
	var levelOrder func([]*TreeNode)
	levelOrder = func(nodeSlice []*TreeNode) {
		if len(nodeSlice) == 0 {
			return
		}
		var nextNodeSlice []*TreeNode
		for i := 0; i < len(nodeSlice); i++ {
			cur := nodeSlice[i]
			res = append(res, cur.Val)

			// 当前node左子节点append到下一层node slice
			if cur.Left != nil {
				nextNodeSlice = append(nextNodeSlice, cur.Left)
			}
			// 当前node右子节点append到下一层node slice
			if cur.Right != nil {
				nextNodeSlice = append(nextNodeSlice, cur.Right)
			}
		}
		levelOrder(nextNodeSlice)
	}

	levelOrder(slice)
	return res
}

// 树的深度优先遍历（前序遍历）
// 需要借助栈实现
func DFS(node *TreeNode) []int {
	res := []int{}
	stack := []*TreeNode{}
	cur := node
	for cur != nil || len(stack) > 0 {
		for cur != nil {
			res = append(res, cur.Val)
			stack = append(stack, cur)
			cur = cur.Left
		}
		// 退出时，说明左边已经遍历完
		// 此时需要弹出栈顶元素，并查看它的右节点
		cur = stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		cur = cur.Right
	}
	return res
}

// 深度-递归
func DFSWithRecursion(node *TreeNode) []int {
	res := []int{}
	if node == nil {
		return res
	}
	res = append(res, node.Val)
	res = append(res, DFSWithRecursion(node.Left)...)
	res = append(res, DFSWithRecursion(node.Right)...)
	return res
}

// LongestUnivaluePath 687.最长同值路径
func LongestUnivaluePath(root *TreeNode) int {
	var ans int
	var dfs func(*TreeNode) int
	// 以当前node节点为中心的同值路径最大长度
	dfs = func(node *TreeNode) int {
		// 1. 边界条件
		if node == nil {
			return 0
		}
		left := dfs(root.Left)
		right := dfs(root.Right)
		leftLen, rightLen := 0, 0
		// 2. 获取左最大和右最大，然后当前为左+右
		if node.Left != nil && node.Left.Val == node.Val {
			leftLen = left + 1
		}
		if node.Right != nil && node.Right.Val != node.Val {
			rightLen = right + 1
		}
		// 3. 更新全局
		ans = max(ans, leftLen+rightLen)
		return max(leftLen, rightLen)
	}
	ans = dfs(root)
	return ans
}

// LowestCommonAncestor 二叉搜索树的最近公共祖先
func LowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	// 1. 临界条件：节点为空或者节点本身为p或q
	// 深度遍历先遍历到的p或q，则先遍历到的节点一定是另一个节点最近祖先
	if root == nil || root.Val == p.Val || root.Val == q.Val {
		return root
	}
	// 2. 分别在左右子树中寻找，如果左边没有“祖先”，则返回右边，反之亦然
	// 如果两边都不为空，则代表两边都找到p和q，则当前节点为最近祖先
	left := LowestCommonAncestor(root.Left, p, q)
	right := LowestCommonAncestor(root.Right, p, q)
	if left == nil {
		return right
	}
	if right == nil {
		return left
	}
	return root
}
