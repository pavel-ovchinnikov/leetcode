package tasks0026

import (
	"reflect"
	"testing"
)

func Test_removeDuplicates(t *testing.T) {
	tests := []struct {
		name     string
		nums     []int
		wantLen  int
		wantNums []int
	}{
		{"test 01", []int{1, 1, 2}, 2, []int{1, 2}},
		{"test 02", []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}, 5, []int{0, 1, 2, 3, 4}},
		{"test 03", []int{}, 0, []int{}},
		{"test 04", []int{1}, 1, []int{1}},
		{"test 05", []int{1, 1, 1}, 1, []int{1}},
		{"test 06", []int{1, 2, 3}, 3, []int{1, 2, 3}},
		{"test 07", []int{1, 1, 2, 3}, 3, []int{1, 2, 3}},
		{"test 08", []int{1, 2, 2, 3}, 3, []int{1, 2, 3}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := removeDuplicates(tt.nums)
			if got != tt.wantLen {
				t.Errorf("removeDuplicates() = %v, want len %v", got, tt.wantLen)
			}
			if !reflect.DeepEqual(tt.nums[:got], tt.wantNums) {
				t.Errorf("nums = %v, want %v", tt.nums[:got], tt.wantNums)
			}
		})
	}
}
