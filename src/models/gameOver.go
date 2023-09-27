package models

import "time"

var gameOverChan = make(chan bool)

func StartGameOverListener(scene SceneShower) {
	go func() {
		<-gameOverChan
		time.Sleep(1500 * time.Millisecond)
		gameOver = false
		score = 0
		scene.Show()
	}()
}

type SceneShower interface {
	Show()
}
