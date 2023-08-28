package main

import "testing"

func Test_simpleArraySum(t *testing.T) {
	type args struct {
		ar []int32
	}
	tests := []struct {
		name    string
		args    args
		wantSum int32
	}{
		{"tc1", args{[]int32{1, 2, 3}}, 6},
		{"tc2", args{[]int32{1, 2, 3, 4, 10, 11}}, 31},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotSum := simpleArraySum(tt.args.ar); gotSum != tt.wantSum {
				t.Errorf("simpleArraySum() = %v, want %v", gotSum, tt.wantSum)
			}
		})
	}
}
