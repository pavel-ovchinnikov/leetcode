package romantointeger

import "testing"

func Test_romanToInt(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		s    string
		want int
	}{
		{
			name: "example 1",
			s:    "III",
			want: 3,
		},
		{
			name: "example 2",
			s:    "IV",
			want: 4,
		},
		{
			name: "example 3",
			s:    "IX",
			want: 9,
		},
		{
			name: "example 4",
			s:    "LVIII",
			want: 58,
		},
		{
			name: "example 5",
			s:    "MCMXCIV",
			want: 1994,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := romanToInt(tt.s)
			// TODO: update the condition below to compare got with tt.want.
			if got != tt.want {
				t.Errorf("romanToInt() = %v, want %v", got, tt.want)
			}
		})
	}
}
