package main

import (
	"log"

	"github.com/bunke/jumper/hitbox"
	"github.com/hajimehoshi/ebiten"
)

type Player struct {
	playerImage  *ebiten.Image
	missileImage *ebiten.Image
	hitbox       hitbox.Hitbox
	speed        float64
	cooldown     int
}

const (
	cooldownTimerLength = 30
)

func NewPlayer(playerImage, missileImage *ebiten.Image, xPos, yPos, speed float64) *Player {
	hitbox := *hitbox.NewHitbox(xPos, yPos, playerImage.Bounds())

	player := &Player{
		playerImage, missileImage, hitbox, speed, 0,
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
			p.hitbox.XPos+(float64(p.playerImage.Bounds().Dx())/2.0)+float64(p.missileImage.Bounds().Dx()),
			p.hitbox.YPos+float64(p.missileImage.Bounds().Dy()),
			0,
			-5,
		))
	}

	return playerAttacks
}

func (p *Player) movePlayer(screenWidth float64, screenHeight float64) {
	if ebiten.IsKeyPressed(ebiten.KeyW) {
		p.hitbox.YPos -= p.speed
	}

	if ebiten.IsKeyPressed(ebiten.KeyS) {
		p.hitbox.YPos += p.speed
	}

	if ebiten.IsKeyPressed(ebiten.KeyA) {
		p.hitbox.XPos -= p.speed
	}

	if ebiten.IsKeyPressed(ebiten.KeyD) {
		p.hitbox.XPos += p.speed
	}

	playerWidth := p.playerImage.Bounds().Dx()
	playerHeight := p.playerImage.Bounds().Dy()

	playerXBound := screenWidth - float64(playerWidth)
	playerYBound := screenHeight - float64(playerHeight)

	if p.hitbox.XPos < 0 {
		p.hitbox.XPos = 0
	}

	if p.hitbox.XPos > playerXBound {
		p.hitbox.XPos = playerXBound
	}

	if p.hitbox.YPos < 0 {
		p.hitbox.YPos = 0
	}

	if p.hitbox.YPos > playerYBound {
		p.hitbox.YPos = playerYBound
	}
}

func (p *Player) collision(e Enemy) bool {
	return p.hitbox.Intersects(e.hitbox)
}
