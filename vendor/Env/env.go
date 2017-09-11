package Env

import (
	"Grid"
	"UAV"
	"fmt"
	"math"
	"os"
	//"Terminal"
	"Point"
	"Terminal"
	"Util"
	"bufio"
	"math/rand"
	"strconv"
	"strings"
)

const (
	GRID_SIZE   int = 60
	TERMIAL_NUM int = 360
	ITERATION   int = 10000
	MAXHEIGHT   int = 5
)

type Environment struct {
	EnvGridSize    int
	EnvTerminalNum int
	EnvGrid        [][]Grid.GridVal
	EnvUav         []UAV.UAVVal
	EnvType        string
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

func (e *Environment) InitTerm(filename string) {
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
		var t = Terminal.TerminalVal{x, y, weight, false, nil, nil}
		//fmt.Println(t)
		e.EnvGrid[x][y].TerminalList = append(e.EnvGrid[x][y].TerminalList, t)
		e.EnvGrid[x][y].TerminalNum += 1
		//fmt.Println(e.EnvGrid[x][y])
		//fmt.Println(e.EnvGrid[x][y].TerminalNum,x,y)
	}
	f.Close()
	fmt.Println("Terminal Initialization Finished")
}
func (e *Environment) InitUAV(uavDistri string, UAVType string) {
	if uavDistri == "" || UAVType == "" {
		err := fmt.Errorf("the argument can't be null")
		check(err)
	}

	pt, pt_num := Point.GetUAVLocs(uavDistri)

	uav := make([]UAV.UAVVal, pt_num)
	for i := 0; i < pt_num; i++ {
		uav[i].Id = i + 1
		uav[i].X = pt[i].X
		uav[i].Y = pt[i].Y
		uav[i].Z = pt[i].Z
		uav[i].LastMove = Point.New(0.0, 0.0, 0.0)
		uav[i].Step = 0.0
		uav[i].LastProfit = 0
	}
	e.EnvUav = uav

	fmt.Println("UAV Initialization Finished")
}
func (e *Environment) BindUAVtoTerm() {
	for i := 0; i < e.EnvGridSize; i++ {
		for j := 0; j < e.EnvGridSize; j++ {
			if e.EnvGrid[i][j].TerminalNum != 0 {
				for k := 0; k < len(e.EnvGrid[i][j].TerminalList); k++ {
					e.EnvGrid[i][j].TerminalList[k].SetUAV(e.EnvUav)
				}
			}

		}
	}
	fmt.Println("Bind to UAV over")
}

func NewEnv(Type string, uavDistri string, UAVType string) *Environment {
	var e Environment
	e.EnvGridSize = GRID_SIZE
	e.EnvTerminalNum = TERMIAL_NUM
	e.EnvType = UAVType
	e.InitTerm(Type)
	e.InitUAV(uavDistri, UAVType)
	e.BindUAVtoTerm()
	return &e
}

func (e *Environment) GetEnv() {
	//fmt.Println(e)
	//fmt.Println(e.EnvTerminalNum, e.EnvGridSize)
	fmt.Println("")
}
func GetRandomSeq(UavNum int) []int {
	seq := make([]int, UavNum)
	//set the number
	for i := range seq {
		seq[i] = i + 1
	}
	//shuffle the seq
	for j := range seq {
		k := rand.Intn(j + 1)
		seq[j], seq[k] = seq[k], seq[j]
	}
	return seq

}

func ld(sir float64) float64 {
	if sir < 0.1 {
		return 0
	} else if sir >= 15 {
		return 4
	} else {
		return math.Log10(sir+1) / math.Log10(2)
	}
}

func (e *Environment) Simulate() {
	for i := 0; i < ITERATION; i++ {
		seq := GetRandomSeq(len(e.EnvUav)) // for each uav that the step is different
		for j := range seq {
			//because i don't understand deeply QQ, so i just write direct , not use call function
			GridSize := e.EnvGridSize
			ServedTerm := 0
			SpectralEfficiency := 0.0
			for k := 0; k < GridSize; k++ {
				for l := 0; l < GridSize; l++ {
					Sir := 0.0
					TermNum := e.EnvGrid[k][l].GetTermNum()
					if TermNum == 0 {
						continue
					}
					t := e.EnvGrid[k][l].GetTerminal()
					point := e.EnvUav[j].LastMove
					Sir = t.PeekSIR(j, point.X, point.Y, point.Z)
					tmp := ld(Sir)
					if tmp > 0 {
						SpectralEfficiency += tmp * float64(TermNum)
						ServedTerm += TermNum
					}

				}
			}
			if ServedTerm > 0 {
				Se := SpectralEfficiency / float64(ServedTerm)
				if e.EnvUav[j].LastProfit > Se {
					e.EnvUav[j].LastMove = e.EnvUav[j].RandomPoint(GridSize)
				}

				e.EnvUav[j].LastProfit = Se
			} else {
				e.EnvUav[j].LastMove = e.EnvUav[j].RandomPoint(GridSize)
			}

			if e.EnvUav[j].LastMove.IsZero() {
				e.EnvUav[j].Step += Util.STEP
			}
			if e.EnvUav[j].LastMove.Z != 0 {
				e.EnvUav[j].MoveByPoint(e.EnvUav[j].LastMove, UAV.MAXHEIGHT)
			} else {
				e.EnvUav[j].MoveByPoint(e.EnvUav[j].LastMove, GridSize)
			}
		}
	}

}
