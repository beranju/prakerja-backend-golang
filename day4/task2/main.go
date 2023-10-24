package main

import (
	"fmt"
)

type Student struct {
	Name  string
	Score int
}

func (students *Student) calculateAvarageScore(list []Student) float64 {
	total := 0
	length := len(list)

	for _, student := range list {
		total += student.Score
	}

	return float64(total) / float64(length)
}

func (student *Student) getMinMaxScore(list []Student) (int, int) {
	maxim := list[0].Score
	minim := list[0].Score
	var maximIdx int
	var minimIdx int
	for i, v := range list {
		if v.Score > maxim {
			maxim = student.Score
			maximIdx = i
		}
		if v.Score < minim {
			minim = v.Score
			minimIdx = i
		}
	}
	return minimIdx, maximIdx
}

func main() {
	var students []Student
	for i := 0; i < 5; i++ {
		var name string
		var score int

		fmt.Printf("Input %d Student's Name ", i+1)
		fmt.Scanln(&name)
		fmt.Printf("Input %d Student's Score ", i+1)
		fmt.Scanln(&score)
		student := Student{Name: name, Score: score}
		students = append(students, student)
	}
	averageScore := students[0].calculateAvarageScore(students)
	fmt.Println("Average Score", averageScore)
	minimIdx, maximIdx := students[0].getMinMaxScore(students)
	fmt.Printf("Max Score of Students : %s (%d)", students[maximIdx].Name, students[maximIdx].Score)
	fmt.Printf("Min Score of Students : %s (%d)", students[minimIdx].Name, students[minimIdx].Score)

}
