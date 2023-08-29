package main

import "testing"

func Test_pickingNumbers(t *testing.T) {
	type args struct {
		a []int32
	}
	tests := []struct {
		name string
		args args
		want int32
	}{
		{"tc1", args{[]int32{4, 6, 5, 3, 3, 1}}, 3},
		{"tc2", args{[]int32{1, 2, 2, 3, 1, 2}}, 5},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := pickingNumbers(tt.args.a); got != tt.want {
				t.Errorf("pickingNumbers() = %v, want %v", got, tt.want)
			}
		})
	}
}
