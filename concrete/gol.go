package concrete

import (
	"bufio"
	"fmt"
	abstraction "gol/abstraction"
	"math/rand"
	"os"
	"time"
)

type Gol struct {
	abstraction.IGol

	lives  [][]abstraction.ILife
	factor float32
}

func buildMatrix(lines int, cols int) [][]abstraction.ILife {
	lives := make([][]abstraction.ILife, lines)
	for i := range lives {
		lives[i] = make([]abstraction.ILife, cols)
		for j := range lives[i] {
			lives[i][j] = NewLife()
		}
	}

	return lives
}

func NewGol(lines int, cols int, factor int) abstraction.IGol {
	lives := buildMatrix(lines, cols)

	return &Gol{
		lives:  lives,
		factor: float32(factor) / 100,
	}
}

func (g *Gol) StartFromFile(filepath string) {
	f, err := os.Open(filepath)
	if err != nil {
		panic(err)
	}
	reader := bufio.NewReader(f)

	var p string
	for i := 0; i < len(g.lives); i++ {
		fmt.Fscanf(reader, "%s\n", &p)
		for pos, char := range p {
			if pos <= len(g.lives[i]) && string(char) == "1" {
				g.lives[i][pos].Toggle()
			}
		}
		fmt.Fscan(reader, "\n")
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

	nextLives := make([][]abstraction.ILife, len(g.lives))

	for i := range g.lives {
		nextLives[i] = make([]abstraction.ILife, len(g.lives[i]))
		for j := range g.lives[i] {
			current := g.lives[i][j]
			nextLives[i][j] = current.NewFrom(current)

			aliveNeighbors := 0

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

			next := nextLives[i][j]
			if current.GetStatus() {
				next.Toggle()
				if aliveNeighbors == 2 || aliveNeighbors == 3 {
					next.Toggle()
				}
			} else if aliveNeighbors == 3 {
				next.Toggle()
			}
		}
	}

	g.lives = nextLives
	for i := range g.lives {
		g.lives[i] = nextLives[i]
		for j := range g.lives[i] {
			g.lives[i][j] = nextLives[i][j]
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
