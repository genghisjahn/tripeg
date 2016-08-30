package tripeg

import (
	"math/rand"
	"time"
)

//Hole struct that contains information
//about a hole in the board, its location
//and whether or not it has a peg in it.
type Hole struct {
	Row   int //make of 5
	Col   int //max of 9
	Peg   bool
	Links []*Hole //Other Holes the hole is connected to
}

//Jump moves a peg from one hole to another
//If it can jump, it removes the peg from the
//overHole hole.
func (h *Hole) Jump(overHole Hole) bool {
	return false
}

//Board contains all the holes that contain the pegs
type Board struct {
	Holes []Hole
}

//BuildBoard makes a board of peg holes.
//All holes have a peg except one randomly assigned.
//The top row has 1, then
//2,3,4,5 for a total of 16 holes.
func BuildBoard() Board {
	var b Board
	s2 := rand.NewSource(time.Now().UnixNano())
	r2 := rand.New(s2)
	empty := r2.Intn(16)
	for r := 1; r < 6; r++ {
		for c := 1; c < r+1; c++ {
			col := 4 - (r) + (c * 2)
			h := Hole{Row: r, Col: col, Peg: true}
			if empty == len(b.Holes) {
				h.Peg = false
			}
			b.Holes = append(b.Holes, h)
		}
	}
	return b
}

func even(number int) bool {
	return number%2 == 0
}
