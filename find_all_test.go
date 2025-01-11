package godash

import (
	"testing"
)

func TestFindAll(t *testing.T) {
	type args struct {
		slice []int
		fn    IteratorFn[int, bool]
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "find all even numbers",
			args: args{
				slice: []int{1, 2, 3, 4, 5, 6},
				fn: func(item int, index int, slice []int) bool {
					return item%2 == 0
				},
			},
			want: []int{2, 4, 6},
		},
		{
			name: "find all numbers greater than 3",
			args: args{
				slice: []int{1, 2, 3, 4, 5, 6},
				fn: func(item int, index int, slice []int) bool {
					return item > 3
				},
			},
			want: []int{4, 5, 6},
		},
		{
			name: "find all negative numbers",
			args: args{
				slice: []int{-1, -2, 3, 4, -5, 6},
				fn: func(item int, index int, slice []int) bool {
					return item < 0
				},
			},
			want: []int{-1, -2, -5},
		},
		{
			name: "find non existent number",
			args: args{
				slice: []int{-1, -2, 3, 4, -5, 6},
				fn: func(item int, index int, slice []int) bool {
					return item > 100
				},
			},
			want: []int{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FindAll(tt.args.slice, tt.args.fn); !equal(got, tt.want) {
				t.Errorf("FindAll() = %v, want %v", got, tt.want)
			}
		})
	}
}

func equal(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i, v := range a {
		if v != b[i] {
			return false
		}
	}
	return true
}
