package main

import "testing"

func Test_extraLongFactorials(t *testing.T) {
	type args struct {
		n int32
	}
	tests := []struct {
		name string
		args args
	}{
		{"tc1", args{1}},
		{"tc2", args{20}},
		{"tc3", args{30}},
		{"tc4", args{50}},
		{"tc5", args{100}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			extraLongFactorials(tt.args.n)
		})
	}
}
