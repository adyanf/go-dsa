package knowing_what_to_track_test

import (
	"testing"

	"go-dsa/knowing_what_to_track"

	"github.com/stretchr/testify/assert"
)

func TestPermutePalindrome(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  bool
	}{
		{
			name:  "Palindrome 1",
			input: "abab",
			want:  true,
		},
		{
			name:  "Palindrome 2",
			input: "racecar",
			want:  true,
		},
		{
			name:  "Palindrome 3",
			input: "baefeab",
			want:  true,
		},
		{
			name:  "Not Palindrome 1",
			input: "peas",
			want:  false,
		},
		{
			name:  "Not Palindrome 2",
			input: "code",
			want:  false,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := knowing_what_to_track.PermutePalindrome(test.input)
			assert.Equal(t, test.want, result)
		})
	}
}

func TestIsAnagram(t *testing.T) {
	tests := []struct {
		name string
		s1   string
		s2   string
		want bool
	}{
		{
			name: "Anagram 1",
			s1:   "anagram",
			s2:   "nagaram",
			want: true,
		},
		{
			name: "Anagram 2",
			s1:   "spare",
			s2:   "pears",
			want: true,
		},
		{
			name: "Anagram 3",
			s1:   "heart",
			s2:   "earth",
			want: true,
		},
		{
			name: "Anagram 4",
			s1:   "triangle",
			s2:   "integral",
			want: true,
		},
		{
			name: "Not Anagram 1",
			s1:   "super",
			s2:   "upper",
			want: false,
		},
		{
			name: "Not Anagram 2",
			s1:   "super",
			s2:   "esupra",
			want: false,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := knowing_what_to_track.IsAnagram(test.s1, test.s2)
			assert.Equal(t, test.want, result)
		})
	}
}
