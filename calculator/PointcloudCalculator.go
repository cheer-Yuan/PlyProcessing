package calculator

import (
	"../mymath"
	"../reader"
	"gonum.org/v1/gonum/mat"
	"log"
	"math/rand"
	"time"
)


/* Plane represents the coefficient of X is standard equation : A * X + B * y + C * z + D = 0 */
type Plane struct {
	A, B, C, D float64
}

/* Point represents the coordinates of a point (X, Y, Z) in 3 dimension )  */
type Point struct {
	X, Y, Z float64
}

/* ZoneCritical represents a subset of the vertices (zone d'intérêt), defined by their coordinates on X and Y axes */
type ZoneCritical struct {
	Xmin, Xmax, Ymin, Ymax float64
	V                      []reader.VertexFormat
	N                      int
}

/* New_p_by_normal Creates a new plane by a point P(x0, y0, z0) and its normal (x, y, z)
 * @param coordinates of the point and the normal of the plane
 * @return The standard equation of the plane
 */
func New_p_by_normal(x0 float64, y0 float64, z0 float64, x float64, y float64, z float64) (*Plane) {
	return &Plane{
		A: x,
		B: y,
		C: z,
		D: - (x0 * x + y0 * y + z0 * z),
	}
}

/* New_p Creates a new plane by 4 coefficients given directly
 * @param 4 coefficients of  the equation
 * @return The standard equation of the plane
 */
func New_p(a, b, c, d float64) (*Plane) {
	return &Plane{
		A: a,
		B: b,
		C: c,
		D: d,
	}
}

/* New_p_by_vertices Creates a new plane by 4 coefficients given directly
 * @param 4 coefficients of  the equation
 * @return The standard equation of the plane
 */
func New_p_by_vertices(v1, v2, v3 reader.VertexFormat) (*Plane) {
	a := (v2.Ply_y - v1.Ply_y) * (v3.Ply_z - v1.Ply_z) - (v3.Ply_y - v1.Ply_y) * (v2.Ply_z - v1.Ply_z)
	b := (v2.Ply_z - v1.Ply_z) * (v3.Ply_x - v1.Ply_x) - (v3.Ply_z - v1.Ply_z) * (v2.Ply_x - v1.Ply_x)
	c := (v2.Ply_x - v1.Ply_x) * (v3.Ply_y - v1.Ply_y) - (v3.Ply_x - v1.Ply_x) * (v2.Ply_y - v1.Ply_y)
	d := -a * v1.Ply_x - b * v1.Ply_y - c * v1.Ply_z
	return &Plane{
		A: a,
		B: b,
		C: c,
		D: d,
	}
}

/* Ifbelongto Decides if a vertex P(x0, y0, z0) belongs to the plane
 * @param X, y, z : coordinates of the vertex to test, A tolerance of the distance which decide whether the vertex belongs to the plane or not
 * @return The distance between the point and the plane,
 */
func (plane *Plane) Ifbelongto(x float64, y float64, z float64, distanceMin float64) (float64, bool) {
	belong := false
	// calculate the real distance from the vertex to the plane
	distance := mymath.DistPointPlane(x, y, z, plane.A, plane.B, plane.C, plane.D)
	// if the distance is inferior to the minimal distance, the vertex will be regarded as belong to the plane
	if distance < distanceMin {belong = true}
	return distance, belong
}

/* DistAvrPointPlane Calculates the average distance of a slice of headers to the plane
 * @param The slice of headers to whose distances will be calculated
 * @return The average distance
 */
func (plane *Plane) DistAvrPointPlane(vlist []reader.VertexFormat) float64 {
	n := len(vlist)
	d := 0.
	for i := 0; i < n; i++ {
		// sum the distance
		d += mymath.DistPointPlane(vlist[i].Ply_x, vlist[i].Ply_y, vlist[i].Ply_z, plane.A, plane.B, plane.C, plane.D)
	}
	return d / float64(n)
}

/* New_z Creates A new critical zone from A .ply file, by X and y axes
 * @param Intervals of X, y to which defines the critic zone.
 * @return The pointer which points to the critic zone
 */
func New_z(x_min float64, x_max float64, y_min float64, y_max float64, points []reader.VertexFormat) (*ZoneCritical) {
	count := 0

	vertex := make([]reader.VertexFormat, 1)

	for i := 0; i < len(points); i++ {
		if mymath.IfInterval(points[i].Ply_x, x_min, x_max) && mymath.IfInterval(points[i].Ply_y, y_min, y_max) {
			vertex = append(vertex, points[i])
			count ++
		}
	}

	return &ZoneCritical{
		Xmin: x_min,
		Xmax: x_max,
		Ymin: y_min,
		Ymax: y_max,
		V:    vertex,
		N:    count,
	}
}

/* Zone_count Returns the size of a critical zone
 * @return The size of a critical zone
 */
func (z *ZoneCritical) Zone_count() int {
	return z.N
}

