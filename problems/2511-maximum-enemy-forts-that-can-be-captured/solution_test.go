package maximumenemyfortsthatcanbecaptured

import "testing"

func Test_captureForts(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		forts []int
		want  int
	}{
		{
			name:  "Test case 1",
			forts: []int{1, 0, 0, -1, 0, 0, 0, 0, 1},
			want:  4,
		},
		{
			name:  "Test case 2",
			forts: []int{0, 0, 1, -1},
			want:  0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := captureForts(tt.forts)
			if got != tt.want {
				t.Errorf("captureForts() = %v, want %v", got, tt.want)
			}
		})
	}
}
