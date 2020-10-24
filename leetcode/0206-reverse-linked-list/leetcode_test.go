package leetcode

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

type give struct {
	head *ListNode
}

func Test_reverseList(t *testing.T) {
	tests := []struct {
		give give
		want *ListNode
	}{
		{
			give: give{
				head: &ListNode{
					Val: 1,
					Next: &ListNode{
						Val: 2,
						Next: &ListNode{
							Val: 3,
							Next: &ListNode{
								Val: 4,
								Next: &ListNode{
									Val: 5,
								},
							},
						},
					},
				},
			},
			want: &ListNode{
				Val: 5,
				Next: &ListNode{
					Val: 4,
					Next: &ListNode{
						Val: 3,
						Next: &ListNode{
							Val: 2,
							Next: &ListNode{
								Val: 1,
							},
						},
					},
				},
			},
		},
	}

	for _, tt := range tests {
		got := reverseList(tt.give.head)
		assert.Equal(t, tt.want, got)
	}
}
