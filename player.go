package main

import (
	"bytes"
	"github.com/hajimehoshi/go-mp3"
	"github.com/hajimehoshi/oto"
	"github.com/mjibson/go-dsp/fft"
	"github.com/mjibson/go-dsp/window"
	"io"
	"math"
	"math/cmplx"
	"os"
)

const (
	kDefaultSmoothingTimeConstant float64 = 0.8
	kDefaultMinDecibels           float64 = -100.0
	kDefaultMaxDecibels           float64 = -30.0
	samplesPreFrame                       = 1152
	channels                              = 2
	bitDepth                              = 2
)

func Player(filename string, analyser func(spectrum [2][]float64, sampleRate int)) error {
	mp3Data, err := os.ReadFile(filename)
	if err != nil {
		return err
	}
	mp3Decoder, err := mp3.NewDecoder(bytes.NewReader(mp3Data))
	if err != nil {
		return err
	}
	bufSize := samplesPreFrame * channels * bitDepth
	oc, err := oto.NewContext(mp3Decoder.SampleRate(), channels, bitDepth, bufSize)
	if err != nil {
		return err
	}
	buf := make([]byte, bufSize)
	player := oc.NewPlayer()
	for {
		_, err := mp3Decoder.Read(buf)
		if err == io.EOF {
			break
		} else if err != nil {
			return err
		}
		samples, err := analyserSamples(buf)

		spectrum := analyserFrequency(samples)
		//analyserFrequency100(spectrum)
		analyser(spectrum, mp3Decoder.SampleRate())
		if err != nil {
			return err
		}
		player.Write(buf)
	}
	return nil
}

func analyserSamples(buf []byte) ([2][]float64, error) {
	var samples [2][]float64
	samples[0] = make([]float64, samplesPreFrame)
	samples[1] = make([]float64, samplesPreFrame)
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
	return samples, nil
}

func analyserFrequency(samples [2][]float64) [2][]float64 {
	var spectrum [2][]float64
	magnitude_scale := float64(1.0) / float64(len(samples[0]))
	smoothing_time := kDefaultSmoothingTimeConstant

	spectrum[0] = make([]float64, samplesPreFrame)
	spectrum[1] = make([]float64, samplesPreFrame)
	for i, chanelSamples := range samples {
		window.Apply(chanelSamples, window.Blackman)
		fftReal := fft.FFTReal(chanelSamples)
		for j := range fftReal {
			scalar_magnitude := cmplx.Abs(fftReal[j]) * float64(magnitude_scale)
			spectrum[i][j] = smoothing_time*spectrum[i][j] + (1-smoothing_time)*scalar_magnitude
			//spectrum[i][j] = LinearToDecibels(spectrum[i][j])
		}

		ConvertToByteData(spectrum[i])
	}
	spectrum[0] = spectrum[0][0 : samplesPreFrame/2]
	spectrum[1] = spectrum[1][0 : samplesPreFrame/2]
	return spectrum
}

//func analyserFrequency100(samples [2][]float64) {
//	var max [2]float64
//	var min [2]float64
//	for i, channelSamples := range samples {
//		max[i] = -1
//		min[i] = 1
//		for j, _ := range channelSamples {
//			if j == 0 {
//				continue
//			}
//			if samples[i][j] > max[i] {
//				max[i] = samples[i][j]
//			}
//			if samples[i][j] < min[i] {
//				min[i] = samples[i][j]
//			}
//		}
//	}
//
//	for i, channelSamples := range samples {
//		dc := channelSamples[0] / float64(len(channelSamples))
//
//		for j, _ := range channelSamples {
//			samples[i][j], _ = helper_graphic.RangeMapper(samples[i][j]/(float64(len(samples[i]))/2), math.MinInt16, math.MaxInt16, 0, 100)
//		}
//		samples[i][0] = dc
//	}
//}

func decode(p []byte) (sample [2]float64) {
	for c := range sample {
		x, n := decodeSample(bitDepth, p)
		sample[c] = Int16ToFloat64(int16(x))
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

func ByteToInt16(val []byte) int16 {
	return int16(val[1]<<8 | val[0])
}

func Int16ToFloat64(val int16) float64 {
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

	range_scale_factor := 1 / (kDefaultMaxDecibels - kDefaultMinDecibels)
	if kDefaultMaxDecibels == kDefaultMinDecibels {
		range_scale_factor = 1
	}
	for i := range samples {
		db_mag := LinearToDecibels(samples[i])
		scaled_value := math.MaxUint8 * (db_mag - kDefaultMinDecibels) * range_scale_factor
		// Clip to valid range.
		if scaled_value < 0 {
			scaled_value = 0
		}
		if scaled_value > math.MaxUint8 {
			scaled_value = math.MaxUint8
		}
		samples[i] = scaled_value
	}

}

func LinearToDecibels(linear float64) float64 {
	ret := 20 * math.Log10(linear)
	if linear > 0 {
		return ret
	}
	return ret
}
