package structs

import (
	"strconv"
	"strings"
)

type LinkedListNode struct {
	Data int
	Next *LinkedListNode
}

func NewLinkedListNode(data int, next *LinkedListNode) *LinkedListNode {
	return &LinkedListNode{
		Data: data,
		Next: next,
	}
}

func InitLinkedListNode(data int) *LinkedListNode {
	return &LinkedListNode{
		Data: data,
		Next: nil,
	}
}

func (ll *LinkedListNode) String() string {
	if ll == nil {
		return ""
	}

	curr := ll
	valueList := []string{}
	for curr != nil {
		valueList = append(valueList, strconv.Itoa(curr.Data))
		curr = curr.Next
	}

	return strings.Join(valueList, " -> ")
}

type LinkedList struct {
	Head *LinkedListNode
}

func CreateLinkedList(lst []int) *LinkedList {
	head := InitLinkedListNode(lst[0])
	currentNode := head
	for i := 1; i < len(lst); i++ {
		newNode := InitLinkedListNode(lst[i])
		currentNode.Next = newNode
		currentNode = newNode
	}

	return &LinkedList{
		Head: head,
	}
}

func (ll *LinkedList) String() string {
	if ll.Head == nil {
		return ""
	}

	curr := ll.Head
	valueList := []string{}
	for curr != nil {
		valueList = append(valueList, strconv.Itoa(curr.Data))
		curr = curr.Next
	}

	return strings.Join(valueList, " -> ")
}
