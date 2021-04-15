package calculator

import (
	"../mymath"
	"../plyfile"
	"../reader"
	"github.com/cheer-Yuan/goNum"
	"math/rand"
	"time"
)


// plane in 3D : ax + by + cz + D = 0
type Plane struct {
	A, B, C, D float64
}

type Point struct {
	a, b, c float64
}

// Create A new plane by A point P(x0, y0, z0) and its normal vector(x, y, z)
func New_p_by_normal(x0 float64, y0 float64, z0 float64, x float64, y float64, z float64) (*Plane) {

	return &Plane{
		A: x,
		B: y,
		C: z,
		D: - (x0 * x + y0 * y + z0 * z),
	}
}

// Create A new plane by 4 coefficient
func New_p(a, b, c, d float64) (*Plane) {

	return &Plane{
		A: a,
		B: b,
		C: c,
		D: d,
	}
}

// Create A new plane from 3 random points from A group of points (critic zone)
func New_p_by_points(critic *Zone_critic) (*Plane) {

	// pick 3 points randomly
	rand.Seed(time.Now().Unix())
	p1 := 0
	p2 := 0
	p3 := 0
	for {
		p1 = rand.Intn(critic.n)
		p2 = rand.Intn(critic.n)
		p3 = rand.Intn(critic.n)
		if p1 != p2 && p1 != p3 && p2 != p3 {
			break
		}
	}

	a, b, c, d := PlaneConstruct(critic.v[p1], critic.v[p2], critic.v[p3])

	return &Plane{
		A: a,
		B: b,
		C: c,
		D: d,
	}
}

/* Pick A face randomly among A slice of faces
 * @param vlist : A slice of the faces each composed by several vertices represented by their index in the slice of vertices
 * @return The index of the face picked
*/
func PickRandomFace(vlist []plyfile.Vertex) int {
	rand.Seed(time.Now().Unix())
	picked := rand.Intn(len(vlist))
	return picked
}

/* Search the faces connected to the given face    TO BE COMPLETED
 * @param flist : faces to be searched; root : the indes of the original face; depth : how far will we do the search
 * @return The list of neighbours
 */
func VoisinFace(flist []plyfile.Face, root int, depth int) []plyfile.Face {
	neighbours := make([]plyfile.Face, 1)
	return neighbours
}


/* Fit several points to a plane, using the least square method : z = ax + by + C, X = |A B C|T, X = (ATA)-1ATb
 * @param vlist : A slice of the faces each composed by several vertices represented by their index in the slice of vertices
 * @return The function of the plane
 */
func PlaneFittingLeastSquare(vlist []reader.VertexFormat) *Plane {
	n := len(vlist)

	// initiate the matrix
	A := goNum.ZeroMatrix(n, 3)
	b := goNum.ZeroMatrix(n, 1)
	for i := 0; i < n; i++ {
		A.Data[i * 3] = vlist[i].Ply_x
		A.Data[i * 3 + 1] = vlist[i].Ply_y
		A.Data[i * 3 + 2] = 1

		b.Data[i] = vlist[i].Ply_z
	}

	// calculate the matrix od solution
	AT := A.Transpose()
	Inversed, _ := goNum.InverseA(goNum.Matrix2ToSlices(mymath.MultMatrix(AT, A)))

	buff := (mymath.MultMatrix(goNum.Slices2ToMatrix(Inversed), AT))
	X := mymath.MultMatrix(buff, b)

	P := New_p(X.Data[0], X.Data[1], -1, X.Data[2])
	return P
}

/* Fit several points to a plane, using the RANSAC method
 * @param vlist : A slice of vertices; ThresholdInline : the difference tolerated when calculating the inlines; ScoreMin : the minimal score which will stop the recurrence; RecurMn : minimal times of recurrence; RecurMax : maximal times of recurrence
 * @return The function of the plane; Index of the points composing that plane
 */
