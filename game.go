package main

import (
	"github.com/hajimehoshi/ebiten/v2"
	"math"
	"spectrum/scene/start"
)

const (
	windowWidth float64 = 800

	barPadding       float64 = 2
	capHeight        float64 = 2
	smoothingTime    float64 = 0.7
	smoothingTimeCap float64 = 0.95
)

var (
	windowHeight = math.Round(windowWidth / 16.0 * 9)
	game         *Game
)

type Band struct {
	lower float64
	upper float64
}

var gSamplesCache [2][]float64
var gCapCache [2][]float64

type GameState int

const (
	gameStateStart   = 0
	gameStatePlaying = 1
)

var startScene *start.Scene

func init() {
	startScene = &start.Scene{}
}

type Game struct {
	state GameState
}

func (g Game) Update() error {
	if g.state == gameStateStart {
		return startScene.Update()
	}
	return nil
}

func (g Game) Draw(screen *ebiten.Image) {
	if g.state == gameStateStart {
		startScene.Draw(screen)
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
