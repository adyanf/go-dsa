package stacks_test

import (
	"testing"

	"go-dsa/stacks"

	"github.com/stretchr/testify/assert"
)

func TestRemoveAdjacentDuplicates(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  string
	}{
		{
			name:  "Remove Adjacent Duplicates 1",
			input: "g",
			want:  "g",
		},
		{
			name:  "Remove Adjacent Duplicates 2",
			input: "ggaabcdeb",
			want:  "bcdeb",
		},
		{
			name:  "Remove Adjacent Duplicates 3",
			input: "abbddaccaaabcd",
			want:  "abcd",
		},
		{
			name:  "Remove Adjacent Duplicates 4",
			input: "aannkwwwkkkwna",
			want:  "kwkwna",
		},
		{
			name:  "Remove Adjacent Duplicates 5",
			input: "abbabccblkklu",
			want:  "u",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := stacks.RemoveAdjacentDuplicates(test.input)

			assert.Equal(t, test.want, result)
		})
	}
}
