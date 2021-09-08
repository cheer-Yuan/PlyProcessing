package calculator

import (
	"Implementation-SVD/svd"
	"dataprocessing/mymath"
	"dataprocessing/plyfile"
	//"dataprocessing/reader"
	"gonum.org/v1/gonum/mat"
	"log"
	"math"
	//"math/rand"
	//"time"
)


/* Plane represents the coefficient of X is standard equation : A * X + B * y + C * z + D = 0 */
type Plane64 struct {
	A, B, C, D float64
}
/* Plane represents the coefficient of X is standard equation : A * X + B * y + C * z + D = 0 */
type Plane32 struct {
	A, B, C, D float32
}

/* Point represents the coordinates of a point (X, Y, Z) in 3 dimension )  */
type Point struct {
	X, Y, Z float64
}

/* AnglesDistribution represents a groups of angles. */
type AnglesDistribution struct {
	//minAngle	float32
	//maxAngle	float32
	numRanges  	int
	distribAVE 	[]int
	distrib		[][]int
}

/* New_AnglesDistribution create a group of analysed angles, suppose numRange = 10
*/
func New_AnglesDistribution(angles []float32) *AnglesDistribution {
	numRanges := 10

	var distrib []int
	for i := 0; i < numRanges; i++ {
		distrib = append(distrib, 0)
	}
	halfpi := 1.5707964

	for _, angle := range angles {
		distr := int(math.Floor(float64(angle) / halfpi * float64(numRanges)))
		distrib[distr] ++
	}

	var buff []int
	buff = append(buff, distrib...)
	return &AnglesDistribution{
		numRanges:  numRanges,
		distribAVE: distrib,
	}
}

/* New_p_by_normal Creates a new plane by a point P(x0, y0, z0) and its normal (x, y, z)
 * @param coordinates of the point and the normal of the plane
 * @return The standard equation of the plane
 */
func New_p_by_normal64(x0 float64, y0 float64, z0 float64, x float64, y float64, z float64) *Plane64 {
	return &Plane64{
		A: x,
		B: y,
		C: z,
		D: -(x0*x + y0*y + z0*z),
	}
}
func New_p_by_normal32(x0 float32, y0 float32, z0 float32, x float32, y float32, z float32) *Plane32 {
	return &Plane32{
		A: x,
		B: y,
		C: z,
		D: -(x0*x + y0*y + z0*z),
	}
}

/* New_p Creates a new plane by 4 coefficients given directly
 * @param 4 coefficients of  the equation
 * @return The standard equation of the plane
 */
func New_p64(a, b, c, d float64) *Plane64 {
	return &Plane64{
		A: a,
		B: b,
		C: c,
		D: d,
	}
}
func New_p32(a, b, c, d float32) *Plane32 {
	return &Plane32{
		A: a,
		B: b,
		C: c,
		D: d,
	}
}

/* New_p_by_verticesMono Creates a new plane by 4 coefficients given directly
 * @param 4 coefficients of  the equation
 * @return The standard equation of the plane
 */
func New_p_by_verticesMono64(v1, v2, v3 plyfile.VertexMono64) *Plane64 {
	a := (v2.Y-v1.Y)*(v3.Z-v1.Z) - (v3.Y-v1.Y)*(v2.Z-v1.Z)
	b := (v2.Z-v1.Z)*(v3.X-v1.X) - (v3.Z-v1.Z)*(v2.X-v1.X)
	c := (v2.X-v1.X)*(v3.Y-v1.Y) - (v3.X-v1.X)*(v2.Y-v1.Y)
	d := -a*v1.X - b*v1.Y - c*v1.Z
	return &Plane64{
		A: a,
		B: b,
		C: c,
		D: d,
	}
}
func New_p_by_verticesMono32(v1, v2, v3 plyfile.VertexMono) *Plane32 {
	a := (v2.Y-v1.Y)*(v3.Z-v1.Z) - (v3.Y-v1.Y)*(v2.Z-v1.Z)
	b := (v2.Z-v1.Z)*(v3.X-v1.X) - (v3.Z-v1.Z)*(v2.X-v1.X)
	c := (v2.X-v1.X)*(v3.Y-v1.Y) - (v3.X-v1.X)*(v2.Y-v1.Y)
	d := -a*v1.X - b*v1.Y - c*v1.Z
	return &Plane32{
		A: a,
		B: b,
		C: c,
		D: d,
	}
}

