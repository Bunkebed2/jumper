package main

import (
	"github.com/hajimehoshi/ebiten"
)

type Attack struct {
	image          *ebiten.Image
	xPos, yPos     float64
	xSpeed, ySpeed float64
}

func NewAttack(image *ebiten.Image, xPos, yPos, xSpeed, ySpeed float64) *Attack {
	attack := &Attack{
		image, xPos, yPos, xSpeed, ySpeed,
	}

	return attack
}

func (a *Attack) move() {
	a.xPos += a.xSpeed
	a.yPos += a.ySpeed
}

func (a *Attack) inBounds(xBound, yBound float64) bool {
	width := a.Dx()
	height := a.image.Bounds().Dy()

	if a.xPos < -float64(width) || a.xPos > xBound || a.yPos < -float64(height) || a.yPos > yBound {
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
	attackRect := a.image.Bounds()
	enemyRect := e.image.Bounds()

	attackRect.Min.X += int(a.xPos)
	attackRect.Max.X += int(a.xPos)
	attackRect.Min.Y += int(a.yPos)
	attackRect.Max.Y += int(a.yPos)

	enemyRect.Min.X += int(e.xPos)
	enemyRect.Max.X += int(e.xPos)
	enemyRect.Min.Y += int(e.yPos)
	enemyRect.Max.Y += int(e.yPos)

	return attackRect.Overlaps(enemyRect)
}
