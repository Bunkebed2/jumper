package main

import (
	"log"

	"github.com/hajimehoshi/ebiten"
)

type Player struct {
	playerImage  *ebiten.Image
	missileImage *ebiten.Image
	xPos, yPos   float64
	speed        float64
	cooldown     int
}

const (
	cooldownTimerLength = 30
)

func NewPlayer(playerImage, missileImage *ebiten.Image, xPos, yPos, speed float64) *Player {
	player := &Player{
		playerImage, missileImage, xPos, yPos, speed, 0,
	}

	return player
}

func (p *Player) fireMissile(playerAttacks []Attack) []Attack {
	if p.cooldown > 0 {
		p.cooldown--
	}

	if p.cooldown == 0 && ebiten.IsKeyPressed(ebiten.KeyE) {
		log.Println(float64(p.playerImage.Bounds().Dx()) / 2.0)
		p.cooldown = cooldownTimerLength
		playerAttacks = append(playerAttacks, *NewAttack(p.missileImage,
			p.xPos+(float64(p.playerImage.Bounds().Dx())/2.0)+float64(p.missileImage.Bounds().Dx()),
			p.yPos+float64(p.missileImage.Bounds().Dy()),
			0,
			-5,
		))
	}

	return playerAttacks
}

func (p *Player) movePlayer(screenWidth float64, screenHeight float64) {
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

	playerWidth := p.playerImage.Bounds().Dx()
	playerHeight := p.playerImage.Bounds().Dy()

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

func (p *Player) intersects(e Enemy) bool {
	playerRect := p.playerImage.Bounds()
	enemyRect := e.image.Bounds()

	playerRect.Min.X += int(p.xPos)
	playerRect.Max.X += int(p.xPos)
	playerRect.Min.Y += int(p.yPos)
	playerRect.Max.Y += int(p.yPos)

	enemyRect.Min.X += int(e.xPos)
	enemyRect.Max.X += int(e.xPos)
	enemyRect.Min.Y += int(e.yPos)
	enemyRect.Max.Y += int(e.yPos)

	return playerRect.Overlaps(enemyRect)
}
