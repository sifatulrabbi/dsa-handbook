package singly

import (
	"testing"

	"github.com/sifatulrabbi/dsa-handbook/linkedlists"
)

func TestInsertAtTheBeginning(t *testing.T) {
	var (
		arr  = []int{1, 2, 3, 4, 5, 6, 7, 8}
		head = linkedlists.FromArrToSinglyLinkedList(arr)
	)
	newHead := InsertAtTheBeginning(head, 0)
	result := linkedlists.ConvertSinglyLinkedListToArr(newHead)
	if result[0] != 0 {
		t.Errorf("Failed to insert at the beginning of the linked list. result: %v\n", result)
	}
}

func TestInsertAtTheEnd(t *testing.T) {
	var (
		arr  = []int{1, 2, 3, 4, 5, 6, 7, 8}
		head = linkedlists.FromArrToSinglyLinkedList(arr)
	)
	InsertAtTheEnd(head, 9)
	result := linkedlists.ConvertSinglyLinkedListToArr(head)
	if result[len(result)-1] != 9 {
		t.Errorf("Failed to insert at the beginning of the linked list. result: %v\n", result)
	}
}

func TestInsertAfterNode(t *testing.T) {
	var (
		arr  = []int{1, 2, 3, 4, 5, 6, 7, 8}
		head = linkedlists.FromArrToSinglyLinkedList(arr)
	)
	InsertAfterNode(head.Next.Next, 0)
	targetNode := head.Next.Next
	if targetNode.Next.Val != 0 && targetNode.Next.Next.Val != 4 {
		t.Errorf("incorrectly updated, result: %v\n", linkedlists.ConvertSinglyLinkedListToArr(head))
	}
}
