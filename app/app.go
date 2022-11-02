package app

import (
	"github.com/hajimehoshi/ebiten/v2/examples/resources/fonts"
	"golang.org/x/image/font"
	"golang.org/x/image/font/opentype"
	"spectrum/sdk"
)

var FftSamples [2][]float64
var SampleRate int
var PlayerManager *sdk.PlayList
var FontMPlus1pRegular font.Face

func init() {
	tt, err := opentype.Parse(fonts.PressStart2P_ttf)
	if err != nil {
		panic(err)
	}
	const dpi = 72
	FontMPlus1pRegular, err = opentype.NewFace(tt, &opentype.FaceOptions{
		Size:    32,
		DPI:     dpi,
		Hinting: font.HintingFull,
	})
	if err != nil {
		panic(err)
	}
}
