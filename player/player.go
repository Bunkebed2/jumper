package player

import (
	"github.com/bunke/jumper/attack"
	"github.com/bunke/jumper/hitbox"
	"github.com/hajimehoshi/ebiten/v2"
)

type Player struct {
	PlayerImage  *ebiten.Image
	missileImage *ebiten.Image
	Hitbox       *hitbox.Hitbox
	speed        float64
	cooldown     int
}

const (
	cooldownTimerLength = 30
)

func NewPlayer(playerImage, missileImage *ebiten.Image, xPos, yPos, speed float64) *Player {
	hitbox := hitbox.NewHitbox(xPos, yPos, playerImage.Bounds())

	player := &Player{
		playerImage, missileImage, hitbox, speed, 0,
	}

	return player
}

func (p *Player) FireMissile(playerAttacks []attack.Attack) []attack.Attack {
	if p.cooldown > 0 {
		p.cooldown--
	} else if ebiten.IsKeyPressed(ebiten.KeyE) {
		p.cooldown = cooldownTimerLength
		playerAttacks = append(playerAttacks, *attack.NewAttack(p.missileImage,
			p.Hitbox.XPos+(float64(p.PlayerImage.Bounds().Dx())/2.0)+float64(p.missileImage.Bounds().Dx()),
			p.Hitbox.YPos+float64(p.missileImage.Bounds().Dy()),
			0,
			-5,
		))
	}

	return playerAttacks
}

func (p *Player) MovePlayer(screenWidth float64, screenHeight float64) {
	if ebiten.IsKeyPressed(ebiten.KeyW) {
		p.Hitbox.YPos -= p.speed
	}

	if ebiten.IsKeyPressed(ebiten.KeyS) {
		p.Hitbox.YPos += p.speed
	}

	if ebiten.IsKeyPressed(ebiten.KeyA) {
		p.Hitbox.XPos -= p.speed
	}

	if ebiten.IsKeyPressed(ebiten.KeyD) {
		p.Hitbox.XPos += p.speed
	}

	playerWidth := p.PlayerImage.Bounds().Dx()
	playerHeight := p.PlayerImage.Bounds().Dy()

	playerXBound := screenWidth - float64(playerWidth)
	playerYBound := screenHeight - float64(playerHeight)

	if p.Hitbox.XPos < 0 {
		p.Hitbox.XPos = 0
	}

	if p.Hitbox.XPos > playerXBound {
		p.Hitbox.XPos = playerXBound
	}

	if p.Hitbox.YPos < 0 {
		p.Hitbox.YPos = 0
	}

	if p.Hitbox.YPos > playerYBound {
		p.Hitbox.YPos = playerYBound
	}
}

func (p *Player) Collision(hb *hitbox.Hitbox) bool {
	return p.Hitbox.Intersects(hb)
}
