package validator

import (
	"testing"
)

func TestLcFirst(t *testing.T) {
	var tests = []struct {
		input    string
		expected string
	}{
		{"Roshan", "roshan"},
		{"Rana", "rana"},
		{"Test", "test"},
	}
	for _, test := range tests {
		if output := LcFirst(test.input); output != test.expected {
			t.Error("Test Failed: {} inputted, {} expected, received: {}", test.input, test.expected, output)
		}
	}
}

func TestUcFirst(t *testing.T) {
	var tests = []struct {
		input    string
		expected string
	}{
		{"roshan", "Roshan"},
		{"rana", "Rana"},
		{"test", "Test"},
	}
	for _, test := range tests {
		if output := UcFirst(test.input); output != test.expected {
			t.Error("Test Failed: {} inputted, {} expected, received: {}", test.input, test.expected, output)
		}
	}
}

func TestSplit(t *testing.T) {
	var tests = []struct {
		input    string
		expected string
	}{
		{"myName", "My name"},
		{"testDriven", "Test driven"},
		{"myOutput", "My output"},
	}
	for _, test := range tests {
		if output := Split(test.input); output != test.expected {
			t.Error("Test Failed: {} inputted, {} expected, received: {}", test.input, test.expected, output)
		}
	}
}
