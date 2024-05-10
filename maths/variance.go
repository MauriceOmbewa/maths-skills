package maths

func Variance(data []float64, mean float64) float64 {
	var tempsq float64
	for i := 0; i < len(data); i++ {
		temp := mean - data[i]
		tempsq += temp * temp

	}

	return tempsq / float64(len(data))
}