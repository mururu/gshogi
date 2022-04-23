package gshogi

import (
	"strconv"
	"strings"
)

type Color int

const (
	BLACK Color = iota
	WHITE
)

var COLORS = []Color{BLACK, WHITE}

type Square int

const (
	A9 Square = iota
	A8
	A7
	A6
	A5
	A4
	A3
	A2
	A1
	B9
	B8
	B7
	B6
	B5
	B4
	B3
	B2
	B1
	C9
	C8
	C7
	C6
	C5
	C4
	C3
	C2
	C1
	D9
	D8
	D7
	D6
	D5
	D4
	D3
	D2
	D1
	E9
	E8
	E7
	E6
	E5
	E4
	E3
	E2
	E1
	F9
	F8
	F7
	F6
	F5
	F4
	F3
	F2
	F1
	G9
	G8
	G7
	G6
	G5
	G4
	G3
	G2
	G1
	H9
	H8
	H7
	H6
	H5
	H4
	H3
	H2
	H1
	I9
	I8
	I7
	I6
	I5
	I4
	I3
	I2
	I1
)

var SQUARES = [81]Square{
	A9, A8, A7, A6, A5, A4, A3, A2, A1,
	B9, B8, B7, B6, B5, B4, B3, B2, B1,
	C9, C8, C7, C6, C5, C4, C3, C2, C1,
	D9, D8, D7, D6, D5, D4, D3, D2, D1,
	E9, E8, E7, E6, E5, E4, E3, E2, E1,
	F9, F8, F7, F6, F5, F4, F3, F2, F1,
	G9, G8, G7, G6, G5, G4, G3, G2, G1,
	H9, H8, H7, H6, H5, H4, H3, H2, H1,
	I9, I8, I7, I6, I5, I4, I3, I2, I1,
}

var PIECE_PROMOTED = map[PieceType]PieceType{
	PIECE_TYPE_PAWN:   PIECE_TYPE_PROM_PAWN,
	PIECE_TYPE_LANCE:  PIECE_TYPE_PROM_LANCE,
	PIECE_TYPE_KNIGHT: PIECE_TYPE_PROM_KNIGHT,
	PIECE_TYPE_SILVER: PIECE_TYPE_PROM_SILVER,
	PIECE_TYPE_BISHOP: PIECE_TYPE_PROM_BISHOP,
	PIECE_TYPE_ROOK:   PIECE_TYPE_PROM_ROOK,
}

var PIECE_PROMOTED_REVERSE = map[PieceType]PieceType{
	PIECE_TYPE_PROM_PAWN:   PIECE_TYPE_PAWN,
	PIECE_TYPE_PROM_LANCE:  PIECE_TYPE_LANCE,
	PIECE_TYPE_PROM_KNIGHT: PIECE_TYPE_KNIGHT,
	PIECE_TYPE_PROM_SILVER: PIECE_TYPE_SILVER,
	PIECE_TYPE_PROM_BISHOP: PIECE_TYPE_BISHOP,
	PIECE_TYPE_PROM_ROOK:   PIECE_TYPE_ROOK,
}

var PIECE_TYPES = []PieceType{
	PIECE_TYPE_PAWN,
	PIECE_TYPE_LANCE,
	PIECE_TYPE_KNIGHT,
	PIECE_TYPE_SILVER,
	PIECE_TYPE_GOLD,
	PIECE_TYPE_BISHOP,
	PIECE_TYPE_ROOK,
	PIECE_TYPE_KING,
	PIECE_TYPE_PROM_PAWN,
	PIECE_TYPE_PROM_LANCE,
	PIECE_TYPE_PROM_KNIGHT,
	PIECE_TYPE_PROM_SILVER,
	PIECE_TYPE_PROM_BISHOP,
	PIECE_TYPE_PROM_ROOK,
}

type Board struct {
	Pieces             [81]*Piece
	PiecesInHand       [2]map[PieceType]int
	Turn               Color
	MoveNumber         int
	MoveStack          []*Move
	CapturedPieceStack []PieceType
}

func NewBoard() *Board {
	b := &Board{}
	b.Reset()
	return b
}

