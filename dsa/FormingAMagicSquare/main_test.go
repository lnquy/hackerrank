package main

import "testing"

func Test_formingMagicSquare(t *testing.T) {
	type args struct {
		s [][]int32
	}
	tests := []struct {
		name string
		args args
		want int32
	}{
		{"tc1", args{[][]int32{{5, 3, 4}, {1, 5, 8}, {6, 4, 2}}}, 7},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := formingMagicSquare(tt.args.s); got != tt.want {
				t.Errorf("formingMagicSquare() = %v, want %v", got, tt.want)
			}
		})
	}
}
