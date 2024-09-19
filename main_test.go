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
			input:    "Some http://link.com/ and more http://links.net!",
			expected: "Some http://********* and more http://**********",
		},
		{
			input:    "No links here, just text.",
			expected: "No links here, just text.",
		},
		{
			input:    "Check out http://example.com and http://test.net now!",
			expected: "Check out http://*********** and http://******** now!",
		},
		{
			input:    "http://startoftext.com",
			expected: "http://***************",
		},
		{
			input:    "http://onlylink.comHAHAHA",
			expected: "http://******************",
		},
		{
			input:    "",
			expected: "",
		},
		{
			input:    "http://short",
			expected: "http://*****",
		},
		{
			input:    "Mixed http://example.com and regular text",
			expected: "Mixed http://*********** and regular text",
		},
		{
			input:    "Here's my spammy page: hTTp://youth-elixir.com",
			expected: "Here's my spammy page: hTTp://youth-elixir.com",
		},
		{
			input:    "Mixed http:// example.com with space after link",
			expected: "Mixed http:// example.com with space after link",
		},
	}

	for _, test := range tests {
		output := spam(test.input)
		if output != test.expected {
			t.Errorf("spam(%q) = %q; expected %q", test.input, output, test.expected)
		}
	}
}
