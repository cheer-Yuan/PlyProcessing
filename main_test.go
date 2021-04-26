package main

import (
	"./calculator"
	"./mymath"
	"./reader"
	"fmt"
	"log"
	"testing"
	"time"
)

func TestExp1(t *testing.T) {
	// this ply file is composed by 3 surfaces : the cabinet, the shredder and the floor, perpendicular one to another (please check the .png file under the same directory)
	filename := "./data/cabinet_paperdestroyer_ground/2021-04-15-12:10:03-.ply"

	// Reading the file using adopted library plyfile
	vlist, _ := reader.ReadPLY(filename)

	// V1 : the surface of the cabinet
	// V1 := calculator.New_z(-0.37, -0.25, -0.15, -0.05, vlist)
	V1 := calculator.New_z(-0.37, -0.25, 0.18, 999, vlist)
	fmt.Println("Number of points in V1 : ", V1.Zone_count())
	// V2 : the surface of the shredder
	V2 := calculator.New_z(0.2, 0.31, 0.15, 999, vlist)
	fmt.Println("Number of points in V2 : ", V2.Zone_count())
	// V3 : the surface of the floor
	V3 := calculator.New_z(-999, 999, -999, -0.142, vlist)
	fmt.Println("Number of points in V3 : ", V3.Zone_count())
	// fit the vertices in V1 to a plane using the RANSAC method
	PlaneRANSAC1, _, hs1 := calculator.PlaneFittingRANSAC(V1.V , 0.005, 0.98,500, 4 * len(V1.V))
	fmt.Println("Plane fitted by RANSAC for group of dots V1 : ", PlaneRANSAC1)
	fmt.Println("Average distance between the vertices and the plane is : ", PlaneRANSAC1.DistAvrPointPlane(V1.V), "  meters")
	log.Println("The high score is : ", hs1)
	fmt.Print("\n")
	// fit the vertices in V1 to a plane using the least square method
	PlaneLS1 := calculator.PlaneFittingLeastSquare(V1.V)
	fmt.Println("Plane fitted by least square for group of dots V1 : ", PlaneLS1)
	fmt.Println("Average distance between the vertices and the plane is : ", PlaneLS1.DistAvrPointPlane(V1.V), "  meters")
	fmt.Print("\n")
	fmt.Print("\n")

	// fit the vertices in V2 to a plane using the 2 methods
	PlaneRANSAC2, _, hs2 := calculator.PlaneFittingRANSAC(V2.V, 0.005, 0.98,500, 4 * len(V2.V))
	fmt.Println("Plane fitted by RANSAC for group of dots V2 : ", PlaneRANSAC2)
	fmt.Println("Average distance between the vertices and the plane is : ", PlaneRANSAC2.DistAvrPointPlane(V2.V), "  meters")
	log.Println("The high score is : ", hs2)
	fmt.Print("\n")
	PlaneLS2 := calculator.PlaneFittingLeastSquare(V2.V)
	fmt.Println("Plane fitted by least square for group of dots V2 : ", PlaneLS2)
	fmt.Println("Average distance between the vertices and the plane is : ", PlaneLS2.DistAvrPointPlane(V2.V), "  meters")
	fmt.Print("\n")
	fmt.Print("\n")

	// fit the vertices in V3 to a plane using the 2 methods
	PlaneRANSAC3, _, hs3 := calculator.PlaneFittingRANSAC(V3.V, 0.005, 0.98,500, 4 * len(V3.V))
	fmt.Println("Plane fitted by RANSAC for group of dots V3 : ", PlaneRANSAC3)
	fmt.Println("Average distance between the vertices and the plane is : ", PlaneRANSAC3.DistAvrPointPlane(V3.V), "  meters")
	log.Println("The high score is : ", hs3)
	fmt.Print("\n")
	PlaneLS3 := calculator.PlaneFittingLeastSquare(V3.V)
	fmt.Println("Plane fitted by least square for group of dots V3 : ", PlaneLS3)
	fmt.Println("Average distance between the vertices and the plane is : ", PlaneLS3.DistAvrPointPlane(V3.V), "  meters")
	fmt.Print("\n")
	fmt.Print("\n")

	// compute the angle formed by the plane from V1 and the plane from V2
	fmt.Println("The angle formed by the plane representing V2 and the one representing V1 (RANSAC) is : ", mymath.VectorsAngle(PlaneRANSAC1.A, PlaneRANSAC1.B, PlaneRANSAC1.C, PlaneRANSAC2.A, PlaneRANSAC2.B, PlaneRANSAC2.C))
	fmt.Print("\n")
	fmt.Println("The angle formed by the plane representing V2 and the one representing V1 (LS) is : ", mymath.VectorsAngle(PlaneLS1.A, PlaneLS1.B, PlaneLS1.C, PlaneLS2.A, PlaneLS2.B, PlaneLS2.C))
	fmt.Print("\n")
	fmt.Print("\n")

	// compute the angle formed by the plane from V1 and the plane from V3
	fmt.Println("The angle formed by the plane representing V3 and the one representing V1 (RANSAC) is : ", mymath.VectorsAngle(PlaneRANSAC1.A, PlaneRANSAC1.B, PlaneRANSAC1.C, PlaneRANSAC3.A, PlaneRANSAC3.B, PlaneRANSAC3.C))
	fmt.Print("\n")
	fmt.Println("The angle formed by the plane representing V3 and the one representing V1 (LS) is : ", mymath.VectorsAngle(PlaneLS1.A, PlaneLS1.B, PlaneLS1.C, PlaneLS3.A, PlaneLS3.B, PlaneLS3.C))
	fmt.Print("\n")
	fmt.Print("\n")

	// compute the angle formed by the plane from V2 and the plane from V3
	fmt.Println("The angle formed by the plane representing V3 and the one representing V2 (RANSAC) is : ", mymath.VectorsAngle(PlaneRANSAC2.A, PlaneRANSAC2.B, PlaneRANSAC2.C, PlaneRANSAC3.A, PlaneRANSAC3.B, PlaneRANSAC3.C))
	fmt.Print("\n")
	fmt.Println("The angle formed by the plane representing V3 and the one representing V2 (LS) is : ", mymath.VectorsAngle(PlaneLS2.A, PlaneLS2.B, PlaneLS2.C, PlaneLS3.A, PlaneLS3.B, PlaneLS3.C))

}

