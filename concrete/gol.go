package concrete

import (
	"fmt"
	abstraction "gol/abstraction"
	"math/rand"
	"time"
)

type Gol struct {
	abstraction.IGol

	lives  [][]abstraction.ILife
	factor float32
}

func NewGol(lines int, cols int, factor int) abstraction.IGol {
	lives := make([][]abstraction.ILife, cols)
	for i := range lives {
		lives[i] = make([]abstraction.ILife, lines)
		for j := range lives[i] {
			lives[i][j] = NewLife()
		}
	}

	return &Gol{
		lives:  lives,
		factor: float32(factor) / 100,
	}
}

func (g *Gol) Start() {
	rand.Seed(time.Now().UTC().UnixNano())

	for i := range g.lives {
		for j := range g.lives[i] {
			rand := g.factor - rand.Float32()
			if rand >= 0 {
				g.lives[i][j].Toggle()
			}
		}
	}
}

func (g *Gol) Next() {
	for i := range g.lives {
		for j := range g.lives[i] {
			aliveNeighbors := 0
			current := g.lives[i][j]

			prev_i := i - 1
			if prev_i < 0 {
				prev_i = len(g.lives) - 1
			}

			prev_j := j - 1
			if prev_j < 0 {
				prev_j = len(g.lives[i]) - 1
			}

			next_i := (i + 1) % len(g.lives)
			next_j := (j + 1) % len(g.lives[i])

			if g.lives[i][prev_j].GetStatus() {
				aliveNeighbors++
			}

			if g.lives[i][next_j].GetStatus() {
				aliveNeighbors++
			}

			if g.lives[prev_i][prev_j].GetStatus() {
				aliveNeighbors++
			}

			if g.lives[prev_i][j].GetStatus() {
				aliveNeighbors++
			}

			if g.lives[prev_i][next_j].GetStatus() {
				aliveNeighbors++
			}

			if g.lives[next_i][j].GetStatus() {
				aliveNeighbors++
			}

			if g.lives[next_i][prev_j].GetStatus() {
				aliveNeighbors++
			}

			if g.lives[next_i][next_j].GetStatus() {
				aliveNeighbors++
			}

			if current.GetStatus() {
				current.Toggle()
				if aliveNeighbors == 2 || aliveNeighbors == 3 {
					current.Toggle()
				}
			} else if aliveNeighbors == 3 {
				current.Toggle()
			}
		}
	}
}

func (g *Gol) Print() {
	for i := range g.lives {
		for j := range g.lives[i] {
			if g.lives[i][j].GetStatus() {
				fmt.Print("#")
			} else {
				fmt.Print(" ")
			}
		}
		fmt.Printf("\n")
	}
}
