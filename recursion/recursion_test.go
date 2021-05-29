package recursion

import "testing"

func TestFactorial(t *testing.T) {
	tests := []struct {
		input    int
		expected int
	}{
		{input: 5, expected: 120},
		{input: 4, expected: 24},
		{input: 3, expected: 6},
	}

	for _, tt := range tests {
		result := Factorial(tt.input)
		if result != tt.expected {
			t.Errorf("wrong result want=%d, got=%d", tt.expected, result)
		}
	}
}
