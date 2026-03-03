package removeduplicatesfromsortedarray

import "testing"

func Test_removeDuplicates(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		nums []int
		want int
	}{
		{name: "test 1", nums: []int{1, 1, 2}, want: 2},
		{name: "test 2", nums: []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}, want: 5},
		{name: "test 3", nums: []int{1}, want: 1},
		{name: "test 4", nums: []int{1, 1}, want: 1},
		{name: "test 5", nums: []int{}, want: 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := removeDuplicates(tt.nums)
			if got != tt.want {
				t.Errorf("removeDuplicates() = %v, want %v", got, tt.want)
			}
		})
	}
}
