package main

import (
	"errors"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/ncruces/zenity"
	"math"
	"spectrum/app"
	"spectrum/scene/led"
	"spectrum/scene/start"
)

const (
	windowWidth float64 = 800
)

var (
	windowHeight = math.Round(windowWidth / 16.0 * 9)
	game         *Game
)

var startScene *start.Scene
var ledScene *led.Scene

func init() {
	startScene = &start.Scene{}
	ledScene = &led.Scene{}
}

type Game struct {
}

func (g Game) trigger(key ebiten.Key, handler func() error) {
	if inpututil.IsKeyJustPressed(key) {
		go handler()
	}
}

func (g Game) Update() error {

	g.trigger(ebiten.KeyO, func() error {
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
		app.GameState = app.GameStatePlaying
		return nil
	})
	g.trigger(ebiten.KeyArrowLeft, func() error {
		app.PlayerManager.Prev()
		return nil
	})
	g.trigger(ebiten.KeyArrowRight, func() error {
		app.PlayerManager.Next()
		return nil
	})
	g.trigger(ebiten.KeySpace, func() error {

		app.PlayerManager.Playing(!app.PlayerManager.Player.IsPlay)
		return nil
	})

	if app.GameState == app.GameStateStart {
		return startScene.Update()
	} else {
		return ledScene.Update()
	}
	return nil
}

func (g Game) Draw(screen *ebiten.Image) {
	if app.GameState == app.GameStateStart {
		startScene.Draw(screen)
	} else {
		ledScene.Draw(screen)
	}
}

func (g Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	s := ebiten.DeviceScaleFactor()
	return int(float64(outsideWidth) * s), int(float64(outsideHeight) * s)
}

func NewWin() error {

	ebiten.SetWindowSize(int(windowWidth), int(windowHeight))
	ebiten.SetWindowTitle("chihuo-mp3-spectrum")

	if err := ebiten.RunGame(game); err != nil {
		return err
	}
	return nil
}
