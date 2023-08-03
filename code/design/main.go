package main

import factory "interview/code/design/1.1_factory"

func main() {
	// ============ 继承1 =============
	// myDog := factory.Dog{factory.Animal{"Tom"}}
	// myDog.Eat()
	// myDog.Bark()

	// ============ 继承2 =============
	// myCat := Cat{"Ketty"}
	// myPerson := Person{"Bob"}
	// myCat.Speak()
	// myPerson.Speak()
	// Speak(&myCat)
	// Speak(&myPerson)

	// ============ 简单工厂 =============
	// bmw3 := factory.BMW{}
	// res := bmw3.Run("三系")
	// mercedes := factory.Benz{}
	// res := mercedes.Run("梅赛德斯")
	// fmt.Println("res: ", res)

	// ============ 工厂方法 =============
	// // 需要一个具体的productA对象
	// // 1-创建一个具体的生产productA的工厂
	// factoryA := &factory.FactoryA{}
	// // 2-生产具体的productA
	// productA := factoryA.Create()
	// res1 := productA.Show("cookie")

	// factoryB := &factory.FactoryB{}
	// productB := factoryB.Create()
	// res2 := productB.Show("shoe")
	// fmt.Println("res1: ", res1)
	// fmt.Println("res2: ", res2)

	// ============ 抽象工厂方法 =============
	// // 需要中国的苹果、香蕉
	// // 1-创建中国工厂
	// cFac := &factory.ChinaFactory{}
	// // 2-创建中国苹果
	// cApple := cFac.CreateApple()
	// cApple.ShowApple()
	// // 3-创建中国香蕉
	// cBanana := cFac.CreateBanana()
	// cBanana.ShowBanana()

	// ============ 单例模式 ==============
	// ************ 饿汉式 *************
	// s1 := factory.GetInstance()
	// s1.Get()
	// // ************ 懒汉式 *************
	// s2 := factory.GetInstance2()
	// s2.Get()

	// ============ 外观模式 ==============
	// homePlayer := new(factory.HomePlayerFacade)

	// homePlayer.DoKTV()
	// fmt.Println("------------")
	// homePlayer.DoKTV()

	// ============ 适配器模式 ==============
	// iphone := factory.NewPhone(factory.NewAdapter(new(factory.V220)))
	// iphone.Charge()

	// ============ 观察者模式 ==============
	s1 := &factory.Student1{"打王者"}
	s2 := &factory.Student1{"吃饭"}

	monitor := &factory.Monitor{}
	monitor.AddListener(s1)
	monitor.AddListener(s2)

	s1.Do()
	s2.Do()
	monitor.Notify()
}
