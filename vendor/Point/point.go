package Point

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type PointVal struct {
	X float64
	Y float64
	Z float64
}

func New(x float64, y float64, z float64) PointVal {
	p := PointVal{X: x, Y: y, Z: z}
	return p
}

func GetUAVLocs(uavDistri string) ([]PointVal, int) {
	pwd, _ := os.Getwd() //get the pwd path
	fmt.Println(pwd)

	f, err := os.Open(pwd + "/" + uavDistri)
	if err != nil {
		fmt.Println("can't load file")
	}
	sc := bufio.NewScanner(f)
	sc.Scan()
	uav_num, _ := strconv.Atoi(sc.Text())

	pt := make([]PointVal, uav_num)
	for i := 0; i < uav_num; i++ {
		sc.Scan()
		point_parse := strings.Split(sc.Text(), " ")
		pt[i].X, _ = strconv.ParseFloat(point_parse[0], 32)
		pt[i].Y, _ = strconv.ParseFloat(point_parse[1], 32)
		pt[i].Z, _ = strconv.ParseFloat(point_parse[2], 32)
	}
	f.Close()
	return pt, uav_num

}