func (b *Board) Reset() {
	b.Pieces = [81]*Piece{
		{PieceType: PIECE_TYPE_LANCE, Color: WHITE}, {PieceType: PIECE_TYPE_KNIGHT, Color: WHITE}, {PieceType: PIECE_TYPE_SILVER, Color: WHITE}, {PieceType: PIECE_TYPE_GOLD, Color: WHITE}, {PieceType: PIECE_TYPE_KING, Color: WHITE}, {PieceType: PIECE_TYPE_GOLD, Color: WHITE}, {PieceType: PIECE_TYPE_SILVER, Color: WHITE}, {PieceType: PIECE_TYPE_KNIGHT, Color: WHITE}, {PieceType: PIECE_TYPE_LANCE, Color: WHITE},
		nil, {PieceType: PIECE_TYPE_ROOK, Color: WHITE}, nil, nil, nil, nil, nil, {PieceType: PIECE_TYPE_BISHOP, Color: WHITE}, nil,
		{PieceType: PIECE_TYPE_PAWN, Color: WHITE}, {PieceType: PIECE_TYPE_PAWN, Color: WHITE}, {PieceType: PIECE_TYPE_PAWN, Color: WHITE}, {PieceType: PIECE_TYPE_PAWN, Color: WHITE}, {PieceType: PIECE_TYPE_PAWN, Color: WHITE}, {PieceType: PIECE_TYPE_PAWN, Color: WHITE}, {PieceType: PIECE_TYPE_PAWN, Color: WHITE}, {PieceType: PIECE_TYPE_PAWN, Color: WHITE}, {PieceType: PIECE_TYPE_PAWN, Color: WHITE},
		nil, nil, nil, nil, nil, nil, nil, nil, nil,
		nil, nil, nil, nil, nil, nil, nil, nil, nil,
		nil, nil, nil, nil, nil, nil, nil, nil, nil,
		{PieceType: PIECE_TYPE_PAWN, Color: BLACK}, {PieceType: PIECE_TYPE_PAWN, Color: BLACK}, {PieceType: PIECE_TYPE_PAWN, Color: BLACK}, {PieceType: PIECE_TYPE_PAWN, Color: BLACK}, {PieceType: PIECE_TYPE_PAWN, Color: BLACK}, {PieceType: PIECE_TYPE_PAWN, Color: BLACK}, {PieceType: PIECE_TYPE_PAWN, Color: BLACK}, {PieceType: PIECE_TYPE_PAWN, Color: BLACK}, {PieceType: PIECE_TYPE_PAWN, Color: BLACK},
		nil, {PieceType: PIECE_TYPE_BISHOP, Color: BLACK}, nil, nil, nil, nil, nil, {PieceType: PIECE_TYPE_ROOK, Color: BLACK}, nil,
		{PieceType: PIECE_TYPE_LANCE, Color: BLACK}, {PieceType: PIECE_TYPE_KNIGHT, Color: BLACK}, {PieceType: PIECE_TYPE_SILVER, Color: BLACK}, {PieceType: PIECE_TYPE_GOLD, Color: BLACK}, {PieceType: PIECE_TYPE_KING, Color: BLACK}, {PieceType: PIECE_TYPE_GOLD, Color: BLACK}, {PieceType: PIECE_TYPE_SILVER, Color: BLACK}, {PieceType: PIECE_TYPE_KNIGHT, Color: BLACK}, {PieceType: PIECE_TYPE_LANCE, Color: BLACK},
	}
	b.PiecesInHand[0] = map[PieceType]int{}
	b.PiecesInHand[1] = map[PieceType]int{}
	b.Turn = BLACK
	b.MoveNumber = 1
	b.MoveStack = []*Move{}
	b.CapturedPieceStack = []PieceType{}
}

