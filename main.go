package main

import (
	//"fmt"
	"Env"
)

func main() {
	e:=Env.NewEnv("hi","hi","hi")
	e.GetEnv()
}
