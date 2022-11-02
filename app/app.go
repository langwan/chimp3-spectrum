package app

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/examples/resources/fonts"
	"golang.org/x/image/font"
	"golang.org/x/image/font/opentype"
	"image/color"
	"spectrum/sdk"
)

type _GameState int

const (
	GameStateStart   = 0
	GameStatePlaying = 1
)

var (
	WhiteImage         = ebiten.NewImage(3, 3)
	SampleRate         int
	PlayerManager      *sdk.PlayList
	FontMPlus1pRegular font.Face
	FrequencyFrame     sdk.FrequencyFrame
	GameState          _GameState
)

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
	WhiteImage.Fill(color.White)
}
