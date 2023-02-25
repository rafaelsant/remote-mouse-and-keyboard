package models

type BodyMouseResponse struct {
	X  int `json:"x"`
	Y  int `json:"y"`
	Dx int `json:"dx"`
	Dy int `json:"dy"`
}

type BodyKeyboardResponse struct {
	Key string `json:"key"`
}

type BodyMouseClick struct {
	Click string `json:"click"`
}
