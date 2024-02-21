package bus

import "fmt"

type bus struct{}

func (bus) BuildRouter() {
	fmt.Println("Router for bus")
}

func NewBus() *bus {
	return &bus{}
}
