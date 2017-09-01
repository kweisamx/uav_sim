package Env

import (
	"fmt"
	"Grid"
	//"Terminal"
	"UAV"
)
const(
	GRID_SIZE int = 60
	TERMIAL_NUM  int = 360
)

type Environment  struct{
	EnvGridSize int
	EnvTerminalNum int
	EnvGrid Grid.GridVal
	EnvUav []UAV.UAVVal

	
}

func NewEnv(Type string, uavDistri string, UAVtype string)Environment{
	var e Environment
	e.EnvGridSize = GRID_SIZE
	e.EnvTerminalNum = TERMIAL_NUM
	return e	
}

func(e Environment) GetEnv(){
	fmt.Println(e.EnvTerminalNum,e.EnvGridSize)
}

