package Util

import (
	"Point"
	"Strategy"
)

const (
	STEP float64 = 0.1
)

func RandomPoint() Point.PointVal {
	st := Strategy.RandomStrategy()
	switch st {

	case "FORWARD":
		return Point.New(0, STEP, 0)
	case "BACKWARD":
		return Point.New(0, -STEP, 0)
	case "RIGHT":
		return Point.New(STEP, 0, 0)
	case "LEFT":
		return Point.New(-STEP, 0, 0)
	default:
		return Point.New(0, 0, 0)
	}
}
