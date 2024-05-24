package testutil

import "go-dsa/structs"

func CreateCycleLinkedList(data []int, pos int) *structs.LinkedList {
	ll := structs.CreateLinkedList(data)

	if pos == -1 {
		return ll
	}

	var currentNode, cycleNode *structs.LinkedListNode

	currentNode = ll.Head
	for i := 0; i < len(data)-1; i++ {
		if i == pos {
			cycleNode = currentNode
		}
		currentNode = currentNode.Next
	}

	if cycleNode != nil {
		currentNode.Next = cycleNode
	}

	return ll
}
