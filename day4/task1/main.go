package main

import (
	"errors"
	"fmt"
)

type Car struct {
	Type   string
	FuelIn float64
}

func (car Car) estimateDistance() (float64, error) {
	var litterPerMil float64 = 1.5
	if car.FuelIn < 0.0 {
		return 0, errors.New("fuel cannot be negative")
	}
	distance := car.FuelIn / litterPerMil
	return distance, nil
}

func main() {
	toyota := Car{
		Type:   "Sport",
		FuelIn: -100.60,
	}
	distance, err := toyota.estimateDistance()
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("You can reach %f mil with %f L of fuel", distance, toyota.FuelIn)
	}

}