/* IfZonebelongto Calculates the percentage of points in the set that belong to the plane
 * @param Struct ZoneCritical defined here below : the groupe of points,  and A tolerance of the distance which decide whether the point belongs to the plane or not
 * @return The percentage of the points in the group which belong to the plane
 */
func (plane *Plane) IfZonebelongto(Zc *ZoneCritical, bias float64) float64 {
	count := 0
	for i := 0; i < Zc.N; i++ {
		distance := mymath.DistPointPlane(Zc.V[i].Ply_x, Zc.V[i].Ply_y, Zc.V[i].Ply_z, plane.A, plane.B, plane.C, plane.D)
		if distance < bias {count++}
	}
	return float64(count) / float64(Zc.N)
}

/* DistVertex Calculates the euclidean distance between 2 vertices
 * @param 2 vertices
 * @return The distance between 2 vertices
 */
func DistVertex(v1, v2 reader.VertexFormat) float64 {
	return mymath.Dist3D(v1.Ply_x, v1.Ply_y, v1.Ply_z, v2.Ply_x, v2.Ply_y, v2.Ply_z)
}

///* PickRandomFace Pick A face randomly among A slice of faces
// * @param vlist : A slice of the faces each composed by several vertices represented by their index in the slice of vertices
// * @return The index of the face picked
//*/
//func PickRandomFace(vlist []plyfile.Vertex) int {
//	rand.Seed(time.Now().Unix())
//	picked := rand.Intn(len(vlist))
//	return picked
//}

///* Search the faces connected to the given face    TO BE COMPLETED
// * @param flist : faces to be searched; root : the indes of the original face; depth : how far will we do the search
// * @return The list of neighbours
// */
//func VoisinFace(flist []plyfile.Face, root int, depth int) []plyfile.Face {
//	neighbours := make([]plyfile.Face, 1)
//	return neighbours
//}


/* PlaneFittingLeastSquare Fit several points to X plane, using the least square method by solving a linear system Ax=b
 * @param vlist : A slice of the faces each composed by several vertices represented by their index in the slice of vertices
 * @return The standard equation of the plane
 */
func PlaneFittingLeastSquare(vlist []reader.VertexFormat) *Plane {
	n := len(vlist)

	// A : matrix of the size n rows and 3 columns. Rows are (xi, yi, 1) in which xi and yi are coordinates x and y of the points
	A := mat.NewDense(n, 3, nil)
	// B : matrix of the size n rows and 1 columns, (zi), the coordinates z of the points
	b := mat.NewDense(n, 1, nil)
	// initializing
	for i := 0; i < n; i++ {
		A.Set(i, 0, vlist[i].Ply_x)
		A.Set(i, 1, vlist[i].Ply_y)
		A.Set(i, 2, 1)
		b.Set(i, 0, vlist[i].Ply_z)
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

	P := New_p(X.At(0, 0), X.At(1, 0), -1, X.At(2, 0))
	return P
}

/* Fit several points to X plane, using the RANSAC method
 * @param vlist : A slice of vertices; ThresholdInline : the difference tolerated when calculating the inlines; ScoreMin : the minimal score which will stop the recurrence; RecurMn : minimal times of recurrence; RecurMax : maximal times of recurrence
 * @return The standard equation of the plane; Index of the points composing that plane
 */
func PlaneFittingRANSAC(vlist []reader.VertexFormat, ThresholdInline float64, ScoreMin float64, RecurMin int, RecurMax int) (*Plane, []int, float64) {
	n := len(vlist)
	highscore := 0.
	BestVertices := make([]int, 3)
	NumIteration := 0
	var P *Plane

	rand.Seed(time.Now().Unix())
	for {
		InlineCount := 0.

		// pick 3 vertices randomly
		p1 := 0
		p2 := 0
		p3 := 0
		for {
			// check if the 3 vertices are different one from another
			p1 = rand.Intn(n)
			p2 = rand.Intn(n)
			p3 = rand.Intn(n)
			if p1 != p2 && p1 != p3 && p2 != p3 {
				break
			}
		}

		// calculate the inline points for the plane formed by these 3 vertices (if the distance to the plane is inferior to the given threshold, this point would be regarded as inline) and take the score
		p := New_p_by_vertices(vlist[p1], vlist[p2], vlist[p3])
		for _, vertex := range vlist {
			if mymath.DistPointPlane(vertex.Ply_x, vertex.Ply_y, vertex.Ply_z, p.A, p.B, p.C, p.D) < ThresholdInline {
				InlineCount += 1
			}
		}
		NumIteration += 1

		// update the score (inline rate). If there is a higher score, note down the plane and 3 vertices which forms it
		score := InlineCount / float64(n)
		if score > highscore {
			highscore = score
			P = New_p(p.A, p.B, p.C, p.D)
			BestVertices[0] = p1
			BestVertices[1] = p2
			BestVertices[2] = p3
		}

		// quit condition : the high score surpass the given parameter, or the number of iteration reach the maximal given in parameter too
		if NumIteration > RecurMin && highscore > ScoreMin || NumIteration >= RecurMax {
			break
		}
	}

	return P, BestVertices, highscore
}




