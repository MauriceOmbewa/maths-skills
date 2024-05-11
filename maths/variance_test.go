package maths

import "testing"

func TestVariance(t *testing.T) {
	type args struct {
		data []float64
		mean float64
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		{
			name: "single element",
			args: args{data: []float64{1}, mean: 1},
			want: 0,
		},
		{
			name: "multiple elements",
			args: args{data: []float64{1, 2, 3, 4, 5}, mean: 3},
			want: 2,
		},
		{
			name: "mean is equal to one element",
			args: args{data: []float64{1, 2, 3, 4, 5}, mean: 3},
			want: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Variance(tt.args.data, tt.args.mean); got != tt.want {
				t.Errorf("Variance() = %v, want %v", got, tt.want)
			}
		})
	}
}
