package godash

import (
	"errors"
	"testing"
)

func TestEveryoneAllPass(t *testing.T) {
	result, err := Everyone([]int{2, 4, 6, 8}, func(num int) (bool, error) {
		return num%2 == 0, nil
	})

	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	if !result {
		t.Errorf("expected true, got %v", result)
	}
}

func TestEveryoneNotAllPass(t *testing.T) {
	result, err := Everyone([]int{2, 4, 5, 8}, func(num int) (bool, error) {
		return num%2 == 0, nil
	})

	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	if result {
		t.Errorf("expected false, got %v", result)
	}
}

func TestEveryoneEmptySlice(t *testing.T) {
	result, err := Everyone([]int{}, func(num int) (bool, error) {
		return true, nil
	})

	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	if !result {
		t.Errorf("expected true, got %v", result)
	}
}

func TestEveryoneWithError(t *testing.T) {
	_, err := Everyone([]int{1, 2, 3}, func(num int) (bool, error) {
		if num == 2 {
			return false, errors.New("iterator error")
		}
		return true, nil
	})

	if err == nil {
		t.Error("expected error, got nil")
	}
}

func TestEveryoneDifferentTypes(t *testing.T) {
	result, err := Everyone([]string{"a", "b", "c"}, func(s string) (bool, error) {
		return len(s) == 1, nil
	})

	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	if !result {
		t.Errorf("expected true, got %v", result)
	}
}
