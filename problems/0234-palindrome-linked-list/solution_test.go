package palindromelinkedlist

import "testing"

func Test_isPalindrome(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		head *ListNode
		want bool
	}{
		{
			name: "Test case 1: Palindrome linked list",
			head: &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 2, Next: &ListNode{Val: 1}}}},
			want: true,
		},
		{
			name: "Test case 2: Non-palindrome linked list",
			head: &ListNode{Val: 1, Next: &ListNode{Val: 2}},
			want: false,
		},
		{
			name: "Test case 3: Single element linked list",
			head: &ListNode{Val: 1},
			want: true,
		},
		{
			name: "Test case 4: Empty linked list",
			head: nil,
			want: true,
		},
		{
			name: "Test case 5: Even length palindrome linked list",
			head: &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 3, Next: &ListNode{Val: 2, Next: &ListNode{Val: 1}}}}},
			want: true,
		},
		{
			name: "Test case 6: Odd length non-palindrome linked list",
			head: &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 3}}},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := isPalindrome(tt.head)
			if got != tt.want {
				t.Errorf("isPalindrome() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_isPalindromev2(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		head *ListNode
		want bool
	}{
		{
			name: "Test case 1: Palindrome linked list",
			head: &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 2, Next: &ListNode{Val: 1}}}},
			want: true,
		},
		{
			name: "Test case 2: Non-palindrome linked list",
			head: &ListNode{Val: 1, Next: &ListNode{Val: 2}},
			want: false,
		},
		{
			name: "Test case 3: Single element linked list",
			head: &ListNode{Val: 1},
			want: true,
		},
		{
			name: "Test case 4: Empty linked list",
			head: nil,
			want: true,
		},
		{
			name: "Test case 5: Even length palindrome linked list",
			head: &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 3, Next: &ListNode{Val: 2, Next: &ListNode{Val: 1}}}}},
			want: true,
		},
		{
			name: "Test case 6: Odd length non-palindrome linked list",
			head: &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 3}}},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := isPalindromev2(tt.head)
			if got != tt.want {
				t.Errorf("isPalindromev2() = %v, want %v", got, tt.want)
			}
		})
	}
}
