package divisibleandnondivisiblesumsdifference

import "testing"

func Test_differenceOfSums(t *testing.T) {
	type args struct {
		n int
		m int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"test 1", args{10, 3}, 19},
		{"test 2", args{5, 6}, 15},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := differenceOfSums(tt.args.n, tt.args.m); got != tt.want {
				t.Errorf("differenceOfSums() = %v, want %v", got, tt.want)
			}
		})
	}
}
