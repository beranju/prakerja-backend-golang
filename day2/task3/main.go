package main

import "fmt"

func main() {
	// a := 5.0
	// b := 4.0
	// t := 8.0

	var a float32
	var b float32
	var t float32

	fmt.Println("Masukkan 3 angka berurutan a b t =")
	fmt.Scanln(&a, &b, &t)

	area := 0.5 * (a + b) * t
	fmt.Printf("area = %.2f", area)

}
