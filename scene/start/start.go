package start

import (
	"errors"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/hajimehoshi/ebiten/v2/text"
	"github.com/ncruces/zenity"
	"image/color"
	"spectrum/app"
)

type Scene struct {
	triggerKeys map[ebiten.Key]bool
}

func (s Scene) Update() error {
	s.trigger(ebiten.KeyO, s.SelectFilesHandler)
	return nil
}

func (s Scene) trigger(key ebiten.Key, handler func() error) {
	if inpututil.IsKeyJustPressed(key) {
		go handler()
	}
}

func (s Scene) SelectFilesHandler() error {
	files, err := zenity.SelectFileMultiple(
		zenity.FileFilters{
			{"选择mp3文件", []string{"*.mp3"}},
		})
	if err != nil {
		return err
	}

	if err != nil {
		return err
	} else if len(files) < 1 {
		return errors.New("selected files is empty")
	}
	err = app.PlayerManager.PlayList(files)
	if err != nil {
		return err
	}
	app.PlayerManager.Play(0)
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
