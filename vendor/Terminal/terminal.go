package Terminal

import (
	"UAV"
	"fmt"
	//"fmt"
	"math"
)

type TerminalVal struct {
	X              int
	Y              int
	Weight         int
	IsCovered      bool
	TermUAV        []UAV.UAVVal
	SignalStrength []float64
}

const (
	NOT_INSTIALIZED float64 = -1
	TRANSMIT_POWER  float64 = 46.0
	MDSP_THRESHOLD  bool    = false
	MDSP            float64 = -100
)

func (t *TerminalVal) AngleToUAV(uavX float64, uavY float64, uavZ float64) float64 {
	if float64(t.X) == uavX && float64(t.Y) == uavY {
		return 90
	} else {
		distance2D := math.Hypot(float64(t.X)-uavX, float64(t.Y)-uavY)
		return math.Atan2(uavZ, distance2D) * 180 / math.Pi
	}

}

func PathLoss(degree float64, distance float64, uav UAV.UAVVal) float64 {
	if degree > 90 || degree < 0 {
		fmt.Println("some error with degree")
		return 0.0
	}
	if degree >= 0 && degree < 10 {
		return 98.4 + 20*math.Log10(distance) + ((2.55 + degree) / (0.0594 + 0.0406*degree))
	} else if degree >= 10 && degree <= 90 {
		return 98.4 + 20*math.Log10(distance) + ((-94.2 + degree) / (-3.44 + 0.0318*degree))
	} else {
		fmt.Println("something wrong")
		return 0.0
	}
}
func dBmToMiliWatt(dbm float64) float64 {
	return math.Pow(10, dbm/10)
}

func (t *TerminalVal) GetSignalStrenth(uav UAV.UAVVal, mx float64, my float64, mz float64) float64 {
	uavX := uav.X + mx
	uavY := uav.Y + my
	uavZ := uav.Z + mz
	if uavZ < 0.0 {
		uavZ = 0
	}
	degree := t.AngleToUAV(uavX, uavY, uavZ)
	if degree < 0 {
		fmt.Println("degree has problem")
		return 0.0
	}
	distance2D := math.Hypot(float64(t.X)-uavX, float64(t.Y)-uavY)
	distance := math.Hypot(distance2D, uavZ)
	power := TRANSMIT_POWER - PathLoss(degree, distance, uav)

	if MDSP_THRESHOLD && (power < MDSP) {
		return 0.0
	}
	return dBmToMiliWatt(power)
}

func (t *TerminalVal) CollectITF(uavID int) float64 {
	itf := 0.0
	for i := range t.TermUAV {
		if i == uavID {
			continue
		}
		if t.SignalStrength[i] == NOT_INSTIALIZED {
			t.SignalStrength[i] = t.GetSignalStrenth(t.TermUAV[i], 0, 0, 0)
		}
		itf += t.SignalStrength[i]
	}
	return itf
}
func (t *TerminalVal) PeekSIR(uavID int , mx float64 , my float64 , mz float64){
	interference := t.CollectITF(uavID)
	t.SignalStrength[uavID] = t.GetSignalStrenth(t.TermUAV[uavID],mx,my,mz)
	tmp := t.SignalStrength[uavID]
	if 
}

func (t *TerminalVal) SetUAV(e []UAV.UAVVal) {
	t.SignalStrength = make([]float64, len(e))
	/*array fill*/
	for i := 0; i < len(t.SignalStrength); i++ {
		t.SignalStrength[i] = -1
	}
	//fmt.Println(t.SignalStrength)

}
