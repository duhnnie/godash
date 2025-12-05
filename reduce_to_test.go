package godash

import (
	"fmt"
	"testing"
)

func TestReduceTo_IntSlice_Sum(t *testing.T) {
	input := []int{1, 2, 3, 4, 5}
	result, err := ReduceTo(input, func(acc int, current int) (int, error) {
		return acc + current, nil
	}, 0)

	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if result != 15 {
		t.Errorf("expected 15, got %d", result)
	}
}

func TestReduceTo_IntSlice_Product(t *testing.T) {
	input := []int{1, 2, 3, 4}
	result, err := ReduceTo(input, func(acc int, current int) (int, error) {
		return acc * current, nil
	}, 1)

	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if result != 24 {
		t.Errorf("expected 24, got %d", result)
	}
}

func TestReduceTo_StringSlice_Concatenate(t *testing.T) {
	input := []string{"Hello", " ", "World"}
	result, err := ReduceTo(input, func(acc string, current string) (string, error) {
		return acc + current, nil
	}, "")

	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if result != "Hello World" {
		t.Errorf("expected 'Hello World', got %q", result)
	}
}

func TestReduceTo_IntSlice_ToStringConversion(t *testing.T) {
	input := []int{1, 2, 3}
	result, err := ReduceTo(input, func(acc string, current int) (string, error) {
		return acc + fmt.Sprintf("%d,", current), nil
	}, "")

	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if result != "1,2,3," {
		t.Errorf("expected '1,2,3,', got %q", result)
	}
}

func TestReduceTo_EmptySlice(t *testing.T) {
	input := []int{}
	result, err := ReduceTo(input, func(acc int, current int) (int, error) {
		return acc + current, nil
	}, 42)

	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if result != 42 {
		t.Errorf("expected 42, got %d", result)
	}
}

func TestReduceTo_ErrorFromReducer(t *testing.T) {
	input := []int{1, 2, 3}
	testErr := fmt.Errorf("reducer error")

	_, err := ReduceTo(input, func(acc int, current int) (int, error) {
		if current == 2 {
			return 0, testErr
		}
		return acc + current, nil
	}, 0)

	if err == nil {
		t.Fatal("expected error, got nil")
	}
	if err != testErr {
		t.Errorf("expected error %v, got %v", testErr, err)
	}
}

func TestReduceTo_SingleElement(t *testing.T) {
	input := []int{5}
	result, err := ReduceTo(input, func(acc int, current int) (int, error) {
		return acc + current, nil
	}, 10)

	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if result != 15 {
		t.Errorf("expected 15, got %d", result)
	}
}
