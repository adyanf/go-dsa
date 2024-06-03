package binary_search_test

import (
	"testing"

	"go-dsa/binary_search"

	"github.com/stretchr/testify/assert"
)

func TestBinarySearch(t *testing.T) {
	tests := []struct {
		name   string
		nums   []int
		target int
		want   int
	}{
		{
			name:   "Binary Search 1",
			nums:   []int{1, 6, 8, 10},
			target: 1,
			want:   0,
		},
		{
			name:   "Binary Search 2",
			nums:   []int{11, 22, 33, 44, 55, 66, 77},
			target: 33,
			want:   2,
		},
		{
			name:   "Binary Search 3",
			nums:   []int{-3, -1, 0, 11, 15},
			target: 0,
			want:   2,
		},
		{
			name:   "Binary Search 4",
			nums:   []int{-30, -27, -8, -6, -1},
			target: -1,
			want:   4,
		},
		{
			name:   "Binary Search 5",
			nums:   []int{1, 2, 3, 4, 5},
			target: 0,
			want:   -1,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := binary_search.BinarySearch(test.nums, test.target)
			assert.Equal(t, test.want, result)
		})
	}
}

func TestBinarySearchRotated(t *testing.T) {
	tests := []struct {
		name   string
		nums   []int
		target int
		want   int
	}{
		{
			name:   "Binary Search Rotated 1",
			nums:   []int{6, 7, 1, 2, 3, 4, 5},
			target: 3,
			want:   4,
		},
		{
			name:   "Binary Search Rotated 2",
			nums:   []int{6, 7, 1, 2, 3, 4, 5},
			target: 1,
			want:   2,
		},
		{
			name:   "Binary Search Rotated 3",
			nums:   []int{4, 5, 6, 1, 2, 3},
			target: 3,
			want:   5,
		},
		{
			name:   "Binary Search Rotated 4",
			nums:   []int{4, 5, 6, 1, 2, 3},
			target: 6,
			want:   2,
		},
		{
			name:   "Binary Search Rotated 5",
			nums:   []int{4},
			target: 1,
			want:   -1,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := binary_search.BinarySearchRotated(test.nums, test.target)
			assert.Equal(t, test.want, result)
		})
	}
}
