package Env

import (
	"Grid"
	"fmt"
	"os"
	//"Terminal"
	"UAV"
	"bufio"
	//"reflect"
	"Terminal"
	"strconv"
	"strings"
)

const (
	GRID_SIZE   int = 60
	TERMIAL_NUM int = 360
)

type Environment struct {
	EnvGridSize    int
	EnvTerminalNum int
	EnvGrid        [][]Grid.GridVal
	EnvUav         []UAV.UAVVal
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}
func (e *Environment) InitGrid(grid_size int) {
	g := make([][]Grid.GridVal, grid_size)

	for i := range g {
		g[i] = make([]Grid.GridVal, grid_size)
	}
	e.EnvGrid = g
	//fmt.Println(e.EnvGrid)
	//fmt.Println(len(e.EnvGrid))
}

func (e *Environment) ReadTermsConfig(filename string) {
	pwd, _ := os.Getwd() //get the pwd path
	fmt.Println(pwd)

	f, err := os.Open(pwd + "/" + filename)
	check(err)
	sc := bufio.NewScanner(f)
	sc.Scan()                                   //for scan the infomation
	info_parse := strings.Split(sc.Text(), " ") //parse
	e.EnvTerminalNum, err = strconv.Atoi(info_parse[0])
	e.EnvGridSize, err = strconv.Atoi(info_parse[1])

	//init grid
	e.InitGrid(e.EnvGridSize)
	
	//init Terminal
	for i := 0; i < e.EnvTerminalNum; i++ {
		sc.Scan()
		term_parse := strings.Split(sc.Text(), " ")

		x, _ := strconv.Atoi(term_parse[0])
		y, _ := strconv.Atoi(term_parse[1])
		weight, _ := strconv.Atoi(term_parse[2])
		//fmt.Println(x,y,weight)
		var t = Terminal.TerminalVal{x, y, weight, false}
		//fmt.Println(t)
		e.EnvGrid[x][y].TerminalList = append(e.EnvGrid[x][y].TerminalList,t)
		//fmt.Println(e.EnvGrid[x][y])
	}
	fmt.Println("Terminal Initialization Finished")
}

func NewEnv(Type string, uavDistri string, UAVtype string) *Environment {
	var e Environment
	e.EnvGridSize = GRID_SIZE
	e.EnvTerminalNum = TERMIAL_NUM
	e.ReadTermsConfig(Type)
	//fmt.Println(e.EnvTest[0][0])
	return &e
}

func (e *Environment) GetEnv() {
	//fmt.Println(e)
	fmt.Println(e.EnvTerminalNum, e.EnvGridSize)
	//fmt.Println("")
}
