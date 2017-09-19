package gradingStudents

import (
	"fmt"
	"reflect"
	"testing"
)

func TestGradeStudents(t *testing.T) {
	testCases := []struct {
		in  []int
		out []int
	}{
		{[]int{73, 67, 38, 33}, []int{75, 67, 40, 33}},
	}

	for _, tc := range testCases {
		actual := gradeStudents(tc.in, len(tc.in))
		if !reflect.DeepEqual(actual, tc.out) {
			fmt.Println("Got", actual)
			fmt.Println("Wanted", tc.out)
		}
	}
}
