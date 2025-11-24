package convertbinarynumberinalinkedlisttointeger

import (
	"testing"

	. "github.com/pavel-ovchinnikov/leetcode/tools/list_node"
)

func Test_getDecimalValue(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		head *ListNode
		want int
	}{
		{
			name: "Test Case 1",
			head: ConvertArrayToList([]int{1, 0, 1}),
			want: 5,
		},
		{
			name: "Test Case 2",
			head: ConvertArrayToList([]int{0}),
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := getDecimalValue(tt.head)
			if got != tt.want {
				t.Errorf("getDecimalValue() = %v, want %v", got, tt.want)
			}
		})
	}
}