func NewBoardFromSfen(sfen string) *Board {
	b := &Board{}

	parts := strings.Split(sfen, " ")

	squareIndex := 0
	previousWasPlus := false
	for _, r := range parts[0] {
		switch r {
		case '1', '2', '3', '4', '5', '6', '7', '8', '9':
			squareIndex += (int(r) - 48)
		case '+':
			previousWasPlus = true
		case '/':
			continue
		default:
			pieceSymbol := string(r)
			if previousWasPlus {
				pieceSymbol = "+" + pieceSymbol
			}
			b.SetPieceAt(Square(squareIndex), NewPieceFromSymbol(pieceSymbol))
			squareIndex += 1
			previousWasPlus = false
		}
	}

	if parts[1] == "w" {
		b.Turn = WHITE
	} else {
		b.Turn = BLACK
	}

	b.PiecesInHand[0] = map[PieceType]int{}
	b.PiecesInHand[1] = map[PieceType]int{}
	if parts[2] != "-" {
		pieceCount := 0
		for _, r := range parts[2] {
			switch r {
			case '1', '2', '3', '4', '5', '6', '7', '8', '9':
				pieceCount *= 10
				pieceCount += (int(r) - 48)
			default:
				p := NewPieceFromSymbol(string(r))
				if pieceCount == 0 {
					pieceCount = 1
				}
				for i := 0; i < pieceCount; i++ {
					b.AddPieceIntoHand(p.PieceType, p.Color)
				}
				pieceCount = 0
			}
		}
	}

	i, _ := strconv.Atoi(parts[3])
	b.MoveNumber = i
	b.MoveStack = []*Move{}
	b.CapturedPieceStack = []PieceType{}

	return b
}

func (b *Board) PushUSI(u string) *Move {
	m := NewMoveFromUSI(u)
	b.Push(m)
	return m
}

func (b *Board) Push(m *Move) {
	// Increment move number
	b.MoveNumber += 1

	b.MoveStack = append(b.MoveStack, m)
	if m == nil {
		b.CapturedPieceStack = append(b.CapturedPieceStack, PIECE_TYPE_NONE)
	} else if b.Pieces[m.ToSquare] == nil {
		b.CapturedPieceStack = append(b.CapturedPieceStack, PIECE_TYPE_NONE)
	} else {
		b.CapturedPieceStack = append(b.CapturedPieceStack, b.PieceTypeAt(m.ToSquare))
	}

	// If move is nil, just swap turn
	if m == nil {
		b.Turn ^= 1
		return
	}

	var pieceType PieceType
	if m.IsFromHand() {
		pieceType = *m.DropPieceType
		b.RemovePieceFromHand(pieceType, b.Turn)
	} else {
		pieceType = b.Pieces[*m.FromSquare].PieceType
		if m.Promotion {
			pieceType = PIECE_PROMOTED[pieceType]
		}
		if p := b.Pieces[m.ToSquare]; p != nil {
			intoHandPieceType := p.PieceType
			if IsPieceTypePromoted(intoHandPieceType) {
				intoHandPieceType = PIECE_PROMOTED_REVERSE[intoHandPieceType]
			}
			b.AddPieceIntoHand(intoHandPieceType, b.Turn)
		}
		b.RemovePieceAt(*m.FromSquare)
	}

	piece := NewPiece(pieceType, b.Turn)

	b.SetPieceAt(m.ToSquare, piece)

	// Swap turn
	b.Turn ^= 1
}

func (b *Board) Pop() *Move {
	if len(b.MoveStack) == 0 {
		return nil
	}

	m := b.MoveStack[len(b.MoveStack)-1]
	b.MoveStack = b.MoveStack[:len(b.MoveStack)-1]

	capturedPieceType := b.CapturedPieceStack[len(b.CapturedPieceStack)-1]
	b.CapturedPieceStack = b.CapturedPieceStack[:len(b.CapturedPieceStack)-1]
	capturedPieceColor := b.Turn

	b.MoveNumber -= 1

	if m == nil {
		b.Turn ^= 1
		return m
	}

	pieceType := b.PieceTypeAt(m.ToSquare)
	if m.Promotion {
		pieceType = PIECE_PROMOTED_REVERSE[pieceType]
	}

	if m.IsFromHand() {
		b.AddPieceIntoHand(pieceType, b.Turn^1)
	} else {
		b.SetPieceAt(*m.FromSquare, NewPiece(pieceType, b.Turn^1))
	}

	if capturedPieceType == PIECE_TYPE_NONE {
		b.RemovePieceAt(m.ToSquare)
	} else {
		b.RemovePieceFromHand(capturedPieceType, capturedPieceColor^1)
		b.SetPieceAt(m.ToSquare, NewPiece(capturedPieceType, capturedPieceColor))
	}

	b.Turn ^= 1

	return m
}

