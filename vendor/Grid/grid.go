package Grid

import (
	//"fmt"
	"Terminal"
)

type GridVal struct {
	TerminalNum  int
	TerminalList []Terminal.TerminalVal
}

func (g *GridVal) GetTermNum() int {
	return g.TerminalNum
}

func (g *GridVal) GetTerminal() Terminal.TerminalVal {
	return g.TerminalList[0]
}
