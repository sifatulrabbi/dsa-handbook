package singly

import (
	"github.com/sifatulrabbi/dsa-handbook/linkedlists"
)

func InsertAtTheBeginning(h *linkedlists.SinglyLinkedNode, val int) *linkedlists.SinglyLinkedNode {
	newH := linkedlists.SinglyLinkedNode{Val: val, Next: h}
	return &newH
}

func InsertAtTheEnd(n *linkedlists.SinglyLinkedNode, val int) {
	// approach 1: recursion
	if n.Next == nil {
		n.Next = &linkedlists.SinglyLinkedNode{Val: val, Next: nil}
	} else {
		InsertAtTheEnd(n.Next, val)
	}

	// approach 2: for loop
	// for cur := n; cur != nil; cur = cur.Next {
	// 	if cur.Next == nil {
	// 		cur.Next = &linkedlists.SinglyLinkedNode{Val: val, Next: nil}
	// 		break
	// 	}
	// }
}

func InsertAfterNode(n *linkedlists.SinglyLinkedNode, val int) {
	next := n.Next
	n.Next = &linkedlists.SinglyLinkedNode{Val: val, Next: next}
}

func InsertAtithPos(h *linkedlists.SinglyLinkedNode, i int, val int) {
}
