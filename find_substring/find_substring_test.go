package main

import (
	"testing"
)

func TestFindSubstring(t *testing.T) {
	tests := []struct {
		str    string
		substr string
		found  bool
	}{
		{"Hello, World!", "World", true},
		{"Hello, World!", "Hello", true},
		{"Hello, World!", "Hello, Universe!", false},
		{"", "Hello", false},
		{"Hello, World!", "", true}, // Substring is an empty string
		{"Hello, World!", "!", true}, // Substring is a single character
	}

	for _, test := range tests {
		found := findSubstring(test.str, test.substr)
		if found != test.found {
			t.Errorf("For string '%s' and substring '%s', expected %v but got %v", test.str, test.substr, test.found, found)
		}
	}
}
