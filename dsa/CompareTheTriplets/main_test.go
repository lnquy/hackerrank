package main

import (
	"reflect"
	"testing"
)

func Test_compareTriplets(t *testing.T) {
	type args struct {
		a []int32
		b []int32
	}
	tests := []struct {
		name       string
		args       args
		wantResult []int32
	}{
		{"tc1", args{[]int32{1, 2, 3}, []int32{3, 2, 1}}, []int32{1, 1}},
		{"tc2", args{[]int32{5, 6, 7}, []int32{3, 6, 10}}, []int32{1, 1}},
		{"tc3", args{[]int32{17, 28, 30}, []int32{99, 16, 8}}, []int32{2, 1}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotResult := compareTriplets(tt.args.a, tt.args.b); !reflect.DeepEqual(gotResult, tt.wantResult) {
				t.Errorf("compareTriplets() = %v, want %v", gotResult, tt.wantResult)
			}
		})
	}
}
