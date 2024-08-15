package enemy

import "math/rand/v2"

type EnemyGenerator struct {
	enemyT1 Enemy
	enemyT2 Enemy
	enemyT3 Enemy
}

func NewEnemyGenerator(enemyT1, enemyT2, enemyT3 Enemy) *EnemyGenerator {
	enemyGen := &EnemyGenerator{enemyT1, enemyT2, enemyT3}
	return enemyGen
}

func (eg *EnemyGenerator) GenerateEnemies(screenWidth int) []Enemy {
	r := rand.IntN(1000)
	enemies := make([]Enemy, 0)
	if 0 < r && r < 8 {
		e := *NewEnemy(eg.enemyT1.Image, 0, 0, eg.enemyT1.speed, eg.enemyT1.HP)
		e.Hitbox.XPos = float64(rand.IntN(screenWidth - e.Dx()))
		enemies = append(enemies, e)
	} else if 8 <= r && r < 12 {
		e := *NewEnemy(eg.enemyT2.Image, 0, 0, eg.enemyT2.speed, eg.enemyT2.HP)
		e.Hitbox.XPos = float64(rand.IntN(screenWidth - e.Dx()))
		enemies = append(enemies, e)
	} else if 12 <= r && r < 14 {
		e := *NewEnemy(eg.enemyT3.Image, 0, 0, eg.enemyT3.speed, eg.enemyT3.HP)
		e.Hitbox.XPos = float64(rand.IntN(screenWidth - e.Dx()))
		enemies = append(enemies, e)
	}

	return enemies
}
