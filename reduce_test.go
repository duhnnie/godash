package godash

import (
	"fmt"
	"testing"
)

func TestReduce(t *testing.T) {
	persons := []person{
		{"Kurt", "Cobain"},
		{"Dave", "Grohl"},
		{"Krist", "Novoselic"},
	}

	expected := "Kurt Cobain, Dave Grohl, Krist Novoselic."

	var reducerFn ReducerFn[person, string] = func(acc string, element person, index int, collection []person) string {
		if index == len(collection)-1 {
			return fmt.Sprintf("%s%s %s.", acc, element.FirstName, element.LastName)
		} else {
			return fmt.Sprintf("%s%s %s, ", acc, element.FirstName, element.LastName)
		}
	}

	reduced := Reduce(persons, reducerFn, "")

	if reduced != expected {
		t.Fatalf("unexpected output for Reduce(): %s | expecting: %s ", reduced, expected)
	}
}
