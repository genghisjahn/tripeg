package tripeg

import (
	"fmt"
	"math"
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
func (h *Hole) Jump(b *Board, overHole *Hole) bool {
	if !overHole.Peg {
		//If there is no peg in the overHole, no jump possible
		return false
	}

	rDif := h.Row - overHole.Row
	cDif := overHole.Col - h.Col
	if cDif == 0 && rDif == 0 {
		//Holes are the same, not valid
		return false
	}
	if math.Abs(float64(rDif)) > 1 {
		//You can't jump over more than 1 row horizontally
		return false
	}
	if rDif > 0 && math.Abs(float64(cDif)) > 1 {
		//You can't jump over more than 1 col vertically
		return false
	}
	if rDif == 0 && math.Abs(float64(cDif)) > 2 {
		return false
		//You can't jump more than 2 cols horizontally
	}
	targetR := 0
	targetC := 0
	if rDif == 0 {
		//This is a horizontal jump
		targetR = h.Row
	}
	if rDif > 0 {
		targetR = overHole.Row - 1
		//This is a up
	}
	if rDif < 0 {
		targetR = overHole.Row + 1
		//This is a jump down
	}
	if cDif < 0 {
		x := 1
		if rDif == 0 {
			x = 2
		}
		targetC = overHole.Col - x
		//This is a jump left
	}
	if cDif > 0 {
		x := 1
		if rDif == 0 {
			x = 2
		}
		targetC = overHole.Col + x
		//This is a jump right
	}
	targetHole := b.GetHole(targetR, targetC)
	if targetHole == nil {
		return false
	}
	if targetHole.Peg {
		return false
	}
	h.Peg = false
	overHole.Peg = false
	targetHole.Peg = true
	return true
}

//Board contains all the holes that contain the pegs
type Board struct {
	Holes   []*Hole
	MoveLog []string
}

//GetHole gets a pointer to a hole based on the row,col coordinates
func (b Board) GetHole(r, c int) *Hole {
	if r < 0 || r > 6 || c < 0 || c > 9 {
		return nil
	}
	for _, v := range b.Holes {
		if v.Col == c && v.Row == r {
			return v
		}
	}
	return nil
}

//BuildBoard makes a board of peg holes.
//All holes have a peg except one randomly assigned.
//The top row has 1, then
//2,3,4,5 for a total of 16 holes.
func BuildBoard(empty int) Board {
	var b Board
	s2 := rand.NewSource(time.Now().UnixNano())
	r2 := rand.New(s2)
	if empty == 0 {
		empty = r2.Intn(15)
	} else {
		empty--
	}

	for r := 1; r < 6; r++ {
		for c := 1; c < r+1; c++ {
			col := 4 - (r) + (c * 2)
			h := Hole{Row: r, Col: col, Peg: true}
			if empty == len(b.Holes) {
				h.Peg = false
			}
			b.Holes = append(b.Holes, &h)
		}
	}
	return b
}

func (b *Board) Solve() {
	b.MoveLog = []string{}
	s2 := rand.NewSource(time.Now().UnixNano())
	r2 := rand.New(s2)
	p1, p2 := 0, 0
	h1 := &Hole{}
	h2 := &Hole{}
	h1, h2 = nil, nil
	for m := 0; m < 4; m++ {
		for { //Main try loop
			p1 = r2.Intn(15)
			p2 = r2.Intn(15)
			for k := range b.Holes {
				if k == p1 {
					h1 = b.Holes[k]
					if h2 != nil {
						break
					}
				}
				if k == p2 {
					h2 = b.Holes[k]
					if h1 != nil {
						break
					}
				}
			}
			if h1.Jump(b, h2) {
				b.MoveLog = append(b.MoveLog, fmt.Sprintf("OK %v %v %v\n", h1, h2, b))
				break
			}
		}
	}
}

func (b Board) String() string {
	result := "\n"
	for r := 1; r < 6; r++ {
		for c := 1; c < 10; c++ {
			h := b.GetHole(r, c)
			mark := " "
			if h != nil {
				mark = "O"
				if h.Peg {
					mark = "*"
				}
			}
			result += mark
		}
		result += "\n"
	}
	return result
}
func even(number int) bool {
	return number%2 == 0
}
