package enemy

import (
	"github.com/bunke/jumper/attack"
	"github.com/bunke/jumper/hitbox"
	"github.com/hajimehoshi/ebiten/v2"
)

type Enemy interface {
	Move()
	Attack(attacks []attack.Attack) []attack.Attack
	InBounds(xBound, yBound float64) bool
	Hitbox() hitbox.Hitbox
	Image() *ebiten.Image
	Dx() int
	Dy() int
	HP() int
	SetHP(hp int)
}

type BasicEnemy struct {
	image  *ebiten.Image
	hitbox hitbox.Hitbox
	speed  float64
	hp     int
}

func NewBasicEnemy(image *ebiten.Image, xPos, yPos, speed float64, hp int) *BasicEnemy {
	hitbox := *hitbox.NewHitbox(xPos, yPos, image.Bounds())

	enemy := &BasicEnemy{
		image, hitbox, speed, hp,
	}

	return enemy
}

func (e *BasicEnemy) Move() {
	e.hitbox.YPos += e.speed
}

func (e *BasicEnemy) Attack(attacks []attack.Attack) []attack.Attack {
	return attacks
}

func (e *BasicEnemy) InBounds(xBound, yBound float64) bool {
	width := e.Dx()
	height := e.image.Bounds().Dy()

	if e.hitbox.XPos < -float64(width) || e.hitbox.XPos > xBound || e.hitbox.YPos < -float64(height) || e.hitbox.YPos > yBound {
		return false
	}

	return true
}

func (e *BasicEnemy) Dx() int {
	return e.hitbox.Dx()
}

func (e *BasicEnemy) Dy() int {
	return e.hitbox.Dy()
}

func (e *BasicEnemy) Hitbox() hitbox.Hitbox {
	return e.hitbox
}

func (e *BasicEnemy) HP() int {
	return e.hp
}

func (e *BasicEnemy) SetHP(hp int) {
	e.hp = hp
}

func (e *BasicEnemy) Image() *ebiten.Image {
	return e.image
}
