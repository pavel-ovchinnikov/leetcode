package tasks0066

import (
	"reflect"
	"testing"
)

func Test_plusOne(t *testing.T) {
	tests := []struct {
		name   string
		digits []int
		want   []int
	}{
		{"test 01", []int{1, 2, 3}, []int{1, 2, 4}},
		{"test 02", []int{4, 3, 2, 1}, []int{4, 3, 2, 2}},
		{"test 03", []int{9}, []int{1, 0}},
		{"test 04", []int{9, 9}, []int{1, 0, 0}},
		{"test 05", []int{1, 9, 9}, []int{2, 0, 0}},
		{"test 06", []int{0}, []int{1}},
		{"test 07", []int{1, 0}, []int{1, 1}},
		{"test 08", []int{8, 9, 9}, []int{9, 0, 0}},
		{"test 09", []int{9, 8, 9}, []int{9, 9, 0}},
		{"test 10", []int{2, 9, 9, 9}, []int{3, 0, 0, 0}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := plusOne(tt.digits); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("plusOne(%v) = %v, want %v", tt.digits, got, tt.want)
			}
		})
	}
}
