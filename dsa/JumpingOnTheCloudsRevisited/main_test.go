package main

import "testing"

func Test_jumpingOnClouds(t *testing.T) {
	type args struct {
		c []int32
		k int32
	}
	tests := []struct {
		name string
		args args
		want int32
	}{
		{"tc1", args{[]int32{0, 0, 1, 0}, 2}, 96},
		{"tc2", args{[]int32{0, 0, 1, 0, 0, 1, 1, 0}, 2}, 92},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := jumpingOnClouds(tt.args.c, tt.args.k); got != tt.want {
				t.Errorf("jumpingOnClouds() = %v, want %v", got, tt.want)
			}
		})
	}
}