/* Ifbelongto Decides if a vertex P(x0, y0, z0) belongs to the plane
 * @param x, y, z : coordinates of the vertex to test, A tolerance of the distance which decide whether the vertex belongs to the plane or not
 * @return The distance between the point and the plane,
 */
func (plane *Plane64) Ifbelongto64(x float64, y float64, z float64, distanceMin float64) (float64, bool) {
	belong := false
	// calculate the real distance from the vertex to the plane
	distance := mymath.DistPointPlane64(x, y, z, plane.A, plane.B, plane.C, plane.D)
	// if the distance is inferior to the minimal distance, the vertex will be regarded as belong to the plane
	if distance < distanceMin {
		belong = true
	}
	return distance, belong
}
func (plane *Plane32) Ifbelongto32(x float32, y float32, z float32, distanceMin float32) (float32, bool) {
	belong := false
	// calculate the real distance from the vertex to the plane
	distance := mymath.DistPointPlane32(x, y, z, plane.A, plane.B, plane.C, plane.D)
	// if the distance is inferior to the minimal distance, the vertex will be regarded as belong to the plane
	if distance < distanceMin {
		belong = true
	}
	return distance, belong
}

/* DistAvrPointPlaneMono64 Calculates the average distance of a slice of headers to the plane
 * @param The slice of headers to whose distances will be calculated
 * @return The average distance
 */
func (plane *Plane64) DistAvrPointPlaneMono64(vlist  []plyfile.VertexMono64) float64 {
	n := len(vlist)
	d := 0.
	for i := 0; i < n; i++ {
		// sum the distance
		d += mymath.DistPointPlane64(vlist[i].X, vlist[i].Y, vlist[i].Z, plane.A, plane.B, plane.C, plane.D)
	}
	return d / float64(n)
}
func (plane *Plane32) DistAvrPointPlaneMono32(vlist []plyfile.VertexMono) float32 {
	n := len(vlist)
	d := float32(0.)
	for i := 0; i < n; i++ {
		// sum the distance
		d += mymath.DistPointPlane32(vlist[i].X, vlist[i].Y, vlist[i].Z, plane.A, plane.B, plane.C, plane.D)
	}
	return d / float32(n)
}

/* PlaneMonoFittingLeastSquare Fit several points to x plane, using the least square method by solving a linear system Ax=b
 * @param vlist : A slice of the faces each composed by several vertices represented by their index in the slice of vertices
 * @return The standard equation of the plane
 */
