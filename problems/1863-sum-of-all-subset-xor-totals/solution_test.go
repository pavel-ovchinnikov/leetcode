package sumofallsubsetxortotals

import "testing"

func Test_subsetXORSum(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		nums []int
		want int
	}{
		{
			name: "Example 1",
			nums: []int{1, 3},
			want: 6,
		},
		{
			name: "Example 2",
			nums: []int{5, 1, 6},
			want: 28,
		},
		{
			name: "Example 3",
			nums: []int{3, 4, 5, 6, 7, 8},
			want: 480,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := subsetXORSum(tt.nums)

			if got != tt.want {
				t.Errorf("subsetXORSum() = %v, want %v", got, tt.want)
			}
		})
	}
}
