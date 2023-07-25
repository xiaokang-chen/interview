package main

import (
	"fmt"
	"interview/code/algorithm"
)

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

	// // N叉树非线性遍历
	// root2 := &TreeNode2{
	// 	val: 1,
	// 	children: []*TreeNode2{
	// 		&TreeNode2{
	// 			val: 2,
	// 		},
	// 		&TreeNode2{
	// 			val: 3,
	// 		},
	// 		&TreeNode2{
	// 			val: 4,
	// 		},
	// 	},
	// }
	// traverse2(root2)

	// list1 := &list.ListNode{
	// 	Val: 1,
	// 	Next: &list.ListNode{
	// 		Val: 2,
	// 		Next: &list.ListNode{
	// 			Val: 4,
	// 		},
	// 	},
	// }
	// list2 := &list.ListNode{
	// 	Val: 1,
	// 	Next: &list.ListNode{
	// 		Val: 3,
	// 		Next: &list.ListNode{
	// 			Val: 4,
	// 		},
	// 	},
	// }
	// node := list.MergeTwoLists(list1, list2)
	// fmt.Println(node)

	// head := &list.ListNode{
	// 	Val: 1,
	// 	Next: &list.ListNode{
	// 		Val: 4,
	// 		Next: &list.ListNode{
	// 			Val: 3,
	// 			Next: &list.ListNode{
	// 				Val: 2,
	// 				Next: &list.ListNode{
	// 					Val: 5,
	// 					Next: &list.ListNode{
	// 						Val: 2,
	// 					},
	// 				},
	// 			},
	// 		},
	// 	},
	// }
	// // resHead := list.Partition(head, 3)
	// resHead := list.ReverseList(head)
	// fmt.Println("res: ", resHead)

	// 排序
	arr := []int{8, 4, 5, 7, 1, 3, 6, 2}

	res := algorithm.QuickSort(arr, 0, len(arr)-1)
	fmt.Println("res: ", res)

	// // 查找
	// res1 := algorithm.BinarySearch(arr, 38)
	// fmt.Println("res1: ", res1)

	// p1 := &list.ListNode{
	// 	Val: 1,
	// 	Next: &list.ListNode{
	// 		Val: 4,
	// 		Next: &list.ListNode{
	// 			Val: 5,
	// 		},
	// 	},
	// }
	// p2 := &list.ListNode{
	// 	Val: 1,
	// 	Next: &list.ListNode{
	// 		Val: 3,
	// 		Next: &list.ListNode{
	// 			Val: 4,
	// 		},
	// 	},
	// }
	// p3 := &list.ListNode{
	// 	Val: 2,
	// 	Next: &list.ListNode{
	// 		Val: 6,
	// 	},
	// }
	// lists := []*list.ListNode{
	// 	p1, p2, p3,
	// }
	// // res := list.MergeKLists(lists)
	// res := heap.MergeKLists(lists)
	// fmt.Println("res1: ", res)
}
