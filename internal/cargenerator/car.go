package cargenerator

import (
	"github.com/VxVxN/game/pkg/rectangle"
	"github.com/hajimehoshi/ebiten/v2"
	"math/rand/v2"
)

type Car struct {
	screenHeight float64
	startRoad    float64
	*rectangle.Rectangle
	image *ebiten.Image
}

type roadLane int

const (
	FirstLane roadLane = iota
	SecondLane
	ThirdLane
	FourthLane
	FifthLane
)

func newCar(image *ebiten.Image, screenHeight, startRoad float64) *Car {
	return &Car{
		Rectangle:    rectangle.New(startRoad+65, -500, float64(image.Bounds().Dx()), float64(image.Bounds().Dy())),
		screenHeight: screenHeight,
		startRoad:    startRoad,
		image:        image,
	}
}

func (car *Car) Update(scrollSpeed float64) {
	car.Y += scrollSpeed
	if car.Y > car.screenHeight {
		car.Y = float64(rand.IntN(400) - 930)
		switch roadLane(rand.IntN(5)) {
		case FirstLane:
			car.X = car.startRoad + 65
		case SecondLane:
			car.X = car.startRoad + 265
		case ThirdLane:
			car.X = car.startRoad + 465
		case FourthLane:
			car.X = car.startRoad + 655
		case FifthLane:
			car.X = car.startRoad + 855
		}
	}
}

func (car *Car) Draw(screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(car.X, car.Y)
	screen.DrawImage(car.image, op)
}

func (car *Car) Reset() {
	car.Y = -500
}
