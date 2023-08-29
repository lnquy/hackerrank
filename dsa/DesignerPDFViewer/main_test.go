package main

import "testing"

func Test_designerPdfViewer(t *testing.T) {
	type args struct {
		h    []int32
		word string
	}
	tests := []struct {
		name string
		args args
		want int32
	}{
		{"tc1", args{[]int32{1, 3, 1, 3, 1, 4, 1, 3, 2, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5}, "abc"}, 9},
		{"tc2", args{[]int32{1, 3, 1, 3, 1, 4, 1, 3, 2, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 7}, "zaba"}, 28},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := designerPdfViewer(tt.args.h, tt.args.word); got != tt.want {
				t.Errorf("designerPdfViewer() = %v, want %v", got, tt.want)
			}
		})
	}
}
