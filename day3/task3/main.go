package main

import (
	"fmt"
	"strconv"
	"strings"
)

func munculSekali(value string) []int {

	chars := strings.Split(value, "")

	countMap := make(map[int]int)
	for _, char := range chars {
		num, err := strconv.Atoi(char)
		if err == nil {
			countMap[num]++
		}
	}

	uniqueNumber := []int{}
	for num, count := range countMap {
		if count == 1 {
			uniqueNumber = append(uniqueNumber, num)
		}
	}
	return uniqueNumber
}

func main() {

	numString := "128731897410927492"
	result := munculSekali(numString)
	fmt.Println(result)

}
