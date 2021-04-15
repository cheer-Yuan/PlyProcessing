package mymath

import (
	"math"
	"github.com/cheer-Yuan/goNum"
)

func IfInterval(a float64, b float64, c float64) bool {
	return (a > b) && (a < c)
}

// DistPointPlane calculate the euclidean distance between a point (x, y, z) and a plane
func DistPointPlane(x, y, z, a, b, c, d float64) float64 {
	return math.Abs(float64( x * a + y * b + z * c + d))/ math.Sqrt(float64(a * a + b * b + c * c))
}

// MultMatrix Multiplication of 2 matrixm, the structure is define in the package goNum
func MultMatrix(A, B goNum.Matrix) goNum.Matrix {
	if (A.Columns != B.Rows) {
		panic("goNum.Matrix.Mul: A and B does not match")
	}
	AmulB := goNum.ZeroMatrix(A.Rows, B.Columns)
	for i := 0; i < A.Rows; i++ {
		for j := 0; j < B.Columns; j++ {
			buff := 0.
			for k := 0; k < A.Columns; k++ {
				buff += A.GetFromMatrix(i, k) * B.GetFromMatrix(k, j)
			}
			AmulB.SetMatrix(i, j, buff)
		}
	}
	return AmulB
}

// VectorsAngle calculate angles formed by 2 3-d vectors，oput : [0, 1/2 pi]
func VectorsAngle(V1_x, V1_y, V1_z , V2_x, V2_y, V2_z float64) float64 {
	a := V1_x * V2_x + V1_y * V2_y + V1_z * V2_z
	b := math.Sqrt(math.Pow(V1_x,2) + math.Pow(V1_y,2) + math.Pow(V1_z,2)) * math.Sqrt(math.Pow(V2_x,2) + math.Pow(V2_y,2) + math.Pow(V2_z,2))

	angle := math.Acos(a / b)

	if angle < 0 || angle > 0.5 * math.Pi {
		return VectorsAngle(V1_x, V1_y, V1_z , -V2_x, -V2_y, -V2_z)
	}
	return angle
}