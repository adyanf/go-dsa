package stacks

import (
	"go-dsa/structs"
)

func RemoveAdjacentDuplicates(str string) string {
	stack := structs.NewStack()

	for _, char := range str {
		top := stack.Top()
		if top == int(char) {
			stack.Pop()
		} else {
			stack.Push(int(char))
		}

	}

	result := ""
	for !stack.IsEmpty() {
		top := stack.Pop()
		if top != -1 {
			result = string(int32(top)) + result
		}
	}

	return result
}