func TestExp2(t *testing.T) {
	// this ply file is composed by 3 surfaces : the cabinet, the shredder and the floor, perpendicular one to another (please check the .png file under the same directory)
	filename := "./data/c_p_g_ofD455_highview/2021-04-16-17:00:24-.ply"

	// Reading the file using adopted library plyfile
	vlist, _ := reader.ReadPLY(filename)

	// get the boundaries of the points on the x and y axes manually
	xmin, xmax, xrange, ymin, ymax, yrange, _, _, _ := calculator.GetBoundaries(vlist)


	// V1 : the surface of the cabinet
	// V1 := calculator.New_z(xmin, xmin + 0.15 * xrange, ymin, ymin + 0.15 * yrange, vlist)
	V1 := calculator.New_z(xmin, xmin + 0.15 * xrange, ymax - 0.5 * yrange, ymax, vlist)
	fmt.Println("Number of points in V1 : ", V1.Zone_count())
	// V2 : the surface of the shredder
	V2 := calculator.New_z(xmax - 0.2 * xrange, xmax - 0.1 * xrange, ymax - 0.35 * yrange, ymax, vlist)
	fmt.Println("Number of points in V2 : ", V2.Zone_count())
	// V3 : the surface of the floor
	V3 := calculator.New_z(xmin + 0.45 * xrange, xmin + 0.65 * xrange, ymin,  ymin + 0.08 * yrange, vlist)
	fmt.Println("Number of points in V3 : ", V3.Zone_count())

	// fit the vertices in V1 to a plane using the RANSAC method
	PlaneRANSAC1, _, hs1 := calculator.PlaneFittingRANSAC(V1.V , 0.005, 0.98,500, 4 * len(V1.V))
	fmt.Println("Plane fitted by RANSAC for group of dots V1 : ", PlaneRANSAC1)
	fmt.Println("Average distance between the vertices and the plane is : ", PlaneRANSAC1.DistAvrPointPlane(V1.V), "  meters")
	fmt.Println("The high score is : ", hs1)
	fmt.Print("\n")
	// fit the vertices in V1 to a plane using the least square method
	PlaneLS1 := calculator.PlaneFittingLeastSquare(V1.V)
	fmt.Println("Plane fitted by least square for group of dots V1 : ", PlaneLS1)
	fmt.Println("Average distance between the vertices and the plane is : ", PlaneLS1.DistAvrPointPlane(V1.V), "  meters")
	fmt.Print("\n")
	fmt.Print("\n")

	// fit the vertices in V2 to a plane using the 2 methods
	PlaneRANSAC2, _, hs2 := calculator.PlaneFittingRANSAC(V2.V, 0.005, 0.98,500, 4 * len(V2.V))
	fmt.Println("Plane fitted by RANSAC for group of dots V2 : ", PlaneRANSAC2)
	fmt.Println("Average distance between the vertices and the plane is : ", PlaneRANSAC2.DistAvrPointPlane(V2.V), "  meters")
	fmt.Println("The high score is : ", hs2)
	fmt.Print("\n")
	PlaneLS2 := calculator.PlaneFittingLeastSquare(V2.V)
	fmt.Println("Plane fitted by least square for group of dots V2 : ", PlaneLS2)
	fmt.Println("Average distance between the vertices and the plane is : ", PlaneLS2.DistAvrPointPlane(V2.V), "  meters")
	fmt.Print("\n")
	fmt.Print("\n")

	// fit the vertices in V3 to a plane using the 2 methods
	PlaneRANSAC3, _, hs3 := calculator.PlaneFittingRANSAC(V3.V, 0.005, 0.98,500, 4 * len(V3.V))
	fmt.Println("Plane fitted by RANSAC for group of dots V3 : ", PlaneRANSAC3)
	fmt.Println("Average distance between the vertices and the plane is : ", PlaneRANSAC3.DistAvrPointPlane(V3.V), "  meters")
	log.Println("Thgit qe high score is : ", hs3)
	fmt.Print("\n")
	PlaneLS3 := calculator.PlaneFittingLeastSquare(V3.V)
	fmt.Println("Plane fitted by least square for group of dots V3 : ", PlaneLS3)
	fmt.Println("Average distance between the vertices and the plane is : ", PlaneLS3.DistAvrPointPlane(V3.V), "  meters")
	fmt.Print("\n")
	fmt.Print("\n")

	// compute the angle formed by the plane from V1 and the plane from V2
	fmt.Println("The angle formed by the plane representing V2 and the one representing V1 (RANSAC) is : ", mymath.VectorsAngle(PlaneRANSAC1.A, PlaneRANSAC1.B, PlaneRANSAC1.C, PlaneRANSAC2.A, PlaneRANSAC2.B, PlaneRANSAC2.C))
	fmt.Print("\n")
	fmt.Println("The angle formed by the plane representing V2 and the one representing V1 (LS) is : ", mymath.VectorsAngle(PlaneLS1.A, PlaneLS1.B, PlaneLS1.C, PlaneLS2.A, PlaneLS2.B, PlaneLS2.C))
	fmt.Print("\n")
	fmt.Print("\n")

	// compute the angle formed by the plane from V1 and the plane from V3
	fmt.Println("The angle formed by the plane representing V3 and the one representing V1 (RANSAC) is : ", mymath.VectorsAngle(PlaneRANSAC1.A, PlaneRANSAC1.B, PlaneRANSAC1.C, PlaneRANSAC3.A, PlaneRANSAC3.B, PlaneRANSAC3.C))
	fmt.Print("\n")
	fmt.Println("The angle formed by the plane representing V3 and the one representing V1 (LS) is : ", mymath.VectorsAngle(PlaneLS1.A, PlaneLS1.B, PlaneLS1.C, PlaneLS3.A, PlaneLS3.B, PlaneLS3.C))
	fmt.Print("\n")
	fmt.Print("\n")

	// compute the angle formed by the plane from V2 and the plane from V3
	fmt.Println("The angle formed by the plane representing V3 and the one representing V2 (RANSAC) is : ", mymath.VectorsAngle(PlaneRANSAC2.A, PlaneRANSAC2.B, PlaneRANSAC2.C, PlaneRANSAC3.A, PlaneRANSAC3.B, PlaneRANSAC3.C))
	fmt.Print("\n")
	fmt.Println("The angle formed by the plane representing V3 and the one representing V2 (LS) is : ", mymath.VectorsAngle(PlaneLS2.A, PlaneLS2.B, PlaneLS2.C, PlaneLS3.A, PlaneLS3.B, PlaneLS3.C))
}


