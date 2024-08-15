package attack

import (
	"github.com/bunke/jumper/hitbox"
	"github.com/hajimehoshi/ebiten/v2"
)

type Attack struct {
	Image          *ebiten.Image
	Hitbox         hitbox.Hitbox
	xSpeed, ySpeed float64
}

func NewAttack(image *ebiten.Image, xPos, yPos, xSpeed, ySpeed float64) *Attack {
	hitbox := *hitbox.NewHitbox(xPos, yPos, image.Bounds())

	attack := &Attack{
		image, hitbox, xSpeed, ySpeed,
	}

	return attack
}

func (a *Attack) Move() {
	a.Hitbox.XPos += a.xSpeed
	a.Hitbox.YPos += a.ySpeed
}

func (a *Attack) InBounds(xBound, yBound float64) bool {
	width := a.Dx()
	height := a.Image.Bounds().Dy()

	if a.Hitbox.XPos < -float64(width) || a.Hitbox.XPos > xBound || a.Hitbox.YPos < -float64(height) || a.Hitbox.YPos > yBound {
		return false
	}

	return true
}

func (a *Attack) Dx() int {
	return a.Image.Bounds().Dx()
}

func (a *Attack) Dy() int {
	return a.Image.Bounds().Dy()
}

func (a *Attack) Intersects(hb hitbox.Hitbox) bool {
	return a.Hitbox.Intersects(hb)
}
