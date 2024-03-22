package main

// 一个链表，实现push，pop，len

// DoubleListNode 双链表节点
type DoubleListNode struct {
	Val       int
	pre, next *DoubleListNode
}

// DoubleList 双链表
type DoubleList struct {
	len        int
	head, tail *DoubleListNode
}

func New(arr []int) *DoubleList {
	n := len(arr)
	dl := &DoubleList{len: n}
	if n == 0 {
		return dl
	}
	curNode := &DoubleListNode{
		Val: arr[0],
	}
	dl.head = curNode
	for i := 1; i < n; i++ {
		node := &DoubleListNode{
			Val: arr[i],
		}
		if i == n-1 {
			dl.tail = node
		}
		node.pre = curNode
		curNode.next = node
		curNode = curNode.next
	}
	return dl
}

func (d *DoubleList) Push(v int) {
	node := &DoubleListNode{
		Val:  v,
		pre:  d.tail,
		next: nil,
	}
	d.tail.next = node
	node.pre = d.tail
	d.tail = node
}

// 1,2,3
func (d *DoubleList) Pop() {
	temp := d.tail.pre
	d.tail.pre.next = nil
	d.tail = temp
}

func (d *DoubleList) Len() int {
	return d.len
}