func (b *Board) PieceTypeAt(s Square) PieceType {
	return b.Pieces[s].PieceType
}

func (b *Board) RemovePieceAt(s Square) {
	b.Pieces[s] = nil
}

func (b *Board) RemovePieceFromHand(pieceType PieceType, color Color) {
	if IsPieceTypePromoted(pieceType) {
		pieceType = PIECE_PROMOTED_REVERSE[pieceType]
	}

	_, ok := b.PiecesInHand[color][pieceType]
	if !ok && b.PiecesInHand[color][pieceType] <= 0 {
		return
	}

	b.PiecesInHand[color][pieceType] -= 1

	if b.PiecesInHand[color][pieceType] == 0 {
		delete(b.PiecesInHand[color], pieceType)
	}
}

func (b *Board) SetPieceAt(s Square, p *Piece) {
	b.Pieces[s] = p
}

func (b *Board) AddPieceIntoHand(pieceType PieceType, color Color) {
	if IsPieceTypePromoted(pieceType) {
		pieceType = PIECE_PROMOTED_REVERSE[pieceType]
	}

	_, ok := b.PiecesInHand[color][pieceType]
	if !ok {
		b.PiecesInHand[color][pieceType] = 0
	}

	b.PiecesInHand[color][pieceType] += 1
}

func (b *Board) LegalMoves() []*Move {
	moves := []*Move{}
	for _, m := range b.PseudoLegalMoves() {
		if !b.IsSuicideOrCheckByDroppingPawn(m) {
			moves = append(moves, m)
		}
	}
	return moves
}

func (b *Board) IsSuicideOrCheckByDroppingPawn(m *Move) bool {
	b.Push(m)
	s := b.WasSuicide()
	d := b.WasCheckByDroppingPawn(m)
	b.Pop()
	return s || d
}

func (b *Board) WasSuicide() bool {
	return b.IsAttackedBy(b.Turn, b.KingAt(b.Turn^1))
}

func (b *Board) KingAt(color Color) Square {
	for i, p := range b.Pieces {
		if p != nil && p.PieceType == PIECE_TYPE_KING && p.Color == color {
			return Square(i)
		}
	}
	return Square(-1)
}

func (b *Board) IsAttackedBy(color Color, s Square) bool {
	for _, pieceType := range PIECE_TYPES {
		for _, s := range b.AttacksFrom(pieceType, s, color^1) {
			p := b.Pieces[s]
			if p != nil && p.PieceType == pieceType && p.Color == color {
				return true
			}
		}
	}

	return false
}

func (b *Board) WasCheckByDroppingPawn(m *Move) bool {
	// TODO

	if m == nil || !m.IsFromHand() || *m.DropPieceType != PIECE_TYPE_PAWN {
		return false
	}

	return false
}

func (b *Board) PseudoLegalMoves() []*Move {
	moves := []*Move{}
	for i, p := range b.Pieces {
		s := Square(i)
		if p != nil && p.Color == b.Turn {
			toSquares := b.AttacksFrom(p.PieceType, s, b.Turn)
			for _, s2 := range toSquares {
				if b.Pieces[s2] != nil && b.Pieces[s2].Color == b.Turn {
					continue
				}

				if CanMoveWithoutPromotion(s2, p.PieceType, b.Turn) {
					moves = append(moves, NewMove(s, s2, false))
				}

				if CanPromote(s, p.PieceType, b.Turn) || CanPromote(s, p.PieceType, b.Turn) {
					moves = append(moves, NewMove(s, s2, true))
				}
			}
		}
	}

	emptySquares := []Square{}
	for i, p := range b.Pieces {
		if p == nil {
			emptySquares = append(emptySquares, Square(i))
		}
	}
	for pieceType, num := range b.PiecesInHand[b.Turn] {
		if num <= 0 {
			continue
		}
		for _, s := range emptySquares {
			if CanMoveWithoutPromotion(s, pieceType, b.Turn) && !b.isDoublePawn(s, pieceType) {
				moves = append(moves, NewMoveFromHand(s, pieceType))
			}
		}
	}
	return moves
}

