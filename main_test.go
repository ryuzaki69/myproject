package main

import (
	"testing"
)

// Функция тестирования
func TestSpam(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{
			input:    "Some https://link.com/ and more https://links.net!",
			expected: "Some https://********* and more https://**********",
		},
		{
			input:    "No links here, just text.",
			expected: "No links here, just text.",
		},
		{
			input:    "Check out https://example.com and https://test.net now!",
			expected: "Check out https://*********** and https://******** now!",
		},
		{
			input:    "https://startoftext.com",
			expected: "https://***************",
		},
		{
			input:    "https://onlylink.comHAHAHA",
			expected: "https://******************",
		},
		{
			input:    "",
			expected: "",
		},
		{
			input:    "https://short",
			expected: "https://*****",
		},
		{
			input:    "Mixed https://example.com and regular text",
			expected: "Mixed https://*********** and regular text",
		},
		{
			input:    "Here's my spammy page: hTTp://youth-elixir.com",
			expected: "Here's my spammy page: hTTp://youth-elixir.com",
		},
		{
			input:    "Mixed https:// example.com with space after link",
			expected: "Mixed https:// example.com with space after link",
		},
	}

	for _, test := range tests {
		output := spam(test.input)
		if output != test.expected {
			t.Errorf("spam(%q) = %q; expected %q", test.input, output, test.expected)
		}
	}
}
