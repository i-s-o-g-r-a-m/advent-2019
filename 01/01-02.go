package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func getInput() (masses []float64) {
	input, err := ioutil.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}
	lines := strings.Split(strings.TrimSpace(string(input)), "\n")
	for _, value := range lines {
		mass, err := strconv.ParseFloat(value, 64)
		if err != nil {
			panic(err)
		}
		masses = append(masses, mass)
	}
	return
}

func computeFuelForMass(mass float64) int {
	return int(mass/3.0) - 2
}

func computeFuelFuel(fuel int, accum int) int {
	moreFuel := computeFuelForMass(float64(fuel))
	if moreFuel <= 0 {
		return accum
	} else {
		return computeFuelFuel(moreFuel, accum+moreFuel)
	}
}

func main() {
	masses := getInput()

	// part 1
	totalFuelReq := 0
	for _, mass := range masses {
		totalFuelReq += computeFuelForMass(mass)
	}
	if totalFuelReq != 3457681 {
		panic("part 1: wrong answer")
	} else {
		fmt.Println("part 1:", totalFuelReq)
	}

	// part 2
	totalFuelReq = 0
	for _, mass := range masses {
		massFuel := computeFuelForMass(mass)
		fuelFuel := computeFuelFuel(massFuel, 0)
		totalFuelReq += massFuel + fuelFuel
	}
	if totalFuelReq != 5183653 {
		panic("part 2: wrong answer")
	} else {
		fmt.Println("part 2:", totalFuelReq)
	}
}
