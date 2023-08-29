package main

import (
	"reflect"
	"testing"
)

func Test_circularArrayRotation(t *testing.T) {
	type args struct {
		a       []int32
		k       int32
		queries []int32
	}
	tests := []struct {
		name       string
		args       args
		wantResult []int32
	}{
		{"tc1", args{[]int32{1, 2, 3}, 2, []int32{0, 1, 2}}, []int32{2, 3, 1}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotResult := circularArrayRotation(tt.args.a, tt.args.k, tt.args.queries); !reflect.DeepEqual(gotResult, tt.wantResult) {
				t.Errorf("circularArrayRotation() = %v, want %v", gotResult, tt.wantResult)
			}
		})
	}
}
