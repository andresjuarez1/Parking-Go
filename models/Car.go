package models

import (
	"fmt"
	"math/rand"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/storage"
)

type Car struct {
	park      *Park
	I         int
	skin      *canvas.Image
	ParkSpace int
}

func NewCar(p *Park, s *canvas.Image) *Car {
	return &Car{
		park: p,
		skin: s,
	}
}

func GenerateCar(n int, park *Park) {
	park.Space <- true
	for i := 0; i < n; i++ {
		CarImg := canvas.NewImageFromURI(storage.NewFileURI("./assets/rayo.png"))
		CarImg.Resize(fyne.NewSize(50, 100))
		CarImg.Move(fyne.NewPos(10, 500))

		NewCar := NewCar(park, CarImg)
		NewCar.I = i + 1

		park.DrawCar <- CarImg
		time.Sleep(time.Millisecond * 200)

		go NewCar.RunCar()
		Wait := rand.Intn(100-50+1) + 50
		time.Sleep(time.Duration(Wait) * time.Millisecond)
	}
}

func (v *Car) RunCar() {
	v.park.Space <- true
	v.park.mutex.Lock()
	for i := 0; i < len(v.park.ParkSpaces); i++ {
		if !v.park.ParkSpaces[i].occupied {
			v.skin.Move(fyne.NewPos(v.park.ParkSpaces[i].x, v.park.ParkSpaces[i].y))
			v.skin.Refresh()
			v.ParkSpace = i
			v.park.ParkSpaces[i].occupied = true
			break
		}
	}

	fmt.Println("Carro ", v.I, " entra al estacionamiento")
	time.Sleep(300 * time.Millisecond)
	v.park.mutex.Unlock()

	Wait := rand.Intn(5-1+1) + 1
	time.Sleep(time.Duration(Wait) * time.Second)

	v.park.mutex.Lock()
	<-v.park.Space
	v.park.ParkSpaces[v.ParkSpace].occupied = false
	v.skin.Move(fyne.NewPos(4600, 5000))
	fmt.Println("Carro ", v.I, " sale del estacionamiento")
	time.Sleep(200 * time.Millisecond)
	v.park.mutex.Unlock()
}
