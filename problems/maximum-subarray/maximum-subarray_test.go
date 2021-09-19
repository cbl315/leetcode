package main

import "testing"

func Test_getMaxSubArrayByDP(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "example-1",
			args: args{
				nums: []int{-2, 1, -3, 4, -1, 2, 1, -5, 4},
			},
			want: 6,
		},
		{
			name: "example-2",
			args: args{
				nums: []int{1},
			},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getMaxSubArrayByDP(tt.args.nums); got != tt.want {
				t.Errorf("getMaxSubArrayByDP() = %v, want %v", got, tt.want)
			}
		})
	}
}
