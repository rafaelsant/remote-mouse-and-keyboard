package services

import (
	"github.com/go-vgo/robotgo"
	models "github.com/rafaelsant/remote-mouse-and-keyboard/models"
)

func MoveMouse(m *models.BodyMouseResponse) {
	sx, sy := robotgo.GetScreenSize()
	mX := ((sx / m.Dx) * m.X)
	mY := ((sy / m.Dy) * m.Y)

	robotgo.MoveSmoothRelative(mX, mY)
}
func MouseClick(m *models.BodyMouseClick) {
	robotgo.Click(m.Click)
}
