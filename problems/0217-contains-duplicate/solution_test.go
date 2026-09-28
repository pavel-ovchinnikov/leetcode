package task0217

import "testing"

func Test_containsDuplicate(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want bool
	}{
		{name: "empty slice", nums: []int{}, want: false},
		{name: "single number", nums: []int{1}, want: false},
		{name: "distinct numbers", nums: []int{1, 2, 3, 4}, want: false},
		{name: "adjacent duplicates", nums: []int{1, 1}, want: true},
		{name: "separated duplicates", nums: []int{1, 2, 3, 1}, want: true},
		{name: "negative duplicate", nums: []int{-3, 2, -3}, want: true},
		{name: "zero duplicate", nums: []int{0, -1, 0}, want: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := containsDuplicate(tt.nums)
			if got != tt.want {
				t.Errorf("containsDuplicate() = %v, want %v", got, tt.want)
			}
		})
	}
}
