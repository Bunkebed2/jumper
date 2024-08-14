package main

import (
	"github.com/bunke/jumper/hitbox"
	"github.com/hajimehoshi/ebiten/v2"
)

type Attack struct {
	image          *ebiten.Image
	hitbox         hitbox.Hitbox
	xSpeed, ySpeed float64
}

func NewAttack(image *ebiten.Image, xPos, yPos, xSpeed, ySpeed float64) *Attack {
	hitbox := *hitbox.NewHitbox(xPos, yPos, image.Bounds())

	attack := &Attack{
		image, hitbox, xSpeed, ySpeed,
	}

	return attack
}

func (a *Attack) move() {
	a.hitbox.XPos += a.xSpeed
	a.hitbox.YPos += a.ySpeed
}

func (a *Attack) inBounds(xBound, yBound float64) bool {
	width := a.Dx()
	height := a.image.Bounds().Dy()

	if a.hitbox.XPos < -float64(width) || a.hitbox.XPos > xBound || a.hitbox.YPos < -float64(height) || a.hitbox.YPos > yBound {
		return false
	}

	return true
}

func (a *Attack) Dx() int {
	return a.image.Bounds().Dx()
}

func (a *Attack) Dy() int {
	return a.image.Bounds().Dy()
}

func (a *Attack) intersects(e Enemy) bool {
	return a.hitbox.Intersects(e.hitbox)
}
