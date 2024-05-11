package maths

import "testing"

func TestMedian(t *testing.T) {
	type args struct {
		data []float64
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		{
			name: "Test Median with odd number of elements",
			args: args{
				data: []float64{1, 3, 5, 7, 9},
			},
			want: 5,
		},

		{
			name: "Test Median with even number of elements",
			args: args{
				data: []float64{1, 3, 5, 7, 8, 9},
			},
			want: 6,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Median(tt.args.data); got != tt.want {
				t.Errorf("Median() = %v, want %v", got, tt.want)
			}
		})
	}
}
