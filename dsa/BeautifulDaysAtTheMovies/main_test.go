package main

import "testing"

func Test_beautifulDays(t *testing.T) {
	type args struct {
		i int32
		j int32
		k int32
	}
	tests := []struct {
		name      string
		args      args
		wantBDays int32
	}{
		{"tc1", args{20, 23, 6}, 2},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotBDays := beautifulDays(tt.args.i, tt.args.j, tt.args.k); gotBDays != tt.wantBDays {
				t.Errorf("beautifulDays() = %v, want %v", gotBDays, tt.wantBDays)
			}
		})
	}
}
