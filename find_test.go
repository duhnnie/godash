package godash

import (
	"testing"
)

func TestFind(t *testing.T) {
	type person struct {
		FirstName string
		LastName  string
	}

	persons := []person{
		{"Kurt", "Cobain"},
		{"Dave", "Grohl"},
		{"Krist", "Novoselic"},
	}

	tests := []struct {
		name     string
		slice    []person
		fn       IteratorFn[person, bool]
		expected person
		found    bool
	}{
		{
			name:  "Find existing person",
			slice: persons,
			fn: func(p person, index int, collection []person) bool {
				return p.FirstName == "Dave"
			},
			expected: person{"Dave", "Grohl"},
			found:    true,
		},
		{
			name:  "Find non-existing person",
			slice: persons,
			fn: func(p person, index int, collection []person) bool {
				return p.FirstName == "John"
			},
			expected: person{},
			found:    false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, found := Find(tt.slice, tt.fn)
			if found != tt.found || result != tt.expected {
				t.Fatalf("unexpected result for Find(): got (%v, %v) | expecting (%v, %v)", result, found, tt.expected, tt.found)
			}
		})
	}
}
