package search

import (
	"testing"
)

func TestBinarySearch(t *testing.T) {
	tests := []struct {
		input    []int
		expected int
		item     int
	}{
		{[]int{1, 3, 4}, 2, 4},
		{[]int{3, 4, 5}, -1, 7},
	}

	for _, tt := range tests {
		pos := BinarySearch(tt.input, tt.item)
		if pos != tt.expected {
			t.Fatalf("Expected and resulting position doesn't match want=%d, got=%d", tt.expected, pos)
		}
	}
}
