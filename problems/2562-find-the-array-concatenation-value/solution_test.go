package findthearrayconcatenationvalue

import "testing"

func Test_findTheArrayConcVal(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		nums []int
		want int64
	}{
		{
			name: "Test 1",
			nums: []int{7, 52, 2, 4},
			want: 596,
		},
		{
			name: "Test 2",
			nums: []int{5, 14, 13, 8, 12},
			want: 673,
		},
		{
			name: "Test 3",
			nums: []int{5},
			want: 5,
		},
		{
			name: "Test 34",
			nums: []int{},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := findTheArrayConcVal(tt.nums)
			if got != tt.want {
				t.Errorf("findTheArrayConcVal() = %v, want %v", got, tt.want)
			}
		})
	}
}
