package main

import (
	"fmt"
)

func main() {
	n := 6
	var numbers = make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Printf("Masukkan bilangan ke-%d: ", i+1)
		_, err := fmt.Scanln(&numbers[i])
		if err != nil {
			fmt.Println("Terjadi kesalahan input", err)
			break
		}
	}
	min, max := findMinMax(numbers)
	fmt.Println("min ", min, ", max ", max)

}

func findMinMax(input []int) (min, max int) {

	if len(input) == 0 {
		fmt.Println("Slice is empty")
	}

	minim := input[0]
	for _, val := range input {
		if val < minim {
			minim = val
		}
	}
	maxim := input[0]
	for _, val := range input {
		if val > maxim {
			maxim = val
		}
	}
	return minim, maxim
}
