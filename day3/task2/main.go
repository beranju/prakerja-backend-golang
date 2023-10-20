package main

import (
	"fmt"
)

func Mapping(slice []string) map[string]int {
	countMap := make(map[string]int)
	for _, v := range slice {
		countMap[v]++
	}
	return countMap
}

func main() {
	words := []string{"apple", "banana", "apple", "Cherry", "banana", "apple", "apple", "Cherry", "apple"}
	wordCount := Mapping(words)
	fmt.Println(wordCount)
}
