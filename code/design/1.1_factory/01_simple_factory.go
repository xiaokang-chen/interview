package factory

import "fmt"

// Car is interface
type Car interface {
	Run(name string) string
}

// BMW、Benz is one of API implement
type BMW struct{}
type Benz struct{}

func New(t int) Car {
	if t == 1 {
		return &BMW{}
	} else {
		return &Benz{}
	}
}

func (b BMW) Run(name string) string {
	return fmt.Sprintf("Hi, %s", name)
}

func (b Benz) Run(name string) string {
	return fmt.Sprintf("Hello, %s", name)
}
