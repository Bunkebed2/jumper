package main

import (
	"bytes"
	"fmt"
	"image/color"
	"log"

	"github.com/bunke/jumper/enemy"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/examples/resources/fonts"
	"github.com/hajimehoshi/ebiten/v2/text/v2"
)

// Create our empty vars
var (
	background      *ebiten.Image
	spaceShip       *ebiten.Image
	enemyShipSprite *ebiten.Image
	playerOne       Player
	enemies         []enemy.Enemy
	playerAttacks   []Attack
	isPlayerAlive   bool
	mplusFaceSource *text.GoTextFaceSource
	eg              *enemy.EnemyGenerator
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
	// Loading the font face source with the data from the font
	ff, err := text.NewGoTextFaceSource(bytes.NewReader(fonts.MPlus1pRegular_ttf))
	if err != nil {
		log.Fatal("error loading font", err)
	}

	mplusFaceSource = ff

	background = loadImage("assets/spacebackground.png")
	spaceShip = loadImage("assets/spaceship.png")
	enemyShipSprite = loadImage("assets/enemyship.png")
	missile := loadImage("assets/missile.png")
	playerOne = *NewPlayer(spaceShip, missile, screenWidth/2.0, screenHeight/2.0, 6)
	isPlayerAlive = true

	e1 := *enemy.NewEnemy(enemyShipSprite, 0, 0, 2)
	e2 := *enemy.NewEnemy(loadImage("assets/enemyFighter.png"), 0, 0, 3)

	enemies = make([]enemy.Enemy, 0)
	enemies = append(enemies, e1)

	eg = enemy.NewEnemyGenerator(e1, e2, e1)

	playerAttacks = make([]Attack, 0)
}

func (g *Game) Update() error {
	if isPlayerAlive {
		playerOne.movePlayer(screenWidth, screenHeight)
		playerAttacks = playerOne.fireMissile(playerAttacks)
	}

	enemies = append(enemies, eg.GenerateEnemies(screenWidth)...)

	for j, _ := range enemies {
		enemies[j].Move()
	}

	for k, _ := range playerAttacks {
		playerAttacks[k].move()
	}

	i := 0
	for j, _ := range enemies {
		if enemies[j].InBounds(screenWidth, screenHeight) {
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
				g.score++
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

func (g *Game) DrawHUD(screen *ebiten.Image) {
	op := &text.DrawOptions{}
	op.GeoM.Translate(10, 10)
	op.ColorScale.ScaleWithColor(color.White)
	text.Draw(screen, fmt.Sprintf("Score: %d", g.score), &text.GoTextFace{Source: mplusFaceSource, Size: 24}, op)
}

func (g *Game) Draw(screen *ebiten.Image) {
	draw(screen, background, 0, 0)

	if isPlayerAlive {
		draw(screen, playerOne.playerImage, playerOne.hitbox.XPos, playerOne.hitbox.YPos)
	}

	for _, e := range enemies {
		draw(screen, e.Image, e.Hitbox.XPos, e.Hitbox.YPos)
	}

	for _, a := range playerAttacks {
		draw(screen, a.image, a.hitbox.XPos, a.hitbox.YPos)
	}

	g.DrawHUD(screen)
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
