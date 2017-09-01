package Env

import (
	"Grid"
	"fmt"
	"os"
	//"Terminal"
	"UAV"
	"bufio"
)

const (
	GRID_SIZE   int = 60
	TERMIAL_NUM int = 360
)

type Environment struct {
	EnvGridSize    int
	EnvTerminalNum int
	EnvGrid        Grid.GridVal
	EnvUav         []UAV.UAVVal
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}
func (e Environment) ReadTermsConfig(filename string) {
	pwd, _ := os.Getwd() //get the pwd path
	fmt.Println(pwd)
	//data, err := ioutil.ReadFile(pwd + "/" + filename)
	//check(err)
	//fmt.Println(string(data))
	f, err := os.Open(pwd + "/" + filename)
	check(err)
	sc := bufio.NewScanner(f)
	for sc.Scan() {
		fmt.Println(sc.Text())
	}
	if err := sc.Err(); err != nil {
		check(err)
	}

}

func NewEnv(Type string, uavDistri string, UAVtype string) Environment {
	var e Environment
	e.EnvGridSize = GRID_SIZE
	e.EnvTerminalNum = TERMIAL_NUM
	e.ReadTermsConfig(Type)
	return e
}

func (e Environment) GetEnv() {
	fmt.Println(e.EnvTerminalNum, e.EnvGridSize)
}
