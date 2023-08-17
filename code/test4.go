package main

import (
	"fmt"
	"sync"
	"time"
)

// Print 无缓冲
// 阻塞通道
func Print() {
	ch := make(chan string)
	wg := sync.WaitGroup{}
	wg.Add(1)
	go func() {
		defer wg.Done()
		s := <-ch
		fmt.Println("str", s)
	}()
	ch <- "A"
	wg.Wait()
}

// Print1 有缓冲
func Print1() {
	ch := make(chan string, 1)
	ch <- "A"
	fmt.Println(<-ch)
}

// Print2 无缓冲，主输出子输出
func Print2() {
	ch := make(chan string)
	go func() {
		ch <- "A"
	}()
	// println(<-ch)
	s := <-ch
	fmt.Println(s)
}

// Print3 无缓冲，主输入子输出
func Print3() {
	ch := make(chan string)
	wg := sync.WaitGroup{}
	wg.Add(1)
	go func() {
		defer wg.Done()
		s := <-ch
		fmt.Println(s)
	}()
	ch <- "A"
	wg.Wait()
}

// Print4 多路复用解决无缓冲阻塞问题
func Print4() {
	ch := make(chan string)
	wg := sync.WaitGroup{}
	wg.Add(1)
	go func() {
		defer wg.Done()
		// 问题：这里循环去读的话，会出现读不到的阻塞问题，造成死锁deadlock
		// for {
		// 	s := <-ch
		// 	fmt.Println("s", s)
		// }
		// 解决方案：使用select的default去处理异常
		select {
		case s := <-ch:
			fmt.Println("s", s)
		default:
			// time.Sleep(time.Second)
			fmt.Println("no data")
		}
	}()
	ch <- "A"
	wg.Wait()
}

// Print5 定时器
// time.Ticker
func Print5() {
	// 每隔1s执行的任务器
	tasker := NewTicker(time.Second)

	// 模拟发送停止信号
	go func(tt *Tasker) {
		time.Sleep(5 * time.Second)
		tt.stopchan <- struct{}{}
	}(tasker)

	for {
		// 多路复用
		select {
		case <-tasker.t.C:
			fmt.Printf("Task:%s\n", time.Now().Format("2006-01-02 15:04:05"))
		case <-tasker.stopchan:
			fmt.Println("停止任务")
			return
		}
	}
}

type Tasker struct {
	t        *time.Ticker
	stopchan chan struct{}
}

func NewTicker(d time.Duration) *Tasker {
	return &Tasker{
		t:        time.NewTicker(d),
		stopchan: make(chan struct{}),
	}
}

// CloseChan 关闭管道
// 下面两种情况会panic:
// 1. 关闭nil通道
// 2. 关闭通道后，继续发送或读取
func CloseChan() {
	ch := make(chan string)
	tk := time.NewTicker(time.Second)
	go func() {
		for range tk.C {
			ch <- "A"
			fmt.Println("插入A...")
		}
	}()
	go func() {
		time.Sleep(2 * time.Second)
		close(ch)
		fmt.Println("关闭通道")
	}()
	for item := range ch {
		time.Sleep(time.Second)
		fmt.Println(item)
	}
}
