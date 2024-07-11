package linkedlists

type SinglyLinkedNode struct {
	Val  int
	Next *SinglyLinkedNode
}

type DoublyLinkedNode struct {
	Val  int
	Prev *DoublyLinkedNode
	Next *DoublyLinkedNode
}

func arrToSinglyLinkedList(nums []int) *SinglyLinkedNode {
	var (
		head                   = SinglyLinkedNode{nums[0], nil}
		curr *SinglyLinkedNode = &head
	)
	for i := 1; i < len(nums); i++ {
		curr.Next = &SinglyLinkedNode{nums[i], nil}
		curr = curr.Next
	}
	return &head
}

func arrToDoublyLinkedList(nums []int) *DoublyLinkedNode {
	var (
		head                   = DoublyLinkedNode{nums[0], nil, nil}
		curr *DoublyLinkedNode = &head
	)
	for i := 1; i < len(nums); i++ {
		curr.Next = &DoublyLinkedNode{
			Val:  nums[i],
			Prev: curr,
			Next: nil,
		}
		curr = curr.Next
	}
	return &head
}
