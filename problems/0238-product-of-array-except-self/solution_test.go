package task0238

import (
	"reflect"
	"testing"
)

func Test_productExceptSelf(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want []int
	}{
		{name: "empty input", nums: []int{}, want: []int{}},
		{name: "single number", nums: []int{5}, want: []int{1}},
		{name: "two numbers", nums: []int{2, 3}, want: []int{3, 2}},
		{name: "positive numbers", nums: []int{1, 2, 3, 4}, want: []int{24, 12, 8, 6}},
		{name: "negative numbers", nums: []int{-1, 2, -3, 4}, want: []int{-24, 12, -8, 6}},
		{name: "one zero", nums: []int{1, 2, 0, 4}, want: []int{0, 0, 8, 0}},
		{name: "two zeros", nums: []int{0, 2, 0, 4}, want: []int{0, 0, 0, 0}},
		{name: "mixed signs and zero", nums: []int{-1, 1, 0, -3, 3}, want: []int{0, 0, 9, 0, 0}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := productExceptSelf(tt.nums)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("productExceptSelf() = %v, want %v", got, tt.want)
			}
		})
	}
}
