package task0344

import (
	"reflect"
	"testing"
)

func Test_reverseString(t *testing.T) {
	type args struct {
		s []byte
	}
	tests := []struct {
		name string
		args args
		want []byte
	}{
		{"test 01", args{[]byte("hello")}, []byte("olleh")},
		{"test 02", args{[]byte("Hannah")}, []byte("hannaH")},
		{"test 03", args{[]byte("a")}, []byte("a")},
		{"test 04", args{[]byte{}}, []byte{}},
		{"test 05", args{[]byte("ab")}, []byte("ba")},
		{"test 06", args{[]byte("A man, a plan, a canal: Panama")}, []byte("amanaP :lanac a ,nalp a ,nam A")},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			reverseString(tt.args.s)
			if !reflect.DeepEqual(tt.args.s, tt.want) {
				t.Errorf("reverseString() = %v, want %v", tt.args.s, tt.want)
			}
		})
	}
}