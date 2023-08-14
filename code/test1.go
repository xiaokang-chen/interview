package main

// ListNode 链表节点
type ListNode struct {
	Val  int
	Next *ListNode
}

// ListReverse 链表反转
// node(3)
// 1->2->3->4`->`5 ListReverse(4)=5
// 1->2->3`<-`4<-5
func ListReverse(root *ListNode) *ListNode {
	// 临界条件
	if root == nil {
		return nil
	}
	newHead := ListReverse(root.Next)
	root.Next.Next = root
	root.Next = nil
	return newHead
}

// LRUCache 最近未使用缓存
type LRUCache struct {
	cap        int // 缓存容量
	size       int // 链表节点数量
	cache      map[int]*LinkedListNode
	head, tail *LinkedListNode
}

// LinkedListNode 链表节点
type LinkedListNode struct {
	pre, next *LinkedListNode
	key, val  int
}

// // InsertVal 插入数据到缓存
// func (l *LRUCache)InsertVal(key, val int) {
// 	// 1. 先判断容量是否满了
// 	if l.cap == l.lruList.len {
// 		// 1.1 如果满了，需要将最近未使用的删除
// 	} else {
// 		// 1.2 如果没满，直接插入
// 		l.lruList
// 	}
// }

// // ReadVal 读取数据
// func ReadVal(key int) int {

// }

var domainList = []string{"baidu.com", "ishumei.com", "fp.ishumei.com"}

// match
// www.ishumei.com => ishumei.com
// log.fp.ishumei.com => fp.ishumei.com
func Match(s string) map[string]int {
	matchMap := make(map[string]int) // str:macth_len
	for _, item := range domainList {
		i := 0
		for j := 0; i < len(item) && j < len(s); {
			// 如果匹配，继续向后找
			if item[i] == s[j] {
				i++
			}
			j++
		}
		// 如果i最后走到尽头，则更新全局长度
		if i == len(item) {
			matchMap[item] = i
		}
	}
	// 查出map中最长的
	return matchMap
}
