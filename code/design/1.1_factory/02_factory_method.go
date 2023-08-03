package factory

import "fmt"

// 包含三个角色：
// 1. 抽象工厂：工厂方法模式的核心，任何工厂类都必须实现这个接口
// 2. 具体工厂：具体工厂类是抽象工厂的一个实现，负责实例化产品对象
// 3. 抽象产品：它为每种产品声明接口，在抽象产品中声明了产品所具有的业务方法
// 4. 具体产品：创建的具体实例对象

// Factory is factory
// 将创建对象的方法封装到Factory中
type Factory interface {
	Create() Product
}

// Product is interface
type Product interface {
	Show(name string) string
}

type ProductA struct{}
type ProductB struct{}

func (a ProductA) Show(name string) string {
	return fmt.Sprintf("i am %s, type is a", name)
}

func (b ProductB) Show(name string) string {
	return fmt.Sprintf("i am %s, type is b", name)
}

// 设立两个工厂，并实现工厂的方法
// FactoryA 专门创建 ProductA
// FactoryB 专门创建 ProductB
// ******* 和简单工厂的区别：具体工厂会创建多个，不同的工厂成产不同的产品 ********
type FactoryA struct{}
type FactoryB struct{}

func (f *FactoryA) Create() Product {
	return &ProductA{}
}

func (f *FactoryB) Create() Product {
	return &ProductB{}
}
