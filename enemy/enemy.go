package enemy

import (
	"github.com/bunke/jumper/hitbox"
	"github.com/hajimehoshi/ebiten/v2"
)

type Enemy struct {
	Image  *ebiten.Image
	Hitbox hitbox.Hitbox
	speed  float64
	HP     int
}

func NewEnemy(image *ebiten.Image, xPos, yPos, speed float64, hp int) *Enemy {
	hitbox := *hitbox.NewHitbox(xPos, yPos, image.Bounds())

	enemy := &Enemy{
		image, hitbox, speed, hp,
	}

	return enemy
}

func (e *Enemy) Move() {
	e.Hitbox.YPos += e.speed
}

func (e *Enemy) InBounds(xBound, yBound float64) bool {
	width := e.Dx()
	height := e.Image.Bounds().Dy()

	if e.Hitbox.XPos < -float64(width) || e.Hitbox.XPos > xBound || e.Hitbox.YPos < -float64(height) || e.Hitbox.YPos > yBound {
		return false
	}

	return true
}

func (e *Enemy) Dx() int {
	return e.Hitbox.Dx()
}

func (e *Enemy) Dy() int {
	return e.Hitbox.Dy()
}
