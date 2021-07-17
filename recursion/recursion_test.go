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

func TestRecursiveBinarySearch(t *testing.T) {
	tests := []struct {
		input    []int
		expected int
		item     int
	}{
		{[]int{1, 3, 4}, 0, 1},
		{[]int{0, 1, 2, 5, 8, 9, 10}, 6, 10},
		{[]int{3, 4, 5}, -1, 7},
	}

	for _, tt := range tests {
		binary := BinarySearch(tt.input, tt.item)
		pos := binary()
		if pos != tt.expected {
			t.Fatalf("Expected and resulting position doesn't match want=%d, got=%d", tt.expected, pos)
		}
	}
}
