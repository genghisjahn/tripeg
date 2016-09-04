package tripeg

import (
	"fmt"
	"strings"
	"testing"
)

func TestValidJumpVertical(t *testing.T) {
	b := BuildBoard(1)
	h, err1 := b.GetHole(3, 3)
	if err1 != nil {
		t.Errorf("Can't find hole 3,3")
	}
	o, err2 := b.GetHole(2, 4)
	if err2 != nil {
		t.Errorf("Can't find hole, 2,4")
	}
	_, err := b.Jump(h, o)
	if err != nil {
		t.Fatal("Should have been successful.")
	}
}

func TestValidJumpHorizontal(t *testing.T) {
	b := BuildBoard(6)
	h, err1 := b.GetHole(3, 3)
	if err1 != nil {
		t.Errorf("Can't find hole 3,3")
	}
	o, err2 := b.GetHole(3, 5)
	if err2 != nil {
		t.Errorf("Can't find hole, 3,5")
	}
	_, err := b.Jump(h, o)
	if err != nil {
		t.Fatal("Should have been successful.")
	}
}

func TestInvavidJumpOverHasNoPeg(t *testing.T) {
	b := BuildBoard(6)
	h, err1 := b.GetHole(2, 6)
	if err1 != nil {
		t.Errorf("Can't find hole 2,6")
	}
	o, err2 := b.GetHole(3, 7)
	if err2 != nil {
		t.Errorf("Can't find hole, 3,7")
	}
	_, err := b.Jump(h, o)
	if err != nil {
		if !strings.HasPrefix(err.Error(), "No Peg in over hole") {
			fmt.Println(err)
			t.Fatal("Should have failed with No Peg in over hole.")
		}
	}
}

func TestInvavidTargetPegFull(t *testing.T) {
	b := BuildBoard(6)
	h, err1 := b.GetHole(3, 3)
	if err1 != nil {
		t.Errorf("Can't find hole 3,3")
	}
	o, err2 := b.GetHole(2, 4)
	if err2 != nil {
		t.Errorf("Can't find hole, 2,4")
	}
	_, err := b.Jump(h, o)
	if err != nil {
		if err.Error() != "Target hole(1,5) has a peg in it\n" {
			t.Fatal("Should have failed Target hole(1,5) has a peg in it")
		}
	}
}

func TestInvalidHolePegEmpty(t *testing.T) {

}
