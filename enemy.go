package main

import (
	"github.com/hajimehoshi/ebiten"
)

type Enemy struct {
	image      *ebiten.Image
	xPos, yPos float64
	speed      float64
}

func NewEnemy(image *ebiten.Image, xPos, yPos, speed float64) *Enemy {
	enemy := &Enemy{
		image, xPos, yPos, speed,
	}

	return enemy
}

func (e *Enemy) move() {
	e.yPos += e.speed
}

func (e *Enemy) inBounds(xBound, yBound float64) bool {
	width := e.Dx()
	height := e.image.Bounds().Dy()

	if e.xPos < -float64(width) || e.xPos > xBound || e.yPos < -float64(height) || e.yPos > yBound {
		return false
	}

	return true
}

func (e *Enemy) Dx() int {
	return e.image.Bounds().Dx()
}

func (e *Enemy) Dy() int {
	return e.image.Bounds().Dy()
}
