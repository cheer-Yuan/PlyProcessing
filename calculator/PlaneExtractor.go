package calculator

import (
	"../mymath"
	"../reader"
	"fmt"
	"math"
	"math/rand"
	"time"
)

//func VoxelSort(vlist []reader.VertexFormat, arr []int) ([]reader.VertexFormat, []int) {
//	return QuickSort(vlist, arr, 0, len(arr)-1)
//}

func QuickSort(vlist []reader.VertexFormat, arr []int, left, right int) ([]reader.VertexFormat, []int) {
	if left < right {
		partitionIndex := partition(vlist, arr, left, right)
		QuickSort(vlist, arr, left, partitionIndex - 1)
		QuickSort(vlist, arr, partitionIndex + 1, right)
	}
	return vlist, arr
}

func partition(vlist []reader.VertexFormat, arr []int, left, right int) int {
	pivot := left
	index := pivot + 1

	for i := index; i <= right; i++ {
		if arr[i] < arr[pivot] {
			arr[i], arr[index] = arr[index], arr[i]
			vlist[i], vlist[index] = vlist[index], vlist[i]
			index += 1
		}
	}
	arr[pivot], arr[index-1] = arr[index-1], arr[pivot]
	vlist[pivot], vlist[index-1] = vlist[index-1], vlist[pivot]
	return index - 1
}

func VertexAverage(vlist []reader.VertexFormat) reader.VertexFormat {
	x, y, z := make([]float64, 0), make([]float64, 0), make([]float64, 0)
	for _, v := range vlist {
		x = append(x, v.Ply_x)
		y = append(y, v.Ply_y)
		z = append(z, v.Ply_z)
	}
	var vertexBuff reader.VertexFormat
	vertexBuff.Ply_x = mymath.Average(x)
	vertexBuff.Ply_y = mymath.Average(y)
	vertexBuff.Ply_z = mymath.Average(z)
	return vertexBuff
}


/* VoxelDownsample Downsample the point cloud file to reduce the computing cost, bu creating voxel grids and calculating the surrounded points' centroid
 * @param vlist : A slice of the vertices; The size of the grid (in meters)
 * @return
 */
func VoxelDownsample(vlist []reader.VertexFormat, gridsize float64) []reader.VertexFormat {
	// the space will be divised into voxels, here we create these grids
	xmin, _, xrange, ymin, _, yrange, zmin, _, _ := GetBoundaries(vlist)
	numX := xrange / gridsize	// how many grids are there in the axis X
	numY := yrange / gridsize	// how many grids are there in the axis Y
	//numZ := zrange / gridsize

	// index of the grid in 1-dimension which contains this v
	//
	//ertex
	gridIndex := make([]int, 0)

	// iterate a vertex to decide its voxel grid position
	for _, v := range vlist {
		indexX := math.Floor(math.Abs(v.Ply_x - xmin) / gridsize)
		indexY := math.Floor(math.Abs(v.Ply_y - ymin) / gridsize)
		indexZ := math.Floor(math.Abs(v.Ply_z - zmin) / gridsize)
		Index := indexX + indexY * numX + indexZ * numY * numX
		gridIndex = append(gridIndex, int(Index))
	}
	fmt.Println(gridIndex[len(vlist) - 1])

	// reorder the points according to its grid container
	vlist, gridIndex = QuickSort(vlist, gridIndex, 0, len(gridIndex)-1)

	// apply the filter
	vertexFiltered := make([]reader.VertexFormat, 0)
	indexMark := 0

	for i := 0; i < len(vlist) - 1; i++ {
		if gridIndex[i] == gridIndex[i + 1] {
			continue
		} else {
			vertextoAverage := vlist[indexMark : i + 1]
			vertexFiltered = append(vertexFiltered, VertexAverage(vertextoAverage))
			indexMark = i
		}
	}

	return vertexFiltered
}

/* PlaneSeqRANSAC Do a sequence of RANSAC fitting on the group of the vertices
 * @param vlist : A slice of the vertices; minDistance : The minimal distance to count a vetex as inline; minVforPlane : The minimal number of inlines to form a plne; maxVtoQuit : the maximal number of remained vertices to terminate the algorithm; iterMax : the maximal number of the iterations
 * @return the slice which contains the planes
 */
func PlaneSeqRANSAC(vlist []reader.VertexFormat, minDistance float64, minVforPlane int, maxVtoQuit int, iterMax int) []Plane {
	rand.Seed(time.Now().Unix())
	ifContinue := true
	planes := make([]Plane, 0)
	numIter := 0

	for ifContinue {
		fmt.Println("Calculating plane No. ", len(planes))
		// pick 3 vertices randomly
		n := len(vlist)
		p1 := 0
		p2 := 0
		p3 := 0

		// check if the 3 vertices are different one from another
		for {
			p1 = rand.Intn(n)
			p2 = rand.Intn(n)
			p3 = rand.Intn(n)
			if p1 != p2 && p1 != p3 && p2 != p3 {
				break
			}
		}

		// build a plane
		p := New_p_by_vertices(vlist[p1], vlist[p2], vlist[p3])

		// count how many vertices would be considered inline regard to this plane
		vetexInline := make([]reader.VertexFormat, 0)
		indexInline := make([]int, 0)
		inlineCount := 0
		for vindex, vertex := range vlist {
			if mymath.DistPointPlane(vertex.Ply_x, vertex.Ply_y, vertex.Ply_z, p.A, p.B, p.C, p.D) < minDistance {
				inlineCount += 1
				vetexInline = append(vetexInline, vertex)
				indexInline = append(indexInline, vindex)
			}
		}

		// if the number of inline vertices surpass the given value, we consider that it is a plane, thus we will add it to the slice. Also we need to delete these inline vertices for the next iteration
		if inlineCount > minVforPlane {

			// adjust the plane and add it to the list of the planes
			newP := PlaneFittingLeastSquare(vetexInline)
			planes = append(planes, *newP)

			// switch the elements to stack the undesired ones at the end of the slice, and finally delete them
			for i, index := range indexInline {
				vlist[index], vlist[len(vlist) - 1 - i] = vlist[len(vlist) - 1 - i], vlist[index]
			}
			vlist = append(vlist[:0], vlist[0:len(vlist) - inlineCount]...)
			numIter = 0
		}

		// decide whether the iteration continues
		numIter++
		fmt.Println(numIter)
		if len(vlist) < maxVtoQuit || numIter > iterMax {
			ifContinue = false
		}
	}

	return planes
}

