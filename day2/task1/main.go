package main

import "fmt"

func main() {
	number := 6

	if number > 1 {
		for i := 2; i <= number; i++ {
			if number%i == 0 {
				fmt.Println(number, "Bukan prima")
				return
			}
		}
		fmt.Println(number, "Prima")
		return
	}
}
