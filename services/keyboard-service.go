package services

import (
	"github.com/go-vgo/robotgo"
	"github.com/rafaelsant/remote-mouse-and-keyboard/models"
)

func SendKey(k *models.BodyKeyboardResponse) {
	robotgo.KeyTap(k.Key)
}
