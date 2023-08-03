package main

import (
	"fmt"
	"interview/code/list"
	"runtime"
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

// 线程池
type GoPool struct {
	wg  sync.WaitGroup
	arr []chan byte
}

// 交替打印26个字母
func PrintLetter() {
	count := 26
	wg := sync.WaitGroup{}
	// 切片，len为100
	gp := &GoPool{wg, make([]chan byte, 100)}

	for i := 0; i < count; i++ {
		gp.arr[i] = make(chan byte)
		// 最后一个chan初始化，作为启动，等待被读
		if i == count-1 {
			go func(i int) {
				gp.arr[i] <- 'A'
			}(i)
		}
	}
	fmt.Printf("协程数量：%d\n", runtime.NumGoroutine())
	gp.wg.Add(count)
	for i := 0; i < count; i++ {
		lastChan := make(chan byte)
		curChan := make(chan byte)
		if i == 0 {
			lastChan = gp.arr[count-1]
		} else {
			lastChan = gp.arr[i-1]
		}
		go func(lastChan, curChan chan byte) {
			defer gp.wg.Done()
			for i := 0; i < 5; i++ {
				s := <-lastChan
				fmt.Printf("%c\n", s)
				curChan <- 'A' + byte(i)
			}
		}(lastChan, curChan)
		fmt.Printf("%d, 协程数量：%d\n", i, runtime.NumGoroutine())
	}
	wg.Wait()
	fmt.Printf("结束, 协程数量：%d\n", runtime.NumGoroutine())
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
