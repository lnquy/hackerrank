package main

import "testing"

func Test_aVeryBigSum(t *testing.T) {
	type args struct {
		ar []int64
	}
	tests := []struct {
		name    string
		args    args
		wantSum int64
	}{
		{"tc1", args{[]int64{1000000001, 1000000002, 1000000003, 1000000004, 1000000005}}, 5000000015},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotSum := aVeryBigSum(tt.args.ar); gotSum != tt.wantSum {
				t.Errorf("aVeryBigSum() = %v, want %v", gotSum, tt.wantSum)
			}
		})
	}
}
