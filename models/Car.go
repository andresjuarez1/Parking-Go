package models

import (
	"fmt"
	"math/rand"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/storage"
)

type Vehicle struct {
	parking *Parking
	I       int
	skin    *canvas.Image
}

func NewVehicle(p *Parking, s *canvas.Image) *Vehicle {
	return &Vehicle{
		parking: p,
		skin:    s,
	}
}

func (v *Vehicle) RunVehicle() {
	v.parking.Space <- true
	v.parking.mutex.Lock()
	x := float32(rand.Intn(650 - 150 + 1))
	y := float32(rand.Intn(300 - 50 + 1))
	v.skin.Move(fyne.NewPos(x, y))
	fmt.Println("Carro ", v.I, " Entra")
	time.Sleep(200 * time.Millisecond)
	v.parking.mutex.Unlock()

	Wait := rand.Intn(5-1+1) + 1
	time.Sleep(time.Duration(Wait) * time.Second)

	v.parking.mutex.Lock()
	<-v.parking.Space
	v.skin.Move(fyne.NewPos(0, 0))
	fmt.Println("Carro ", v.I, " Sale")
	time.Sleep(200 * time.Millisecond)
	v.parking.mutex.Unlock()
}

func GenerateVehicle(n int, parking *Parking) {
	parking.Space <- true
	for i := 0; i < n; i++ {
		VehicleImg := canvas.NewImageFromURI(storage.NewFileURI("./assets/vehicle.png"))
		VehicleImg.Resize(fyne.NewSize(50, 100))
		x := rand.Intn(700-100+1) + 1
		VehicleImg.Move(fyne.NewPos(float32(x), 500))

		NewVehicle := NewVehicle(parking, VehicleImg)
		NewVehicle.I = i + 1

		parking.DrawVehicle <- VehicleImg
		go NewVehicle.RunVehicle()
		Wait := rand.Intn(700-100+1) + 1
		time.Sleep(time.Duration(Wait) * time.Millisecond)
	}
}
