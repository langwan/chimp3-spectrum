package main

import (
	"spectrum/app"
	"spectrum/sdk"
)

func main() {
	app.PlayerManager = &sdk.PlayList{Player: sdk.New()}
	app.PlayerManager.Player.PlayList = app.PlayerManager
	app.PlayerManager.Player.UpdateFrequencyFrame = func(p *sdk.Player, frame sdk.FrequencyFrame) {
		app.FrequencyFrame = frame
	}
	game = &Game{}
	err := NewWin()
	if err != nil {
		panic(err)
	}
}
