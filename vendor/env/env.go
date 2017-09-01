package env

import (
	"fmt"
)
const(
	GRID_SIZE int = 60
	TERMIAL_NUM  int = 360
)

type Environment  struct{
	GridSize int
	TerminalNum int
}

func NewEnv()Environment{
	var e Environment
	e.GridSize = GRID_SIZE
	e.TerminalNum = TERMIAL_NUM
	return e	
}

func(e Environment) GetEnv(){
	fmt.Println(e.TerminalNum,e.GridSize)
}

