package removeelement

import "testing"

func Test_removeElement(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		nums []int
		val  int
		want int
	}{
		{name: "test 1", nums: []int{3, 2, 2, 3}, val: 3, want: 2},
		{name: "test 2", nums: []int{0, 1, 2, 2, 3, 0, 4, 2}, val: 2, want: 5},
		{name: "test 3", nums: []int{0, 1, 0, 0}, val: 0, want: 1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := removeElement(tt.nums, tt.val)
			if got != tt.want {
				t.Errorf("removeElement() = %v, want %v", got, tt.want)
			}
		})
	}
}
