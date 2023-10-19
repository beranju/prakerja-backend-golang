package main

import "fmt"

func main() {
	// gunakan ini jika value belum diketahui
	var ageHuman uint = 10
	fmt.Println(ageHuman)
	var name = "beranju"
	fmt.Println(name)

	// gunakan ini  ketika valuenya sudah di ketahui
	phi := 3.14
	fmt.Println(phi)

	// bitwise
	a := 10
	result := a << 1
	fmt.Println(result)

	// condition
	age := 18
	if age > 17 {
		fmt.Println("Kamu dewasa")
	} else {
		fmt.Println("kamu anak")
	}

	// looping
	for i := 0; i < 10; i++ {
		fmt.Println("hello")
	}

	/*
	****
	****
	****
	 */
	for i := 1; i <= 5; i++ {
		for j := 1; j <= 5; j++ {
			fmt.Print("*")
		}
		fmt.Println()
	}

	/*
	*
	**
	***
	 */
	for i := 1; i <= 5; i++ {
		for j := 1; j <= i; j++ {
			fmt.Print("*")
		}
		fmt.Println()
	}

	/*
		  *
		 ***
		*****
	*/
	n := 5
	for i := 1; i <= n; i++ {

		// print sebelah kiri
		for k := n - 1; k >= i; k-- {
			fmt.Print(" ")
		}

		// print sebelah kanan
		for j := 1; j <= i; j++ {
			fmt.Print("* ")
		}
		fmt.Println()
	}

}
