package linked_list_test

import (
	"testing"

	"go-dsa/linked_list"
	"go-dsa/structs"

	"github.com/stretchr/testify/assert"
)

func TestReverse(t *testing.T) {
	tests := []struct {
		name  string
		input *structs.LinkedListNode
		want  *structs.LinkedListNode
	}{
		{
			name:  "Reverse 1",
			input: structs.CreateLinkedList([]int{1, -2, 3, 4, -5, 4, 3, -2, 1}).Head,
			want:  structs.CreateLinkedList([]int{1, -2, 3, 4, -5, 4, 3, -2, 1}).Head,
		},
		{
			name:  "Reverse 2",
			input: structs.CreateLinkedList([]int{-1, -5, -3, -7, -8, -6, -2}).Head,
			want:  structs.CreateLinkedList([]int{-2, -6, -8, -7, -3, -5, -1}).Head,
		},
		{
			name:  "Reverse 3",
			input: structs.CreateLinkedList([]int{-1, 2, -3, 4}).Head,
			want:  structs.CreateLinkedList([]int{4, -3, 2, -1}).Head,
		},
		{
			name:  "Reverse 4",
			input: structs.CreateLinkedList([]int{1, -1, -2, 3, -4, 5}).Head,
			want:  structs.CreateLinkedList([]int{5, -4, 3, -2, -1, 1}).Head,
		},
		{
			name:  "Reverse 5",
			input: structs.CreateLinkedList([]int{28, 21, 14, 7}).Head,
			want:  structs.CreateLinkedList([]int{7, 14, 21, 28}).Head,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := linked_list.Reverse(test.input)
			assert.Equal(t, test.want.String(), result.String())
		})
	}
}

func TestReverseBetween(t *testing.T) {
	tests := []struct {
		name  string
		input *structs.LinkedListNode
		left  int
		right int
		want  *structs.LinkedListNode
	}{
		{
			name:  "Reverse Between 1",
			input: structs.CreateLinkedList([]int{6, 8, 7}).Head,
			left:  1,
			right: 2,
			want:  structs.CreateLinkedList([]int{8, 6, 7}).Head,
		},
		{
			name:  "Reverse Between 2",
			input: structs.CreateLinkedList([]int{9, 0, 8, 2}).Head,
			left:  2,
			right: 4,
			want:  structs.CreateLinkedList([]int{9, 2, 8, 0}).Head,
		},
		{
			name:  "Reverse Between 3",
			input: structs.CreateLinkedList([]int{1, 2, 3, 4, 5}).Head,
			left:  1,
			right: 5,
			want:  structs.CreateLinkedList([]int{5, 4, 3, 2, 1}).Head,
		},
		{
			name:  "Reverse Between 4",
			input: structs.CreateLinkedList([]int{7, 4, 6, 1, 5, 8}).Head,
			left:  2,
			right: 5,
			want:  structs.CreateLinkedList([]int{7, 5, 1, 6, 4, 8}).Head,
		},
		{
			name:  "Reverse Between 5",
			input: structs.CreateLinkedList([]int{0, 8, 3, 1, 9, 2, 7}).Head,
			left:  1,
			right: 6,
			want:  structs.CreateLinkedList([]int{2, 9, 1, 3, 8, 0, 7}).Head,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := linked_list.ReverseBetween(test.input, test.left, test.right)
			assert.Equal(t, test.want.String(), result.String())
		})
	}
}
