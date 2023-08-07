package producer

import (
	"fmt"
	"time"
)

// 1. 定义一个链表
// 2. 实现一个链表反转

type ListNode struct {
	Val  int
	Next *ListNode
}

// Reverse 链表反转
func Reverse(root *ListNode) *ListNode {
	if root == nil || root.Next == nil {
		return root
	}
	var pre *ListNode = nil
	for root != nil {
		// 1. 记录next
		next := root.Next
		// 2. root Next调整到pre
		root.Next = pre
		// 3. pre更新
		pre = root
		// 4. root后移
		root = next
	}
	return pre
}

// Reverse2 递归实现链表反转
// root=2
// 1->2->3->4->5
// 1->2<-3<-4<-5 Reverse2(2) = 5
// 1<-2<-3<-4<-5 Reverse2(1) = 5
func Reverse2(root *ListNode) *ListNode {
	if root == nil {
		return root
	}
	head := Reverse2(root.Next)
	root.Next.Next = root
	root.Next = nil
	return head
}

// ProducerAndConsumer 生产者消费者模式
func ProducerAndConsumer() {
	// 带缓冲的channel
	message := make(chan int, 10)
	// 结束chan
	done := make(chan bool)

	defer close(message)

	// 消费者
	go func() {
		ticker := time.NewTicker(1 * time.Second)
		for range ticker.C {
			select {
			case <-done:
				fmt.Println("child process interrupt...")
				return
			default:
				fmt.Printf("receive message:%d\n", <-message)
			}
		}
	}()

	// 生产者
	for i := 0; i < 10; i++ {
		message <- i
	}

	// 5S后主线程关闭done通道
	time.Sleep(5 * time.Second)
	close(done)
	time.Sleep(1 * time.Second)
	fmt.Println("main process exit!")
}
