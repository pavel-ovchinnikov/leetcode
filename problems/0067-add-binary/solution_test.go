package tasks0067

import "testing"

func Test_addBinary(t *testing.T) {
	tests := []struct {
		name string
		a    string
		b    string
		want string
	}{
		{"test 01", "11", "1", "100"},
		{"test 02", "1010", "1011", "10101"},
		{"test 03", "0", "0", "0"},
		{"test 04", "1", "0", "1"},
		{"test 05", "0", "1", "1"},
		{"test 06", "1", "1", "10"},
		{"test 07", "100", "100", "1000"},
		{"test 08", "111", "111", "1110"},
		{"test 09", "1010", "101", "1111"},
		{"test 10", "1", "111", "1000"},
		{"test 11", "1111", "1", "10000"},
		{"test 12", "10", "1", "11"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := addBinary(tt.a, tt.b); got != tt.want {
				t.Errorf("addBinary(%q, %q) = %q, want %q", tt.a, tt.b, got, tt.want)
			}
		})
	}
}
