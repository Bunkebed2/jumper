package hitbox

import "image"

type Hitbox struct {
	XPos, YPos float64
	box        image.Rectangle
}

func NewHitbox(xPos, yPos float64, box image.Rectangle) *Hitbox {
	hitbox := &Hitbox{xPos, yPos, box}
	return hitbox
}

func (h Hitbox) Intersects(h2 *Hitbox) bool {
	playerRect := h.box
	enemyRect := h2.box

	playerRect.Min.X += int(h.XPos)
	playerRect.Max.X += int(h.XPos)
	playerRect.Min.Y += int(h.YPos)
	playerRect.Max.Y += int(h.YPos)

	enemyRect.Min.X += int(h2.XPos)
	enemyRect.Max.X += int(h2.XPos)
	enemyRect.Min.Y += int(h2.YPos)
	enemyRect.Max.Y += int(h2.YPos)

	return playerRect.Overlaps(enemyRect)
}

func (h *Hitbox) Dx() int {
	return h.box.Dx()
}

func (h *Hitbox) Dy() int {
	return h.box.Dy()
}
