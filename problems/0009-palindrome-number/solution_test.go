package palindromenumber

import "testing"

func Test_isPalindrome(t *testing.T) {
	type args struct {
		x int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{name: "test1", args: args{x: 121}, want: true},
		{name: "test2", args: args{x: -121}, want: false},
		{name: "test3", args: args{x: 10}, want: false},
		{name: "test4", args: args{x: -101}, want: false},
		{name: "test5", args: args{x: 0}, want: true},
		{name: "test6", args: args{x: 12321}, want: true},
		{name: "test7", args: args{x: 123321}, want: true},
		{name: "test8", args: args{x: 1000021}, want: false},
		{name: "test9", args: args{x: 1001}, want: true},
		{name: "test10", args: args{x: 100}, want: false},
		{name: "test11", args: args{x: 2147447412}, want: true},
		{name: "test12", args: args{x: 2147483647}, want: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isPalindrome(tt.args.x); got != tt.want {
				t.Errorf("isPalindrome() = %v, want %v", got, tt.want)
			}
		})
	}
}
