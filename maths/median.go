package maths

func Median(data []float64) float64 {
	for i := 0; i < len(data); i++ {
		for j := i + 1; j < len(data); j++ {
			if data[i] > data[j] {
				data[i], data[j] = data[j], data[i]
			}
		}
	}
	var median float64
	mid := len(data) / 2

	if len(data)%2 == 0 {
		median = (data[mid] + data[mid-1]) / 2
	} else {
		median = data[mid]
	}
	return median
}