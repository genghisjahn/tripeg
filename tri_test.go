package tripeg

import (
	"fmt"
	"strings"
	"testing"
)

func TestValidJumpVertical(t *testing.T) {
	b, errb := BuildBoard(5, 1)
	if errb != nil {
		t.Fatal(errb)
	}
	h, err1 := b.GetHole(3, 3)
	if err1 != nil {
		t.Errorf("Can't find hole 3,3")
	}
	o, err2 := b.GetHole(2, 4)
	if err2 != nil {
		t.Errorf("Can't find hole, 2,4")
	}
	_, _, err := b.Jump(h, o)
	if err != nil {
		t.Fatal("Should have been successful.", err)
	}
}

func TestValidJumpHorizontal(t *testing.T) {
	b, errb := BuildBoard(5, 6)
	if errb != nil {
		t.Fatal(errb)
	}
	h, err1 := b.GetHole(3, 3)
	if err1 != nil {
		t.Errorf("Can't find hole 3,3")
	}
	o, err2 := b.GetHole(3, 5)
	if err2 != nil {
		t.Errorf("Can't find hole, 3,5")
	}
	_, _, err := b.Jump(h, o)
	if err != nil {
		t.Fatal("Should have been successful.")
	}
}

func TestInvavidJumpOverHasNoPeg(t *testing.T) {
	b, errb := BuildBoard(5, 6)
	if errb != nil {
		t.Fatal(errb)
	}
	h, err1 := b.GetHole(2, 6)
	if err1 != nil {
		t.Errorf("Can't find hole 2,6")
	}
	o, err2 := b.GetHole(3, 7)
	if err2 != nil {
		t.Errorf("Can't find hole, 3,7")
	}
	_, _, err := b.Jump(h, o)
	if err != nil {
		if !strings.HasPrefix(err.Error(), "No Peg in over hole") {
			fmt.Println(err)
			t.Fatal("Should have failed with No Peg in over hole.")
		}
	}
}

func TestInvavidTargetPegFull(t *testing.T) {
	b, errb := BuildBoard(5, 6)
	if errb != nil {
		t.Fatal(errb)
	}
	h, err1 := b.GetHole(3, 3)
	if err1 != nil {
		t.Errorf("Can't find hole 3,3")
	}
	o, err2 := b.GetHole(2, 4)
	if err2 != nil {
		t.Errorf("Can't find hole, 2,4")
	}
	_, _, err := b.Jump(h, o)
	if err != nil {
		if err.Error() != "Target hole(1,5) has a peg in it\n" {
			t.Fatal("Should have failed Target hole(1,5) has a peg in it")
		}
	}
}

func TestInvalidHolePegEmpty(t *testing.T) {
	b, errb := BuildBoard(5, 6)
	if errb != nil {
		t.Fatal(errb)
	}
	h, err1 := b.GetHole(3, 7)
	if err1 != nil {
		t.Errorf("Can't find hole 3,7")
	}
	o, err2 := b.GetHole(2, 6)
	if err2 != nil {
		t.Errorf("Can't find hole, 2,6")
	}
	_, _, err := b.Jump(h, o)
	if err != nil {
		if err.Error() != "No Peg in move hole 2,6\n" {
			t.Fatal("Should have failed with No Peg in move hole 2,6")
		}
	}
}

func TestErrorsTwo(t *testing.T) {
	ea := ErrorArray{}
	for r := 0; r < 3; r++ {
		ea.Add(fmt.Errorf("Error %d", r+1))
	}
	expected := "Error 1\nError 2\nError 3"
	received := ea.Error()
	if received != expected {
		t.Fatalf("Expected:\n%v\nReceived:\n%v\n", expected, received)
	}
}

func TestErrorsTwelve(t *testing.T) {
	ea := ErrorArray{}
	for r := 0; r < 13; r++ {
		ea.Add(fmt.Errorf("Error %d", r+1))
	}
	expected := "Error 1\nError 2\nError 3\nError 4\nError 5\nError 6\nError 7\nError 8\nError 9\nError 10\nToo many errors! Count: 12"
	received := ea.Error()
	if received != expected {
		t.Fatalf("Expected:\n%v\nReceived:\n%v\n", expected, received)
	}
}

func TestValidSolve(t *testing.T) {
	b, err := BuildBoard(5, 6)
	if err != nil {
		t.Fatal(err)
	}
	b.Solve()
	mlen := len(b.MoveChart)
	if mlen != 13 {
		t.Fatalf("Expected 13 moves, got %d\n", mlen)
	}
}

func TestGetLastPegForSixRows(t *testing.T) {
	b, err := BuildBoard(6, 1)
	if err != nil {
		t.Fatal(err)
	}
	if _, bErr := b.GetHole(6, 11); bErr != nil {
		t.Fatal(bErr)
	}
}
