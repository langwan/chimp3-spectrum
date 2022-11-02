package sdk

import (
	helper_graphic "github.com/langwan/langgo/helpers/graphic"
	"github.com/mjibson/go-dsp/fft"
	"math"
)

func Mode0(ware *[2][]float64) *[2][]float64 {
	wareReal := fft.FFTReal(ware[0])
	var max float64 = 0
	for i := 0; i < len(ware[0]); i++ {
		fr := real(wareReal[i])
		fi := imag(wareReal[i])
		magnitude := math.Sqrt(fr*fr + fi*fi)
		ware[0][i] = magnitude

		if magnitude > max {
			max = magnitude
		}
	}
	for i := 0; i < len(ware[0]); i++ {
		ware[0][i], _ = helper_graphic.RangeMapper(ware[0][i], 0, max, 0, 1)
	}

	wareReal = fft.FFTReal(ware[1])
	max = 0
	for i := 0; i < len(ware[1]); i++ {
		fr := real(wareReal[i])
		fi := imag(wareReal[i])
		magnitude := math.Sqrt(fr*fr + fi*fi)
		ware[1][i] = magnitude
		if magnitude > max {
			max = magnitude
		}
	}

	for i := 0; i < len(ware[1]); i++ {
		ware[1][i], _ = helper_graphic.RangeMapper(ware[1][i], 0, max, 0, 1)
	}
	return ware
}
