package UAV

import (
	"Point"
	"Util"
)

const (
	MAXHEIGHT int = 5
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

func (u *UAVVal) CheckBoundry(pt Point.PointVal, GridSize int, MaxHeight int) bool {
	if (u.X+pt.X) < 0 || (u.X+pt.X) > float64(GridSize) {
		return false
	}
	if (u.Y+pt.Y) < 0 || (u.Y+pt.Y) > float64(GridSize) {
		return false
	}
	if (u.Z+pt.Z) < 0 || (u.Z+pt.Z) > float64(MaxHeight) {
		return false
	}
	return true

}
func (u *UAVVal) CheckBoundary(coordinate float64, boundary int) bool {
	if coordinate < 0 || coordinate > float64(boundary) {
		return false
	}
	return true
}
func (u *UAVVal) Move(x float64, y float64, z float64, boundary int) {
	if u.CheckBoundary(u.X+x, boundary) {
		u.X += x
	}
	if u.CheckBoundary(u.Y+y, boundary) {
		u.Y += y
	}
	if u.CheckBoundary(u.Z+z, boundary) {
		u.Z += z
	}
}
func (u *UAVVal) MoveByPoint(pt Point.PointVal, boundary int) {
	u.Move(pt.X, pt.Y, pt.Z, boundary)
}
func (u *UAVVal) RandomPoint(GridSize int) Point.PointVal {
	pt := Point.New(0, 0, 0)
	for {
		pt.Set(Util.RandomPoint())
		if u.CheckBoundry(pt, GridSize, MAXHEIGHT) {
			return pt
		}
	}
}
