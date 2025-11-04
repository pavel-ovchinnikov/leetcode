package findthemaximumachievablenumber

import "testing"

func Test_theMaximumAchievableX(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		num  int
		t    int
		want int
	}{
		{
			name: "Example 1",
			num:  4,
			t:    1,
			want: 6,
		},
		{
			name: "Example 2",
			num:  3,
			t:    2,
			want: 7,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := theMaximumAchievableX(tt.num, tt.t)
			if got != tt.want {
				t.Errorf("theMaximumAchievableX() = %v, want %v", got, tt.want)
			}
		})
	}
}
