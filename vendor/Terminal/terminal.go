package Terminal

import (
	"UAV"
	//"fmt"
)

type TerminalVal struct {
	X              int
	Y              int
	Weight         int
	IsCovered      bool
	TermUAV        []UAV.UAVVal
	SignalStrength []float64
}

func (t *TerminalVal) SetUAV(e []UAV.UAVVal) {
	t.SignalStrength = make([]float64,len(e))
	/*array fill*/
	for i := 0; i < len(t.SignalStrength); i++ {
		t.SignalStrength[i] = -1
	}
	//fmt.Println(t.SignalStrength)
	

}