func (b *Board) isDoublePawn(to Square, pieceType PieceType) bool {
	if pieceType != PIECE_TYPE_PAWN {
		return false
	}
	for i := int(to); i < 81; i += 9 {
		if p := b.Pieces[i]; p != nil && p.Color == b.Turn && p.PieceType == PIECE_TYPE_PAWN {
			return true
		}
	}

	for i := int(to); i >= 0; i -= 9 {
		if p := b.Pieces[i]; p != nil && p.Color == b.Turn && p.PieceType == PIECE_TYPE_PAWN {
			return true
		}
	}

	return false
}

func (b *Board) AttacksFrom(pieceType PieceType, s Square, color Color) []Square {
	if color == BLACK {
		switch pieceType {
		case PIECE_TYPE_NONE:
			return []Square{}
		case PIECE_TYPE_PAWN:
			return []Square{s - 9}
		case PIECE_TYPE_LANCE:
			ss := []Square{}
			for i := int(s) - 9; i >= 0; i -= 9 {
				p := b.Pieces[i]
				if p == nil {
					ss = append(ss, Square(i))
					continue
				}
				if p.Color != color {
					ss = append(ss, Square(i))
				}
				break
			}
			return ss
		case PIECE_TYPE_KNIGHT:
			return []Square{s - 17, s - 19}
		case PIECE_TYPE_SILVER:
			f := fileIndex(s)
			r := rankIndex(s)
			switch f {
			case 0:
				// left side
				switch r {
				case 0:
					return []Square{s + 10}
				case 8:
					return []Square{s - 9, s - 8}
				default:
					return []Square{s - 9, s - 8, s + 10}
				}
			case 8:
				// right side
				switch r {
				case 0:
					return []Square{s + 8}
				case 8:
					return []Square{s - 10, s - 9}
				default:
					return []Square{s - 10, s - 9, s + 8}
				}
			default:
				switch r {
				case 0:
					return []Square{s + 8, s + 10}
				case 8:
					return []Square{s - 10, s - 9, s - 8}
				default:
					return []Square{s - 10, s - 9, s - 8, s + 8, s + 10}
				}
			}
		case PIECE_TYPE_GOLD:
			return attacksBGold(s)
		case PIECE_TYPE_BISHOP:
			return b.attacksBishop(s, color)
		case PIECE_TYPE_ROOK:
			return b.attacksRook(s, color)
		case PIECE_TYPE_KING:
			f := fileIndex(s)
			r := rankIndex(s)
			switch f {
			case 0:
				// left side
				switch r {
				case 0:
					return []Square{s + 1, s + 9, s + 10}
				case 8:
					return []Square{s - 9, s - 8, s + 1}
				default:
					return []Square{s - 9, s - 8, s + 1, s + 9, s + 10}
				}
			case 8:
				// right side
				switch r {
				case 0:
					return []Square{s - 1, s + 8, s + 9}
				case 8:
					return []Square{s - 10, s - 9, s - 1}
				default:
					return []Square{s - 10, s - 9, s - 1, s + 8, s + 9}
				}
			default:
				switch r {
				case 0:
					return []Square{s - 1, s + 1, s + 8, s + 9, s + 10}
				case 8:
					return []Square{s - 10, s - 9, s - 8, s - 1, s + 1}
				default:
					return []Square{s - 10, s - 9, s - 8, s - 1, s + 1, s + 8, s + 9, s + 10}
				}
			}
		case PIECE_TYPE_PROM_PAWN:
			return attacksBGold(s)
		case PIECE_TYPE_PROM_LANCE:
			return attacksBGold(s)
		case PIECE_TYPE_PROM_KNIGHT:
			return attacksBGold(s)
		case PIECE_TYPE_PROM_SILVER:
			return attacksBGold(s)
		case PIECE_TYPE_PROM_BISHOP:
			return b.attacksPromBishop(s, color)
		case PIECE_TYPE_PROM_ROOK:
			return b.attacksPromRook(s, color)
		}
	} else {
		switch pieceType {
		case PIECE_TYPE_NONE:
			return []Square{}
		case PIECE_TYPE_PAWN:
			return []Square{s + 9}
		case PIECE_TYPE_LANCE:
			ss := []Square{}
			for i := int(s) + 9; i >= 0; i += 9 {
				p := b.Pieces[i]
				if p == nil {
					ss = append(ss, Square(i))
					continue
				}
				if p.Color != color {
					ss = append(ss, Square(i))
				}
				break
			}
			return ss
		case PIECE_TYPE_KNIGHT:
			return []Square{s + 17, s + 19}
		case PIECE_TYPE_SILVER:
			f := fileIndex(s)
			r := rankIndex(s)
			switch f {
			case 0:
				// left side
				switch r {
				case 0:
					return []Square{s + 9, s + 10}
				case 8:
					return []Square{s - 8}
				default:
					return []Square{s - 8, s + 9, s + 10}
				}
			case 8:
				// right side
				switch r {
				case 0:
					return []Square{s + 8, s + 9}
				case 8:
					return []Square{s - 10}
				default:
					return []Square{s - 10, s + 8, s + 9}
				}
			default:
				switch r {
				case 0:
					return []Square{s + 8, s + 9, s + 10}
				case 8:
					return []Square{s - 10, s - 8}
				default:
					return []Square{s - 10, s - 8, s + 8, s + 9, s + 10}
				}
			}
		case PIECE_TYPE_GOLD:
			return attacksWGold(s)
		case PIECE_TYPE_BISHOP:
			return b.attacksBishop(s, color)
		case PIECE_TYPE_ROOK:
			return b.attacksRook(s, color)
		case PIECE_TYPE_KING:
			f := fileIndex(s)
			r := rankIndex(s)
			switch f {
			case 0:
				// left side
				switch r {
				case 0:
					return []Square{s + 1, s + 9, s + 10}
				case 8:
					return []Square{s - 9, s - 8, s + 1}
				default:
					return []Square{s - 9, s - 8, s + 1, s + 9, s + 10}
				}
			case 8:
				// right side
				switch r {
				case 0:
					return []Square{s - 1, s + 8, s + 9}
				case 8:
					return []Square{s - 10, s - 9, s - 1}
				default:
					return []Square{s - 10, s - 9, s - 1, s + 8, s + 9}
				}
			default:
				switch r {
				case 0:
					return []Square{s - 1, s + 1, s + 8, s + 9, s + 10}
				case 8:
					return []Square{s - 10, s - 9, s - 8, s - 1, s + 1}
				default:
					return []Square{s - 10, s - 9, s - 8, s - 1, s + 1, s + 8, s + 9, s + 10}
				}
			}
		case PIECE_TYPE_PROM_PAWN:
			return attacksWGold(s)
		case PIECE_TYPE_PROM_LANCE:
			return attacksWGold(s)
		case PIECE_TYPE_PROM_KNIGHT:
			return attacksWGold(s)
		case PIECE_TYPE_PROM_SILVER:
			return attacksWGold(s)
		case PIECE_TYPE_PROM_BISHOP:
			return b.attacksPromBishop(s, color)
		case PIECE_TYPE_PROM_ROOK:
			return b.attacksPromRook(s, color)
		}
	}
	return []Square{}
}

