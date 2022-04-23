package gshogi

import (
	"math/rand"
	"time"
)

type Engine interface {
	Next(*Board) *Move
}

type DefaultEngine struct{}

func (*DefaultEngine) Next(b *Board) *Move {
	ms := b.LegalMoves()
	rand.Seed(time.Now().UnixNano())
	return ms[rand.Intn(len(ms))]
}
