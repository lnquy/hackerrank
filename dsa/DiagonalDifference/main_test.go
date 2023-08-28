package main

import "testing"

func Test_diagonalDifference(t *testing.T) {
	type args struct {
		arr [][]int32
	}
	tests := []struct {
		name string
		args args
		want int32
	}{
		{"tc1", args{[][]int32{{1, 2, 3}, {4, 5, 6}, {9, 8, 9}}}, 2},
		{"tc2", args{[][]int32{{11, 2, 4}, {4, 5, 6}, {10, 8, -12}}}, 15},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := diagonalDifference(tt.args.arr); got != tt.want {
				t.Errorf("diagonalDifference() = %v, want %v", got, tt.want)
			}
		})
	}
}
