package bicycle

import "fmt"

type bicycle struct{}

func (bicycle) BuildRouter() {
	fmt.Println("Router for bicycle")
}

func NewBicycle() *bicycle {
	return &bicycle{}
}
