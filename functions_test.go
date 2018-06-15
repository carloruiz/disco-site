package myserver

import "testing"

func TestSquare(t *testing.T) {
	var ans int = 64
	if Square(8) != ans {
		t.Error("square(8) evaluated to ", Square(8))
	}
}
