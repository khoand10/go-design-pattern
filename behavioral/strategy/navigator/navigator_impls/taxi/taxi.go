package taxi

import "fmt"

type taxi struct{}

func (taxi) BuildRouter() {
	fmt.Println("Router for taxi")
}

func NewTaxi() *taxi {
	return &taxi{}
}
