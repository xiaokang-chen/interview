package main

import (
	"fmt"
	"interview/code/list"
	"sync"
)

// 交替打印a,b
func PrintAB() {
	wg := sync.WaitGroup{}
	chanA := make(chan string)
	chanB := make(chan string)
	count := 5
	wg.Add(2)
	go func() {
		defer wg.Done()
		for i := 0; i < count; i++ {
			str := <-chanA
			fmt.Println(str)
			chanB <- "B"
		}
	}()
	go func() {
		defer wg.Done()
		for i := 0; i < count; i++ {
			str := <-chanB
			fmt.Println(str)
			// 死锁的坑，最后一个B打印完后，不用再给A发信号了
			if i == count-1 {
				return
			}
			chanA <- "A"
		}
	}()

	// 同步先执行开启
	chanA <- "AT"
	wg.Wait()
}

// 链表判断环
// 快慢指针
func IsCircle(node *list.ListNode) bool {
	slow, fast := node, node
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next

		// 快慢指针相遇，存在环
		if fast == slow {
			return true
		}
	}
	return false
}
