package structs

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
