package models

import (
	"sync"

	"fyne.io/fyne/v2/canvas"
)

type Park struct {
	Space       chan bool
	DrawCar chan *canvas.Image
	mutex       sync.Mutex
}

func NewPark(nS int) *Park {
	return &Park{
		Space:       make(chan bool, nS+1),
		DrawCar: make(chan *canvas.Image, 1),
	}
}
