package factory

import "fmt"

// 结构型：外观模式
// 包含两个角色
// 1. 外观：为调用方定义简单的调用接口
// 2. 子系统：功能提供者

// 案例举例：
// 在家庭影院系统中，包含”KTV模式“、”游戏模式“
// 切换到不同的模式，会打开对应的设备
// KTV模式：电视机、投影仪、麦克风、灯光、音箱
// 游戏模式：电视机、灯光、游戏机

// 子系统-电视机
type TV struct{}

// 子系统-音箱
type VoiceBox struct{}

// 子系统-灯光
type Light struct{}

// 子系统-游戏机
type Xbox struct{}

// 子系统-麦克风
type MicroPhone struct{}

// 子系统-投影仪
type Projector struct{}

func (t *TV) On() {
	fmt.Println("打开 电视机")
}
func (v *VoiceBox) On() {
	fmt.Println("打开 音箱")
}
func (l *Light) On() {
	fmt.Println("打开 灯光")
}
func (x *Xbox) On() {
	fmt.Println("打开 游戏机")
}
func (m *MicroPhone) On() {
	fmt.Println("打开 麦克风")
}
func (p *Projector) On() {
	fmt.Println("打开 投影仪")
}

// 外观：家庭影院
type HomePlayerFacade struct {
	tv    TV
	vb    VoiceBox
	light Light
	xbox  Xbox
	mp    MicroPhone
	pro   Projector
}

// KTV模式
func (hp *HomePlayerFacade) DoKTV() {
	fmt.Println("家庭影院进入KTV模式")
	hp.tv.On()
	hp.pro.On()
	hp.mp.On()
	hp.vb.On()
}

// 游戏模式
func (hp *HomePlayerFacade) DoGame() {
	fmt.Println("家庭影院进入Game模式")
	hp.tv.On()
	hp.light.On()
	hp.xbox.On()
}
