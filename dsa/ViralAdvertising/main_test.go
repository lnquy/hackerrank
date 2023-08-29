package main

import "testing"

func Test_viralAdvertising(t *testing.T) {
	type args struct {
		n int32
	}
	tests := []struct {
		name      string
		args      args
		wantLikes int32
	}{
		{"tc1", args{1}, 2},
		{"tc2", args{2}, 5},
		{"tc3", args{3}, 9},
		{"tc4", args{4}, 15},
		{"tc5", args{5}, 24},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotLikes := viralAdvertising(tt.args.n); gotLikes != tt.wantLikes {
				t.Errorf("viralAdvertising() = %v, want %v", gotLikes, tt.wantLikes)
			}
		})
	}
}
