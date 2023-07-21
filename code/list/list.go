package list

// ListNode 单链表
type ListNode struct {
	Val  int
	Next *ListNode
}

// MergeTwoLists 21.合并两个有序列表
func MergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	// 虚拟头节点
	resNode := &ListNode{-1, nil}
	p := resNode
	p1, p2 := list1, list2
	for p1 != nil && p2 != nil {
		if p1.Val > p2.Val {
			p.Next = p2
			p2 = p2.Next
		} else {
			p.Next = p1
			p1 = p1.Next
		}
		// p指针不断前进
		p = p.Next
	}

	if p1 != nil {
		p.Next = p1
	}
	if p2 != nil {
		p.Next = p2
	}
	return resNode.Next
}

// MergeTwoLists2 递归
// 返回值小的节点，并每次设置值小的节点的Next
func MergeTwoLists2(list1 *ListNode, list2 *ListNode) *ListNode {
	if list1 == nil {
		return list2
	}
	if list2 == nil {
		return list1
	}
	// 如果List1小
	// 1. 连接list1
	// 返回值小的节点-list1
	if list1.Val < list2.Val {
		list1.Next = MergeTwoLists2(list1.Next, list2)
		return list1
	} else {
		list2.Next = MergeTwoLists2(list1, list2.Next)
		return list2
	}
}

// Partition 86.分隔链表
// 问题简化：需要获取2个链表，一个链表元素全小于x，另一个链表元素全大于x
func Partition(head *ListNode, x int) *ListNode {
	// 存放小于x的节点
	dummy1 := &ListNode{-1, nil}
	// 存放大于x的节点
	dummy2 := &ListNode{-1, nil}
	// p1，p2指针负责生成结果链表
	p1, p2 := dummy1, dummy2
	// p负责遍历原链表
	p := head
	for p != nil {
		if p.Val < x {
			p1.Next = p
			p1 = p1.Next
		} else {
			p2.Next = p
			p2 = p2.Next
		}
		temp := p.Next
		p.Next = nil
		p = temp
	}
	// 连接两个链表
	p1.Next = dummy2.Next
	return dummy1.Next
}

// MergeKLists 合并K个升序链表
func MergeKLists(lists []*ListNode) *ListNode {
	return nil
}

// ReverseList 链表反转
func ReverseList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	var (
		pre *ListNode = nil
		cur           = head
	)
	for cur != nil {
		next := cur.Next
		cur.Next = pre
		pre = cur
		cur = next
	}
	return pre
}
