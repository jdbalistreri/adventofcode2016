package main

import (
	"testing"
)

func TestParseLine(t *testing.T) {
	_, sector, checksum := parseLine("aaaaa-bbb-z-y-x-123[abxyz]")
	if sector != 123 {
		t.Errorf("incorrect sector %d", sector)
	}
	if checksum != "abxyz" {
		t.Errorf("incorrect checksum %s", checksum)
	}
}

func TestMakeChecksum(t *testing.T) {
	checksum := makeChecksum([]string{"aaaaa", "bbb", "z", "y", "x"})
	if checksum != "abxyz" {
		t.Errorf("incorrect checksum %s", checksum)
	}
}

var fibTests = []struct {
	n        int // input
	expected int // expected result
}{
	{1, 1},
	{2, 1},
	{3, 2},
	{4, 3},
	{5, 5},
	{6, 8},
	{7, 13},
}

var checksumTests = []struct {
	line     string
	expected bool
}{
	{"aaaaa-bbb-z-y-x-123[abxyz]", true},
	{"a-b-c-d-e-f-g-h-987[abcde]", true},
	{"not-a-real-room-404[oarel]", true},
	{"totally-real-room-200[decoy]", false},
}

func TestIsValidLine(t *testing.T) {
	for _, tt := range checksumTests {
		_, actual := isValidLine(tt.line)
		if actual != tt.expected {
			t.Errorf("isValidLine(%s): %t", tt.line, actual)
		}
	}
}

var shiftTests = []struct {
	input    []string
	n        int
	expected string
}{
	{[]string{"abc"}, 1, "bcd "},
	{[]string{"zzz"}, 2, "bbb "},
}

func TestShiftWords(t *testing.T) {
	for _, tt := range shiftTests {
		actual := shiftWords(tt.input, tt.n)
		if actual != tt.expected {
			t.Errorf("shiftWords(%v): expected %s, actual %s", tt.input, tt.expected, actual)
		}
	}
}
