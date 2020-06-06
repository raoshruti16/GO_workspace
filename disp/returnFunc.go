package main

import (
	"fmt"
	"math"
	"os"
)

func main() {

	var acceleration, velocity, displacement, time, exit float64

	// Repeatedly do displacement computation with user input unless exits
	for {

		getParam(&acceleration, "acceleration")
		getParam(&velocity, "velocity")
		getParam(&displacement, "displacement")

		fmt.Println("\nCreating function to enter time parameter......\n")
		funcTime := GenDisplaceFn(acceleration,
			velocity,
			displacement)

		getParam(&time, "time")
		fmt.Println("\nCalculating displacment.....:\n")
		fmt.Println("Final displacement = " , funcTime(time))

		getParam(&exit, "0 to exit or 1 to continue again")
		if exit == 0 {
			break
		}
	}

}

// Takes in float parameter from user displaying message
func getParam(param *float64, paramName string) {
	fmt.Printf("\nEnter %s: \n", paramName)
	_,err :=  fmt.Scan(param)
	if err != nil {
		panic(err)
	}
	if *param < 0 {
		fmt.Println("Velocity, displacement, time, acceleration or exit code must be non-negative")
        os.Exit(1)
	}
}

// Takes in parameters returning function that can take in time parameter
func GenDisplaceFn(acceleration,
					velocity,
					displacement float64) func(t float64) float64 {
	f := func(t float64) float64 {
		return acceleration*0.5*math.Pow(t,2) + velocity*t + displacement
	}

	return f
}

