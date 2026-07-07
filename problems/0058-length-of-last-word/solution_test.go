package tasks0058

import "testing"

func Test_lengthOfLastWord(t *testing.T) {
	tests := []struct {
		name string
		s    string
		want int
	}{
		{"test 01", "Hello World", 5},
		{"test 02", "   fly me   to   the moon  ", 4},
		{"test 03", "luffy is still joyboy", 6},
		{"test 04", "a", 1},
		{"test 05", "a ", 1},
		{"test 06", "word", 4},
		{"test 07", "   word", 4},
		{"test 08", "word   ", 4},
		{"test 09", "   word   ", 4},
		{"test 10", "today is a good  day", 3},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := lengthOfLastWord(tt.s); got != tt.want {
				t.Errorf("lengthOfLastWord(%q) = %v, want %v", tt.s, got, tt.want)
			}
		})
	}
}
