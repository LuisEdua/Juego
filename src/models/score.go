package models

import (
	"fmt"
	"fyne.io/fyne/v2/widget"
)

var score int
var scoreChan = make(chan int)

func CreateScore() *widget.Label {
	label := widget.NewLabel(fmt.Sprintf("Puntaje: %d", 0))
	return label
}

func StartScoreUpdater(scoreLabel *widget.Label) {
	go func() {
		for points := range scoreChan {
			score += points
			scoreLabel.SetText(fmt.Sprintf("Puntaje: %d", score))
		}
	}()
}

func UpdateScore(points int) {
	scoreChan <- points
}
