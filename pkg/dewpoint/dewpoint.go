package dewpoint

import (
	"math"
)
func Calculate(T float64, H float64) float64 {
	const a float64 = 17.62
	const b float64 = 243.12
	alpha := math.Log(H/100) + a*T/(b+T)
	return math.Round(((b*alpha)/(a-alpha))*100) / 100
}
