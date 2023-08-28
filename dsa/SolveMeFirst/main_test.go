package main

import "testing"

func Test_solveMeFirst(t *testing.T) {
	type args struct {
		a uint32
		b uint32
	}
	tests := []struct {
		name string
		args args
		want uint32
	}{
		{"tc1", args{0, 0}, 0},
		{"tc2", args{2, 3}, 5},
		{"tc3", args{10, 5}, 15},
		{"tc4", args{100, 1000}, 1100},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := solveMeFirst(tt.args.a, tt.args.b); got != tt.want {
				t.Errorf("solveMeFirst() = %v, want %v", got, tt.want)
			}
		})
	}
}
