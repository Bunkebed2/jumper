package main

import (
	"github.com/hajimehoshi/ebiten"
)

type Player struct {
	image      *ebiten.Image
	xPos, yPos float64
	speed      float64
}

func (p Player) movePlayer(screenWidth float64, screenHeight float64) {
	if ebiten.IsKeyPressed(ebiten.KeyW) {
		p.yPos -= p.speed
	}

	if ebiten.IsKeyPressed(ebiten.KeyS) {
		p.yPos += p.speed
	}

	if ebiten.IsKeyPressed(ebiten.KeyA) {
		p.xPos -= p.speed
	}

	if ebiten.IsKeyPressed(ebiten.KeyD) {
		p.xPos += p.speed
	}

	playerWidth := p.image.Bounds().Dx()
	playerHeight := p.image.Bounds().Dy()

	playerXBound := screenWidth - float64(playerWidth)
	playerYBound := screenHeight - float64(playerHeight)

	if p.xPos < 0 {
		p.xPos = 0
	}

	if p.xPos > playerXBound {
		p.xPos = playerXBound
	}

	if p.yPos < 0 {
		p.yPos = 0
	}

	if p.yPos > playerYBound {
		p.yPos = playerYBound
	}
}
