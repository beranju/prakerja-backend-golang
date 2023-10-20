package main

import "fmt"

func ArrayMerge(arrayA, arrayB []string) []string {
	return append(arrayA, arrayB...)
}

func main() {
	arrayA := []string{"ayam", "burung", "cicak"}
	arrayB := []string{"domba", "entok", "fanta"}
	fmt.Println(ArrayMerge(arrayA, arrayB))

	arrayC := []string{"ayam", "bebek"}
	arrayD := []string{}
	fmt.Println(ArrayMerge(arrayC, arrayD))
}
