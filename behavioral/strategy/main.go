package main

import (
	"go-design-pattern/behavioral/strategy/navigator"
	"go-design-pattern/behavioral/strategy/navigator/impls/bicycle"
)

func main() {
	bicycleNav := bicycle.NewBicycle()
	// bus := bus.NewBus()
	// taxi := taxi.NewTaxi()

	nav := navigator.NewNavigator(bicycleNav)
	nav.Navigator.BuildRouter()
}
