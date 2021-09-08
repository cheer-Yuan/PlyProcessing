package mymath

import (
	"math"
)

func IfInterval(a float64, b float64, c float64) bool {
	return (a > b) && (a < c)
}


// SqrtRootFloat64 Compute the sqrt using the fast square root method, for 64bits float
func SqrtRootFloat64(number float64) float64 {
	var i uint64
	var x, y float64
	f := 1.5
	x = number * 0.5
	y = number
	i = math.Float64bits(y)
	i = 0x5fe6ec85e7de30da - (i >> 1)
	y = math.Float64frombits(i)
	y = y * (f - (x * y * y))
	y = y * (f - (x * y * y))
	return number * y
}
func SqrtRootFloat32(number float32) float32 {
	var i uint32
	var x, y float32
	f := float32(1.5)
	x = number * float32(0.5)
	y = number
	i = math.Float32bits(y)
	i = 0x5f3759df - (i >> 1)
	y = math.Float32frombits(i)
	y = y * (f - (x * y * y))
	y = y * (f - (x * y * y))
	return number * y
}


// VectorsAngle calculate angles formed by 2 3-d vectors，output in in the interval [0, 1/2 * pi]
func VectorsAngle64(V1_x, V1_y, V1_z , V2_x, V2_y, V2_z float64) float64 {
	a := V1_x * V2_x + V1_y * V2_y + V1_z * V2_z
	b := SqrtRootFloat64(math.Pow(V1_x,2) + math.Pow(V1_y,2) + math.Pow(V1_z,2)) * SqrtRootFloat64(math.Pow(V2_x,2) + math.Pow(V2_y,2) + math.Pow(V2_z,2))

	angle := math.Acos(a / b)

	if angle < 0 || angle > 0.5 * math.Pi {
		return VectorsAngle64(V1_x, V1_y, V1_z , -V2_x, -V2_y, -V2_z)
	}
	return angle
}
func VectorsAngle32(V1_x, V1_y, V1_z , V2_x, V2_y, V2_z float32) float32 {
	a := V1_x * V2_x + V1_y * V2_y + V1_z * V2_z
	b := SqrtRootFloat32(V1_x * V1_x + V1_y * V1_y + V1_z * V1_z) * SqrtRootFloat32(V2_x * V2_x + V2_y * V2_y + V2_z * V2_z)

	angle := math.Acos(float64(a / b))

	if angle < 0 || angle > 0.5 * math.Pi {
		return VectorsAngle32(V1_x, V1_y, V1_z , -V2_x, -V2_y, -V2_z)
	}
	return float32(angle)
}

func Dist3D(V1_x, V1_y, V1_z , V2_x, V2_y, V2_z float64) float64 {
	return math.Sqrt(math.Pow(V1_x - V2_x,2) + math.Pow(V1_y - V2_y,2) + math.Pow(V1_z - V2_z,2))
}

func Average(slice []float64) float64 {
	buff := 0.
	for _, i := range slice {
		buff += i
	}
	buff = buff / float64(len(slice))
	return buff
}

// check if the variable is already in the list
func ExistIntList(list []int, val int) bool {
	if len(list) == 0 {
		return false
	}
	for _, i := range list {
		if i == val {
			return true
		}
	}
	return false
}
func ExistIntList32(list []int32, val int32) bool {
	if len(list) == 0 {
		return false
	}
	for _, i := range list {
		if i == val {
			return true
		}
	}
	return false
}

// DistPointPlane4 32bit calculation except for sqrt
func DistPointPlane32(x, y, z, a, b, c, d float32) float32 {
	devised := x * a + y * b + z * c + d
	devisor := a * a + b * b + c * c
	if devised > 0 {
		return devised / SqrtRootFloat32(devisor)
	} else {
		return -devised / SqrtRootFloat32(devisor)
	}
}
func DistPointPlane64(x, y, z, a, b, c, d float64) float64 {
	devised := x * a + y * b + z * c + d
	devisor := a * a + b * b + c * c
	if devised > 0 {
		return devised / SqrtRootFloat64(devisor)
	} else {
		return -devised / SqrtRootFloat64(devisor)
	}
}

