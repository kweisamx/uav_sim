package UAV

import (
	"Point"
)

type UAVVal struct {
	Id         int
	X          float64
	Y          float64
	Z          float64
	LastMove   Point.PointVal
	Step       float64
	LastProfit float64
}
