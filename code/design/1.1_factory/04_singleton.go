package factory

import (
	"fmt"
	"sync"
)

// 单例模式要解决的问题：
// 保证一个类只有一个对象，且该对象的功能依然能被其他模块使用
// ···
// 三个要点：
// 1. 类只能有一个实例
// 2. 该实例必须在类内自行创建
// 3. 类需要向外部提供访问该实例的方法

// **************************** 类型1：饿汉式 => 在初始化单例唯一指针的时候，就已经提前开辟好了一个对象，申请了内存
// 优点：不会出现线程并发创建，导致多个单例的出现（线程安全）
// 缺点：如果这个单例对象在业务逻辑中没有被使用，也会客观创建一块内存
// ···
// 1. 确保类非公有化，外界不能通过该类创建对象，所以需要小写
type singleton struct{}

// 2. 设置一个指向唯一对象的指针
var instance *singleton = new(singleton)

// 3. 对外提供一个方法获取这个唯一实例对象
func GetInstance() *singleton {
	return instance
}

func (s *singleton) Get() {
	fmt.Println("饿汉式单例对象方法")
}

// **************************** 类型2：懒汉式 => 只有首次”获取单例“方法被调用，才会生成单例
type singleton2 struct{}

// 重点：没有初始化
var instance2 *singleton2

// 重点：只有在第一次方法调用的时候才创建instance，所以叫”懒“汉式
func GetInstance2() *singleton2 {
	// 只有首次方法被调用，才会生成这个实例的单例
	if instance2 == nil {
		instance2 = new(singleton2)
		return instance2
	}

	// 非首次直接返回已经创建好的示例
	return instance2
}

func (s *singleton2) Get() {
	fmt.Println("懒汉式单例对象方法")
}

// **************************** 类型3：线程安全的懒汉式 => 多个线程同时调用”创建实例“方法，会导致多个实例被创建，需要解决线程安全的问题
// ================ 方式1：互斥锁 ===============
// 优点：解决了线程安全问题
// 缺点：加锁极大影响性能
// 定义锁
var lock sync.Mutex

type singleton3 struct{}

var instance3 *singleton3

func GetInstance3() *singleton3 {
	// 为了线程安全，增加互斥锁
	lock.Lock()
	defer lock.Unlock()

	if instance3 == nil {
		instance3 = new(singleton3)
	}
	return instance3
}

// ================ 方式2：内存的状态存留 ===============
// sync.Once => Do(fn) 判断是否执行过该方法，如果执行过则不执行
var once sync.Once

type singleton4 struct{}

var instance4 *singleton4

func GetInstance4() *singleton4 {
	once.Do(func() {
		instance4 = new(singleton4)
	})

	return instance4
}
