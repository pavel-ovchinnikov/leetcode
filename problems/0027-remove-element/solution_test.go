package tasks0027

import (
	"reflect"
	"testing"
)

func Test_removeElement(t *testing.T) {
	tests := []struct {
		name     string
		nums     []int
		val      int
		wantK    int
		wantNums []int
	}{
		{"test 01", []int{3, 2, 2, 3}, 3, 2, []int{2, 2}},
		{"test 02", []int{0, 1, 2, 2, 3, 0, 4, 2}, 2, 5, []int{0, 1, 3, 0, 4}},
		{"test 03", []int{}, 0, 0, []int{}},
		{"test 04", []int{1}, 1, 0, []int{}},
		{"test 05", []int{1}, 2, 1, []int{1}},
		{"test 06", []int{1, 1, 1}, 1, 0, []int{}},
		{"test 07", []int{1, 2, 3}, 0, 3, []int{1, 2, 3}},
		{"test 08", []int{4, 5}, 4, 1, []int{5}},
		{"test 09", []int{4, 5}, 5, 1, []int{4}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := removeElement(tt.nums, tt.val)
			if got != tt.wantK {
				t.Errorf("removeElement() = %v, want len %v", got, tt.wantK)
			}
			if !reflect.DeepEqual(tt.nums[:got], tt.wantNums) {
				t.Errorf("nums = %v, want %v", tt.nums[:got], tt.wantNums)
			}
		})
	}
}
