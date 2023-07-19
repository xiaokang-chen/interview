package main

func main() {
	// 数组线性遍历
	// var arr = []int{1, 2, 3, 4}
	// traverse(arr)

	// 二叉树非线性遍历
	// root1 := &TreeNode1{
	// 	val: 1,
	// 	left: &TreeNode1{
	// 		val: 2,
	// 	},
	// 	right: &TreeNode1{
	// 		val: 3,
	// 	},
	// }
	// traverse1(root1)

	// N叉树非线性遍历
	root2 := &TreeNode2{
		val: 1,
		children: []*TreeNode2{
			&TreeNode2{
				val: 2,
			},
			&TreeNode2{
				val: 3,
			},
			&TreeNode2{
				val: 4,
			},
		},
	}
	traverse2(root2)
}
