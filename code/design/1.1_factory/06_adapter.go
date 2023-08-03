package factory

import "fmt"

// 结构型：适配器模式
// 将一个类的接口转换成客户希望的另外一个接口，使得原本由于接口不兼容的类可以一起工作
// 包含三个角色
// 1. 目标抽象类：
// 2. 适配器类：
// 3. 适配者类：

// 场景：给支持5V电压的iphone手机使用220V电压充电

// 抽象类：适配目标V5
type V5 interface {
	Use5V()
}

// 业务实现
type Phone struct {
	v5 V5
}

func NewPhone(v5 V5) *Phone {
	return &Phone{v5}
}

func (p Phone) Charge() {
	fmt.Println("给手机充电开始，使用5V电压...")

	p.v5.Use5V()
}

// 适配者
type V220 struct{}

func (v V220) Use220V() {
	fmt.Println("使用220V电压")
}

// 适配器
type Adapter struct {
	v220 *V220
}

func (a *Adapter) Use5V() {
	fmt.Println("使用适配器充电，转换电压...")
	a.v220.Use220V()
}

func NewAdapter(v220 *V220) *Adapter {
	return &Adapter{v220}
}
