package gradingStudents

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	var n int
	io := bufio.NewReader(os.Stdin)
	_, err := fmt.Fscan(io, &n)
	if err != nil {
		fmt.Println(err)
	}

	grades := make([]int, n)

	for i := 0; i < n; i++ {
		fmt.Fscan(io, &grades[i])
	}

	newGrades := gradeStudents(grades, n)

	fmt.Println(newGrades)
}

func gradeStudents(grades []int, n int) []int {

	newGrades := make([]int, n)

	for i, grade := range grades {
		newGrades[i] = adjustGrade(grade)
	}

	return newGrades
}

func adjustGrade(grade int) int {
	diff := grade % 5

	if grade < 38 {
		return grade
	}

	if 2 < diff {
		return grade + (5 - diff)
	}

	return grade
}
