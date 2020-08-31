package dewpoint

import (
	"errors"
	"math"	
)
func Calculate(T float64, H float64) (float64, error) {
	if T < -45 || T > 60 {
		return 0, errors.New("Temperatur nicht im gültigen Bereich (-45 - +60°C)")
	} else {
		if H < 0 || H > 100 {
			return 0, errors.New("Feuchtigkeit nicht im gültigen Bereich (0 - 100%)")
		} else {
			const a float64 = 17.62
			const b float64 = 243.12
			alpha := math.Log(H/100) + a*T/(b+T)
			return math.Round(((b*alpha)/(a-alpha))*100) / 100, nil
		}
	}

}
