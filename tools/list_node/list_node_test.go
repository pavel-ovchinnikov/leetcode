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
			list := convertArrayToList(tt.values)
			if !reflect.DeepEqual(convertListToArray(list), tt.want) {
				t.Errorf("convertArrayToList() = %v, want %v", convertListToArray(list), tt.want)
			}
		})
	}
}
