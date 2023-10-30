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
	park *Park

	I       int
	skin    *canvas.Image
}

func NewCar(p *Park, s *canvas.Image) *Car {
	return &Car{
		park: p,
		skin:    s,
	}
}

func GenerateCar(n int, park *Park) {
	park.Space <- true
	for i := 0; i < n; i++ {
		CarImg := canvas.NewImageFromURI(storage.NewFileURI("./assets/rayo.png"))
		CarImg.Resize(fyne.NewSize(50, 100))
			x := rand.Intn(700-100+1) + 1
		CarImg.Move(fyne.NewPos(float32(x), 500))

		NewCar := NewCar(park, CarImg)
		NewCar.I = i + 1

		park.DrawCar <- CarImg
		go NewCar.RunCar()
		Wait := rand.Intn(700-100+1) + 1
			time.Sleep(time.Duration(Wait) * time.Millisecond)
	}
}

func (v *Car) RunCar() {
	v.park.Space <- true
	v.park.mutex.Lock()
	x := float32(rand.Intn(650 - 150 + 1))
	y := float32(rand.Intn(300 - 50 + 1))
	v.skin.Move(fyne.NewPos(x, y))
	fmt.Println("Carro ", v.I, " Entra")
	time.Sleep(200 * time.Millisecond)
	v.park.mutex.Unlock()

	Wait := rand.Intn(5-1+1) + 1
	time.Sleep(time.Duration(Wait) * time.Second)

	v.park.mutex.Lock()
	<-v.park.Space
	v.skin.Move(fyne.NewPos(0, 0))
	fmt.Println("Carro ", v.I, " Sale")
	time.Sleep(200 * time.Millisecond)
	v.park.mutex.Unlock()
}
