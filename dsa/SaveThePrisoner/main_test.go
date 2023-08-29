package main

import "testing"

func Test_saveThePrisoner(t *testing.T) {
	type args struct {
		n int32
		m int32
		s int32
	}
	tests := []struct {
		name string
		args args
		want int32
	}{
		{"tc1", args{4, 6, 2}, 3},
		{"tc2", args{5, 2, 1}, 2},
		{"tc3", args{5, 2, 2}, 3},
		{"tc4", args{7, 19, 2}, 6},
		{"tc5", args{3, 7, 3}, 3},
		{"tc10", args{46934, 543563655, 46743}, 20809},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := saveThePrisoner(tt.args.n, tt.args.m, tt.args.s); got != tt.want {
				t.Errorf("saveThePrisoner() = %v, want %v", got, tt.want)
			}
		})
	}
}
