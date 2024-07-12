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
		t.FailNow()
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
		t.FailNow()
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
		t.FailNow()
	}
}

func TestInsertAtIthPosition(t *testing.T) {
	var (
		arr  = []int{1, 2, 3, 4, 5, 6, 7, 8}
		head = linkedlists.FromArrToSinglyLinkedList(arr)
	)
	InsertAtIthPos(head.Next.Next, 4, 0)
	targetNode := head.Next.Next
	if targetNode.Next.Val != 0 && targetNode.Next.Next.Val != 5 {
		t.Errorf("incorrectly updated, result: %v\n", linkedlists.ConvertSinglyLinkedListToArr(head))
		t.FailNow()
	}
}

func TestDeleteHead(t *testing.T) {
	var (
		arr  = []int{1, 2, 3, 4, 5, 6, 7, 8}
		head = linkedlists.FromArrToSinglyLinkedList(arr)
	)
	newHead := DeleteHead(head)
	if newHead.Val != 2 {
		t.Errorf("failed to remove the head, result: %v\n",
			linkedlists.ConvertSinglyLinkedListToArr(newHead))
		t.FailNow()
	}
}

func TestDeleteTail(t *testing.T) {
	var (
		arr  = []int{1, 2, 3, 4, 5, 6, 7, 8}
		head = linkedlists.FromArrToSinglyLinkedList(arr)
	)
	DeleteTail(head)
	result := linkedlists.ConvertSinglyLinkedListToArr(head)
	if result[len(result)-1] != 7 {
		t.Errorf("failed to remove the tail, result: %v\n", result)
		t.FailNow()
	}
}

func TestDeleteANode(t *testing.T) {
	var (
		arr  = []int{1, 2, 3, 4, 5, 6, 7, 8}
		head = linkedlists.FromArrToSinglyLinkedList(arr)
	)
	DeleteNode(head.Next.Next)
	result := linkedlists.ConvertSinglyLinkedListToArr(head)
	if result[2] != 4 {
		t.Errorf("unable to delete node(3), result: %v\n", result)
		t.FailNow()
	}
}

func TestDeleteAtIthPosition(t *testing.T) {
	var (
		arr  = []int{1, 2, 3, 4, 5, 6, 7, 8}
		head = linkedlists.FromArrToSinglyLinkedList(arr)
	)
	DeleteAtIthPos(head, 3)
	result := linkedlists.ConvertSinglyLinkedListToArr(head)
	if result[2] != 4 {
		t.Errorf("unable to delete node(3), result: %v\n", result)
		t.FailNow()
	}
}
