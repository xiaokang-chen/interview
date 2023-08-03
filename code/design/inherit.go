package main

import "fmt"

// `Go` 中不存在继承和多态 所以通过下面方法实现：
// 1. 继承：使用【匿名组合】-嵌套结构体来实现
// 2. 多态：接口实现

// ======================== 继承：嵌套结构体 =========================
type Animal struct {
	Name string
}

func (a *Animal) Eat() {
	fmt.Printf("%s is eating\n", a.Name)
}

type Dog struct {
	Animal
}

func (d *Dog) Bark() {
	fmt.Printf("%s is barking\n", d.Name)
}

// ======================== 嵌套结构体 END =========================

// ======================== 多态：接口实现 =========================

// 1. 定义一个能发出声音的接口
type Speaker interface {
	Speak()
}

// 2. 定义一个人和一个猫的结构体，他们都实现了Speak()方法
type Person struct {
	Name string
}

func (p *Person) Speak() {
	fmt.Printf("%s is speaking\n", p.Name)
}

type Cat struct {
	Name string
}

func (c *Cat) Speak() {
	fmt.Printf("%s is cat, and speaking\n", c.Name)
}

// 3. 定义一个Speak()函数，该函数的参数是Speak接口类型，该函数可以根据不同的Speak实例，输出不同语句
func Speak(s Speaker) {
	s.Speak()
}
