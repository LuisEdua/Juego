package models

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

var gameOver bool

func CreateCookie() fyne.CanvasObject {
	buttonSize := fyne.NewSize(400, 400)
	cookieButton := NewImageButton("src/assets/cookie.png", buttonSize, func() {
		if !gameOver {
			UpdateScore(1)
		}
	})

	return container.NewMax(cookieButton)
}

func NewImageButton(imagePath string, size fyne.Size, tapped func()) fyne.CanvasObject {
	img := canvas.NewImageFromFile(imagePath)
	img.FillMode = canvas.ImageFillContain
	btn := widget.NewButton("", tapped)
	return container.NewMax(btn, img)
}
