package main

import (
	"fmt"
	factory "interview/code/design/1.1_factory"
)

func main() {
	// bmw3 := factory.BMW{}
	// res := bmw3.Run("三系")
	// mercedes := factory.Benz{}
	// res := mercedes.Run("梅赛德斯")
	// fmt.Println("res: ", res)

	factoryA := &factory.FactoryA{}
	productA := factoryA.Create()
	res1 := productA.Show("cookie")

	factoryB := &factory.FactoryB{}
	productB := factoryB.Create()
	res2 := productB.Show("shoe")
	fmt.Println("res1: ", res1)
	fmt.Println("res2: ", res2)
}
