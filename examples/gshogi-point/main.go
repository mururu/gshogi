package main

import (
	"bufio"
	"log"
	"math/rand"
	"os"
	"time"

	"github.com/mururu/gshogi"
)

type PointEngine struct{}

func (engine *PointEngine) Next(b *gshogi.Board) *gshogi.Move {
	maxScore := 0
	candidates := []*gshogi.Move{}

	for _, move := range b.LegalMoves() {
		b.Push(move)
		score := engine.calculatePoint(b)
		if score == maxScore {
			candidates = append(candidates, move)
		}
		if score > maxScore {
			maxScore = score
			candidates = []*gshogi.Move{move}
		}
		b.Pop()
	}

	rand.Seed(time.Now().UnixNano())
	return candidates[rand.Intn(len(candidates))]
}

func (*PointEngine) calculatePoint(b *gshogi.Board) int {
	point := 0

	pointMap := map[gshogi.PieceType]int{
		gshogi.PIECE_TYPE_PAWN:        100,
		gshogi.PIECE_TYPE_LANCE:       200,
		gshogi.PIECE_TYPE_KNIGHT:      200,
		gshogi.PIECE_TYPE_SILVER:      300,
		gshogi.PIECE_TYPE_GOLD:        400,
		gshogi.PIECE_TYPE_BISHOP:      500,
		gshogi.PIECE_TYPE_ROOK:        600,
		gshogi.PIECE_TYPE_KING:        200,
		gshogi.PIECE_TYPE_PROM_PAWN:   200,
		gshogi.PIECE_TYPE_PROM_LANCE:  300,
		gshogi.PIECE_TYPE_PROM_KNIGHT: 400,
		gshogi.PIECE_TYPE_PROM_SILVER: 500,
		gshogi.PIECE_TYPE_PROM_BISHOP: 600,
		gshogi.PIECE_TYPE_PROM_ROOK:   700,
	}

	for _, piece := range b.Pieces {
		if piece != nil && piece.Color == b.Turn^1 {
			point += pointMap[piece.PieceType]
		}
	}

	for pieceType, num := range b.PiecesInHand[b.Turn^1] {
		point += pointMap[pieceType] * num
	}

	return point
}

func main() {
	engine := &PointEngine{}
	handler := gshogi.NewUSIHandler(engine, os.Stdout)
	s := bufio.NewScanner(os.Stdin)

	for s.Scan() {
		t := s.Text()
		if err := handler.Handle(t); err != nil {
			break
		}
	}
	if s.Err() != nil {
		log.Fatal(s.Err())
	}
}
