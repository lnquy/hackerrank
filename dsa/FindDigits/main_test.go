package main

import "testing"

func Test_findDigits(t *testing.T) {
	type args struct {
		n int32
	}
	tests := []struct {
		name         string
		args         args
		wantDivisors int32
	}{
		{"tc1", args{124}, 3},
		{"tc2", args{111}, 3},
		{"tc3", args{12}, 2},
		{"tc4", args{1012}, 3},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotDivisors := findDigits(tt.args.n); gotDivisors != tt.wantDivisors {
				t.Errorf("findDigits() = %v, want %v", gotDivisors, tt.wantDivisors)
			}
		})
	}
}
