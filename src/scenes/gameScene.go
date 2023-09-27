package scenes

import (
	"Juego/src/models"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/theme"
	"time"
)

type GameScene struct {
	window fyne.Window
}

func NewGameScene(window fyne.Window) *GameScene {
	return &GameScene{window: window}
}

func (s *GameScene) Show() {
	background := canvas.NewImageFromFile("src/assets/cocina.jpg")
	background.FillMode = canvas.ImageFillOriginal

	scoreLabel := models.CreateScore()
	models.StartScoreUpdater(scoreLabel)

	cookie := models.CreateCookie()

	timeLabel, startTimer := models.CreateTimer(10 * time.Second)
	go startTimer()

	menuScene := NewMenuScene(s.window)
	models.StartGameOverListener(menuScene)

	vbox := container.NewVBox(
		container.NewMax(canvas.NewRectangle(theme.BackgroundColor()), scoreLabel),
		container.NewMax(canvas.NewRectangle(theme.BackgroundColor()), timeLabel),
		container.NewMax(canvas.NewRectangle(theme.BackgroundColor()), cookie),
	)

	centeredVBox := container.NewCenter(vbox)

	s.window.SetContent(
		container.NewMax(
			background,   // Esto actúa como fondo
			centeredVBox, // Este es el contenido centrado
		),
	)
}
