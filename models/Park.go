package models

import (
	"sync"

	"fyne.io/fyne/v2/canvas"
)

type Park struct {
	Space      chan bool
	DrawCar    chan *canvas.Image
	mutex      sync.Mutex
	ParkSpaces []Slot
}

type Slot struct {
	x        float32
	y        float32
	occupied bool
}

func NewPark(nS int) *Park {
	return &Park{
		Space: make(chan bool, nS+1),

		DrawCar: make(chan *canvas.Image, 1),

		ParkSpaces: []Slot{
			{x: 10, y: 190, occupied: false},
			{x: 90, y: 190, occupied: false},
			{x: 170, y: 190, occupied: false},
			{x: 245, y: 190, occupied: false},
			{x: 325, y: 190, occupied: false},
			{x: 405, y: 190, occupied: false},
			{x: 495, y: 190, occupied: false},
			{x: 570, y: 190, occupied: false},
			{x: 645, y: 190, occupied: false},
			{x: 730, y: 190, occupied: false},
			{x: 10, y: 320, occupied: false},
			{x: 90, y: 320, occupied: false},
			{x: 170, y: 320, occupied: false},
			{x: 245, y: 320, occupied: false},
			{x: 325, y: 320, occupied: false},
			{x: 405, y: 320, occupied: false},
			{x: 495, y: 320, occupied: false},
			{x: 570, y: 320, occupied: false},
			{x: 645, y: 320, occupied: false},
			{x: 730, y: 320, occupied: false},
		},
		
	}
}
