package godash

import (
	"fmt"
	"testing"
)

func TestFindFirst_IntSlice_FindEvenNumber(t *testing.T) {
	input := []int{1, 3, 4, 5, 7}
	result, found, err := FindFirst(input, func(item int) (bool, error) {
		return item%2 == 0, nil
	})

	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if !found {
		t.Errorf("expected found=true, got false")
	}
	if result != 4 {
		t.Errorf("expected 4, got %d", result)
	}
}

func TestFindFirst_IntSlice_NoMatch(t *testing.T) {
	input := []int{1, 3, 5, 7, 9}
	result, found, err := FindFirst(input, func(item int) (bool, error) {
		return item%2 == 0, nil
	})

	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if found {
		t.Errorf("expected found=false, got true")
	}
	if result != 0 {
		t.Errorf("expected 0, got %d", result)
	}
}

func TestFindFirst_StringSlice_LengthGreaterThanThree(t *testing.T) {
	input := []string{"go", "a", "ok", "dashboard"}
	result, found, err := FindFirst(input, func(item string) (bool, error) {
		return len(item) > 3, nil
	})

	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if !found {
		t.Errorf("expected found=true, got false")
	}
	if result != "dashboard" {
		t.Errorf("expected dashboard, got %s", result)
	}
}

func TestFindFirst_EmptySlice(t *testing.T) {
	input := []int{}
	result, found, err := FindFirst(input, func(item int) (bool, error) {
		return item > 0, nil
	})

	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if found {
		t.Errorf("expected found=false, got true")
	}
	if result != 0 {
		t.Errorf("expected 0, got %d", result)
	}
}

func TestFindFirst_ErrorFromIterator(t *testing.T) {
	input := []int{1, 2, 3}
	testErr := fmt.Errorf("iterator error")

	_, _, err := FindFirst(input, func(item int) (bool, error) {
		if item == 2 {
			return false, testErr
		}
		return false, nil
	})

	if err == nil {
		t.Fatal("expected error, got nil")
	}
	if err != testErr {
		t.Errorf("expected error %v, got %v", testErr, err)
	}
}

func TestFindFirst_EarlyTermination(t *testing.T) {
	input := []int{1, 2, 3, 4, 5}
	callCount := 0

	result, found, err := FindFirst(input, func(item int) (bool, error) {
		callCount++
		return item == 2, nil
	})

	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if !found {
		t.Errorf("expected found=true, got false")
	}
	if result != 2 {
		t.Errorf("expected 2, got %d", result)
	}
	if callCount != 2 {
		t.Errorf("expected 2 calls, got %d", callCount)
	}
}
