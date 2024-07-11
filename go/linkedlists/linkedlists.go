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

func FromArrToSinglyLinkedList(nums []int) *SinglyLinkedNode {
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

func FromArrToDoublyLinkedList(nums []int) *DoublyLinkedNode {
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

func ConvertSinglyLinkedListToArr(head *SinglyLinkedNode) []int {
	var (
		arr = []int{}
		cur = head
	)
	for cur != nil {
		arr = append(arr, cur.Val)
		cur = cur.Next
	}
	return arr
}

func ConvertDoublyLinkedListToArr(head *DoublyLinkedNode) []int {
	var (
		arr = []int{}
		cur = head
	)
	for cur != nil {
		arr = append(arr, cur.Val)
		cur = cur.Next
	}
	return arr
}