func PlaneFittingRANSAC(vlist []reader.VertexFormat, ThresholdInline float64, ScoreMin float64, RecurMin int, RecurMax int) (*Plane, []int) {
	n := len(vlist)
	highscore := 0.
	BestVertices := make([]int, 3)
	NumIteration := 0
	var P *Plane
	rand.Seed(time.Now().Unix())

	for {
		InlineCount := 0.

		// pick 3 points randomly
		p1 := 0
		p2 := 0
		p3 := 0
		for {
			p1 = rand.Intn(n)
			p2 = rand.Intn(n)
			p3 = rand.Intn(n)
			if p1 != p2 && p1 != p3 && p2 != p3 {
				break
			}
		}

		a, b, c, d := PlaneConstruct(vlist[p1], vlist[p2], vlist[p3])

		// calculate the inlines (if the distance to the plane is inferior to the given threshold, this point would be regarded as inline) and take the score
		for _, vertex := range vlist {
			if mymath.DistPointPlane(vertex.Ply_x, vertex.Ply_y, vertex.Ply_z, a, b, c, d) < ThresholdInline {
				InlineCount += 1
			}
		}
		NumIteration += 1

		// record update
		score := InlineCount / float64(n)
		if score > highscore {
			highscore = score
			P = New_p(a, b, c, d)
			BestVertices[0] = p1
			BestVertices[1] = p2
			BestVertices[2] = p3
		}

		if NumIteration > RecurMin && highscore > ScoreMin || NumIteration >= RecurMax {
			break
		}
	}

	return P, BestVertices
}




/** Decide if A point P(x0, y0, z0) belongs to the plane
 * @param x, y, z : coordinates of the point to test, A tolerance of the distance which decide whether the point belongs to the plane or not
 * @return The distance between the point and the plane
 */
func (plane *Plane) Ifbelongto(x float64, y float64, z float64, bias float64) (float64, bool) {
	belong := false
	distance := mymath.DistPointPlane(x, y, z, plane.A, plane.B, plane.C, plane.D)
	if distance < bias {belong = true}
	return distance, belong
}

/** Decide if the percentage that A groupe of points belong to the plane
 * @param Struct Zone_critic defined here below : the groupe of points,  and A tolerance of the distance which decide whether the point belongs to the plane or not
 * @return The percentage of the points in the group which belong to the plane
 */
func (plane *Plane) IfZonebelongto(Zc *Zone_critic, bias float64) float64 {
	count := 0
	for i := 0; i < Zc.n; i++ {
		distance := mymath.DistPointPlane(Zc.v[i].Ply_x, Zc.v[i].Ply_y, Zc.v[i].Ply_z, plane.A, plane.B, plane.C, plane.D)
		if distance < bias {count++}
	}
	return float64(count) / float64(Zc.n)
}

/** Calculate the average distance of a slice of headers to this plane
 * @param The slice of headers to whose distances will be calculated
 * @return The average distance
 */
func (plane *Plane) DistAvrPointPlane(vlist []reader.VertexFormat) float64 {
	n := len(vlist)
	d := 0.
	for i := 0; i < n; i++ {
		d += mymath.DistPointPlane(vlist[i].Ply_x, vlist[i].Ply_y, vlist[i].Ply_z, plane.A, plane.B, plane.C, plane.D)
	}
	return d / float64(n)
}


type Zone_critic struct {

	x_min, x_max, y_min, y_max float64
	v []reader.VertexFormat
	n int

}

/** Create A new critical zone from A .ply file, by x and y axes
 * @param Intervals of x, y to which defines the critic zone.
 * @return The pointer which points to the critic zone
 */
func New_z(x_min float64, x_max float64, y_min float64, y_max float64, points []reader.VertexFormat) (*Zone_critic) {
	count := 0

	vertex := make([]reader.VertexFormat, 1)

	for i := 0; i < len(points); i++ {
		if mymath.IfInterval(points[i].Ply_x, x_min, x_max) && mymath.IfInterval(points[i].Ply_y, y_min, y_max) {
			vertex = append(vertex, points[i])
			count ++
		}
	}

	return &Zone_critic{
		x_min: x_min,
		x_max: x_max,
		y_min: y_min,
		y_max: y_max,
		v : vertex,
		n: count,
	}
}

func (z *Zone_critic) Zone_count() int {
	return z.n
}

func (z *Zone_critic) Give_point(num int) reader.VertexFormat {
	return z.v[num]
}

func PlaneConstruct(v1, v2, v3 reader.VertexFormat) (float64, float64, float64, float64) {
	a := (v2.Ply_y - v1.Ply_y) * (v3.Ply_z - v1.Ply_z) - (v3.Ply_y - v1.Ply_y) * (v2.Ply_z - v1.Ply_z)
	b := (v2.Ply_z - v1.Ply_z) * (v3.Ply_x - v1.Ply_x) - (v3.Ply_z - v1.Ply_z) * (v2.Ply_x - v1.Ply_x)
	c := (v2.Ply_x - v1.Ply_x) * (v3.Ply_y - v1.Ply_y) - (v3.Ply_x - v1.Ply_x) * (v2.Ply_y - v1.Ply_y)
	d := -a * v1.Ply_x - b * v1.Ply_y - c * v1.Ply_z
	return a, b, c, d
}