package main

import (
	"testing"
)

var maxLetterInputs = []struct {
	letters  map[string]int
	expected string
}{
	{map[string]int{"d": 10, "a": 1, "c": 4}, "d"},
}

func TestMaxLetter(t *testing.T) {
	for _, tt := range maxLetterInputs {
		actual := maxLetter(tt.letters)
		if actual != tt.expected {
			t.Errorf("maxLetter(%v): expected %s, actual %s", tt.letters, tt.expected, actual)
		}
	}
}

var minLetterInputs = []struct {
	letters  map[string]int
	expected string
}{
	{map[string]int{"d": 10, "a": 1, "c": 4}, "a"},
}

func TestMinLetter(t *testing.T) {
	for _, tt := range minLetterInputs {
		actual := minLetter(tt.letters)
		if actual != tt.expected {
			t.Errorf("minLetter(%v): expected %s, actual %s", tt.letters, tt.expected, actual)
		}
	}
}
