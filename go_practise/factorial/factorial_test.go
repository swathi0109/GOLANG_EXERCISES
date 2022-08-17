package main

import (
	"testing"
)

func TestFactorial(t *testing.T) {
	if factorial(3) != 6 {
		t.Error("Expected factorial(3) to be 6")
	}
}

func TestTableFactorial(t *testing.T) {
	tests := []struct {
		input    int32
		expected int32
	}{
		{3, 6},
		{4, 24},
		{6, 720},
	}
	for _, tc := range tests {
		r := factorial(tc.input)
		if r != tc.expected {
			t.Errorf("Factorial of %d is wrong", tc.input)
		}
	}

}
