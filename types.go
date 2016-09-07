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
	Row int //make of 5
	Col int //max of 9
	Peg bool
}

//Jump from the Board struct type
func (b Board) Jump(m, o Hole) (Board, Hole, error) {
	result := Board{}
	thole := Hole{}
	for _, r := range b.Holes {
		result.Holes = append(result.Holes, r)
	}
	if !m.Peg {
		//If there is no peg in the moveHole, no jump possible
		return result, thole, fmt.Errorf("No Peg in move hole %d,%d\n", o.Row, o.Col)
	}
	if !o.Peg {
		//If there is no peg in the overHole, no jump possible
		return result, thole, fmt.Errorf("No Peg in over hole %d,%d\n", o.Row, o.Col)
	}
	rDif := m.Row - o.Row
	cDif := o.Col - m.Col
	if cDif == 0 && rDif == 0 {
		//Holes are the same, not valid
		return result, thole, fmt.Errorf("Jump peg and over hole are the same\n")
	}
	if math.Abs(float64(rDif)) > 1 {
		//You can't jump over more than 1 row horizontally
		return result, thole, fmt.Errorf("Invalid horizonal movement %d\n", rDif)
	}
	if rDif > 0 && math.Abs(float64(cDif)) > 1 {
		//You can't jump over more than 1 col vertically
		return result, thole, fmt.Errorf("Invalid vertical movement %d\n", cDif)
	}
	if rDif == 0 && math.Abs(float64(cDif)) > 2 {
		return result, thole, fmt.Errorf("Invalid horizantal movement %d\n", rDif)
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
	targetHole, err := b.GetHole(targetR, targetC)
	if err != nil {
		return result, thole, fmt.Errorf("Target hole(%d,%d) does not exist\n", targetR, targetC)
	}
	if targetHole.Peg {
		return result, thole, fmt.Errorf("Target hole(%d,%d) has a peg in it\n", targetHole.Row, targetHole.Col)
	}
	for k, bh := range result.Holes {
		if bh.Row == m.Row && bh.Col == m.Col {
			result.Holes[k].Peg = false
		}
		if bh.Row == o.Row && bh.Col == o.Col {
			result.Holes[k].Peg = false
		}
		if bh.Row == targetHole.Row && bh.Col == targetHole.Col {
			result.Holes[k].Peg = true
		}
	}
	return result, targetHole, nil
}

//Board contains all the holes that contain the pegs
type Board struct {
	Holes   []Hole
	MoveLog []string
}

//GetHole gets a pointer to a hole based on the row,col coordinates
func (b Board) GetHole(r, c int) (Hole, error) {
	if r < 0 || r > 6 || c < 0 || c > 9 {
		return Hole{}, fmt.Errorf("Hole %d,%d does not exist\n", r, c)
	}
	for _, v := range b.Holes {
		if v.Col == c && v.Row == r {
			return v, nil
		}
	}
	return Hole{}, fmt.Errorf("Hole %d,%d does not exist\n", r, c)
}

//BuildBoard makes a board of peg holes.
//All holes have a peg except one randomly assigned.
//The top row has 1, then
//2,3,4,5 for a total of 16 holes.
func BuildBoard(empty int) (Board, error) {
	var b Board
	if empty < 0 || empty > 15 {
		return b, fmt.Errorf("1st parameter must be >=0 or <=15, you supplied %d", empty)
	}
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
			b.Holes = append(b.Holes, h)
		}
	}
	return b, nil
}

type move struct {
	H Hole
	O Hole
	T Hole
}

func (m move) String() string {
	return fmt.Sprintf("[%d,%d] over [%d,%d] to [%d,%d]", m.H.Row, m.H.Col, m.O.Row, m.O.Col, m.T.Row, m.T.Col)
}

//Solve does a brute force solving of the game
func (b *Board) Solve() error {
	s2 := rand.NewSource(time.Now().UnixNano())
	r2 := rand.New(s2)
	var newBoard = b
	validMove := 0
	for {
		aMoves := []move{}
		o := Hole{}
		var err error
		for _, v := range newBoard.Holes {
			if v.Peg {
				//upleft
				o, err = newBoard.GetHole(v.Row-1, v.Col-1)
				if err == nil {
					_, t, errJ := newBoard.Jump(v, o)
					if errJ == nil {
						aMoves = append(aMoves, move{H: v, O: o, T: t})
					}
				}

				//upright
				o, err = newBoard.GetHole(v.Row-1, v.Col+1)
				if err == nil {
					_, t, errJ := newBoard.Jump(v, o)
					if errJ == nil {
						aMoves = append(aMoves, move{H: v, O: o, T: t})
					}
				}

				//left
				o, err = newBoard.GetHole(v.Row, v.Col-2)
				if err == nil {
					_, t, errJ := newBoard.Jump(v, o)
					if errJ == nil {
						aMoves = append(aMoves, move{H: v, O: o, T: t})
					}
				}
				//right
				o, err = newBoard.GetHole(v.Row, v.Col+2)
				if err == nil {
					_, t, errJ := newBoard.Jump(v, o)
					if errJ == nil {
						aMoves = append(aMoves, move{H: v, O: o, T: t})
					}
				}

				//downleft
				o, err = newBoard.GetHole(v.Row+1, v.Col-1)
				if err == nil {
					_, t, errJ := newBoard.Jump(v, o)
					if errJ == nil {
						aMoves = append(aMoves, move{H: v, O: o, T: t})
					}
				}

				//downright
				o, err = newBoard.GetHole(v.Row+1, v.Col+1)
				if err == nil {
					_, t, errJ := newBoard.Jump(v, o)
					if errJ == nil {
						aMoves = append(aMoves, move{H: v, O: o, T: t})
					}
				}
			}
		}
		if len(aMoves) == 0 {
			//No legal moves left
			newBoard = b
			validMove = 0
			b.MoveLog = []string{}
			continue
		}
		available := r2.Intn(len(aMoves))
		cBoard, _, errN := newBoard.Jump(aMoves[available].H, aMoves[available].O)
		if errN != nil {
			return errN
		}
		validMove++
		b.MoveLog = append(b.MoveLog, fmt.Sprintf("%v", aMoves[available]))
		newBoard = &cBoard
		if validMove == 13 {
			break
		}
	}
	return nil
}

func (b Board) String() string {
	result := "\n"
	for r := 1; r < 6; r++ {
		for c := 1; c < 10; c++ {
			h, err := b.GetHole(r, c)
			mark := " "
			if err == nil {
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
