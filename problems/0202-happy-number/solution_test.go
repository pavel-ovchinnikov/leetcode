package task0202

import "testing"

func Test_isHappy(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"test 01", args{1}, true},
		{"test 02", args{19}, true},
		{"test 03", args{7}, true},
		{"test 04", args{10}, true},
		{"test 05", args{13}, true},
		{"test 06", args{68}, true},
		{"test 07", args{100}, true},
		{"test 08", args{2}, false},
		{"test 09", args{3}, false},
		{"test 10", args{4}, false},
		{"test 11", args{5}, false},
		{"test 12", args{6}, false},
		{"test 13", args{8}, false},
		{"test 14", args{9}, false},
		{"test 15", args{20}, false},
		{"test 16", args{123}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isHappy(tt.args.n); got != tt.want {
				t.Errorf("isHappy() = %v, want %v", got, tt.want)
			}
		})
	}
}
