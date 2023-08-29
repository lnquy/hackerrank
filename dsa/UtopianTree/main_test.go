package main

import "testing"

func Test_utopianTree(t *testing.T) {
	type args struct {
		n int32
	}
	tests := []struct {
		name string
		args args
		want int32
	}{
		{"tc1", args{0}, 1},
		{"tc2", args{1}, 2},
		{"tc3", args{2}, 3},
		{"tc4", args{3}, 6},
		{"tc5", args{4}, 7},
		{"tc6", args{5}, 14},
		{"tc7", args{6}, 15},
		{"tc8", args{7}, 30},
		{"tc9", args{8}, 31},
		{"tc10", args{9}, 62},
		{"tc11", args{10}, 63},
		{"tc12", args{11}, 126},
		{"tc13", args{12}, 127},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := utopianTree(tt.args.n); got != tt.want {
				t.Errorf("utopianTree() = %v, want %v", got, tt.want)
			}
		})
	}
}

var _res int32

func BenchmarkUtopianTree(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_res = utopianTree(60)
	}
}
