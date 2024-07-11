class SinglyLinkedList:
    val: int
    next = None

    def __init__(self, val: int):
        self.val = val
        self.next = None

    @staticmethod
    def from_arr(arr: list[int]):
        head = SinglyLinkedList(arr[0])
        curr = head
        i = 1
        while i < len(arr):
            curr.next = SinglyLinkedList(arr[i])
            curr = curr.next
            i += 1
        return head


class DoublyLinkedList:
    val: int
    prev = None
    next = None

    def __init__(self, val: int):
        self.val = val
        self.next = None

    @staticmethod
    def from_arr(arr: list[int]):
        head = DoublyLinkedList(arr[0])
        curr = head
        i = 1
        while i < len(arr):
            curr.next = DoublyLinkedList(arr[i])
            curr = curr.next
            i += 1
        return head
