package gshogi

import "strings"

var SQUARE_NAMES = [81]string{
	"9a", "8a", "7a", "6a", "5a", "4a", "3a", "2a", "1a",
	"9b", "8b", "7b", "6b", "5b", "4b", "3b", "2b", "1b",
	"9c", "8c", "7c", "6c", "5c", "4c", "3c", "2c", "1c",
	"9d", "8d", "7d", "6d", "5d", "4d", "3d", "2d", "1d",
	"9e", "8e", "7e", "6e", "5e", "4e", "3e", "2e", "1e",
	"9f", "8f", "7f", "6f", "5f", "4f", "3f", "2f", "1f",
	"9g", "8g", "7g", "6g", "5g", "4g", "3g", "2g", "1g",
	"9h", "8h", "7h", "6h", "5h", "4h", "3h", "2h", "1h",
	"9i", "8i", "7i", "6i", "5i", "4i", "3i", "2i", "1i",
}

type Move struct {
	FromSquare    *Square
	ToSquare      Square
	Promotion     bool
	DropPieceType *PieceType
}

func NewMove(from Square, to Square, promotion bool) *Move {
	return &Move{
		FromSquare:    &from,
		ToSquare:      to,
		Promotion:     promotion,
		DropPieceType: nil,
	}
}

func NewMoveFromHand(to Square, dropPieceType PieceType) *Move {
	return &Move{
		FromSquare:    nil,
		ToSquare:      to,
		Promotion:     false,
		DropPieceType: &dropPieceType,
	}
}

func nameToSquare(s string) Square {
	for i, n := range SQUARE_NAMES {
		if s == n {
			return Square(i)
		}
	}
	return Square(-1)
}

func NewMoveFromUSI(u string) *Move {
	if len(u) == 4 {
		if u[1] == '*' {
			p := NewPieceFromSymbol(u[0:1])
			return NewMoveFromHand(nameToSquare(u[2:4]), p.PieceType)
		} else {
			return NewMove(nameToSquare(u[0:2]), nameToSquare(u[2:4]), false)
		}
	} else {
		// u[4] must be '+'
		return NewMove(nameToSquare(u[0:2]), nameToSquare(u[2:4]), true)
	}
}

func (m *Move) USI() string {
	if m.IsFromHand() {
		return strings.ToUpper(PIECE_SYMBOLS[*m.DropPieceType]) + "*" + SQUARE_NAMES[m.ToSquare]
	} else {
		u := SQUARE_NAMES[*m.FromSquare] + SQUARE_NAMES[m.ToSquare]
		if m.Promotion {
			u += "+"
		}
		return u
	}
}

func (m *Move) IsFromHand() bool {
	return m.FromSquare == nil
}
