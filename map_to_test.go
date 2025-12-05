package godash

import (
	"errors"
	"fmt"
	"reflect"
	"testing"
)

func TestMapTo(t *testing.T) {
	tests := []struct {
		name          string
		collection    []int
		mapFunction   ElementIteratorFn[int, string]
		expected      []string
		expectedError bool
	}{
		{
			name:       "successfully transform integers to strings",
			collection: []int{1, 2, 3},
			mapFunction: func(element int) (string, error) {
				return fmt.Sprintf("num_%d", element), nil
			},
			expected:      []string{"num_1", "num_2", "num_3"},
			expectedError: false,
		},
		{
			name:       "empty collection",
			collection: []int{},
			mapFunction: func(element int) (string, error) {
				return fmt.Sprintf("%d", element), nil
			},
			expected:      []string{},
			expectedError: false,
		},
		{
			name:       "transformation returns error",
			collection: []int{1, 2, 3},
			mapFunction: func(element int) (string, error) {
				if element == 2 {
					return "", errors.New("transformation failed")
				}
				return fmt.Sprintf("%d", element), nil
			},
			expected:      []string{"1", "", ""},
			expectedError: true,
		},
		{
			name:       "single element",
			collection: []int{42},
			mapFunction: func(element int) (string, error) {
				return fmt.Sprintf("value: %d", element), nil
			},
			expected:      []string{"value: 42"},
			expectedError: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := MapTo(tt.collection, tt.mapFunction)

			if (err != nil) != tt.expectedError {
				t.Errorf("expected error: %v, got: %v", tt.expectedError, err != nil)
			}

			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("expected %v, got %v", tt.expected, result)
			}
		})
	}
}
