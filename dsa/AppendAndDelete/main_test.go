package main

import "testing"

func Test_appendAndDelete(t *testing.T) {
	type args struct {
		s string
		t string
		k int32
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"tc1", args{"abc", "def", 6}, "Yes"},
		{"tc2", args{"hackerhappy", "hackerrank", 9}, "Yes"},
		{"tc3", args{"aba", "aba", 7}, "Yes"},
		{"tc4", args{"ashley", "ash", 2}, "No"},
		{"tc5", args{"aaaaaaaaaa", "aaaaa", 7}, "Yes"},
		{"tc6", args{"abcd", "abcdert", 10}, "No"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := appendAndDelete(tt.args.s, tt.args.t, tt.args.k); got != tt.want {
				t.Errorf("appendAndDelete() = %v, want %v", got, tt.want)
			}
		})
	}
}
