package gshogi

import "strings"

type PieceType int

const (
	PIECE_TYPE_NONE PieceType = iota
	PIECE_TYPE_PAWN
	PIECE_TYPE_LANCE
	PIECE_TYPE_KNIGHT
	PIECE_TYPE_SILVER
	PIECE_TYPE_GOLD
	PIECE_TYPE_BISHOP
	PIECE_TYPE_ROOK
	PIECE_TYPE_KING
	PIECE_TYPE_PROM_PAWN
	PIECE_TYPE_PROM_LANCE
	PIECE_TYPE_PROM_KNIGHT
	PIECE_TYPE_PROM_SILVER
	PIECE_TYPE_PROM_BISHOP
	PIECE_TYPE_PROM_ROOK
)

var PIECE_SYMBOLS = []string{"", "p", "l", "n", "s", "g", "b", "r", "k", "+p", "+l", "+n", "+s", "+b", "+r"}
var PIECE_JAPANESE_SYMBOLS = []string{"", "歩", "香", "桂", "銀", "金", "角", "飛", "玉", "と", "杏", "圭", "全", "馬", "龍"}

type Piece struct {
	PieceType PieceType
	Color     Color
}

func NewPiece(pieceType PieceType, color Color) *Piece {
	return &Piece{
		PieceType: pieceType,
		Color:     color,
	}
}

func NewPieceFromSymbol(s string) *Piece {
	if s == strings.ToLower(s) {
		return NewPiece(symbolToPieceType(s), WHITE)
	} else {
		return NewPiece(symbolToPieceType(strings.ToLower(s)), BLACK)
	}
}

func symbolToPieceType(s string) PieceType {
	for i, sym := range PIECE_SYMBOLS {
		if s == sym {
			return PieceType(i)
		}
	}
	return PieceType(-1)
}

func IsPieceTypePromoted(pieceType PieceType) bool {
	return pieceType >= PIECE_TYPE_PROM_PAWN
}

func (p *Piece) IsPromoted() bool {
	return IsPieceTypePromoted(p.PieceType)
}

func (p *Piece) Symbol() string {
	if p.Color == BLACK {
		return strings.ToUpper(PIECE_SYMBOLS[p.PieceType])
	} else {
		return PIECE_SYMBOLS[p.PieceType]
	}
}
