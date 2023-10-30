package screen

import (
	"park/models"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/storage"

	"fyne.io/fyne/v2/container"
)

type GameScene struct {
	window  fyne.Window
	content *fyne.Container
}

func (s *GameScene) Render() {
	backgroundImage := canvas.NewImageFromURI(storage.NewFileURI("./assets/park2.png"))

	backgroundImage.Resize(fyne.NewSize(800, 600))
	backgroundImage.Move(fyne.NewPos(0, 0))

	s.content = container.NewWithoutLayout(
		backgroundImage,
	)
	s.window.SetContent(s.content)
	s.StartGame()

}

func NewScene(window fyne.Window) *GameScene {
	scene := &GameScene{window: window}
	scene.Render()
	return scene
}

func (s *GameScene) StartGame() {
	e := models.NewPark(20)
	go models.GenerateVehicle(100, e)
	go s.DrawVehicles(e)
}

func (s *GameScene) DrawVehicles(e *models.Parking) {
	for {
		imagen := <-e.DrawVehicle
		s.content.Add(imagen)
		s.window.Canvas().Refresh(s.content)
	}
}
