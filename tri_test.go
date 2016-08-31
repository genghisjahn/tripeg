package tripeg

import "testing"

func TestValidJumpVertical(t *testing.T) {
	b := BuildBoard(1)
	h := b.GetHole(3, 3)
	o := b.GetHole(2, 4)
	if !h.Jump(&b, o) {
		t.Fatal("Should have been successful.")
	}
}

func TestValidJumpHorizontal(t *testing.T) {
	b := BuildBoard(6)
	h := b.GetHole(3, 3)
	o := b.GetHole(3, 5)
	if !h.Jump(&b, o) {
		t.Fatal("Should have been successful.")
	}
}

func TestInvavidJumpOverHasNoPeg(t *testing.T) {
	b := BuildBoard(6)
	h := b.GetHole(3, 3)
	o := b.GetHole(3, 5)
	if !h.Jump(&b, o) {
		t.Fatal("Should have been successful.")
	}
	h = b.GetHole(3, 7)
	o = b.GetHole(3, 5)
	if h.Jump(&b, o) {
		t.Fatal("Should have failed.")
	}
}

func TestInvavidTargetPegFull(t *testing.T) {
	b := BuildBoard(6)
	h := b.GetHole(4, 2)
	o := b.GetHole(4, 4)
	if h.Jump(&b, o) {
		t.Fatal("Should have failed.")
	}
}

func TestInvalidHolePegEmpty(t *testing.T) {
	b := BuildBoard(6)
	h := b.GetHole(3, 7)
	o := b.GetHole(3, 5)
	if h.Jump(&b, o) {
		t.Fatal("Should have failed.")
	}
}
