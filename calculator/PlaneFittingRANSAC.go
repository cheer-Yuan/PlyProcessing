package calculator

import (
	"dataprocessing/mymath"
	"dataprocessing/plyfile"
	"math/rand"
	"time"
)

/* PlaneMonoFittingRANSAC32 Fit several points to x plane, using the RANSAC method
 * @param vlist : A slice of vertices; ThresholdInline : the difference tolerated when calculating the inlines;
 * @param minScore : the minimal score which will stop the recurrence; RecurMn : minimal times of recurrence; RecurMax : maximal times of recurrence
 * @return The standard equation of the plane; Index of the points composing that plane; the percentage of the inline vertices
 */
func PlaneMonoFittingRANSAC32(vlist []plyfile.VertexMono, minDistance float32, minScore float32, minNforLoop int, maxNforLoop int) (*Plane32, []int, float32) {

	n := len(vlist)
	highscore := float32(0)
	BestVertices := make([]int, 3)
	NumIteration := 0
	var P *Plane32

	rand.Seed(time.Now().Unix())
	for {
		InlineCount := float32(0)

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

		// calculate the inlines for the plane formed by these 3 vertices (if the distance to the plane is inferior to the given threshold, this point would be regarded as inline) and take the score
		p := New_p_by_verticesMono32(vlist[p1], vlist[p2], vlist[p3])
		for _, vertex := range vlist {
			if mymath.DistPointPlane32(vertex.X, vertex.Y, vertex.Z, p.A, p.B, p.C, p.D) < minDistance {
				InlineCount += 1
			}
		}
		NumIteration += 1

		// update the score (inline rate). If there is a higher score, note down the plane and 3 vertices which forms it
		score := InlineCount / float32(n)
		if score > highscore {
			highscore = score
			P = New_p32(p.A, p.B, p.C, p.D)
			BestVertices[0] = p1
			BestVertices[1] = p2
			BestVertices[2] = p3
		}

		// quit condition : the high score surpass the given parameter, or the number of iteration reach the maximal given in parameter too
		if NumIteration > minNforLoop && highscore > minScore || NumIteration >= maxNforLoop {
			break
		}
	}

	return P, BestVertices, highscore
}

func PlaneMonoFittingRANSAC64(vlist []plyfile.VertexMono64, minDistance float64, minScore float64, minNforLoop int, maxNforLoop int) (*Plane64, []int, float64) {

	n := len(vlist)
	highscore := float64(0)
	BestVertices := make([]int, 3)
	NumIteration := 0
	var P *Plane64

	rand.Seed(time.Now().Unix())
	for {
		InlineCount := float64(0)

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

		// calculate the inlines for the plane formed by these 3 vertices (if the distance to the plane is inferior to the given threshold, this point would be regarded as inline) and take the score
		p := New_p_by_verticesMono64(vlist[p1], vlist[p2], vlist[p3])
		for _, vertex := range vlist {
			if mymath.DistPointPlane64(vertex.X, vertex.Y, vertex.Z, p.A, p.B, p.C, p.D) < minDistance {
				InlineCount += 1
			}
		}
		NumIteration += 1

		// update the score (inline rate). If there is a higher score, note down the plane and 3 vertices which forms it
		score := InlineCount / float64(n)
		if score > highscore {
			highscore = score
			P = New_p64(p.A, p.B, p.C, p.D)
			BestVertices[0] = p1
			BestVertices[1] = p2
			BestVertices[2] = p3
		}

		// quit condition : the high score surpass the given parameter, or the number of iteration reach the maximal given in parameter too
		if NumIteration > minNforLoop && highscore > minScore || NumIteration >= maxNforLoop {
			break
		}
	}

	return P, BestVertices, highscore
}