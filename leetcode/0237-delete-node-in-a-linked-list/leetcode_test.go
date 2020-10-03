package leetcode

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_deleteNode(t *testing.T) {
	t.Run("delete first node", func(t *testing.T) {
		expected := &ListNode{
			Val: 5,
			Next: &ListNode{
				Val: 1,
				Next: &ListNode{
					Val: 9,
				},
			},
		}

		head := createListNode()
		deleteNode(head)
		assert.Equal(t, expected, head)
	})

	t.Run("delete second node", func(t *testing.T) {
		expected := &ListNode{
			Val: 4,
			Next: &ListNode{
				Val: 1,
				Next: &ListNode{
					Val: 9,
				},
			},
		}

		head := createListNode()
		deleteNode(head.Next)
		assert.Equal(t, expected, head)
	})

	t.Run("delete third node", func(t *testing.T) {
		expected := &ListNode{
			Val: 4,
			Next: &ListNode{
				Val: 5,
				Next: &ListNode{
					Val: 9,
				},
			},
		}

		head := createListNode()
		deleteNode(head.Next.Next)
		assert.Equal(t, expected, head)
	})
}

func createListNode() *ListNode {
	return &ListNode{
		Val: 4,
		Next: &ListNode{
			Val: 5,
			Next: &ListNode{
				Val: 1,
				Next: &ListNode{
					Val: 9,
				},
			},
		},
	}
}
