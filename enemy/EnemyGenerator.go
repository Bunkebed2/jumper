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
	if 0 < r && r < 10 {
		e := *NewEnemy(eg.enemyT1.Image, 0, 0, eg.enemyT1.speed)
		e.Hitbox.XPos = float64(rand.IntN(screenWidth - e.Dx()))
		enemies = append(enemies, e)
	} else if 10 <= r && r < 15 {
		e := *NewEnemy(eg.enemyT2.Image, 0, 0, eg.enemyT2.speed)
		e.Hitbox.XPos = float64(rand.IntN(screenWidth - e.Dx()))
		enemies = append(enemies, e)
	}

	return enemies
}
