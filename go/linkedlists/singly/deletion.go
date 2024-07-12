package singly

import (
	"github.com/sifatulrabbi/dsa-handbook/linkedlists"
)

func DeleteHead(h *linkedlists.SinglyLinkedNode) *linkedlists.SinglyLinkedNode {
	return h.Next
}

func DeleteTail(n *linkedlists.SinglyLinkedNode) {
	if n.Next == nil {
		n = nil
		return
	}
	for cur := n; cur != nil; cur = cur.Next {
		if cur.Next.Next == nil {
			cur.Next = nil
			break
		}
	}
}

func DeleteNode(n *linkedlists.SinglyLinkedNode) {
	for cur := n; cur.Next != nil; cur = cur.Next {
		cur.Val = cur.Next.Val
		if cur.Next.Next == nil {
			cur.Next = nil
			break
		}
	}
}

func DeleteAtIthPos(h *linkedlists.SinglyLinkedNode, idx int) {
	var (
		curr = h
		i    = 1
	)
	for curr != nil {
		if i == idx {
			DeleteNode(curr)
			break
		}

		i++
		curr = curr.Next
	}
}
