package octavebands

import "math"

type Band struct {
	Min float64
	Mid float64
	Max float64
}

const (
	defaultCenter      = 1000.0
	defaultSpectrumMin = 15.0
	defaultSpectrumMax = 21000.0
)

func GenBanks(fraction float64) []Band {
	factor := math.Pow(2, fraction)
	factor2 := math.Pow(math.Sqrt2, fraction)
	var bands []Band
	for c := defaultCenter; c >= defaultSpectrumMin; c /= factor {
		bands = append([]Band{Band{
			Min: c / factor2,
			Mid: c,
			Max: c * factor2,
		}}, bands...)
	}
	for c := defaultCenter * factor; c <= defaultSpectrumMax; c *= factor {
		bands = append(bands, Band{
			Min: c / factor2,
			Mid: c,
			Max: c * factor2,
		})
	}
	return bands
}
