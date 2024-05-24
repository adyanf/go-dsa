package two_pointers_test

import (
	"testing"

	"go-dsa/two_pointers"

	"github.com/stretchr/testify/assert"
)

func TestIsPalindrome(t *testing.T) {
	tests := []struct {
		name        string
		inputString string
		want        bool
	}{
		{
			name:        "Palindrome 1",
			inputString: "kaYak",
			want:        true,
		},
		{
			name:        "Palindrome 2",
			inputString: "A",
			want:        true,
		},
		{
			name:        "Not Palindrome 1",
			inputString: "hello",
			want:        false,
		},
		{
			name:        "Not Palindrome 2",
			inputString: "RaCEACAR",
			want:        false,
		},
		{
			name:        "Not Palindrome 3",
			inputString: "ABCDABCD",
			want:        false,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := two_pointers.IsPalindrome(test.inputString)

			assert.Equal(t, test.want, result)
		})
	}
}

func TestSumOfThreeValues(t *testing.T) {
	tests := []struct {
		name   string
		nums   []int
		target int
		want   bool
	}{
		{
			name:   "Target achieved 1",
			nums:   []int{3, 7, 1, 2, 8, 4, 5},
			target: 10,
			want:   true,
		},
		{
			name:   "Target achieved 2",
			nums:   []int{-1, 2, 1, -4, 5, -3},
			target: -8,
			want:   true,
		},
		{
			name:   "Target achieved 3",
			nums:   []int{-1, 2, 1, -4, 5, -3},
			target: 0,
			want:   true,
		},
		{
			name:   "Target not achieved 1",
			nums:   []int{1, 0, -1},
			target: -1,
			want:   false,
		},
		{
			name:   "Target not achieved 2",
			nums:   []int{3, 7, 1, 2, 8, 4, 5},
			target: 21,
			want:   false,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := two_pointers.SumOfThreeValues(test.nums, test.target)

			assert.Equal(t, test.want, result)
		})
	}
}
