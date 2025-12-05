package godash

import (
	"fmt"
	"testing"
)

func TestFindAll_IntSlice_EvenNumbers(t *testing.T) {
	t.Parallel()
	input := []int{1, 2, 3, 4, 5, 6}
	expected := []int{2, 4, 6}

	result, err := FindAll(input, func(item int, index int, slice []int) (bool, error) {
		return item%2 == 0, nil
	})

	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if !equalIntSlices(result, expected) {
		t.Errorf("expected %v, got %v", expected, result)
	}
}

func TestFindAll_StringSlice_LengthGreaterThanThree(t *testing.T) {
	t.Parallel()
	input := []string{"go", "dash", "test", "a", "example"}
	expected := []string{"dash", "test", "example"}

	result, err := FindAll(input, func(item string, index int, slice []string) (bool, error) {
		return len(item) > 3, nil
	})

	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if !equalStringSlices(result, expected) {
		t.Errorf("expected %v, got %v", expected, result)
	}
}

func TestFindAll_EmptySlice(t *testing.T) {
	t.Parallel()
	input := []int{}
	expected := []int{}

	result, err := FindAll(input, func(item int, index int, slice []int) (bool, error) {
		return item > 0, nil
	})

	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if !equalIntSlices(result, expected) {
		t.Errorf("expected %v, got %v", expected, result)
	}
}

func TestFindAll_ErrorFromPredicate(t *testing.T) {
	t.Parallel()
	input := []int{1, 2, 3}
	testErr := fmt.Errorf("predicate error")

	_, err := FindAll(input, func(item int, index int, slice []int) (bool, error) {
		if item == 2 {
			return false, testErr
		}
		return true, nil
	})

	if err == nil {
		t.Fatal("expected error, got nil")
	}
	if err != testErr {
		t.Errorf("expected error %v, got %v", testErr, err)
	}
}

// Helper functions for slice comparison
func equalIntSlices(a, b []int) bool {
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

func equalStringSlices(a, b []string) bool {
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
