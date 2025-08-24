package main

import (
	"fmt"
	"image/color"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

const (
	G         = 0.1
	Scale     = 1.0
	NumBodies = 3
)

type Body struct {
	x, y   float64
	vx, vy float64
	m      float64
	color  color.Color
}

type Game struct {
	bodies   [NumBodies]*Body
	paused   bool
	selected int
}

func (g *Game) Update() error {
	if ebiten.IsKeyPressed(ebiten.KeyP) {
		g.paused = true
	}
	if ebiten.IsKeyPressed(ebiten.KeyO) {
		g.paused = false
	}
	if ebiten.IsKeyPressed(ebiten.KeyR) {
		g.reset()
	}

	if ebiten.IsKeyPressed(ebiten.Key1) {
		g.selected = 0
	}
	if ebiten.IsKeyPressed(ebiten.Key2) {
		g.selected = 1
	}
	if ebiten.IsKeyPressed(ebiten.Key3) {
		g.selected = 2
	}

	selected := g.bodies[g.selected]
	step := 0.5
	if ebiten.IsKeyPressed(ebiten.KeyArrowUp) {
		selected.vy -= step
	}
	if ebiten.IsKeyPressed(ebiten.KeyArrowDown) {
		selected.vy += step
	}
	if ebiten.IsKeyPressed(ebiten.KeyArrowLeft) {
		selected.vx -= step
	}
	if ebiten.IsKeyPressed(ebiten.KeyArrowRight) {
		selected.vx += step
	}
	if ebiten.IsKeyPressed(ebiten.KeyW) {
		selected.y -= step * 2
	}
	if ebiten.IsKeyPressed(ebiten.KeyS) {
		selected.y += step * 2
	}
	if ebiten.IsKeyPressed(ebiten.KeyA) {
		selected.x -= step * 2
	}
	if ebiten.IsKeyPressed(ebiten.KeyD) {
		selected.x += step * 2
	}

	if g.paused {
		return nil
	}

	for i := range NumBodies {
		b1 := g.bodies[i]
		ax, ay := 0.0, 0.0
		for j := range NumBodies {
			if i == j {
				continue
			}
			b2 := g.bodies[j]
			dx := b2.x - b1.x
			dy := b2.y - b1.y
			distSq := dx*dx + dy*dy + 0.01
			f := G * b2.m / distSq
			ax += f * dx
			ay += f * dy
		}
		b1.vx += ax
		b1.vy += ay
	}

	for i := range NumBodies {
		b := g.bodies[i]
		b.x += b.vx
		b.y += b.vy
	}

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	for _, b := range g.bodies {
		ebitenutil.DrawRect(screen, b.x, b.y, 8, 8, b.color)
	}

	info := "Controls:\n" +
		"[1][2][3] = Select body\n" +
		"Arrows = Change velocity\n" +
		"WASD = Change position\n" +
		"P = Pause, O = Resume, R = Reset\n\n"

	for i, b := range g.bodies {
		tag := ""
		if i == g.selected {
			tag = " <- selected"
		}
		info += fmt.Sprintf("Body %d: x=%.2f y=%.2f vx=%.2f vy=%.2f%s\n",
			i+1, b.x, b.y, b.vx, b.vy, tag)
	}

	if g.paused {
		info += "\nSimulation: PAUSED"
	} else {
		info += "\nSimulation: RUNNING"
	}

	ebitenutil.DebugPrint(screen, info)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return 800, 600
}

func (g *Game) reset() {
	g.bodies = [NumBodies]*Body{
		{x: 300, y: 300, vx: 0, vy: -1.2, m: 10, color: color.RGBA{255, 0, 0, 255}},
		{x: 500, y: 300, vx: 0, vy: 1.2, m: 10, color: color.RGBA{0, 255, 0, 255}},
		{x: 400, y: 450, vx: 1.2, vy: 0, m: 10, color: color.RGBA{0, 0, 255, 255}},
	}
	g.paused = false
	g.selected = 0
}

func main() {
	game := &Game{}
	game.reset()

	ebiten.SetWindowSize(800, 600)
	ebiten.SetWindowTitle("Interactive 3-Body Problem")
	ebiten.SetWindowResizingMode(ebiten.WindowResizingModeEnabled)
	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}

