package main

import "testing"

func Test_birthdayCakeCandles(t *testing.T) {
	type args struct {
		candles []int32
	}
	tests := []struct {
		name           string
		args           args
		wantMaxCounter int32
	}{
		{"tc1", args{[]int32{3, 2, 1, 3}}, 2},
		{"tc2", args{[]int32{4, 4, 1, 3}}, 2},
		{"tc3", args{[]int32{4}}, 1},
		{"tc4", args{[]int32{4, 4, 4, 4}}, 4},
		{"tc4", args{[]int32{4, 3, 2, 1}}, 1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotMaxCounter := birthdayCakeCandles(tt.args.candles); gotMaxCounter != tt.wantMaxCounter {
				t.Errorf("birthdayCakeCandles() = %v, want %v", gotMaxCounter, tt.wantMaxCounter)
			}
		})
	}
}
