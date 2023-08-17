package list

// ZipListNode 压缩列表节点
type ZipListNode struct {
	PreLen   int         // 记录前一节点的长度，目的是可以从后向前遍历
	Encoding string      // 记录当前节点的类型
	Data     interface{} // 记录当前节点的数据
}

// ZipList 压缩列表
type ZipList struct {
	Val []ZipListNode
}

// ListNode 单链表节点
type ListNode struct {
	Val  int
	Next *ListNode
}

// LinkedListNode 双链表节点
type LinkedListNode struct {
	Val  int
	Pre  *LinkedListNode
	Next *LinkedListNode
}

// QuickListNode 快链表节点
type QuickListNode struct {
	Val  *ZipList // 指向压缩数组的指针
	Pre  *QuickListNode
	Next *QuickListNode
}

// QuickList 快链表
type QuickList struct {
	Head *QuickListNode
	Tail *QuickListNode
	// ...
}

// MergeTwoLists 21.合并两个有序列表
func MergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	// 虚拟头节点
	dummyNode := &ListNode{-1, nil}
	p := dummyNode
	p1, p2 := list1, list2
	for p1 != nil && p2 != nil {
		if p1.Val > p2.Val {
			p.Next = p2
			p2 = p2.Next
		} else {
			p.Next = p1
			p1 = p1.Next
		}
		// 向后移动
		p = p.Next
	}

	if p1 != nil {
		p.Next = p1
	}
	if p2 != nil {
		p.Next = p2
	}
	return dummyNode.Next
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

// ReverseList 链表反转
// 关键点：设置一个前驱节点（pre）
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

// ReverseList2 用递归方式实现反转链表
func ReverseList2(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	newHead := ReverseList2(head.Next)
	head.Next.Next = head
	head.Next = nil
	return newHead
}

// MergeKLists 23.合并k个链表
func MergeKLists(lists []*ListNode) *ListNode {
	resNode := &ListNode{-1, nil}
	for i := 0; i < len(lists); i++ {
		resNode = MergeTwoLists(resNode, lists[i])
	}
	return resNode.Next
}

// MergeKLists2 利用分治合并k个链表
// 类似归并排序思想
func MergeKLists2(lists []*ListNode) *ListNode {
	m := len(lists)
	if m == 0 {
		return nil
	}
	if m == 1 {
		return lists[0]
	}
	left := MergeKLists2(lists[:m/2])  // 合并左半部分
	right := MergeKLists2(lists[m/2:]) // 合并右半部分
	return MergeTwoLists(left, right)
}

// FindFromEnd 单链表倒数第k个节点
// 技巧：双指针-快慢指针
// 倒数第k个节点就是正数第n-k+1个节点，从头节点往后走n-k即可到达
func FindFromEnd(head *ListNode, k int) *ListNode {
	p1, p2 := head, head
	// p1先向后走k步
	for i := 0; i < k; i++ {
		p1 = p1.Next
	}
	// p1再向后走n-k，p2与此同时向后移动，此时p1=nil，p2
	// 这轮巧妙之处在于循环靠p1=nil这个条件
	for p1 != nil {
		p1 = p1.Next
		p2 = p2.Next
	}
	return p2
}

// RemoveNthFromEnd 19.删除链表第N个节点
func RemoveNthFromEnd(head *ListNode, k int) *ListNode {
	dummy := &ListNode{-1, head}
	// 这里给p2赋值dummy有效的解决了边界问题
	p1, p2 := head, dummy
	// p1先向后走k步，不走k+1的原因是有可能越界
	for i := 0; i < k; i++ {
		p1 = p1.Next
	}
	for p1 != nil {
		p1 = p1.Next
		p2 = p2.Next
	}
	// p2为待删除节点的前驱
	p2.Next = p2.Next.Next
	return dummy.Next
}

// MiddleNode 876.链表中点
func MiddleNode(head *ListNode) *ListNode {
	slow, fast := head, head
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}
	return slow
}

// HasCycle 141.判断链表成环
func HasCycle(head *ListNode) bool {
	slow, fast := head, head
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
		if fast == slow {
			return true
		}
	}
	return false
}

// DetectCycle 142.返回链接入环节点
func DetectCycle(head *ListNode) *ListNode {
	slow, fast := head, head
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
		if fast == slow {
			break
		}
	}
	// 如果没有环，需要返回Nil
	if fast != nil && fast.Next != nil {
		return nil
	}
	slow = head
	// 二次相遇时在环入口
	for fast != slow {
		slow = slow.Next
		fast = fast.Next
	}
	return slow
}

// DetectCycle2 哈希表寻找入口环节点
// 循环遍历，每次都将遍历节点插入到map中
func DetectCycle2(head *ListNode) *ListNode {
	seen := map[*ListNode]bool{}
	p := head
	for p != nil {
		// 如果存在于map，则返回
		if _, ok := seen[p]; ok {
			return p
		}
		seen[p] = true
		p = p.Next
	}
	return nil
}

// GetIntersectionNode 160.相交链表
// 技巧在于将两个指针以某种方式，能够同时到达相交点
// 可以通过逻辑上连接两个链表，来让两个指针分别遍历的链表长度相同，以达到最后都到达相交点的目的
func GetIntersectionNode(headA, headB *ListNode) *ListNode {
	p1, p2 := headA, headB
	for p1 != p2 {
		// p1走一步，如果走到A末尾，转到B
		if p1 == nil {
			p1 = headB
		} else {
			p1 = p1.Next
		}
		// p2走一步，如果走到B末尾，转到A
		if p2 == nil {
			p2 = headA
		} else {
			p2 = p2.Next
		}
	}
	return p1
}

// SortList 148.排序链表
func SortList(head *ListNode) *ListNode {
	// 归并排序
	var merge func(head1, head2 *ListNode) *ListNode
	var sort func(head *ListNode) *ListNode

	merge = func(head1, head2 *ListNode) *ListNode {
		// 虚拟头节点
		dummy := &ListNode{-1, nil}
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

	sort = func(head *ListNode) *ListNode {
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
