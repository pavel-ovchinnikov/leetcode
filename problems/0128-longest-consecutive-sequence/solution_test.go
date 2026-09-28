package task0128

import "testing"

func Test_longestConsecutive(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want int
	}{
		{name: "nil slice", nums: nil, want: 0},
		{name: "empty slice", nums: []int{}, want: 0},
		{name: "single number", nums: []int{7}, want: 1},
		{name: "no consecutive numbers", nums: []int{1, 3, 5}, want: 1},
		{name: "unordered sequence", nums: []int{100, 4, 200, 1, 3, 2}, want: 4},
		{name: "duplicates in sequence", nums: []int{0, 3, 7, 2, 5, 8, 4, 6, 0, 1}, want: 9},
		{name: "all duplicates", nums: []int{2, 2, 2}, want: 1},
		{name: "sequence across zero", nums: []int{2, -2, 0, -1, 1}, want: 5},
		{name: "longest of multiple sequences", nums: []int{10, 11, 1, 2, 3, 20}, want: 3},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := longestConsecutive(tt.nums); got != tt.want {
				t.Errorf("longestConsecutive(%v) = %d, want %d", tt.nums, got, tt.want)
			}
		})
	}
}
