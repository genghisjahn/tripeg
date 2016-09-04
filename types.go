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

//Jump from the Board struct type
func (b Board) Jump(m, o Hole) (Board, error) {
	result := b
	if !m.Peg {
		//If there is no peg in the moveHole, no jump possible
		return result, fmt.Errorf("No Peg in move hole %d,%d\n", o.Row, o.Col)
	}
	if !o.Peg {
		//If there is no peg in the overHole, no jump possible
		return result, fmt.Errorf("No Peg in over hole %d,%d\n", o.Row, o.Col)
	}
	rDif := m.Row - o.Row
	cDif := o.Col - m.Col
	if cDif == 0 && rDif == 0 {
		//Holes are the same, not valid
		return result, fmt.Errorf("Jump peg and over hole are the same\n")
	}
	if math.Abs(float64(rDif)) > 1 {
		//You can't jump over more than 1 row horizontally
		return result, fmt.Errorf("Invalid horizonal movement %d\n", rDif)
	}
	if rDif > 0 && math.Abs(float64(cDif)) > 1 {
		//You can't jump over more than 1 col vertically
		return result, fmt.Errorf("Invalid vertical movement %d\n", cDif)
	}
	if rDif == 0 && math.Abs(float64(cDif)) > 2 {
		return result, fmt.Errorf("Invalid horizantal movement %d\n", rDif)
		//You can't jump more than 2 cols horizontally
	}
	targetR := 0
	targetC := 0
	if rDif == 0 {
		//This is a horizontal jump
		targetR = m.Row
	}
	if rDif > 0 {
		targetR = o.Row - 1
		//This is a up
	}
	if rDif < 0 {
		targetR = o.Row + 1
		//This is a jump down
	}
	if cDif < 0 {
		x := 1
		if rDif == 0 {
			x = 2
		}
		targetC = o.Col - x
		//This is a jump left
	}
	if cDif > 0 {
		x := 1
		if rDif == 0 {
			x = 2
		}
		targetC = o.Col + x
		//This is a jump right
	}
	targetHole := b.GetHole(targetR, targetC)
	if targetHole == nil {
		return result, fmt.Errorf("Target hole(%d,%d) does not exist\n", targetR, targetC)
	}
	if targetHole.Peg {
		return result, fmt.Errorf("Target hole(%d,%d) has a peg in it\n", targetHole.Row, targetHole.Col)
	}
	for _, bh := range result.Holes {
		if bh.Row == m.Row && bh.Col == m.Col {
			bh.Peg = false
		}
		if bh.Row == o.Row && bh.Col == o.Col {
			bh.Peg = false
		}
		if bh.Row == targetHole.Row && bh.Col == targetHole.Col {
			bh.Peg = true
		}
	}
	return result, nil
}

//Jump moves a peg from one hole to another
//If it can jump, it removes the peg from the
//overHole hole.
//**THIS FUNC DOES NOT WORK***
func (h *Hole) Jump(b Board, overHole *Hole) error {
	if !overHole.Peg {
		//If there is no peg in the overHole, no jump possible
		return fmt.Errorf("No Peg in %d,%d\n", overHole.Row, overHole.Col)
	}

	rDif := h.Row - overHole.Row
	cDif := overHole.Col - h.Col
	if cDif == 0 && rDif == 0 {
		//Holes are the same, not valid
		return fmt.Errorf("Jump peg and over hold are the same\n")
	}
	if math.Abs(float64(rDif)) > 1 {
		//You can't jump over more than 1 row horizontally
		return fmt.Errorf("Invalid horizonal movement %d\n", rDif)
	}
	if rDif > 0 && math.Abs(float64(cDif)) > 1 {
		//You can't jump over more than 1 col vertically
		return fmt.Errorf("Invalid vertical movement %d\n", cDif)
	}
	if rDif == 0 && math.Abs(float64(cDif)) > 2 {
		return fmt.Errorf("Invalid horizantal movement %d\n", rDif)
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
		return fmt.Errorf("Target hole(%d,%d) does not exist\n", targetR, targetC)
	}
	if targetHole.Peg {
		return fmt.Errorf("Target hole(%d,%d) has a peg in it\n", targetHole.Row, targetHole.Col)
	}
	h.Peg = false
	overHole.Peg = false
	targetHole.Peg = true
	return nil
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
	//Find out how many holes
	//can make a legal Move
	//find out how many moves can be made from those hole
	//randomly pick one of those
	//try again until there are no moves to make
	//or 14 legal moves have been made, (winner)
	//Print out all the winning moves
	type cMove struct {
		H *Hole
		O *Hole
	}
	fboard := Board{}
	fboard = *b
	cMoves := []cMove{}
	moves := 0
	holes := b.Holes[0:2]
	for _, v := range holes {
		var bt = fboard
		if v.Peg == true {
			fmt.Println(v.Row, v.Col)
			o := bt.GetHole(v.Row-1, v.Col-1)
			if o != nil {
				if v.Jump(bt, o) != nil {
					//upleft
					fmt.Println("UL")
					cMoves = append(cMoves, cMove{H: v, O: o})
				}
			}
			bt = fboard
			o = bt.GetHole(v.Row-1, v.Col+1)
			if o != nil {
				if v.Jump(bt, o) != nil {
					//upright
					fmt.Println("UR")
					cMoves = append(cMoves, cMove{H: v, O: o})
				}
			}
			bt = fboard
			o = bt.GetHole(v.Row, v.Col+2)
			if o != nil {
				if v.Jump(bt, o) != nil {
					//right
					cMoves = append(cMoves, cMove{H: v, O: o})
				}

			}
			bt = fboard
			o = bt.GetHole(v.Row, v.Col-2)
			if o != nil {
				if v.Jump(bt, o) != nil {
					//left
					cMoves = append(cMoves, cMove{H: v, O: o})
				}
			}
			bt = fboard
			o = bt.GetHole(v.Row+1, v.Col-2)
			if o != nil {
				if v.Jump(bt, o) != nil {
					//downleft
					cMoves = append(cMoves, cMove{H: v, O: o})
				}
			}
			bt = fboard
			o = bt.GetHole(v.Row+1, v.Col+2)
			if o != nil {
				if v.Jump(bt, o) != nil {
					//downright
					cMoves = append(cMoves, cMove{H: v, O: o})
				}
			}
		}
	}
	_ = moves
	for _, mv := range cMoves {
		fmt.Println(mv.H, mv.O)
	}
	return
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
