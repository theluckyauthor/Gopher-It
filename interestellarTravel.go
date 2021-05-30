package main

import "fmt"

// Create the function fuelGauge() here
func fuelGauge(fuel int) {
	fmt.Println("you have", fuel, "fuel left")
}

// Create the function calculateFuel() here
func calculateFuel(planet string) int {
	var fuel int = 0
	switch planet {
	case "Mercury":
		fuel = 500000
	case "Venus":
		fuel = 300000
	case "Mars":
		fuel = 1000000
	}
	return fuel
}

// Create the function greetPlanet() here
func greetPlanet(planet string) {
	fmt.Println("You are going to " + planet)
}

// Create the function cantFly() here
func cantFly() {
	fmt.Println("We do not have the available fuel to fly there.")
}

// Create the function flyToPlanet() here
func flyToPlanet(planet string, fuel int) int {
	var fuelRemaining, fuelCost int

	fuelRemaining = fuel
	fuelCost = calculateFuel(planet)
	if fuelRemaining >= fuelCost {
		greetPlanet(planet)
		fuelRemaining -= fuelCost
	} else {
		cantFly()
	}
	return fuelRemaining
}

func main() {

	// Create `planetChoice` and `fuel`
	var fuel int = 1000000
	var planetChoice string = "Mars"
	// And then liftoff!
	fuel = flyToPlanet(planetChoice, fuel)
	fuelGauge(fuel)

}
