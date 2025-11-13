package minimizemaximumpairsuminarray

import "testing"

func Test_minPairSum(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		nums []int
		want int
	}{
		{
			name: "Example 1",
			nums: []int{3, 5, 2, 3},
			want: 7,
		},
		{
			name: "Example 2",
			nums: []int{3, 5, 4, 2, 4, 6},
			want: 8,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := minPairSum(tt.nums)
			if got != tt.want {
				t.Errorf("minPairSum() = %v, want %v", got, tt.want)
			}
		})
	}
}
