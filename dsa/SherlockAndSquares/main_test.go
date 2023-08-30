package main

import "testing"

func Test_squares(t *testing.T) {
	type args struct {
		a int32
		b int32
	}
	tests := []struct {
		name      string
		args      args
		wantCount int32
	}{
		{"tc1", args{24, 49}, 3},
		{"tc2", args{3, 9}, 2},
		{"tc3", args{17, 24}, 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotCount := squares(tt.args.a, tt.args.b); gotCount != tt.wantCount {
				t.Errorf("squares() = %v, want %v", gotCount, tt.wantCount)
			}
		})
	}
}
