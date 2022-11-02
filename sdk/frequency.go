package sdk

import (
	"bytes"
	"fmt"
	"github.com/mjibson/go-dsp/fft"
	"github.com/mjibson/go-dsp/window"
	"io"
	"math"
	"math/cmplx"
)

const (
	channels      = 2
	bitDepth      = 2
	smoothingTime = 0.8
	minDecibels   = -100.0
	maxDecibels   = -30.0
	FftSize       = 2048
)

type FrequencyFrame struct {
	SampleRate int          `json:"sample_rate"`
	Name       string       `json:"name"`
	Filepath   string       `json:"filepath"`
	IsPlay     bool         `json:"is_player"`
	Spectrum   [2][]float64 `json:"spectrum"`
}

func (f FrequencyFrame) String() string {
	return fmt.Sprintf("spectrum length %d %d", len(f.Spectrum[0]), len(f.Spectrum[1]))
}

func frequencySpectrum(buf []byte, sampleRate int, fftSize int) ([2][]float64, error) {
	var samples [2][]float64
	samples[0] = make([]float64, fftSize)
	samples[1] = make([]float64, fftSize)
	var tmp [4]byte
	reader := bytes.NewReader(buf)
	i := 0
	for {
		_, err := reader.Read(tmp[:])

		if err == io.EOF {
			break
		} else if err != nil {
			return samples, err
		}
		sample := decode(tmp[:])

		samples[0][i] = sample[0]
		samples[1][i] = sample[1]
		i++
	}
	spectrum := frequencyFft(samples)
	return spectrum, nil
}

func frequencyFft(samples [2][]float64) [2][]float64 {
	var spectrum [2][]float64
	magnitudeScale := 1.0 / float64(len(samples[0]))
	spectrum[0] = make([]float64, len(samples[0]))
	spectrum[1] = make([]float64, len(samples[1]))
	for i, chanelSamples := range samples {
		window.Apply(chanelSamples, window.Blackman)
		fftReal := fft.FFTReal(chanelSamples)
		for j := range fftReal {
			scalarMagnitude := cmplx.Abs(fftReal[j]) * magnitudeScale
			spectrum[i][j] = smoothingTime*spectrum[i][j] + (1-smoothingTime)*scalarMagnitude
		}
		ConvertToByteData(spectrum[i])
	}
	spectrum[0] = spectrum[0][0 : len(samples[0])/2]
	spectrum[1] = spectrum[1][0 : len(samples[1])/2]
	return spectrum
}

func decode(p []byte) (sample [2]float64) {
	for c := range sample {
		x, n := decodeSample(bitDepth, p)
		sample[c] = int16ToFloat64(int16(x))
		p = p[n:]
	}
	return sample
}

func decodeSample(precision int, p []byte) (x int64, n int) {
	var val int64
	for i := precision - 1; i >= 0; i-- {
		val <<= 8
		val += int64(p[i])
	}
	return val, precision
}

func int16ToFloat64(val int16) float64 {
	if val == math.MaxInt16 {
		return 1
	} else if val == math.MaxInt16 {
		return -1
	} else if val == 0 {
		return 0
	}
	return float64(val) / float64(math.MaxInt16+1)
}

func ConvertToByteData(samples []float64) {
	rangeScaleFactor := 1 / (maxDecibels - minDecibels)
	if maxDecibels == minDecibels {
		rangeScaleFactor = 1
	}
	for i := range samples {
		dbMag := LinearToDecibels(samples[i])
		scaledValue := math.MaxUint8 * (dbMag - minDecibels) * rangeScaleFactor
		if scaledValue < 0 {
			scaledValue = 0
		}
		if scaledValue > math.MaxUint8 {
			scaledValue = math.MaxUint8
		}
		samples[i] = scaledValue
	}
}

func LinearToDecibels(linear float64) float64 {
	ret := 20 * math.Log10(linear)
	if linear > 0 {
		return ret
	}
	return ret
}
