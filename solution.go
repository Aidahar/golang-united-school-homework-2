package square

import "math"

// Define custom int type to hold sides number and update CalcSquare signature by replacing #yourTypeNameHere#

// Define constants to represent 0, 3 and 4 sides.  Test uses mnemos: SidesTriangle(==3), SidesSquare(==4), SidesCircle(==0)
// it's like:
// CalcSquare(10.0, SidesTriangle)
// CalcSquare(10.0, SidesSquare)
// CalcSquare(10.0, SidesCircle)

const (
	SidesCircle   uint8 = 0
	SidesTriangle uint8 = 3
	SidesSquare   uint8 = 4
)

type SideNums uint8

func CalcSquare(sideLen float64, sidesNum SideNums) float64 {
	switch {
	case SidesCircle == uint8(sidesNum):
		return math.Pi * sideLen * sideLen
	case SidesTriangle == uint8(sidesNum):
	case SidesSquare == uint8(sidesNum):
	}
	return 0.0
}
