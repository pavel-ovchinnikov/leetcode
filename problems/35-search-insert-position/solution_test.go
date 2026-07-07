package tasks0035

import "testing"

func Test_searchInsert(t *testing.T) {
	tests := []struct {
		name   string
		nums   []int
		target int
		want   int
	}{
		{"test 01", []int{1, 3, 5, 6}, 5, 2},
		{"test 02", []int{1, 3, 5, 6}, 2, 1},
		{"test 03", []int{1, 3, 5, 6}, 7, 4},
		{"test 04", []int{1, 3, 5, 6}, 0, 0},
		{"test 05", []int{1, 3, 5, 6}, 3, 1},
		{"test 06", []int{1, 3, 5, 6}, 6, 3},
		{"test 07", []int{1}, 0, 0},
		{"test 08", []int{1}, 2, 1},
		{"test 09", []int{1}, 1, 0},
		{"test 10", []int{}, 5, 0},
		{"test 11", []int{1, 3}, 0, 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := searchInsert(tt.nums, tt.target); got != tt.want {
				t.Errorf("searchInsert(%v, %d) = %v, want %v", tt.nums, tt.target, got, tt.want)
			}
		})
	}
}
