package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	// recoverExample()
	// fmt.Println(unsafe.Sizeof(demo1{})) // 8
	// fmt.Println(unsafe.Sizeof(demo2{})) // 12

	// u := []User{
	// 	{"A", 12},
	// 	{"B", 15},
	// 	{"C", 8},
	// }
	// s1 := make([]*User, 0, len(u))
	// for _, v := range u {
	// 	s1 = append(s1, &v)
	// }
	// for _, item := range s1 {
	// 	fmt.Println(item)
	// }
	// arr := [][]int{{1, 0}, {2, 0}, {3, 1}, {3, 2}}
	// res := getAns(arr)
	// fmt.Println(res)
	// m["c"] = 3
	// smp := sync.Map{}
	// smp.Store("a", 1)
	// smp.Store("b", 2)
	// smp = m
	// fmt.Println(smp.Load("a"))
	// arr := []int{1, 2, 3, 4}
	// list := New(arr)
	// list.Pop()
	// list.Push(5)
	// list.Push(6)
	// len := list.Len()
	// fmt.Println("len: ", len)

	go doprint(1)
	go doprint(2)
	time.Sleep(2 * time.Second)
}

var a string
var done bool
var once sync.Once = sync.Once{}

func setup() {
	a = "hello, world"
	done = true
}

func doprint(hao int) {
	if !done {
		once.Do(setup)
	}
	print("a", a, hao)
}

type NestedIterator struct {
	vals []int
}

type User struct {
	Name string
	age  int
}

// 内存对齐
type demo1 struct {
	a int8
	b int16
	c int32
}
type demo2 struct {
	a int8
	b int32
	c int16
}

// recover()用于捕获panic并处理
func doSomething() {
	defer func() {
		if r := recover(); r != nil {
			// 处理panic
			fmt.Println("Recovered:", r)
		}
	}()
	panic("发生了panic！")
}

func recoverFromPanic() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered from panic:", r)
			fmt.Println(456)
		}
	}()

	fmt.Println(123)
	panic("Something went wrong!")
}

func recoverExample() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered from panic:", r)
		}
	}()

	fmt.Println("Before panic")

	panic("Something went wrong!")

	fmt.Println("After panic") // 这行代码不会被执行
}
