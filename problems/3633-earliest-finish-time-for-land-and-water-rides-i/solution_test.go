package earliestfinishtimeforlandandwaterridesi

import "testing"

func Test_earliestFinishTime(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		landStartTime  []int
		landDuration   []int
		waterStartTime []int
		waterDuration  []int
		want           int
	}{
		{
			name:           "example 1",
			landStartTime:  []int{2, 8},
			landDuration:   []int{4, 1},
			waterStartTime: []int{6},
			waterDuration:  []int{3},
			want:           9,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := earliestFinishTime(tt.landStartTime, tt.landDuration, tt.waterStartTime, tt.waterDuration)
			if got != tt.want {
				t.Errorf("earliestFinishTime() = %v, want %v", got, tt.want)
			}
		})
	}
}
