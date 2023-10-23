package main

import "fmt"

func ArrayMerge(arrayA, arrayB []string) []string {
	result := []string{}
	uniqueMap := map[string]bool{}
	mergeSlice := append(arrayA, arrayB...)
	for _, value := range mergeSlice {
		_, isFound := uniqueMap[value]
		if !isFound {
			result = append(result, value)
			uniqueMap[value] = true
		}
	}
	return result
}

func main() {
	arrayA := []string{"ayam", "burung", "cicak"}
	arrayB := []string{"domba", "entok", "fanta, ayam"}
	fmt.Println(ArrayMerge(arrayA, arrayB))

	arrayC := []string{"ayam", "bebek"}
	arrayD := []string{}
	fmt.Println(ArrayMerge(arrayC, arrayD))
}
