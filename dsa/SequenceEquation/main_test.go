package main

import (
	"reflect"
	"testing"
)

func Test_permutationEquation(t *testing.T) {
	type args struct {
		p []int32
	}
	tests := []struct {
		name       string
		args       args
		wantResult []int32
	}{
		{"tc1", args{[]int32{5, 2, 1, 3, 4}}, []int32{4, 2, 5, 1, 3}},
		{"tc2", args{[]int32{2, 3, 1}}, []int32{2, 3, 1}},
		{"tc3", args{[]int32{4, 3, 5, 1, 2}}, []int32{1, 3, 5, 4, 2}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotResult := permutationEquation(tt.args.p); !reflect.DeepEqual(gotResult, tt.wantResult) {
				t.Errorf("permutationEquation() = %v, want %v", gotResult, tt.wantResult)
			}
		})
	}
}