func PlaneMonoFittingLeastSquare64(vlist []plyfile.VertexMono64) *Plane64 {
	n := len(vlist)

	// A : matrix of the size n rows and 3 columns. Rows are (xi, yi, 1) in which xi and yi are coordinates x and y of the points
	A := mat.NewDense(n, 3, nil)
	// B : matrix of the size n rows and 1 columns, (zi), the coordinates z of the points
	b := mat.NewDense(n, 1, nil)
	// initializing
	for i := 0; i < n; i++ {
		A.Set(i, 0, vlist[i].X)
		A.Set(i, 1, vlist[i].Y)
		A.Set(i, 2, 1)
		b.Set(i, 0, vlist[i].Z)
	}

	//  compute : A * x = b --> ATA * x = AT * b --> x = (ATA)^-1 * AT * b

	AT := A.T()
	// multiply A by AT at the left side to prepare for the inverse
	var ATA mat.Dense
	ATA.Mul(AT, A)

	// compute the inverse
	var ATAInv mat.Dense
	err := ATAInv.Inverse(&ATA)
	if err != nil {
		log.Fatalf("A is not invertible: %v", err)
	}

	var ATAInvAT mat.Dense
	ATAInvAT.Mul(&ATAInv, AT)
	var X mat.Dense
	X.Mul(&ATAInvAT, b)

	P := New_p64(X.At(0, 0), X.At(1, 0), -1, X.At(2, 0))
	return P
}
func PlaneMonoFittingLeastSquare32(vlist []plyfile.VertexMono) *Plane32 {
	n := len(vlist)

	// A : matrix of the size n rows and 3 columns. Rows are (xi, yi, 1) in which xi and yi are coordinates x and y of the points
	A := mat.NewDense(n, 3, nil)
	// B : matrix of the size n rows and 1 columns, (zi), the coordinates z of the points
	b := mat.NewDense(n, 1, nil)
	// initializing
	for i := 0; i < n; i++ {
		A.Set(i, 0, float64(vlist[i].X))
		A.Set(i, 1, float64(vlist[i].Y))
		A.Set(i, 2, 1)
		b.Set(i, 0, float64(vlist[i].Z))
	}

	//  compute : A * x = b --> ATA * x = AT * b --> x = (ATA)^-1 * AT * b

	AT := A.T()
	// multiply A by AT at the left side to prepare for the inverse
	var ATA mat.Dense
	ATA.Mul(AT, A)

	// compute the inverse
	var ATAInv mat.Dense
	err := ATAInv.Inverse(&ATA)
	if err != nil {
		log.Fatalf("A is not invertible: %v", err)
	}

	var ATAInvAT mat.Dense
	ATAInvAT.Mul(&ATAInv, AT)
	var X mat.Dense
	X.Mul(&ATAInvAT, b)

	P := New_p32(float32(X.At(0, 0)), float32(X.At(1, 0)), -1, float32(X.At(2, 0)))
	return P
}


func PlaneMonoFittingLeastSquare32SVDseq(vlist []plyfile.VertexMono) *Plane32 {
	// initiating
	var verticesData, xyzMean []float32
	xyzMean = append(xyzMean, 0, 0, 0)
	for _, vertex := range vlist {
		verticesData = append(verticesData, vertex.X, vertex.Y, vertex.Z)
		xyzMean[0] += vertex.X
		xyzMean[1] += vertex.Y
		xyzMean[2] += vertex.Z
	}

	// compute the average value of columns (x, y and z)
	xyzMean[0] /= float32(len(vlist))
	xyzMean[1] /= float32(len(vlist))
	xyzMean[2] /= float32(len(vlist))

	// remove the average coordinates
	var centeredData []float32
	j := 0
	for i := 0; i < len(verticesData); i++ {
		centeredData = append(centeredData, verticesData[i] - xyzMean[j])
		j++
		if j == 3 {
			j = 0
		}
	}

	_, _, v := svd.SVD_seq32(centeredData, 0.001, true, true, int32(len(vlist)), 3)

	a := v[2]
	b := v[5]
	c := v[8]
	d := - (a * xyzMean[0] + b * xyzMean[1] + c * xyzMean[2])

	P := New_p32(a, b, c, d)
	return P
}
func PlaneMonoFittingLeastSquare64SVDseq(vlist []plyfile.VertexMono64) *Plane64 {
	// initiating
	var verticesData, xyzMean []float64
	xyzMean = append(xyzMean, 0, 0, 0)
	for _, vertex := range vlist {
		verticesData = append(verticesData, vertex.X, vertex.Y, vertex.Z)
		xyzMean[0] += vertex.X
		xyzMean[1] += vertex.Y
		xyzMean[2] += vertex.Z
	}

	// compute the average value of columns (x, y and z)
	xyzMean[0] /= float64(len(vlist))
	xyzMean[1] /= float64(len(vlist))
	xyzMean[2] /= float64(len(vlist))

	// remove the average coordinates
	var centeredData []float64
	j := 0
	for i := 0; i < len(verticesData); i++ {
		centeredData = append(centeredData, float64(verticesData[i] - xyzMean[j]))
		j++
		if j == 3 {
			j = 0
		}
	}

	_, _, v := svd.SVD_seq(centeredData, 0.001, true, true, len(vlist), 3)

	a := v[2]
	b := v[5]
	c := v[8]
	d := - (a * xyzMean[0] + b * xyzMean[1]+ c * xyzMean[2])

	P := New_p64(a, b, c, d)
	return P
}
