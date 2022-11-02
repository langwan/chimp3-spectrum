package main

import (
	"image/color"
	"spectrum/app"
	"spectrum/preset"
)

var led *preset.Led2

func main() {

	game = &Game{
		frequencyBands: 256,
		samples:        make(chan [2][]float64),
	}

	led = &preset.Led2{
		Fraction:       1.0 / 12.0,
		SpacingFactorX: 0.2,
		MinDecibels:    0.0,
		MaxDecibels:    255.0,
		LedCellHeight:  12.0,
		SpacingFactorY: 0.2,
		Color1: color.RGBA{
			R: 0xff,
			G: 0x56,
			B: 0x00,
			A: 200,
		},
		Color2: color.RGBA{
			R: 0x22,
			G: 0xff,
			B: 0x00,
			A: 200,
		},
	}

	go func() {

		Player("./samples/5.mp3", func(spectrum [2][]float64, sampleRate int) {
			//game.samples <- spectrum
			app.SampleRate = sampleRate
			app.FftSamples[0] = spectrum[0]
			app.FftSamples[1] = spectrum[1]

		})
	}()
	err := NewWin()
	if err != nil {
		panic(err)
	}
}
