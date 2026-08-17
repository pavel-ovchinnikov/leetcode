package task0345

import "testing"

func Test_reverseVowels(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"test 01", args{"hello"}, "holle"},
		{"test 02", args{"leetcode"}, "leotcede"},
		{"test 03", args{"aA"}, "Aa"},
		{"test 04", args{"a"}, "a"},
		{"test 05", args{"aeiou"}, "uoiea"},
		{"test 06", args{"AI"}, "IA"},
		{"test 07", args{"bcdfg"}, "bcdfg"},
		{"test 08", args{"a."}, "a."},
		{"test 09", args{" "}, " "},
		{"test 10", args{"ae"}, "ea"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := reverseVowels(tt.args.s); got != tt.want {
				t.Errorf("reverseVowels() = %v, want %v", got, tt.want)
			}
		})
	}
}