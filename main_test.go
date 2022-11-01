package main

import (
	"github.com/mjibson/go-dsp/fft"
	"math"
	"math/cmplx"
	"testing"
)

func TestFFT(t *testing.T) {
	//fft

	samples := []float64{1, 2, 3, 4, 5, 6, 7, 8}
	fftReal := fft.FFTReal(samples)
	t.Log("fftReal", fftReal)

	var fftResult []float64
	//取绝对值
	for i := range fftReal {
		fftResult = append(fftResult, cmplx.Abs(fftReal[i])/float64(len(samples)))
	}
	t.Log("fftResult", fftResult)

	dcComponent := fftResult[0] / float64(len(samples))
	t.Log("dcComponent", dcComponent)

	var spectrum = make([]float64, len(samples)/2+1)

	for i := 0; i < len(samples)/2+1; i++ {
		spectrum[i] = fftResult[i]
	}
	t.Log("spectrum", spectrum)

}

func Test2(t *testing.T) {
	//fft

	a := float64(10)

	b := complex(a, 0)
	t.Log(cmplx.Abs(b))

}

func TestInt16ToFloat64(t *testing.T) {
	t.Log(Int16ToFloat64(math.MaxInt16))
	t.Log(Int16ToFloat64(math.MinInt16))
	t.Log(Int16ToFloat64(100))
	t.Log(Int16ToFloat64(-100))
	t.Log(Int16ToFloat64(0))
	t.Log(complex(Int16ToFloat64(-100), 0))
}
