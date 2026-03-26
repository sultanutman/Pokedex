package main

import (
	"testing"
)

func TestCleanInput(t *testing.T) {
	testCases := []struct {
		input    string
		expected []string
	}{
		{
			input:    "yes no",
			expected: []string{"yes", "no"},
		},
		{
			input:    "Man WoMaN",
			expected: []string{"man", "woman"},
		},
		{
			input:    "  hello  world  ",
			expected: []string{"hello", "world"},
		},
	}

	for _, testCase := range testCases {
		actual := cleanInput(testCase.input)
		if len(actual) != len(testCase.expected) {
			t.Errorf(`
			Failed, length of actual and expected result inequal
			input: %s
			actual result: %v
			expected result: %v
			`, testCase.input, actual, testCase.expected)
		}

		for i := range actual {
			word := actual[i]
			expected_word := testCase.expected[i]

			if len(word) != len(expected_word) {
				t.Errorf(`
				Failed, length of words at index %d does not match
				input: %s
				actual result: %v
				expected result: %v
				`, i, testCase.input, actual, testCase.expected)
			}

			if word != expected_word {
				t.Errorf(`
				Failed, not the same word at index %d 
				input: %s
				actual result: %v
				expected result: %v
				`, i, testCase.input, actual, testCase.expected)
			}
		}

	}
}