func attacksBGold(s Square) []Square {
	f := fileIndex(s)
	r := rankIndex(s)
	switch f {
	case 0:
		// left side
		switch r {
		case 0:
			return []Square{s + 1, s + 9}
		case 8:
			return []Square{s - 9, s - 8, s + 1}
		default:
			return []Square{s - 9, s - 8, s + 1, s + 9}
		}
	case 8:
		// right side
		switch r {
		case 0:
			return []Square{s - 1, s + 9}
		case 8:
			return []Square{s - 10, s - 9, s - 1}
		default:
			return []Square{s - 10, s - 9, s - 1, s + 9}
		}
	default:
		switch r {
		case 0:
			return []Square{s - 1, s + 1, s + 9}
		case 8:
			return []Square{s - 10, s - 9, s - 8, s - 1, s + 1}
		default:
			return []Square{s - 10, s - 9, s - 8, s - 1, s + 1, s + 9}
		}
	}
}

func attacksWGold(s Square) []Square {
	f := fileIndex(s)
	r := rankIndex(s)
	switch f {
	case 0:
		// left side
		switch r {
		case 0:
			return []Square{s + 1, s + 9, s + 10}
		case 8:
			return []Square{s - 9, s + 1}
		default:
			return []Square{s - 9, s + 1, s + 9, s + 10}
		}
	case 8:
		// right side
		switch r {
		case 0:
			return []Square{s - 1, s + 8, s + 9}
		case 8:
			return []Square{s - 9, s - 1}
		default:
			return []Square{s - 9, s - 1, s + 8, s + 9}
		}
	default:
		switch r {
		case 0:
			return []Square{s - 1, s + 1, s + 8, s + 9, s + 10}
		case 8:
			return []Square{s - 9, s - 1, s + 1}
		default:
			return []Square{s - 9, s - 1, s + 1, s + 8, s + 9, s + 10}
		}
	}
}

