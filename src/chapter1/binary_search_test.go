package chapter1

import "testing"

func Test_binarySearch(t *testing.T) {
	type args struct {
		searchKey int
		arr       []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "test case:not exists!",
			args: args{
				searchKey: 0,
				arr:       []int{},
			},
			want: -1,
		},
		{
			name: "test case:left 1",
			args: args{
				searchKey: 1,
				arr:       []int{1, 2, 3},
			},
			want: 0,
		},
		{
			name: "test case:right 1",
			args: args{
				searchKey: 3,
				arr:       []int{1, 2, 3},
			},
			want: 2,
		},
		{
			name: "test case:in mid",
			args: args{
				searchKey: 2,
				arr:       []int{1, 2, 3},
			},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := binarySearch(tt.args.searchKey, tt.args.arr); got != tt.want {
				t.Errorf("binarySearch() = %v, want %v", got, tt.want)
			}
		})
	}
}
