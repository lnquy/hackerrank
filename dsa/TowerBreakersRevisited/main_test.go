package main

import "testing"

func Test_towerBreakers(t *testing.T) {
	type args struct {
		arr []int32
	}
	tests := []struct {
		name string
		args args
		want int32
	}{
		{"tc1", args{[]int32{1, 2}}, 1},
		{"tc2", args{[]int32{1, 2, 3}}, 2},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := towerBreakers(tt.args.arr); got != tt.want {
				t.Errorf("towerBreakers() = %v, want %v", got, tt.want)
			}
		})
	}
}
