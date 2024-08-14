package main

import (
	"log"
	"math/rand/v2"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

// Create our empty vars
var (
	background      *ebiten.Image
	spaceShip       *ebiten.Image
	enemyShipSprite *ebiten.Image
	playerOne       Player
	enemies         []Enemy
	playerAttacks   []Attack
	isPlayerAlive   bool
)

const (
	screenWidth, screenHeight = 1280, 720
)

type Game struct {
	score int
}

func loadImage(imgPath string) *ebiten.Image {
	image, _, err := ebitenutil.NewImageFromFile(imgPath)
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

func init() {
	background = loadImage("assets/spacebackground.png")
	spaceShip = loadImage("assets/spaceship.png")
	enemyShipSprite = loadImage("assets/enemyship.png")
	missile := loadImage("assets/missile.png")
	playerOne = *NewPlayer(spaceShip, missile, screenWidth/2.0, screenHeight/2.0, 6)
	isPlayerAlive = true

	enemies = make([]Enemy, 0)
	enemies = append(enemies, *NewEnemy(enemyShipSprite, 0, 0, 2))

	playerAttacks = make([]Attack, 0)
}

func (g *Game) Update() error {
	if isPlayerAlive {
		playerOne.movePlayer(screenWidth, screenHeight)
		playerAttacks = playerOne.fireMissile(playerAttacks)
	}

	r := rand.IntN(200)
	if r == 10 {
		e := *NewEnemy(enemyShipSprite, 0, 0, 2)
		e.hitbox.XPos = float64(rand.IntN(screenWidth - e.Dx()))
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

		if isPlayerAlive && playerOne.collision(enemies[j]) {
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

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	draw(screen, background, 0, 0)

	if isPlayerAlive {
		draw(screen, playerOne.playerImage, playerOne.hitbox.XPos, playerOne.hitbox.YPos)
	}

	for _, e := range enemies {
		draw(screen, e.image, e.hitbox.XPos, e.hitbox.YPos)
	}

	for _, a := range playerAttacks {
		draw(screen, a.image, a.hitbox.XPos, a.hitbox.YPos)
	}
}

func (g *Game) Layout(ow, oh int) (int, int) {
	return ow, oh
}

func main() {
	ebiten.SetWindowSize(screenWidth, screenHeight)
	ebiten.SetWindowTitle("JUMPER")
	if err := ebiten.RunGame(&Game{0}); err != nil {
		log.Fatal(err)
	}
}
