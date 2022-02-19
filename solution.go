package square

import "math"

// Define custom int type to hold sides number and update CalcSquare signature by replacing #yourTypeNameHere#

// Define constants to represent 0, 3 and 4 sides.  Test uses mnemos: SidesTriangle(==3), SidesSquare(==4), SidesCircle(==0)
// it's like:
// CalcSquare(10.0, SidesTriangle)
// CalcSquare(10.0, SidesSquare)
// CalcSquare(10.0, SidesCircle)

const (
	SidesCircle   = 0
	SidesTriangle = 3
	SidesSquare   = 4
)

type SideNums int

func CalcSquare(sideLen float64, sidesNum SideNums) float64 {
	switch sidesNum {
	case 0:
		return math.Pi * sideLen * sideLen
	case 3:
		var p float64 = (sideLen + sideLen + sideLen) / 2
		return math.Sqrt(p * (p - sideLen) * (p - sideLen) * (p - sideLen))
	case 4:
		ans := float64(sideLen) * float64(sideLen)
		return ans
	}
	return 0.0
}
