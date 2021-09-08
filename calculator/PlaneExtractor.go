package calculator

import (
	"dataprocessing/mymath"
	"dataprocessing/plyfile"
	"io/ioutil"
	"log"
	"math/rand"
	"time"
)

var r = rand.New(rand.NewSource(time.Now().UnixNano()))

func PlaneMonoConsecRANSAC32(vlist []plyfile.VertexMono, maxDistance64 float64, minScoreforRAN64 float64, minVforPlane int, maxAnglePlanes64 float64, maxVtoQuit int, iterMax int, volumeBatch int) ([]Plane32, int, [][]plyfile.VertexMono) {

	maxDistance  := float32(maxDistance64)
	minScoreforRAN := float32(minScoreforRAN64)
	maxAnglePlanes := float32(maxAnglePlanes64)

	ifContinue := true                                   // controls whether the main iteration of the RANSAC will continue
	planes := make([]Plane32, 0)                           // the slice containing the plane obtained
	verticesOfPlanes := make([][]plyfile.VertexMono, 0) // the slice containing the inlines of the planes
	numIter := 0                                         // count the number of the iteration

	for ifContinue {

		// pick a batch of n vertices randomly from the vertex slice
		vertexBatch := make([]plyfile.VertexMono, 0)
		indexBatch := make([]int, 0)
		for {
			indexRandom := r.Intn(len(vlist))


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
		P, _, _ := PlaneMonoFittingRANSAC32(vertexBatch, maxDistance, minScoreforRAN, 100, 2 * volumeBatch)

		// search for inlines to this fitted plane among all vertices
		vetexInline := make([]plyfile.VertexMono, 0)
		indexInline := make([]int, 0)
		inlineCount := 0
		for vindex, vertex := range vlist {
			if mymath.DistPointPlane32(vertex.X, vertex.Y, vertex.Z, P.A, P.B, P.C, P.D) < maxDistance {

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
		newP := PlaneMonoFittingLeastSquare32(vetexInline)

		// check that if the new plane belongs to the same planar facade as one of the planes already obtained. If it is the case, mix the vertices and re-adjust the two planes using least square method
		if len(planes) > 0 {
			for index, planeObtained := range planes {
				// first check : if the angle formed by the two planes are inferior to the given value
				// second check : if the average distance from the points forming the old plane and the new plane is inferior to the given value
				if mymath.VectorsAngle32(planeObtained.A, planeObtained.B, planeObtained.C, newP.A, newP.B, newP.C) < maxAnglePlanes && planeObtained.DistAvrPointPlaneMono32(vetexInline) < 4 * maxDistance {
					//&& newP.DistAvrPointPlane(vetexInline) <	16 * maxDistance
					// the two planes will be considered the same, we re-adjust de plane with all vertices from the two slices
					//fmt.Println("Redundant plane detected")
					newP = PlaneMonoFittingLeastSquare32(append(vetexInline, verticesOfPlanes[index]...))
					verticesOfPlanes[index] = append(verticesOfPlanes[index], vetexInline...)
					planes[index] = *newP
					vetexInline = nil
				}
			}
		}

		// if vetexInline != nil (empty slice) means that no similar planes found, we add the new plane to the slice. Otherwise there was a similar plane
		if vetexInline != nil {
			planes = append(planes, *newP)
			verticesOfPlanes = append(verticesOfPlanes, vetexInline)
		}

		// remove the inline points from the original slice, color these vertices and move them to a new slice to prepare for the painting
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

	// add the rest of the vertices to the matrix
	verticesOfPlanes = append(verticesOfPlanes, vlist)

	// print the results
	//fmt.Println("The number of vertex unprocessed : ", len(vlist) )
	//fmt.Println("The number of planes generated : ", len(planes), "They are : ")
	//for _, i := range planes {
	//	fmt.Println(i)
	//}

	//fmt.Println(numIter)

	return planes, len(planes), verticesOfPlanes
}
func PlaneMonoConsecRANSAC32SVD(vlist []plyfile.VertexMono, maxDistance64 float64, minScoreforRAN64 float32, minVforPlane int, maxAnglePlanes64 float64, maxVtoQuit int, iterMax int, volumeBatch int) ([]Plane32, int, [][]plyfile.VertexMono) {

	maxDistance  := float32(maxDistance64)
	minScoreforRAN := float32(minScoreforRAN64)
	maxAnglePlanes := float32(maxAnglePlanes64)

	ifContinue := true                                   // controls whether the main iteration of the RANSAC will continue
	planes := make([]Plane32, 0)                           // the slice containing the plane obtained
	verticesOfPlanes := make([][]plyfile.VertexMono, 0) // the slice containing the inlines of the planes
	numIter := 0                                         // count the number of the iteration

	for ifContinue {

		// pick a batch of n vertices randomly from the vertex slice
		vertexBatch := make([]plyfile.VertexMono, 0)
		indexBatch := make([]int, 0)
		for {
			indexRandom := r.Intn(len(vlist))


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
		P, _, _ := PlaneMonoFittingRANSAC32(vertexBatch, maxDistance, minScoreforRAN, 100, 2 * volumeBatch)

		// search for inlines to this fitted plane among all vertices
		vetexInline := make([]plyfile.VertexMono, 0)
		indexInline := make([]int, 0)
		inlineCount := 0
		for vindex, vertex := range vlist {
			if mymath.DistPointPlane32(vertex.X, vertex.Y, vertex.Z, P.A, P.B, P.C, P.D) < maxDistance {

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
		newP := PlaneMonoFittingLeastSquare32SVDseq(vetexInline)

		// check that if the new plane belongs to the same planar facade as one of the planes already obtained. If it is the case, mix the vertices and re-adjust the two planes using least square method
		if len(planes) > 0 {
			for index, planeObtained := range planes {
				// first check : if the angle formed by the two planes are inferior to the given value
				// second check : if the average distance from the points forming the old plane and the new plane is inferior to the given value
				if mymath.VectorsAngle32(planeObtained.A, planeObtained.B, planeObtained.C, newP.A, newP.B, newP.C) < maxAnglePlanes && planeObtained.DistAvrPointPlaneMono32(vetexInline) < 4 * maxDistance {
					//&& newP.DistAvrPointPlane(vetexInline) <	16 * maxDistance
					// the two planes will be considered the same, we re-adjust de plane with all vertices from the two slices
					//fmt.Println("Redundant plane detected")
					newP = PlaneMonoFittingLeastSquare32SVDseq(append(vetexInline, verticesOfPlanes[index]...))
					verticesOfPlanes[index] = append(verticesOfPlanes[index], vetexInline...)
					planes[index] = *newP
					vetexInline = nil
				}
			}
		}

		// if vetexInline != nil (empty slice) means that no similar planes found, we add the new plane to the slice. Otherwise there was a similar plane
		if vetexInline != nil {
			planes = append(planes, *newP)
			verticesOfPlanes = append(verticesOfPlanes, vetexInline)
		}

		// remove the inline points from the original slice, color these vertices and move them to a new slice to prepare for the painting
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

	// add the rest of the vertices to the matrix
	verticesOfPlanes = append(verticesOfPlanes, vlist)

	// print the results
	//fmt.Println("The number of vertex unprocessed : ", len(vlist) )
	//fmt.Println("The number of planes generated : ", len(planes), "They are : ")
	//for _, i := range planes {
	//	fmt.Println(i)
	//}

	//fmt.Println(numIter)

	return planes, len(planes), verticesOfPlanes
}
func PlaneMonoConsecRANSAC64(vlist []plyfile.VertexMono64, maxDistance float64, minScoreforRAN float64, minVforPlane int, maxAnglePlanes float64, maxVtoQuit int, iterMax int, volumeBatch int) ([]Plane64, int, [][]plyfile.VertexMono64) {

	ifContinue := true                                   // controls whether the main iteration of the RANSAC will continue
	planes := make([]Plane64, 0)                           // the slice containing the plane obtained
	verticesOfPlanes := make([][]plyfile.VertexMono64, 0) // the slice containing the inlines of the planes
	numIter := 0                                         // count the number of the iteration

	for ifContinue {

		// pick a batch of n vertices randomly from the vertex slice
		vertexBatch := make([]plyfile.VertexMono64, 0)
		indexBatch := make([]int, 0)
		for {
			indexRandom := r.Intn(len(vlist))


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
		P, _, _ := PlaneMonoFittingRANSAC64(vertexBatch, maxDistance, minScoreforRAN, 100, 2 * volumeBatch)

		// search for inlines to this fitted plane among all vertices
		vetexInline := make([]plyfile.VertexMono64, 0)
		indexInline := make([]int, 0)
		inlineCount := 0
		for vindex, vertex := range vlist {
			if mymath.DistPointPlane64(vertex.X, vertex.Y, vertex.Z, P.A, P.B, P.C, P.D) < maxDistance {

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
		newP := PlaneMonoFittingLeastSquare64(vetexInline)

		// check that if the new plane belongs to the same planar facade as one of the planes already obtained. If it is the case, mix the vertices and re-adjust the two planes using least square method
		if len(planes) > 0 {
			for index, planeObtained := range planes {
				// first check : if the angle formed by the two planes are inferior to the given value
				// second check : if the average distance from the points forming the old plane and the new plane is inferior to the given value
				if mymath.VectorsAngle64(planeObtained.A, planeObtained.B, planeObtained.C, newP.A, newP.B, newP.C) < maxAnglePlanes && planeObtained.DistAvrPointPlaneMono64(vetexInline) < 4 * maxDistance {
					//&& newP.DistAvrPointPlane(vetexInline) <	16 * maxDistance
					// the two planes will be considered the same, we re-adjust de plane with all vertices from the two slices
					//fmt.Println("Redundant plane detected")
					newP = PlaneMonoFittingLeastSquare64(append(vetexInline, verticesOfPlanes[index]...))
					verticesOfPlanes[index] = append(verticesOfPlanes[index], vetexInline...)
					planes[index] = *newP
					vetexInline = nil
				}
			}
		}

		// if vetexInline != nil (empty slice) means that no similar planes found, we add the new plane to the slice. Otherwise there was a similar plane
		if vetexInline != nil {
			planes = append(planes, *newP)
			verticesOfPlanes = append(verticesOfPlanes, vetexInline)
		}

		// remove the inline points from the original slice, color these vertices and move them to a new slice to prepare for the painting
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

	// add the rest of the vertices to the matrix
	verticesOfPlanes = append(verticesOfPlanes, vlist)

	// print the results
	//fmt.Println("The number of vertex unprocessed : ", len(vlist) )
	//fmt.Println("The number of planes generated : ", len(planes), "They are : ")
	//for _, i := range planes {
	//	fmt.Println(i)
	//}

	//fmt.Println(numIter)

	return planes, len(planes), verticesOfPlanes
}
func PlaneMonoConsecRANSAC64SVD(vlist []plyfile.VertexMono64, maxDistance float64, minScoreforRAN float64, minVforPlane int, maxAnglePlanes float64, maxVtoQuit int, iterMax int, volumeBatch int) ([]Plane64, int, [][]plyfile.VertexMono64) {

	ifContinue := true                                   // controls whether the main iteration of the RANSAC will continue
	planes := make([]Plane64, 0)                           // the slice containing the plane obtained
	verticesOfPlanes := make([][]plyfile.VertexMono64, 0) // the slice containing the inlines of the planes
	numIter := 0                                         // count the number of the iteration

	for ifContinue {

		// pick a batch of n vertices randomly from the vertex slice
		vertexBatch := make([]plyfile.VertexMono64, 0)
		indexBatch := make([]int, 0)
		for {
			indexRandom := int(r.Intn(len(vlist)))


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
		P, _, _ := PlaneMonoFittingRANSAC64(vertexBatch, maxDistance, minScoreforRAN, 100, 2 * volumeBatch)

		// search for inlines to this fitted plane among all vertices
		vetexInline := make([]plyfile.VertexMono64, 0)
		indexInline := make([]int, 0)
		inlineCount := 0
		for vindex, vertex := range vlist {
			if mymath.DistPointPlane64(vertex.X, vertex.Y, vertex.Z, P.A, P.B, P.C, P.D) < maxDistance {

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
		newP := PlaneMonoFittingLeastSquare64SVDseq(vetexInline)

		// check that if the new plane belongs to the same planar facade as one of the planes already obtained. If it is the case, mix the vertices and re-adjust the two planes using least square method
		if len(planes) > 0 {
			for index, planeObtained := range planes {
				// first check : if the angle formed by the two planes are inferior to the given value
				// second check : if the average distance  the points forming the old plane and the new plane is inferior to the given value
				if !(mymath.VectorsAngle64(planeObtained.A, planeObtained.B, planeObtained.C, newP.A, newP.B, newP.C) > maxAnglePlanes) && planeObtained.DistAvrPointPlaneMono64(vetexInline) < 4 * maxDistance {
					//&& newP.DistAvrPointPlane(vetexInline) <	16 * maxDistance
					// the two planes will be considered the same, we re-adjust de plane with all vertices from the two slices
					//fmt.Println("Redundant plane detected")
					newP = PlaneMonoFittingLeastSquare64SVDseq(append(vetexInline, verticesOfPlanes[index]...))
					verticesOfPlanes[index] = append(verticesOfPlanes[index], vetexInline...)
					planes[index] = *newP
					vetexInline = nil
				}
			}
		}

		// if vetexInline != nil (empty slice) means that no similar planes found, we add the new plane to the slice. Otherwise there was a similar plane
		if vetexInline != nil {
			planes = append(planes, *newP)
			verticesOfPlanes = append(verticesOfPlanes, vetexInline)
		}

		// remove the inline points from the original slice, color these vertices and move them to a new slice to prepare for the painting
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

	// add the rest of the vertices to the matrix
	verticesOfPlanes = append(verticesOfPlanes, vlist)

	// print the results
	//fmt.Println("The number of vertex unprocessed : ", len(vlist) )
	//fmt.Println("The number of planes generated : ", len(planes), "They are : ")
	//for _, i := range planes {
	//	fmt.Println(i)
	//}

	//fmt.Println(numIter)

	return planes, len(planes), verticesOfPlanes
}


/* ListDir Reads all files in the given routine and returns a slice containing them
 * @return []string : all files' names
 */
func ListDir(filename string) []string {
	names := make([]string, 0)
	files, err := ioutil.ReadDir(filename)
	if err != nil {
		log.Fatal(err)
	}

	for _, f := range files {
		names = append(names, f.Name())
	}
	return names
}

func AngleOfPlanes32(planes []Plane32) []float32 {
	var angles []float32
	for i := 0; i < len(planes) - 1; i++ {
		for j := i + 1; j < len(planes); j++ {
			angles = append(angles, mymath.VectorsAngle32(planes[i].A, planes[i].B, planes[i].C, planes[j].A, planes[j].B, planes[j].C))
		}
	}
	return angles
}

func AngleOfPlanes64(planes []Plane64) []float64 {
	var angles []float64
	for i := 0; i < len(planes) - 1; i++ {
		for j := i + 1; j < len(planes); j++ {
			angles = append(angles, mymath.VectorsAngle64(planes[i].A, planes[i].B, planes[i].C, planes[j].A, planes[j].B, planes[j].C))
		}
	}
	return angles
}


