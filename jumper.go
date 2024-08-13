package main

import (
	"log"
	"math/rand/v2"

	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"
)

// Create our empty vars
var (
	err             error
	background      *ebiten.Image
	spaceShip       *ebiten.Image
	enemyShipSprite *ebiten.Image
	playerOne       Player
	enemies         []Enemy
	playerAttacks   []Attack
	isPlayerAlive   bool
)

const (
	screenWidth, screenHeight = 640, 480
)

func loadImage(imgPath string) *ebiten.Image {
	image, _, err := ebitenutil.NewImageFromFile(imgPath, ebiten.FilterDefault)
	if err != nil {
		log.Fatal(err)
	}

	return image
}

func draw(screen *ebiten.Image, image *ebiten.Image, xPos, yPos float64) {
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(xPos, yPos)
	screen.DrawImage(image, op)
}

func remove(s []Enemy, i int) []Enemy {
	s[i] = s[len(s)-1]
	return s[:len(s)-1]
}

// Run this code once at startup
func init() {
	background = loadImage("assets/Space_Background.png")
	spaceShip = loadImage("assets/spaceship.png")
	enemyShipSprite = loadImage("assets/enemy_ship.png")
	missile := loadImage("assets/missile.png")
	playerOne = *NewPlayer(spaceShip, missile, screenWidth/2.0, screenHeight/2.0, 4)
	isPlayerAlive = true

	enemies = make([]Enemy, 0)
	enemies = append(enemies, *NewEnemy(enemyShipSprite, 0, 0, 2))

	playerAttacks = make([]Attack, 0)
}

func update(screen *ebiten.Image) error {
	if isPlayerAlive {
		playerOne.movePlayer(screenWidth, screenHeight)
		playerAttacks = playerOne.fireMissile(playerAttacks)
	}

	r := rand.IntN(200)
	if r == 10 {
		e := *NewEnemy(enemyShipSprite, 0, 0, 2)
		e.xPos = float64(rand.IntN(screenWidth - e.Dx()))
		enemies = append(enemies, e)
	}

	for j, _ := range enemies {
		enemies[j].move()
	}

	for k, _ := range playerAttacks {
		playerAttacks[k].move()
	}

	i := 0
	for j, _ := range enemies {
		if enemies[j].inBounds(screenWidth, screenHeight) {
			enemies[i] = enemies[j]
			i++
		}

		if isPlayerAlive && playerOne.intersects(enemies[j]) {
			log.Println("Player Died")
			isPlayerAlive = false
		}
	}
	enemies = enemies[:i]

	i = 0
	for j, _ := range playerAttacks {
		if playerAttacks[j].inBounds(screenWidth, screenHeight) {
			playerAttacks[i] = playerAttacks[j]
			i++
		}
	}
	playerAttacks = playerAttacks[:i]

	i = 0
	for j, _ := range playerAttacks {
		attackHit := false
		k := 0
		for l, _ := range enemies {
			if !playerAttacks[j].intersects(enemies[l]) {
				enemies[k] = enemies[l]
				k++
			} else {
				attackHit = true
			}
		}
		enemies = enemies[:k]

		if !attackHit {
			playerAttacks[i] = playerAttacks[j]
			i++
		}
	}
	playerAttacks = playerAttacks[:i]

	if ebiten.IsDrawingSkipped() {
		return nil
	}

	draw(screen, background, 0, 0)

	if isPlayerAlive {
		draw(screen, playerOne.playerImage, playerOne.xPos, playerOne.yPos)
	}

	for _, e := range enemies {
		draw(screen, e.image, e.xPos, e.yPos)
	}

	for _, a := range playerAttacks {
		draw(screen, a.image, a.xPos, a.yPos)
	}

	return nil
}

func main() {
	if err := ebiten.Run(update, screenWidth, screenHeight, 1, "JUMPER"); err != nil {
		log.Fatal(err)
	}
}
