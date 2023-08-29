package main

import "testing"

func Test_angryProfessor(t *testing.T) {
	type args struct {
		k int32
		a []int32
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"tc1", args{3, []int32{-1, -3, 4, 2}}, "YES"},
		{"tc2", args{2, []int32{0, -1, 2, 1}}, "NO"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := angryProfessor(tt.args.k, tt.args.a); got != tt.want {
				t.Errorf("angryProfessor() = %v, want %v", got, tt.want)
			}
		})
	}
}
