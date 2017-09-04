package Terminal

import (
	"UAV"
)

type TerminalVal struct {
	X         int
	Y         int
	Weight    int
	IsCovered bool
	TermUAV   []UAV.UAVVal
	SignalStrength []float64
}
