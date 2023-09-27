package models

import (
	"fmt"
	"fyne.io/fyne/v2/widget"
	"time"
)

func CreateTimer(gameTime time.Duration) (*widget.Label, func()) {
	timeLeft := gameTime
	timeLabel := widget.NewLabel(fmt.Sprintf("Tiempo: %.0f", timeLeft.Seconds()))

	startTimer := func() {
		for !gameOver {
			time.Sleep(time.Second)
			if timeLeft > 0 {
				timeLeft -= time.Second
				timeLabel.SetText(fmt.Sprintf("Tiempo: %.0f", timeLeft.Seconds()))
			} else {
				gameOver = true
				gameOverChan <- true
				timeLabel.SetText("Tiempo agotado")

			}
		}
	}

	return timeLabel, startTimer
}
