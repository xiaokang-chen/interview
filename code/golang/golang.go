package golang

import (
	"fmt"
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

// type I1 interface {
// 	add()
// 	set()
// }

// type S1 struct {
// 	Name string
// 	Age  int
// }

// func (s S1) add() {
// 	fmt.Println("s1 add 1")
// }

// func (s *S1) set() {
// 	s.Age++
// 	fmt.Println("s1 add 2")
// }
