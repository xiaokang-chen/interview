package main

// Node 链表节点
type Node struct {
	Val  int
	Next *Node
}

// SortList 排序链表
// 4 2 1 3
func SortList(head *Node) *Node {
	// 归并排序
	var merge func(head1, head2 *Node) *Node
	var sort func(head *Node) *Node

	merge = func(head1, head2 *Node) *Node {
		// 虚拟头节点
		dummy := &Node{-1, nil}
		p, q := head1, head2
		cur := dummy
		for p != nil && q != nil {
			if p.Val < q.Val {
				cur.Next = p
				p = p.Next
			} else {
				cur.Next = q
				q = q.Next
			}
			cur = cur.Next
		}
		// 后续处理
		if p != nil {
			cur.Next = p
		}
		if q != nil {
			cur.Next = q
		}

		return dummy.Next
	}

	sort = func(head *Node) *Node {
		// 1. 边界处理
		if head == nil || head.Next == nil {
			return head
		}
		// 2. 快慢指针，先定位到mid
		slow, fast := head, head.Next
		for fast != nil && fast.Next != nil {
			slow = slow.Next
			fast = fast.Next.Next
		}

		mid := slow
		next := mid.Next
		// 中间断链
		mid.Next = nil
		return merge(sort(head), sort(next))
	}

	return sort(head)
}
