package main

import (
	"go-design-pattern/behavioral/strategy/navigator"
	"go-design-pattern/behavioral/strategy/navigator/navigator_impls/bicycle"
)

func main() {
	bicycle := bicycle.NewBicycle()
	// bus := bus.NewBus()
	// taxi := taxi.NewTaxi()

	nav := navigator.NewNavigator(bicycle)
	nav.Navigator.BuildRouter()
}
