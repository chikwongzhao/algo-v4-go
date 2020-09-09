package chapter1

import (
	"testing"
)

func Test_gcd(t *testing.T) {
	type args struct {
		x int
		y int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "x=-1、y=-1",
			args: args{
				x: -1,
				y: -1,
			},
			want: -1,
		},
		{
			name: "x=-1、y=0",
			args: args{
				x: -1,
				y: 0,
			},
			want: 0,
		},
		{
			name: "x=0、y=-1",
			args: args{
				x: 0,
				y: -1,
			},
			want: 0,
		},
		{
			name: "x=0、y=0",
			args: args{
				x: 0,
				y: 0,
			},
			want: 0,
		},
		{
			name: "x=0、y=1",
			args: args{
				x: 0,
				y: 1,
			},
			want: 0,
		},
		{
			name: "x=1、y=0",
			args: args{
				x: 1,
				y: 0,
			},
			want: 0,
		},
		{
			name: "x=1、y=1",
			args: args{
				x: 1,
				y: 1,
			},
			want: 1,
		},
		{
			name: "x=1、y=2",
			args: args{
				x: 1,
				y: 2,
			},
			want: 1,
		},
		{
			name: "x=2、y=1",
			args: args{
				x: 2,
				y: 1,
			},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := gcd(tt.args.x, tt.args.y); got != tt.want {
				t.Errorf("gcd() = %v, want %v", got, tt.want)
			}
		})
	}
}
