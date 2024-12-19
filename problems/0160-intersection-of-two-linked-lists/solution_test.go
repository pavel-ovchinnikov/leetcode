package intersectionoftwolinkedlists

import (
	"reflect"
	"testing"

	. "github.com/pavel-ovchinnikov/leetcode/tools/structures/linked_list"
)

func Test_getIntersectionNode(t *testing.T) {
	type args struct {
		headA *ListNode
		headB *ListNode
	}
	listA1, listB1, tail1 := makeTwoListAndTail(
		[]int{1, 2},
		[]int{-5, -2, 1},
		[]int{3, 4},
	)

	listA2, listB2, tail2 := makeTwoListAndTail(
		[]int{1},
		[]int{-5, -2, 1},
		[]int{3, 4},
	)

	tests := []struct {
		name string
		args args
		want *ListNode
	}{
		{
			"test 1",
			args{listA1, listB1},
			tail1,
		},
		{
			"test 2",
			args{listA2, listB2},
			tail2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := getIntersectionNode(tt.args.headA, tt.args.headB)
			if !reflect.DeepEqual(ConvertListToArray(got), ConvertListToArray(tt.want)) {
				t.Errorf("getIntersectionNode() = %v, want %v", got, tt.want)
			}
		})
	}
}

func makeTwoListAndTail(headA, headB, tail []int) (*ListNode, *ListNode, *ListNode) {
	listA := ConvertArrayToList(headA)
	listB := ConvertArrayToList(headB)
	listTail := ConvertArrayToList(tail)

	addTail(listA, listTail)
	addTail(listB, listTail)

	return listA, listB, listTail
}

func addTail(head, tail *ListNode) {
	cur := head
	for cur.Next != nil {
		cur = cur.Next
	}
	cur.Next = tail
}

func Test_getIntersectionNode_v2(t *testing.T) {
	type args struct {
		headA *ListNode
		headB *ListNode
	}
	listA1, listB1, tail1 := makeTwoListAndTail(
		[]int{1, 2},
		[]int{-5, -2, 1},
		[]int{3, 4},
	)

	listA2, listB2, tail2 := makeTwoListAndTail(
		[]int{1},
		[]int{-5, -2, 1},
		[]int{3, 4},
	)

	tests := []struct {
		name string
		args args
		want *ListNode
	}{
		{
			"test 1",
			args{listA1, listB1},
			tail1,
		},
		{
			"test 2",
			args{listA2, listB2},
			tail2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := getIntersectionNode_v2(tt.args.headA, tt.args.headB)
			if !reflect.DeepEqual(ConvertListToArray(got), ConvertListToArray(tt.want)) {
				t.Errorf("getIntersectionNode() = %v, want %v", got, tt.want)
			}
		})
	}
}
