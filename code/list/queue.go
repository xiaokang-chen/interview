package list

// PriorityQueue 优先级队列
type PriorityQueue []*ListNode

func (pq PriorityQueue) Len() int {
	return len(pq)
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}

func (pq *PriorityQueue) Push(x interface{}) {
	node := x.(*ListNode)
	*pq = append(*pq, node)
}

func (pq *PriorityQueue) Pop(index int) interface{} {
	old := *pq
	node := old[index]
	for i := index; i < len(old)-1; i++ {
		old[i] = old[i+1]
	}
	return node
}

// MergeKLists 23.合并K个升序链表
func MergeKLists(lists []*ListNode) *ListNode {
	if len(lists) == 0 {
		return nil
	}
	// 虚拟头节点
	dummy := &ListNode{-1, nil}
	p := dummy
	// 优先级队列构造小顶堆
	pq := make(PriorityQueue, 0)
	// 将k个链表头节点加入小顶堆
	for _, head := range lists {
		if head != nil {
			pq.Push(head)
		}
	}

	for pq.Len() > 0 {
		minIndex := 0
		// 获取最小节点，接到结果链表中
		for i := 1; i < pq.Len(); i++ {
			if pq[i].Val < pq[minIndex].Val {
				minIndex = i
			}
		}
		minNode := pq.Pop(minIndex).(*ListNode)
		p.Next = minNode
		if minNode.Next != nil {
			pq.Push(minNode.Next)
		}
		// p指针向前
		p = p.Next
	}
	return dummy.Next
}
