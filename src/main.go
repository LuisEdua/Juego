package main

import (
	"Juego/src/views"
	"fyne.io/fyne/v2/app"
)

func main() {
	myApp := app.New()
	myWindow := myApp.NewWindow("Cookie Clicker")
	myWindow.CenterOnScreen()
	myWindow.SetFixedSize(true)

	gameView := views.NewGameView(myWindow)
	gameView.Show()
	myWindow.ShowAndRun()
}
