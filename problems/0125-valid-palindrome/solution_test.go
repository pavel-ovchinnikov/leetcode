package tasks0125

import "testing"

func Test_isPalindrome(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"test 01", args{"A man, a plan, a canal: Panama"}, true},
		{"test 02", args{"race a car"}, false},
		{"test 03", args{" "}, true},
		{"test 04", args{"."}, true},
		{"test 05", args{".,"}, true},
		{"test 06", args{"a"}, true},
		{"test 07", args{"aa"}, true},
		{"test 08", args{"aba"}, true},
		{"test 09", args{"ab"}, false},
		{"test 10", args{"abca"}, false},
		{"test 11", args{"0P"}, false},
		{"test 12", args{"0P0"}, true},
		{"test 13", args{"Never odd or even"}, true},
		{"test 14", args{"No 'x', in Nixon"}, true},
		{"test 15", args{"Was it a car or a cat I saw?"}, true},
		{"test 16", args{"Madam, I'm Adam"}, true},
		{"test 17", args{"Zeus was deified, saw Suez"}, true},
		{"test 18", args{"a a"}, true},
		{"test 19", args{"a b"}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isPalindrome(tt.args.s); got != tt.want {
				t.Errorf("isPalindrome() = %v, want %v", got, tt.want)
			}
		})
	}
}
