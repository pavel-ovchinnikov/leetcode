package numberofdistinctaverages

import "testing"

func Test_distinctAverages(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		nums []int
		want int
	}{
		{
			name: "Test Case 1",
			nums: []int{4, 1, 4, 0, 3, 5},
			want: 2,
		}, {
			name: "Test Case 2",
			nums: []int{1, 100},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := distinctAverages(tt.nums)
			if got != tt.want {
				t.Errorf("distinctAverages() = %v, want %v", got, tt.want)
			}
		})
	}
}
