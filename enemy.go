package main

import (
	"github.com/bunke/jumper/hitbox"
	"github.com/hajimehoshi/ebiten"
)

type Enemy struct {
	image  *ebiten.Image
	hitbox hitbox.Hitbox
	speed  float64
}

func NewEnemy(image *ebiten.Image, xPos, yPos, speed float64) *Enemy {
	hitbox := *hitbox.NewHitbox(xPos, yPos, image.Bounds())

	enemy := &Enemy{
		image, hitbox, speed,
	}

	return enemy
}

func (e *Enemy) move() {
	e.hitbox.YPos += e.speed
}

func (e *Enemy) inBounds(xBound, yBound float64) bool {
	width := e.Dx()
	height := e.image.Bounds().Dy()

	if e.hitbox.XPos < -float64(width) || e.hitbox.XPos > xBound || e.hitbox.YPos < -float64(height) || e.hitbox.YPos > yBound {
		return false
	}

	return true
}

func (e *Enemy) Dx() int {
	return e.hitbox.Dx()
}

func (e *Enemy) Dy() int {
	return e.hitbox.Dy()
}