func (b *Board) attacksBishop(s Square, color Color) []Square {
	ss := []Square{}
	// upper left
	for i := int(s) - 10; ; i -= 10 {
		if fileIndex(Square(i+18)) == 8 {
			break
		}
		if i < 0 {
			break
		}
		p := b.Pieces[i]
		if p == nil {
			ss = append(ss, Square(i))
			continue
		}
		if p.Color != color {
			ss = append(ss, Square(i))
		}
		break
	}

	// upper right
	for i := int(s) - 8; ; i -= 8 {
		if fileIndex(Square(i+18)) == 0 {
			break
		}
		if i < 0 {
			break
		}
		p := b.Pieces[i]
		if p == nil {
			ss = append(ss, Square(i))
			continue
		}
		if p.Color != color {
			ss = append(ss, Square(i))
		}
		break
	}

	// lower left
	for i := int(s) + 8; ; i += 8 {
		if fileIndex(Square(i)) == 8 {
			break
		}
		if rankIndex(Square(i)) > 8 {
			break
		}
		p := b.Pieces[i]
		if p == nil {
			ss = append(ss, Square(i))
			continue
		}
		if p.Color != color {
			ss = append(ss, Square(i))
		}
		break
	}

	// lower right
	for i := int(s) + 10; ; i += 10 {
		if fileIndex(Square(i)) == 8 {
			break
		}
		if rankIndex(Square(i)) > 8 {
			break
		}
		p := b.Pieces[i]
		if p == nil {
			ss = append(ss, Square(i))
			continue
		}
		if p.Color != color {
			ss = append(ss, Square(i))
		}
		break
	}

	return ss
}

func (b *Board) attacksPromBishop(s Square, color Color) []Square {
	ss := b.attacksBishop(s, color)

	var around []Square
	f := fileIndex(s)
	r := rankIndex(s)
	switch f {
	case 0:
		// left side
		switch r {
		case 0:
			around = []Square{s + 1, s + 9}
		case 8:
			around = []Square{s - 9, s + 1}
		default:
			around = []Square{s - 9, s + 1, s + 9}
		}
	case 8:
		// right side
		switch r {
		case 0:
			around = []Square{s - 1, s + 9}
		case 8:
			around = []Square{s - 9, s - 1}
		default:
			around = []Square{s - 9, s - 1, s + 9}
		}
	default:
		switch r {
		case 0:
			around = []Square{s - 1, s + 1, s + 9}
		case 8:
			around = []Square{s - 9, s - 1, s + 1}
		default:
			around = []Square{s - 9, s - 1, s + 1, s + 9}
		}
	}

	return append(ss, around...)

}

