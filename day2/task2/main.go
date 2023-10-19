package main

import "fmt"

func main() {
	number := 7

	if number%7 == 0 {
		fmt.Println(number, "Bilangan kelipatan 7")
	} else {
		fmt.Println(number, "Bukan bilangan kelipatan 7")
	}
}