/* PlaneConsecRANSAC A new invented version of the previous function. A process of batch compisition is added to the beginning of each iteration. RANSAC will be applied to the vertices in the batch. Then the plane obtained will be passed to the generalization and re-adjustment. Another process of redundance detection is also added to check whether the new plane and an old one belong to the same planar facade.
 * @param vlist : A slice of the vertices; minDistance : The minimal distance to count a vetex as inline; minScoreforRAN : the minimal score to reach for obtaine a plane from the function RANSAC; minVforPlane : The minimal number of inlines to form a plne; maxAnglePlanes : the maximal value of the angle to consider the two planes as different; maxVtoQuit : the maximal number of remained vertices to terminate the algorithm; iterMax : the maximal number of the iterations; volumeBatch :
 * @return the slice which contains the planes
 */
func PlaneConsecRANSAC(vlist []reader.VertexFormat, minDistance float64, minScoreforRAN float64, minVforPlane int, maxAnglePlanes float64, maxVtoQuit int, iterMax int, volumeBatch int) []Plane {
	rand.Seed(time.Now().Unix())
	ifContinue := true                            // controls whether the main iteration of the RANSAC will continue
	planes := make([]Plane, 0)                    // the slice containing the plane obtained
	vofPlanes := make([][]reader.VertexFormat, 0) // the slice containing the inlines of the planes
	numIter := 0                                  // count the number of the iteration

	for ifContinue {

		// pick a batch of n vertices randomly from the vertex slice
		vertexBatch := make([]reader.VertexFormat, 0)
		indexBatch := make([]int, 0)
		for {
			n := len(vlist)
			indexRandom := rand.Intn(n)

			// check if this vertex is already chosen. If not, append the vertex and its index to the batch
			if mymath.ExistIntList(indexBatch, indexRandom) {
				continue
			}
			vertexBatch = append(vertexBatch, vlist[indexRandom])
			indexBatch = append(indexBatch, indexRandom)

			// end condition
			if len(vertexBatch) == volumeBatch {
				break
			}
		}

		// do RANSAC fitting within this batch of the points
		P, _, _ := PlaneFittingRANSAC(vertexBatch, minDistance, minScoreforRAN, 100, 2 * volumeBatch)



		// search for inlines to this fitted plane among all vertices
		vetexInline := make([]reader.VertexFormat, 0)
		indexInline := make([]int, 0)
		inlineCount := 0
		for vindex, vertex := range vlist {
			if mymath.DistPointPlane(vertex.Ply_x, vertex.Ply_y, vertex.Ply_z, P.A, P.B, P.C, P.D) < minDistance {
				inlineCount += 1
				vetexInline = append(vetexInline, vertex)
				indexInline = append(indexInline, vindex)
			}
		}

		// if the number of inline vertices surpass the given value, we consider that it is a plane, otherwise we lance a new iteration
		if inlineCount < minVforPlane {
			numIter++
			if len(vlist) < maxVtoQuit || numIter > iterMax {
				ifContinue = false
			}
			continue
		}

		// re-adjust the plane with all inlines
		newP := PlaneFittingLeastSquare(vetexInline)

		// check that if the new plane belongs to the same planar facade as one of the planes already obtained. If it is the case, mix the vertices and re-adjust the two planes using least square method
		if len(planes) > 0 {
			for index, planeObtained := range planes {
				// first check : if the angle formed by the two planes are inferior to the given value
				// second check : if the average distance from the points forming the old plane and the new plane is inferior to the given value
				if mymath.VectorsAngle(planeObtained.A, planeObtained.B, planeObtained.C, newP.A, newP.B, newP.C) < maxAnglePlanes && newP.DistAvrPointPlane(vetexInline) < minDistance {
					// the two planes will be considered the same, we re-adjust de plane with all vertices from the two slices
					fmt.Println("Redundant plane detected")
					newP = PlaneFittingLeastSquare(append(vetexInline, vofPlanes[index]...))
					vofPlanes[index] = append(vofPlanes[index], vetexInline...)
					vetexInline = nil
					break
				}
			}
		}

		// if vetexInline != nil means that no similar planes found, we add the new plane to the slice
		planes = append(planes, *newP)
		vofPlanes = append(vofPlanes, vetexInline)

		// remove the inline points from the original slice
		for i, index := range indexInline {
			vlist[index], vlist[len(vlist) - 1 - i] = vlist[len(vlist) - 1 - i], vlist[index]
		}
		vlist = append(vlist[:0], vlist[0:len(vlist) - inlineCount]...)

		// decide whether the iteration continues
		numIter++
		if len(vlist) < maxVtoQuit || numIter > iterMax {
			ifContinue = false
		}
	}

	fmt.Println("The number of vertex unprocessed : ", len(vlist) )
	fmt.Println("The number of planes generated : ", len(planes), "They are : ")
	for _, i := range planes {
		fmt.Println(i)
	}

	return planes
}
