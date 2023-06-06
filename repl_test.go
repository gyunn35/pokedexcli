package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCleanInput(t *testing.T) {
	tests := []struct {
		input    string
		expected []string
	}{
		{
			input:    "hello world",
			expected: []string{"hello", "world"},
		},
		{
			input:    "my nickname is gyunn35",
			expected: []string{"my", "nickname", "is", "gyunn35"},
		},
	}

	for _, test := range tests {
		actual := cleanInput(test.input)
		if len(actual) != len(test.expected) {
			assert.Equal(t, len(test.expected), len(actual))
			continue
		}
		for i := range actual {
			assert.Equal(t, test.expected[i], actual[i])
		}
	}
}