func TestExp3(t *testing.T) {
	// this ply file is composed by 3 surfaces : the cabinet, the shredder and the floor, perpendicular one to another (please check the .png file under the same directory)
	filename := "./data/c_p_g_ofD455_highview/2021-04-16-17:00:24-.ply"

	// Reading the file using adopted library plyfile
	vlist, _ := reader.ReadPLY(filename)

	P := calculator.PlaneSeqRANSAC(vlist, 0.1, 50000, 16000, 200)

	// compute the angle formed by the plane from V1 and the plane from V2
	fmt.Println("The angle formed by the plane representing V2 and the one representing V1 is : ", mymath.VectorsAngle(P[0].A, P[0].B, P[0].C, P[1].A, P[1].B, P[1].C))
	fmt.Print("\n")
	fmt.Print("\n")

	// compute the angle formed by the plane from V1 and the plane from V3
	fmt.Println("The angle formed by the plane representing V3 and the one representing V1 is : ", mymath.VectorsAngle(P[0].A, P[0].B, P[0].C, P[2].A, P[2].B, P[2].C))
	fmt.Print("\n")
	fmt.Print("\n")

	// compute the angle formed by the plane from V2 and the plane from V3
	fmt.Println("The angle formed by the plane representing V3 and the one representing V2 is : ", mymath.VectorsAngle(P[1].A, P[1].B, P[1].C, P[2].A, P[2].B, P[2].C))
}


