package linked_list

import "go-dsa/structs"

func Reverse(head *structs.LinkedListNode) *structs.LinkedListNode {
	var prev, next *structs.LinkedListNode
	curr := head

	for curr != nil {
		next = curr.Next
		curr.Next = prev
		prev = curr
		curr = next
	}

	head = prev
	return head
}

func ReverseBetween(head *structs.LinkedListNode, left int, right int) *structs.LinkedListNode {
	if head == nil || left == right {
		return head
	}

	var prev, next, beforeStart, start *structs.LinkedListNode
	curr := head

	i := 1
	for ; i < left; i++ {
		beforeStart = curr
		curr = curr.Next
	}

	start = curr

	for ; i <= right; i++ {
		next = curr.Next
		curr.Next = prev
		prev = curr
		curr = next
	}

	start.Next = curr
	if beforeStart == nil {
		head = prev
	} else {
		beforeStart.Next = prev
	}

	return head
}
