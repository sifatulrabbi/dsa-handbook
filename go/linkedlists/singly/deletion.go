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

func DeleteNode(n *linkedlists.SinglyLinkedNode, val int) {
}

func DeleteAtithPos(n *linkedlists.SinglyLinkedNode, i int, val int) {
}
