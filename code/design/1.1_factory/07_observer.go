package factory

import "fmt"

// 行为型：观察者模式
// 包含2个角色
// 1. 被观察者（抽象主题）：为调用方定义简单的调用接口
// 2. 子系统：功能提供者

// 抽象方法
// 1. 观察者
type Observer interface {
	Update() // 观察者得到通知后触发的动作
}

// 2. 通知者
type Notifier interface {
	AddListener(o Observer)
	RemoveListener(o Observer)
	Notify()
}

// 实现观察者
type Student1 struct {
	Thing string
}

func (s1 Student1) Update() {
	fmt.Println("同学1 更新了 ", s1.Thing)
}

func (s1 Student1) Do() {
	fmt.Println("同学1正在做事")
}

type Student2 struct {
	Thing string
}

func (s2 Student2) Update() {
	fmt.Println("同学2 更新了 ", s2.Thing)
}

func (s2 Student2) Do() {
	fmt.Println("同学2正在做事")
}

// 实现通知者
type Monitor struct {
	ObserverList []Observer
}

func (m *Monitor) AddListener(observer Observer) {
	m.ObserverList = append(m.ObserverList, observer)
}

func (m *Monitor) RemoveListener(observer Observer) {
	for index, l := range m.ObserverList {
		// 找到要删除的元素
		if observer == l {
			// 将元素删除：把删除节点后的元素和前面连接起来
			m.ObserverList = append(m.ObserverList, m.ObserverList[index+1:]...)
			break
		}
	}
}

func (m *Monitor) Notify() {
	for _, l := range m.ObserverList {
		// 依次调用方法
		l.Update()
	}
}
