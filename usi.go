package gshogi

import (
	"fmt"
	"io"
	"strings"
)

type USIHandler struct {
	Engine   Engine
	Board    *Board
	Out      io.Writer
	Name     string
	Author   string
	bestMove *Move
}

func NewUSIHandler(engine Engine, out io.Writer) *USIHandler {
	return &USIHandler{Engine: engine, Board: nil, Out: out, Name: "gshogi", Author: "gshogi"}
}

func (h *USIHandler) send(cmd string) {
	fmt.Fprintln(h.Out, cmd)
}

func (h *USIHandler) debug(msg string) {
	h.send("\n" + msg)
}

func (h *USIHandler) handleUsi(cmds []string) error {
	h.send("id name " + h.Name)
	h.send("id author " + h.Author)
	h.send("usiok")
	return nil
}

func (h *USIHandler) handleIsready(cmds []string) error {
	h.Board = NewBoard()
	h.send("readyok")
	return nil
}

func (h *USIHandler) handleSetoption(cmds []string) error {
	return nil
}

func (h *USIHandler) handleUsinewgame(cmds []string) error {
	return nil
}

func (h *USIHandler) handlePosition(cmds []string) error {
	var sfen string
	switch cmds[1] {
	case "sfen":
		sfen = strings.Join(cmds[2:6], " ")
		cmds = cmds[6:]
	case "startpos":
		sfen = "lnsgkgsnl/1r5b1/ppppppppp/9/9/9/PPPPPPPPP/1B5R1/LNSGKGSNL b - 1"
		cmds = cmds[2:]
	}

	h.Board = NewBoardFromSfen(sfen)

	if len(cmds) == 0 || cmds[0] != "moves" {
		return nil
	}

	for _, c := range cmds[1:] {
		h.Board.PushUSI(c)
	}

	return nil
}

func (h *USIHandler) handleGo(cmds []string) error {
	// TODO: handle cmds

	h.bestMove = h.Engine.Next(h.Board)
	if h.bestMove == nil {
		h.send("bestmove resign")
	} else {
		h.send("bestmove " + h.bestMove.USI())
	}
	return nil
}

func (h *USIHandler) handleStop(cmds []string) error {
	if h.bestMove == nil {
		h.send("bestmove resign")
	} else {
		h.send("bestmove " + h.bestMove.USI())
	}
	return nil
}

func (h *USIHandler) handlePonderhit(cmds []string) error {
	// TODO
	return nil
}

func (h *USIHandler) handleQuit(cmds []string) error {
	return fmt.Errorf("receive quit command")
}

func (h *USIHandler) handleGameover(cmds []string) error {
	// TODO
	return nil
}

func (h *USIHandler) handleDebug(cmds []string) error {
	for _, line := range strings.Split(h.Board.String(), "\n") {
		h.debug(line)
	}
	return nil
}

func (h *USIHandler) Handle(cmd string) error {
	cmds := strings.Split(cmd, " ")
	switch cmds[0] {
	case "usi":
		return h.handleUsi(cmds)
	case "isready":
		return h.handleIsready(cmds)
	case "setoption":
		return h.handleSetoption(cmds)
	case "usinewgame":
		return h.handleUsinewgame(cmds)
	case "position":
		return h.handlePosition(cmds)
	case "go":
		return h.handleGo(cmds)
	case "stop":
		return h.handleStop(cmds)
	case "ponderhit":
		return h.handlePonderhit(cmds)
	case "quit":
		return h.handleQuit(cmds)
	case "gameover":
		return h.handleGameover(cmds)
	case "debug":
		return h.handleDebug(cmds)
	default:
		h.debug("unkown cmd: " + cmd)
		return fmt.Errorf("invalid command: %s", cmd)
	}
}
