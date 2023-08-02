package factory

import "fmt"

// `Go` 中不存在继承 所以使用匿名组合来实现

// factory is factory
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
type FactoryA struct{}
type FactoryB struct{}

func (f *FactoryA) Create() Product {
	return &ProductA{}
}

func (f *FactoryB) Create() Product {
	return &ProductB{}
}
