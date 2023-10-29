package models

import (
	"sync"
	"fyne.io/fyne/v2/canvas"
)

type Parking struct {
	Space chan bool
	DrawVehicle chan *canvas.Image
	mutex sync.Mutex
}


func NewParking(nS int) *Parking {
	return &Parking{
		Space: make(chan bool, nS+1),
		DrawVehicle: make(chan *canvas.Image,1),
	}
}