package godash

import (
	"fmt"
	"testing"
)

func TestFind_IntSlice(t *testing.T) {
	slice := []int{1, 2, 3, 4, 5}
	fn := func(item int, idx int, s []int) (bool, error) {
		return item%2 == 0, nil
	}
	result, found, err := Find(slice, fn)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if !found || result != 2 {
		t.Errorf("expected 2, true; got %v, %v", result, found)
	}
}

func TestFind_StringSlice(t *testing.T) {
	slice := []string{"apple", "banana", "cherry"}
	fn := func(item string, idx int, s []string) (bool, error) {
		return item == "banana", nil
	}
	result, found, err := Find(slice, fn)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if !found || result != "banana" {
		t.Errorf("expected 'banana', true; got %v, %v", result, found)
	}
}

func TestFind_NoMatch(t *testing.T) {
	slice := []int{1, 3, 5}
	fn := func(item int, idx int, s []int) (bool, error) {
		return item%2 == 0, nil
	}
	result, found, err := Find(slice, fn)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if found || result != 0 {
		t.Errorf("expected 0, false; got %v, %v", result, found)
	}
}

func TestFind_EmptySlice(t *testing.T) {
	slice := []int{}
	fn := func(item int, idx int, s []int) (bool, error) {
		return true, nil
	}
	result, found, err := Find(slice, fn)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if found || result != 0 {
		t.Errorf("expected 0, false; got %v, %v", result, found)
	}
}

func TestFind_ErrorInFn(t *testing.T) {
	slice := []int{1, 2, 3}
	expectedErr := fmt.Errorf("test error")
	fn := func(item int, idx int, s []int) (bool, error) {
		if item == 2 {
			return false, expectedErr
		}
		return false, nil
	}
	result, found, err := Find(slice, fn)
	if err != expectedErr {
		t.Errorf("expected error %v, got %v", expectedErr, err)
	}
	if found || result != 0 {
		t.Errorf("expected 0, false; got %v, %v", result, found)
	}
}
