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
				t.Errorf("Selection Sort: actual and expected item doesn't match at index %d, got=%d, want=%d", i, item, tt.expected[i])
			}
		}
	}
}

func TestQuickSort(t *testing.T) {
	tests := []struct {
		input    []int
		expected []int
	}{
		{input: []int{9, 1, 10, 4, 7}, expected: []int{1, 4, 7, 9, 10}},
		{input: []int{2, 1, 8, 9, 3}, expected: []int{1, 2, 3, 8, 9}},
	}

	for _, tt := range tests {
		sorted := QuickSort(tt.input)
		for i, item := range sorted {
			if item != tt.expected[i] {
				t.Errorf("Quick Sort: actual and expected item doesn't match at index %d, got=%d, want=%d", i, item, tt.expected[i])
			}
		}
	}
}
