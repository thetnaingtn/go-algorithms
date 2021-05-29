package sort

import (
	"testing"
)

func TestSelectionSort(t *testing.T) {
	tests := []struct {
		input    []int
		expected []int
	}{
		{
			input:    []int{6, 1, 2},
			expected: []int{1, 2, 6},
		},
		{
			input:    []int{6, 1, 1, 2},
			expected: []int{1, 1, 2, 6},
		},
		{
			input:    []int{9, 10, -1, 5, -2, 3},
			expected: []int{-2, -1, 3, 5, 9, 10},
		},
	}

	for _, tt := range tests {
		sorted := SelectionSort(tt.input)

		for i, item := range sorted {
			if item != tt.expected[i] {
				t.Errorf("actual and expected item doesn't match at index %d, got=%d, want=%d", i, item, tt.expected[i])
			}
		}
	}
}
