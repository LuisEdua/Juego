package views

import (
	"Juego/src/scenes"
	"fyne.io/fyne/v2"
)

var (
	myWindow fyne.Window
)

type GameView struct {
	window fyne.Window
}

func NewGameView(window fyne.Window) *GameView {
	return &GameView{window: window}
}

func (s *GameView) Show() {
	myWindow = s.window
	menuScene := scenes.NewMenuScene(myWindow)
	menuScene.Show()
}
