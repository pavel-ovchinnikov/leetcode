package listnode

import (
	"reflect"
	"testing"
)

func Test_converters(t *testing.T) {
	tests := []struct {
		name   string // description of this test case
		values []int
		want   []int
	}{
		{
			name:   "Test 1",
			values: []int{1, 2, 3, 4, 5, 5, 4, 3, 2, 1},
			want:   []int{1, 2, 3, 4, 5, 5, 4, 3, 2, 1},
		},
		{
			name:   "Test 2",
			values: []int{},
			want:   []int{},
		},
		{
			name:   "Test 3",
			values: []int{1},
			want:   []int{1},
		},
		{
			name:   "Test 4",
			values: []int{1, 1},
			want:   []int{1, 1},
		},
		{
			name:   "Test 5",
			values: []int{1, 4, 1},
			want:   []int{1, 4, 1},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			list := ConvertArrayToList(tt.values)
			if !reflect.DeepEqual(ConvertListToArray(list), tt.want) {
				t.Errorf("convertArrayToList() = %v, want %v", ConvertListToArray(list), tt.want)
			}
		})
	}
}

func TestAreListsEqual(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		l1   *ListNode
		l2   *ListNode
		want bool
	}{
		{
			name: "Test Case 1",
			l1:   ConvertArrayToList([]int{1, 2, 3}),
			l2:   ConvertArrayToList([]int{1, 2, 3}),
			want: true,
		},
		{
			name: "Test Case 2",
			l1:   ConvertArrayToList([]int{1, 2, 3}),
			l2:   ConvertArrayToList([]int{1, 2, 4}),
			want: false,
		},
		{
			name: "Test Case 3",
			l1:   nil,
			l2:   nil,
			want: true,
		},
		{
			name: "Test Case 4",
			l1:   nil,
			l2:   ConvertArrayToList([]int{0}),
			want: false,
		},
		{
			name: "Test Case 5",
			l1:   ConvertArrayToList([]int{1, 2, 3}),
			l2:   ConvertArrayToList([]int{1, 2, 3, 4}),
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := AreListsEqual(tt.l1, tt.l2)
			if got != tt.want {
				t.Errorf("AreListsEqual() = %v, want %v", got, tt.want)
			}
		})
	}
}
