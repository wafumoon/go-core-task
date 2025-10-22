package main

import (
	"testing"
)

func TestSliceExample(t *testing.T) {
	tests := []struct {
		name     string
		input    []int
		expected []int
	}{
		{
			name:     "only even numbers",
			input:    []int{1, 2, 3, 4, 5, 6},
			expected: []int{2, 4, 6},
		},
		{
			name:     "empty slice",
			input:    []int{},
			expected: []int{},
		},
		{
			name:     "only odd numbers",
			input:    []int{1, 3, 5, 7},
			expected: []int{},
		},
		{
			name:     "mixed numbers",
			input:    []int{10, 15, 20, 25},
			expected: []int{10, 20},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := sliceExample(tt.input)
			if !equalSlices(result, tt.expected) {
				t.Errorf("sliceExample(%v) = %v, expected %v", tt.input, result, tt.expected)
			}
		})
	}
}

func TestAddElements(t *testing.T) {
	tests := []struct {
		name     string
		input    []int
		value    int
		expected []int
	}{
		{
			name:     "add to empty slice",
			input:    []int{},
			value:    5,
			expected: []int{5},
		},
		{
			name:     "add to existing slice",
			input:    []int{1, 2, 3},
			value:    4,
			expected: []int{1, 2, 3, 4},
		},
		{
			name:     "add zero",
			input:    []int{10, 20},
			value:    0,
			expected: []int{10, 20, 0},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := addElements(tt.input, tt.value)
			if !equalSlices(result, tt.expected) {
				t.Errorf("addElements(%v, %d) = %v, expected %v", tt.input, tt.value, result, tt.expected)
			}
		})
	}
}

func TestCopySlice(t *testing.T) {
	tests := []struct {
		name  string
		input []int
	}{
		{
			name:  "normal slice",
			input: []int{1, 2, 3, 4, 5},
		},
		{
			name:  "empty slice",
			input: []int{},
		},
		{
			name:  "single element",
			input: []int{42},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			copy := copySlice(tt.input)

			if !equalSlices(copy, tt.input) {
				t.Errorf("copySlice(%v) = %v, expected %v", tt.input, copy, tt.input)
			}

			if len(tt.input) > 0 {

				original := append(tt.input, 999)

				if len(copy) != len(tt.input) {
					t.Errorf("copy was modified when original changed")
				}
				tt.input = original[:len(original)-1]
			}
		})
	}
}

func TestRemoveElement(t *testing.T) {
	tests := []struct {
		name     string
		input    []int
		index    int
		expected []int
	}{
		{
			name:     "remove middle element",
			input:    []int{1, 2, 3, 4, 5},
			index:    2,
			expected: []int{1, 2, 4, 5},
		},
		{
			name:     "remove first element",
			input:    []int{10, 20, 30},
			index:    0,
			expected: []int{20, 30},
		},
		{
			name:     "remove last element",
			input:    []int{1, 2, 3},
			index:    2,
			expected: []int{1, 2},
		},
		{
			name:     "remove from single element",
			input:    []int{100},
			index:    0,
			expected: []int{},
		},
		{
			name:     "index out of bounds - adjusted to last",
			input:    []int{1, 2, 3},
			index:    5,
			expected: []int{1, 2},
		},
		{
			name:     "negative index - adjusted to last",
			input:    []int{1, 2, 3, 4},
			index:    -1,
			expected: []int{1, 2, 3},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			original := make([]int, len(tt.input))
			copy(original, tt.input)

			result := removeElement(tt.input, tt.index)

			if !equalSlices(result, tt.expected) {
				t.Errorf("removeElement(%v, %d) = %v, expected %v", tt.input, tt.index, result, tt.expected)
			}

			if !equalSlices(tt.input, original) {
				t.Errorf("original slice was modified: got %v, expected %v", tt.input, original)
			}
		})
	}
}

func TestIntegration(t *testing.T) {
	original := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	// 2.2 - фильтруем четные
	evens := sliceExample(original)
	if !equalSlices(evens, []int{2, 4, 6, 8, 10}) {
		t.Errorf("sliceExample failed: got %v", evens)
	}

	// 2.3 - добавляем элемент
	withAdded := addElements(original, 25)
	if !equalSlices(withAdded, []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 25}) {
		t.Errorf("addElements failed: got %v", withAdded)
	}

	// 2.4 - копируем
	copy := copySlice(original)
	if !equalSlices(copy, original) {
		t.Errorf("copySlice failed: got %v, expected %v", copy, original)
	}

	// 2.5 - удаляем элемент
	removed := removeElement(original, 4)
	if !equalSlices(removed, []int{1, 2, 3, 4, 6, 7, 8, 9, 10}) {
		t.Errorf("removeElement failed: got %v", removed)
	}
}

func equalSlices(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}
