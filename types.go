package tripeg

import (
	"fmt"
	"math/rand"
	"time"
)

type Hole struct {
	Index int
	Row   int //make of 5
	Col   int //max of 9
	Peg   bool
	Links []*Hole //Other Holes the hole is connected to
}

func (h *Hole) Jump(overHole Hole) bool {
	return false
}

type Board struct {
	//May not need this
	Holes []Hole
}

//BuildBoard makes a board of peg holes.
//All holes are empty.  The top row has 1, then
//2,3,4,5 for a total of 16 holes.
func BuildBoard() Board {
	var b Board
	s2 := rand.NewSource(time.Now().UnixNano())
	r2 := rand.New(s2)
	fmt.Print(r2.Intn(16))
	i := 0
	for r := 1; r < 6; r++ {
		for c := 1; c < r+1; c++ {
			col := 4 - (r) + (c * 2)
			i++
			h := Hole{Index: i, Row: r, Col: col}
			b.Holes = append(b.Holes, h)
		}
	}
	fmt.Println("*", len(b.Holes))
	return b
}

func even(number int) bool {
	return number%2 == 0
}
