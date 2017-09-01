package main

import (
	//"fmt"
	"Env"
)

func main() {
	e := Env.NewEnv("poisson_distribute.txt", "hi", "hi")
	e.GetEnv()
}
