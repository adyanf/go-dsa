package fast_and_slow_pointers_test

import (
	"testing"

	"go-dsa/fast_and_slow_pointers"
	"go-dsa/structs"
	"go-dsa/testutil"

	"github.com/stretchr/testify/assert"
)

func TestIsHappy(t *testing.T) {
	tests := []struct {
		name  string
		input int
		want  bool
	}{
		{
			name:  "Happy Number 1",
			input: 1,
			want:  true,
		},
		{
			name:  "Happy Number 2",
			input: 19,
			want:  true,
		},
		{
			name:  "Happy Number 3",
			input: 7,
			want:  true,
		},
		{
			name:  "Not Happy Number 1",
			input: 2147483646,
			want:  false,
		},
		{
			name:  "Not Happy Number 2",
			input: 8,
			want:  false,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := fast_and_slow_pointers.IsHappy(test.input)

			assert.Equal(t, test.want, result)
		})
	}
}

func TestDetectCycle(t *testing.T) {
	tests := []struct {
		name       string
		linkedList *structs.LinkedList
		want       bool
	}{
		{
			name:       "Cycle 1",
			linkedList: testutil.CreateCycleLinkedList([]int{2, 4, 6, 8, 10}, 2),
			want:       true,
		},
		{
			name:       "Cycle 2",
			linkedList: testutil.CreateCycleLinkedList([]int{1, 2, 3, 4, 5}, 3),
			want:       true,
		},
		{
			name:       "Cycle 3",
			linkedList: testutil.CreateCycleLinkedList([]int{3, 6, 9, 10, 11}, 0),
			want:       true,
		},
		{
			name:       "Not Cycle 1",
			linkedList: testutil.CreateCycleLinkedList([]int{1, 3, 5, 7, 9}, -1),
			want:       false,
		},
		{
			name:       "Not Cycle 2",
			linkedList: testutil.CreateCycleLinkedList([]int{0, 2, 3, 5, 6}, -1),
			want:       false,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := fast_and_slow_pointers.DetectCycle(test.linkedList)

			assert.Equal(t, test.want, result)
		})
	}
}