func (b *Board) attacksRook(s Square, color Color) []Square {
	ss := []Square{}
	for i := int(s) - 9; i >= 0; i -= 9 {
		p := b.Pieces[i]
		if p == nil {
			ss = append(ss, Square(i))
			continue
		}
		if p.Color != color {
			ss = append(ss, Square(i))
		}
		break
	}
	for i := int(s) + 9; i >= 0; i += 9 {
		p := b.Pieces[i]
		if p == nil {
			ss = append(ss, Square(i))
			continue
		}
		if p.Color != color {
			ss = append(ss, Square(i))
		}
		break
	}
	for i := int(s) + 1; ; i += 1 {
		if fileIndex(Square(i)) == 0 {
			break
		}
		p := b.Pieces[i]
		if p == nil {
			ss = append(ss, Square(i))
			continue
		}
		if p.Color != color {
			ss = append(ss, Square(i))
		}
		break
	}
	for i := int(s) - 1; ; i -= 1 {
		if fileIndex(Square(i+9)) == 8 {
			break
		}
		p := b.Pieces[i]
		if p == nil {
			ss = append(ss, Square(i))
			continue
		}
		if p.Color != color {
			ss = append(ss, Square(i))
		}
		break
	}

	return ss
}

func (b *Board) attacksPromRook(s Square, color Color) []Square {
	ss := b.attacksRook(s, color)

	var around []Square
	f := fileIndex(s)
	r := rankIndex(s)
	switch f {
	case 0:
		// left side
		switch r {
		case 0:
			around = []Square{s + 10}
		case 8:
			around = []Square{s - 8}
		default:
			around = []Square{s - 8, s - 10}
		}
	case 8:
		// right side
		switch r {
		case 0:
			around = []Square{s + 8}
		case 8:
			around = []Square{s - 10}
		default:
			around = []Square{s - 10, s + 8}
		}
	default:
		switch r {
		case 0:
			around = []Square{s + 8, s + 10}
		case 8:
			around = []Square{s - 10, s - 8}
		default:
			around = []Square{s - 10, s - 8, s + 8, s + 10}
		}
	}

	return append(ss, around...)
}

func fileIndex(s Square) int {
	return int(s) % 9
}

func rankIndex(s Square) int {
	return int(s) / 9
}

func CanPromote(from Square, pieceType PieceType, color Color) bool {
	switch pieceType {
	case PIECE_TYPE_PAWN, PIECE_TYPE_LANCE, PIECE_TYPE_KNIGHT, PIECE_TYPE_SILVER, PIECE_TYPE_BISHOP, PIECE_TYPE_ROOK:
		if color == BLACK {
			return rankIndex(from) <= 2
		} else {
			return rankIndex(from) >= 6
		}
	default:
		return false
	}
}

func CanMoveWithoutPromotion(to Square, pieceType PieceType, color Color) bool {
	if color == BLACK {
		return (pieceType != PIECE_TYPE_PAWN && pieceType != PIECE_TYPE_LANCE && pieceType != PIECE_TYPE_KNIGHT) ||
			(pieceType == PIECE_TYPE_PAWN && rankIndex(to) > 0) ||
			(pieceType == PIECE_TYPE_LANCE && rankIndex(to) > 0) ||
			(pieceType == PIECE_TYPE_KNIGHT && rankIndex(to) > 1)
	} else {
		return (pieceType != PIECE_TYPE_PAWN && pieceType != PIECE_TYPE_LANCE && pieceType != PIECE_TYPE_KNIGHT) ||
			(pieceType == PIECE_TYPE_PAWN && rankIndex(to) < 8) ||
			(pieceType == PIECE_TYPE_LANCE && rankIndex(to) < 8) ||
			(pieceType == PIECE_TYPE_KNIGHT && rankIndex(to) < 7)
	}
}

func (b *Board) String() string {
	pieces := []string{}
	for _, x := range SQUARES {
		s := ""
		p := b.Pieces[x]
		if p == nil {
			s = " ."
		} else {
			if !p.IsPromoted() {
				s += " "
			}
			s += p.Symbol()
		}
		pieces = append(pieces, s)
	}

	raws := []string{}
	for i := 0; i < 9; i++ {
		raws = append(raws, strings.Join(pieces[i*9:i*9+9], " "))
	}

	buf := strings.Join(raws, "\n")

	if len(b.PiecesInHand[BLACK])+len(b.PiecesInHand[WHITE]) > 0 {
		buf += "\n\n"

		for _, c := range COLORS {
			for k, v := range b.PiecesInHand[c] {
				if v > 0 {
					buf += " "
					buf += NewPiece(k, c).Symbol()
					buf += "*"
					buf += strconv.Itoa(v)
				}
			}
		}
	}

	return buf
}
