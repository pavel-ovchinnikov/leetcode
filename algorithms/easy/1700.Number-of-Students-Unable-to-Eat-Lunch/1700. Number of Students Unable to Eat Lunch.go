package main

import "fmt"

func countStudents(students []int, sandwiches []int) int {
	count := 0
	for len(students) > count {
		if students[0] == sandwiches[0] {
			sandwiches = append(sandwiches[:0], sandwiches[1:]...)
			count = 0
		} else {
			students = append(students, students[0])
			count++
		}
		students = append(students[:0], students[1:]...)
	}
	return len(students)
}

func main() {
	fmt.Println(countStudents(
		[]int{1, 1, 0, 0},
		[]int{0, 1, 0, 1}))

	fmt.Println(countStudents(
		[]int{1, 1, 1, 0, 0, 1},
		[]int{1, 0, 0, 0, 1, 1}))
}
