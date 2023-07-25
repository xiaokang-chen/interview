package heap

import (
	"container/heap"
	"interview/code/list"
)

type PriorityQueue []*list.ListNode

func (pq PriorityQueue) Less(i, j int) bool  { return pq[i].Val < pq[j].Val }
func (pq PriorityQueue) Swap(i, j int)       { pq[i], pq[j] = pq[j], pq[i] }
func (pq PriorityQueue) Len() int            { return len(pq) }
func (pq *PriorityQueue) Push(x interface{}) { *pq = append(*pq, x.(*list.ListNode)) }
func (pq *PriorityQueue) Pop() interface{} {
	n := len(*pq)
	ans := (*pq)[n-1]
	*pq = (*pq)[:n-1]
	return ans
}

// mergeKLists 合并k个元素
func MergeKLists(lists []*list.ListNode) *list.ListNode {
	if len(lists) == 0 {
		return nil
	}
	// 虚拟头节点
	dummy := &list.ListNode{Val: -1}
	p := dummy
	// 优点级队列，最小堆
	pq := make(PriorityQueue, 0)
	heap.Init(&pq)
	// 将k个链表的头节点加入最小堆
	for _, head := range lists {
		if head != nil {
			heap.Push(&pq, head)
		}
	}
	// 每次取出小节点值并给到p，再push到优先级队列，重新构建小顶堆
	for pq.Len() > 0 {
		// 获取最小节点，接到结果链表中
		node := heap.Pop(&pq).(*list.ListNode)
		p.Next = node
		if node.Next != nil {
			heap.Push(&pq, node.Next)
		}
		// p指针不断前进
		p = p.Next
	}
	return dummy.Next
}
