package start

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text"
	"image/color"
	"spectrum/app"
)

type Scene struct {
	triggerKeys map[ebiten.Key]bool
}

func (s Scene) Update() error {

	return nil
}

func (s Scene) Draw(screen *ebiten.Image) {
	sw, sh := screen.Size()
	msg := "chimp3 spectrum\n\n\n[o] select mp3 file\n\n[space] play or pause\n\n[<-] previous\n\n[->] next\n\n\nauthor github.com/langwan"
	msgRect := text.BoundString(app.FontMPlus1pRegular, msg)
	msgX := (sw - msgRect.Dx()) / 2.0
	msgY := (sh - msgRect.Dy()) / 2.0
	text.Draw(screen, msg, app.FontMPlus1pRegular, msgX, msgY, color.White)
}

func (s Scene) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return 0, 0
}
