package godash

import (
	"errors"
	"testing"
)

func TestEvery(t *testing.T) {
	type args struct {
		slice    []int
		iterator IteratorFn[int, bool]
	}
	tests := []struct {
		name    string
		args    args
		want    bool
		wantErr bool
	}{
		{
			name: "all even numbers",
			args: args{
				slice: []int{2, 4, 6, 8},
				iterator: func(num int, idx int, s []int) (bool, error) {
					return num%2 == 0, nil
				},
			},
			want:    true,
			wantErr: false,
		},
		{
			name: "not all even numbers",
			args: args{
				slice: []int{2, 3, 6, 8},
				iterator: func(num int, idx int, s []int) (bool, error) {
					return num%2 == 0, nil
				},
			},
			want:    false,
			wantErr: false,
		},
		{
			name: "empty slice",
			args: args{
				slice: []int{},
				iterator: func(num int, idx int, s []int) (bool, error) {
					return num%2 == 0, nil
				},
			},
			want:    true,
			wantErr: false,
		},
		{
			name: "error in iterator",
			args: args{
				slice: []int{2, 4, 6, 8},
				iterator: func(num int, idx int, s []int) (bool, error) {
					if num == 6 {
						return false, errors.New("test error")
					}
					return true, nil
				},
			},
			want:    false,
			wantErr: true,
		},
		{
			name: "all positive numbers",
			args: args{
				slice: []int{1, 2, 3, 4},
				iterator: func(num int, idx int, s []int) (bool, error) {
					return num > 0, nil
				},
			},
			want:    true,
			wantErr: false,
		},
		{
			name: "contains negative number",
			args: args{
				slice: []int{1, -2, 3, 4},
				iterator: func(num int, idx int, s []int) (bool, error) {
					return num > 0, nil
				},
			},
			want:    false,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			got, err := Every(tt.args.slice, tt.args.iterator)
			if (err != nil) != tt.wantErr {
				t.Errorf("Every() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Every() = %v, want %v", got, tt.want)
			}
		})
	}
}
