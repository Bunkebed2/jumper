package enemy

import (
	"github.com/bunke/jumper/attack"
	"github.com/bunke/jumper/hitbox"
	"github.com/hajimehoshi/ebiten/v2"
)

type AdvancedEnemy struct {
	base        *BasicEnemy
	AttackImage *ebiten.Image
	cooldown    int
}

const (
	cooldownTimerLength = 225
	firstGunportLoc     = 20
	secondGunportLoc    = -16
)

func NewAdvancedEnemy(base *BasicEnemy, attackImage *ebiten.Image) *AdvancedEnemy {
	ae := &AdvancedEnemy{
		base, attackImage, 10,
	}
	return ae
}

func (ae *AdvancedEnemy) Attack(attacks []attack.Attack) []attack.Attack {
	if ae.cooldown > 0 {
		ae.cooldown--
	} else {
		ae.cooldown = cooldownTimerLength
		attacks = append(attacks, *attack.NewAttack(ae.AttackImage,
			ae.Hitbox().XPos+float64(firstGunportLoc),
			ae.Hitbox().YPos+float64(ae.AttackImage.Bounds().Dy()),
			0,
			ae.base.speed+5,
		))
	}
	return ae.base.Attack(attacks)
}

func (ae *AdvancedEnemy) Move() {
	ae.base.Move()
}

func (ae *AdvancedEnemy) InBounds(xBound, yBound float64) bool {
	return ae.base.InBounds(xBound, yBound)
}

func (ae *AdvancedEnemy) Dx() int {
	return ae.base.Dx()
}

func (ae *AdvancedEnemy) Dy() int {
	return ae.base.Dy()
}

func (ae *AdvancedEnemy) Hitbox() *hitbox.Hitbox {
	return ae.base.Hitbox()
}

func (ae *AdvancedEnemy) HP() int {
	return ae.base.HP()
}

func (ae *AdvancedEnemy) SetHP(hp int) {
	ae.base.SetHP(hp)
}

func (ae *AdvancedEnemy) Image() *ebiten.Image {
	return ae.base.Image()
}

func (ae *AdvancedEnemy) basicEnemy() *BasicEnemy {
	return ae.base
}
