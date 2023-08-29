package main

import "testing"

func Test_timeConversion(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"tc1", args{"05:06:07AM"}, "05:06:07"},
		{"tc2", args{"05:06:07PM"}, "17:06:07"},
		{"tc3", args{"12:01:02AM"}, "00:01:02"},
		{"tc4", args{"12:01:02PM"}, "12:01:02"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := timeConversion(tt.args.s); got != tt.want {
				t.Errorf("timeConversion() = %v, want %v", got, tt.want)
			}
		})
	}
}
