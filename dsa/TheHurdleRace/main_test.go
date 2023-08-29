package main

import "testing"

func Test_hurdleRace(t *testing.T) {
	type args struct {
		k      int32
		height []int32
	}
	tests := []struct {
		name string
		args args
		want int32
	}{
		{"tc1", args{4, []int32{1, 6, 3, 5, 2}}, 2},
		{"tc2", args{7, []int32{2, 5, 4, 5, 2}}, 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := hurdleRace(tt.args.k, tt.args.height); got != tt.want {
				t.Errorf("hurdleRace() = %v, want %v", got, tt.want)
			}
		})
	}
}
