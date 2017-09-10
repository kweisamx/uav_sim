package BioInspired

import (
	"Point"
	"UAV"
	"fmt"
)

type BioInspiredUAV struct {
	UAV.UAVVal
	LastProfit int
	LastMove   Point.PointVal
	Step       float64
}

func (b *BioInspiredUAV) Run() {
	fmt.Println("hello world")
}
