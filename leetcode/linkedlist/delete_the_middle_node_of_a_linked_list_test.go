package linkedlist

import (
	"reflect"
	"testing"
)

func TestDeleteMiddle(t *testing.T) {
	testCases := []struct {
		name     string
		head     *ListNode
		expected *ListNode
	}{
		{
			name:     "1",
			head:     &ListNode{Val: 1, Next: &ListNode{Val: 3, Next: &ListNode{Val: 4, Next: &ListNode{Val: 7, Next: &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 6, Next: nil}}}}}}},
			expected: &ListNode{Val: 1, Next: &ListNode{Val: 3, Next: &ListNode{Val: 4, Next: &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 6, Next: nil}}}}}},
		},
		{
			name:     "2",
			head:     &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 3, Next: &ListNode{Val: 4, Next: nil}}}},
			expected: &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 4, Next: nil}}},
		},
		{
			name:     "3",
			head:     &ListNode{Val: 2, Next: &ListNode{Val: 1, Next: nil}},
			expected: &ListNode{Val: 2, Next: nil},
		},
		{
			name:     "4",
			head:     &ListNode{Val: 1, Next: nil},
			expected: nil,
		},
	}

	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			got := DeleteMiddle(tt.head)
			want := tt.expected

			if !reflect.DeepEqual(got, want) {
				t.Errorf("got %v, want %v", got, want)
			}
		})
	}
}