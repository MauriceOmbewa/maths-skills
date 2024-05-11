package maths

import (
	"math"
	"testing"
)

func TestStdDev(t *testing.T) {
	type args struct {
		variance float64
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		{
			name: "valid variance",
			args: args{variance: 4.0},
			want: 2.0,
		},
		{
			name: "zero variance",
			args: args{variance: 0.0},
			want: 0.0,
		},
		{
			name: "very large variance",
			args: args{variance: 1e100},
			want: math.Sqrt(1e100),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := StdDev(tt.args.variance); got != tt.want {
				t.Errorf("StdDev() = %v, want %v", got, tt.want)
			}
		})
	}
}
