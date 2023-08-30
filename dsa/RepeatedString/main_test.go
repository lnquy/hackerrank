package main

import "testing"

func Test_repeatedString(t *testing.T) {
	type args struct {
		s string
		n int64
	}
	tests := []struct {
		name       string
		args       args
		wantResult int64
	}{
		{"tc1", args{"abcac", 10}, 4},
		{"tc2", args{"aba", 10}, 7},
		{"tc3", args{"a", 1000000000000}, 1000000000000},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotResult := repeatedString(tt.args.s, tt.args.n); gotResult != tt.wantResult {
				t.Errorf("repeatedString() = %v, want %v", gotResult, tt.wantResult)
			}
		})
	}
}
