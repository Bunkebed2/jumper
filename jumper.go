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
	playerOne = Player{spaceShip, screenWidth / 2.0, screenHeight / 2.0, 4}
	isPlayerAlive = true

	enemies = make([]Enemy, 0)
	enemies = append(enemies, *NewEnemy(enemyShipSprite, 0, 0, 2))
}

func update(screen *ebiten.Image) error {
	if isPlayerAlive {
		playerOne.movePlayer(screenWidth, screenHeight)
	}

	r := rand.IntN(200)
	if r == 10 {
		e := *NewEnemy(enemyShipSprite, 0, 0, 2)
		e.xPos = float64(rand.IntN(screenWidth - e.Dx()))
		enemies = append(enemies, e)
	}

	i := 0
	for j, _ := range enemies {
		enemies[j].move()
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

	if ebiten.IsDrawingSkipped() {
		return nil
	}

	draw(screen, background, 0, 0)

	if isPlayerAlive {
		draw(screen, playerOne.image, playerOne.xPos, playerOne.yPos)
	}

	for _, e := range enemies {
		draw(screen, e.image, e.xPos, e.yPos)
	}

	return nil
}

func main() {
	if err := ebiten.Run(update, screenWidth, screenHeight, 1, "JUMPER"); err != nil {
		log.Fatal(err)
	}
}
