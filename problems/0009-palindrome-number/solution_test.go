package task0009

import "testing"

func Test_isPalindrome(t *testing.T) {
	tests := []struct {
		name string
		x    int
		want bool
	}{
		{"test 01", -121, false},
		{"test 02", -1, false},

		{"test 03", 0, true},
		{"test 04", 1, true},
		{"test 05", 9, true},

		{"test 06", 11, true},
		{"test 07", 121, true},
		{"test 08", 1221, true},
		{"test 09", 12321, true},
		{"test 10", 1001, true},

		{"test 11", 10, false},
		{"test 12", 123, false},
		{"test 13", 1234, false},

		{"test 14", 10001, true},

		{"test 15", 123456789987654321, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isPalindrome(tt.x); got != tt.want {
				t.Errorf("isPalindrome(%d) = %v, want %v", tt.x, got, tt.want)
			}
		})
	}
}