func TestExp4(t *testing.T) {
	// this ply file is composed by 3 surfaces : the cabinet, the shredder and the floor, perpendicular one to another (please check the .png file under the same directory)
	filename := "./data/c_p_g_ofD455_highview/2021-04-16-17_00_24-.ply"

	// Reading the file using adopted library plyfile
	vlist, _ := reader.ReadPLY(filename)

	P := calculator.PlaneConsecRANSAC(vlist, 0.01, 0.1, 10000, 0.05, 10000, 500, 500)

	// compute all the angles formed by the planes one and another
	for i := 0; i < len(P) - 1; i++ {
		for j := i + 1; j < len(P); j++ {
			fmt.Println("The angle formed by the plane representing V", i, " and the one representing V", j, " is : ", mymath.VectorsAngle(P[i].A, P[i].B, P[i].C, P[j].A, P[j].B, P[j].C))
		}
	}

	// A photo of the desk, more complex than the former one, in which there are boxes, windows, file containers and etc. ...
	filename2 := "./data/desk_static_complex_L515/2021-04-22-12_41_26-.ply"

	// Reading the file using adopted library plyfile
	vlist2, _ := reader.ReadPLY(filename2)

	P2 := calculator.PlaneConsecRANSAC(vlist2, 0.01, 0.1, 10000, 0.05, 10000, 500, 500)

	fmt.Println(len(P2))

	// compute all the angles formed by the planes one and another
	for i := 0; i < len(P2) - 1; i++ {
		for j := i + 1; j < len(P2); j++ {
			fmt.Println("The angle formed by the plane representing V", i, " and the one representing V", j, " is : ", mymath.VectorsAngle(P2[i].A, P2[i].B, P2[i].C, P2[j].A, P2[j].B, P2[j].C))

		}
	}
}

func TestForreport1(t *testing.T) {
	var start time.Time
	var elapsed time.Duration
	no3surface := 0
	anglesfause := 0

	for i := 0; i < 100; i++ {
		// this ply file is composed by 3 surfaces : the cabinet, the shredder and the floor, perpendicular one to another (please check the .png file under the same directory)
		filename := "./data/c_p_g_ofD455_highview/2021-04-16-17_00_24-.ply"



		// Reading the file using adopted library plyfile
		vlist, _ := reader.ReadPLY(filename)

		start = time.Now()

		P := calculator.PlaneConsecRANSAC(vlist, 0.01, 0.1, 10000, 0.1, 10000, 500, 500)

		elapsed += time.Since(start)

		if len(P) != 3 {
			no3surface += 1
		}

		for i := 0; i < len(P) - 1; i++ {
			for j := i + 1; j < len(P); j++ {
				if mymath.VectorsAngle(P[i].A, P[i].B, P[i].C, P[j].A, P[j].B, P[j].C) < 1.5 {
					anglesfause += 1
				}
			}
		}
	}

	fmt.Println("Time used : ", elapsed)
	fmt.Println("Times doesnt detected 3 planes : ", no3surface)
	fmt.Println("Angles < 1.5 : ", anglesfause)
}

func TestForreport2(t *testing.T) {
	var start time.Time
	var elapsed time.Duration
	numSurface := make([]int, 0)

	for i := 0; i < 10; i++ {
		// this ply file is composed by 3 surfaces : the cabinet, the shredder and the floor, perpendicular one to another (please check the .png file under the same directory)
		filename := "./data/office_and_window_complex_D455/2021-04-22-12_08_45-.ply"



		// Reading the file using adopted library plyfile
		vlist, _ := reader.ReadPLY(filename)

		start = time.Now()

		P := calculator.PlaneConsecRANSAC(vlist, 0.01, 0.1, 10000, 0.1, 10000, 500, 500)

		elapsed += time.Since(start)

		numSurface = append(numSurface, len(P))
	}

	fmt.Println("Time used : ", elapsed)
	fmt.Println("Mean num surface : ", mymath.IntMean(numSurface))
	fmt.Println("Variance surface < 1.5 : ", mymath.IntVariance(numSurface, mymath.IntMean(numSurface)))
}