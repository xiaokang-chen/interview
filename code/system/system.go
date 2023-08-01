package system

import "fmt"

// ================================== 哲学家进餐问题 =========================================

// N 哲学家人数
var N = 5

// 叉子，对应的信号量
var signal = [5]int{1, 1, 1, 1, 1}

// think 思考
func think() {
	fmt.Println("think")
}

// eat 进餐
func eat() {
	fmt.Println("eat")
}

// P操作，信号量+1
func P(position int) {
	signal[position] += 1
}

// V操作，信号量+1
func V(position int) {
	signal[position] -= 1
}

// SmartPerson1 哲学家进餐问题-方案1
// i为哲学家编号，0-4
//
// 问题：如果五个哲学家同时拿起左边叉子，就会出现死锁现象
func SmartPerson1(i int) {
	for {
		// 哲学家思考
		think()
		// 拿左边的叉子
		P(i)
		// 拿右边的叉子
		P((i + 1) % N)
		// 哲学家进餐
		eat()
		// 放下左边的叉子
		V(i)
		// 放下右边的叉子
		V((i + 1) % N)
	}
}

// 互斥信号量
var mutex = 1

// SmartPerson1 哲学家进餐问题-方案2
// 拿叉子前，加个互斥信号量
//
// 问题：每次只能一个哲学家进餐，不能并发
func SmartPerson2(i int) {
	for {
		// 哲学家思考
		think()
		// 进入临界区
		P(mutex)
		// 拿左边的叉子
		P(i)
		// 拿右边的叉子
		P((i + 1) % N)
		// 哲学家进餐
		eat()
		// 放下左边的叉子
		V(i)
		// 放下右边的叉子
		V((i + 1) % N)
		// 退出临界区
		V(mutex)
	}
}

// SmartPerson3
// 偶数编号【先拿左边的叉子后拿右边的叉子】，奇数编号【先拿右边的叉子后拿左边的叉子】
//
// 不会出现死锁，也可以两个人同时进餐
func SmartPerson3(i int) {
	for {
		// 哲学家思考
		think()
		// 奇偶判断
		if i%2 == 0 {
			// 拿左边的叉子
			P(i)
			// 拿右边的叉子
			P((i + 1) % N)
		} else {
			// 拿右边的叉子
			P((i + 1) % N)
			// 拿左边的叉子
			P(i)
		}
		// 哲学家进餐
		eat()

		// 放下左边的叉子
		V(i)
		// 放下右边的叉子
		V((i + 1) % N)
	}
}

// ================================== 哲学家进餐问题 END =========================================

// ================================== 读者-写者问题 =========================================

// 控制写操作的互斥信号量，初始值为1
var wMutex = 1

// 正在进行读操作的读者个数
var rCount = 0

// 控制对rCount的互斥修改，初始值为1
var rCountMutex = 1

// 写
func write() {
	fmt.Println("write")
}

// 读
func read() {
	fmt.Println("read")
}

func RW_P(mutexType int) {
	if mutexType == 1 {
		wMutex++
	}
	if mutexType == 2 {
		rCountMutex++
	}
}

func RW_V(mutexType int) {
	if mutexType == 1 {
		wMutex--
	}
	if mutexType == 2 {
		rCountMutex--
	}
}

// 信号量
func writer() {
	for {
		// 进入临界区
		RW_P(1)
		// 读
		write()
		// 离开临界区
		RW_V(1)
	}
}

func reader() {
	for {
		// 进入临界区
		RW_P(2)
		// 如果有写者，则阻塞写者
		if rCount == 0 {
			RW_P(1)
		}
		// 读者数量+1
		rCount++
		// 离开临界区
		RW_V(2)

		read()

		// 进入临界区
		RW_P(2)
		// 读完数据，离开
		rCount--
		// 最后一个读者离开，唤醒写者
		if rCount == 0 {
			RW_V(1)
		}
		// 离开临界区
		RW_V(2)
	}
}

// ================================== 读者-写者问题 END =========================================
