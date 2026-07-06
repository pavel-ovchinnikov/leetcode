package task0013

import "testing"

func Test_romanToInt(t *testing.T) {
	tests := []struct {
		name string
		s    string
		want int
	}{
		{"test 01", "III", 3},
		{"test 02", "LVIII", 58},
		{"test 03", "MCMXCIV", 1994},
		{"test 04", "I", 1},
		{"test 05", "IV", 4},
		{"test 06", "IX", 9},
		{"test 07", "XL", 40},
		{"test 08", "XC", 90},
		{"test 09", "CD", 400},
		{"test 10", "CM", 900},
		{"test 11", "M", 1000},
		{"test 12", "MMMCMXCIX", 3999},
		{"test 13", "XIV", 14},
		{"test 14", "CXL", 140},
		{"test 15", "MMXXIV", 2024},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := romanToInt(tt.s); got != tt.want {
				t.Errorf("romanToInt(%q) = %v, want %v", tt.s, got, tt.want)
			}
		})
	}
}
