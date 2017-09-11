package main

import (
	//"fmt"
	"Env"
)

func main() {
	e := Env.NewEnv("poisson_distribute.txt", "uavConfig_height100m.txt", "BioInspired")
	//fmt.Println(e)
	e.GetEnv()
	e.Simulate()

}
