package minimumaverageofsmallestandlargestelements

import "testing"

func Test_minimumAverage(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		nums []int
		want float64
	}{
		{
			name: "Example 1",
			nums: []int{4, 1, 3, 2},
			want: 2.5,
		},
		{
			name: "Example 2",
			nums: []int{1, 2, 3, 4, 5, 6},
			want: 3.5,
		},
		{
			name: "Example 3",
			nums: []int{1, 100},
			want: 50.5,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := minimumAverage(tt.nums)
			if got != tt.want {
				t.Errorf("minimumAverage() = %v, want %v", got, tt.want)
			}
		})
	}
}
