package scenes

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

type MenuScene struct {
	window fyne.Window
}

func NewMenuScene(window fyne.Window) *MenuScene {
	return &MenuScene{window: window}
}

func (s *MenuScene) Show() {
	background := canvas.NewImageFromFile("src/assets/cocina.jpg")
	background.FillMode = canvas.ImageFillOriginal

	playButton := TransparentButton("Jugar", func() {
		gameScene := NewGameScene(s.window)
		gameScene.Show()
	})

	s.window.SetContent(container.NewMax(
		background,
		container.NewCenter(playButton),
	))

}

func TransparentButton(label string, tapped func()) fyne.CanvasObject {
	btn := widget.NewButton(label, tapped)
	return widget.NewCard("", "", btn)
}
