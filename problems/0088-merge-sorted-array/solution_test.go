package tasks0088

import (
	"reflect"
	"testing"
)

func Test_merge(t *testing.T) {
	tests := []struct {
		name  string
		nums1 []int
		m     int
		nums2 []int
		n     int
		want  []int
	}{
		{"typical", []int{1, 2, 3, 0, 0, 0}, 3, []int{2, 5, 6}, 3, []int{1, 2, 2, 3, 5, 6}},
		{"nums1 empty", []int{0, 0, 0}, 0, []int{1, 2, 3}, 3, []int{1, 2, 3}},
		{"nums2 empty", []int{1, 2, 3}, 3, []int{}, 0, []int{1, 2, 3}},
		{"both empty", []int{}, 0, []int{}, 0, []int{}},
		{"all nums1 less", []int{1, 2, 3, 0, 0, 0}, 3, []int{4, 5, 6}, 3, []int{1, 2, 3, 4, 5, 6}},
		{"all nums2 less", []int{4, 5, 6, 0, 0, 0}, 3, []int{1, 2, 3}, 3, []int{1, 2, 3, 4, 5, 6}},
		{"duplicates across", []int{1, 1, 1, 0, 0, 0}, 3, []int{1, 1, 1}, 3, []int{1, 1, 1, 1, 1, 1}},
		{"single each nums1 less", []int{1, 0}, 1, []int{2}, 1, []int{1, 2}},
		{"single each nums2 less", []int{2, 0}, 1, []int{1}, 1, []int{1, 2}},
		{"negative numbers", []int{-1, 0, 0, 0}, 1, []int{-2, 2, 3}, 3, []int{-2, -1, 2, 3}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			merge(tt.nums1, tt.m, tt.nums2, tt.n)
			if !reflect.DeepEqual(tt.nums1, tt.want) {
				t.Errorf("merge() nums1 = %v, want %v", tt.nums1, tt.want)
			}
		})
	}
}
