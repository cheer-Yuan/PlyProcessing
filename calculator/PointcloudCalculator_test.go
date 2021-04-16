package calculator

import (
	"../reader"
	"fmt"
	"math/rand"
	"testing"
	"../mymath"
	"time"
)

func TestFittingMethods(t *testing.T) {
	filename := "../data/angle_of_board_and_paper_card_on_table/2021-03-31-09:56:18-.ply"
	vlist, _ := reader.ReadPLY(filename)

	// v1 represents the paper board
	V1 := New_z(-0.4, -0.2, -0.23, 0.13, vlist)
	fmt.Println("Number of points in V1 : ", V1.Zone_count())
	V2 := New_z(0.15, 0.35, -0.33, 0.13, vlist)
	fmt.Println("Number of points in V2 : ", V2.Zone_count())
	// V3 represents the bottom board of the desk
	V3 := New_z(-0.2, 0.15, -0.21, 0.11, vlist)
	fmt.Println("Number of points in V3 : ", V3.Zone_count())
	// V4 represents the surface of the desk
	V4 := New_z(0.15, 0.4, 0.25, 999, vlist)
	fmt.Println("Number of points in V4 : ", V4.Zone_count())
	// V5 : small subset of V3
	V5 := New_z(-0.05, 0, -0.03, 0.02, vlist)
	fmt.Println("Number of points in V5 : ", V5.Zone_count())
	// V6 : small subset of V4
	V6 := New_z(0.25, 0.3, 0.35, 0.4, vlist)
	fmt.Println("Number of points in V6 : ", V6.Zone_count())
	fmt.Print("\n")

	//PlaneRANSAC1, _ := PlaneFittingRANSAC(V1.V, 0.1, 0.9,500, len(V1.V))
	//fmt.Println("Plane fitted by RANSAC for group of dots V1 : ", PlaneRANSAC1)
	//fmt.Println("Average distance between the vertices and the plane is : ", PlaneRANSAC1.DistAvrPointPlane(V1.V))
	//fmt.Print("\N")
	//PlaneLS1 := PlaneFittingLeastSquare(V1.V)
	//fmt.Println("Plane fitted by least square for group of dots V1 : ", PlaneLS1)
	//fmt.Println("Average distance between the vertices and the plane is : ", PlaneLS1.DistAvrPointPlane(V1.V))
	//fmt.Print("\N")
	//fmt.Print("\N")
	//
	//PlaneRANSAC3, _ := PlaneFittingRANSAC(V3.V, 0.1, 0.9,500, len(V3.V))
	//fmt.Println("Plane fitted by RANSAC for group of dots V3 : ", PlaneRANSAC3)
	//fmt.Println("Average distance between the vertices and the plane is : ", PlaneRANSAC3.DistAvrPointPlane(V3.V))
	//fmt.Print("\N")
	//PlaneLS3 := PlaneFittingLeastSquare(V3.V)
	//fmt.Println("Plane fitted by least square for group of dots V3 : ", PlaneLS3)
	//fmt.Println("Average distance between the vertices and the plane is : ", PlaneLS3.DistAvrPointPlane(V3.V))
	//fmt.Print("\N")
	//fmt.Print("\N")
	//
	//PlaneRANSAC4, _ := PlaneFittingRANSAC(V4.V, 0.1, 0.9,500, len(V4.V))
	//fmt.Println("Plane fitted by RANSAC for group of dots V4 : ", PlaneRANSAC4)
	//fmt.Println("Average distance between the vertices and the plane is : ", PlaneRANSAC4.DistAvrPointPlane(V4.V))
	//fmt.Print("\N")
	//PlaneLS4 := PlaneFittingLeastSquare(V4.V)
	//fmt.Println("Plane fitted by least square for group of dots V4 : ", PlaneLS4)
	//fmt.Println("Average distance between the vertices and the plane is : ", PlaneLS4.DistAvrPointPlane(V4.V))
	//fmt.Print("\N")
	//fmt.Print("\N")

	PlaneRANSAC5, _, _ := PlaneFittingRANSAC(V5.V, 0.01, 0.98,500, len(V5.V))
	fmt.Println("Plane fitted by RANSAC for group of dots V5 : ", PlaneRANSAC5)
	fmt.Println("Average distance between the vertices and the plane is : ", PlaneRANSAC5.DistAvrPointPlane(V5.V))
	fmt.Print("\n")
	PlaneLS5 := PlaneFittingLeastSquare(V5.V)
	fmt.Println("Plane fitted by least square for group of dots V5 : ", PlaneLS5)
	fmt.Println("Average distance between the vertices and the plane is : ", PlaneLS5.DistAvrPointPlane(V5.V))
	fmt.Print("\n")
	fmt.Print("\n")

	PlaneRANSAC6, _, _ := PlaneFittingRANSAC(V6.V, 0.01, 0.98,500, len(V6.V))
	fmt.Println("Plane fitted by RANSAC for group of dots V6 : ", PlaneRANSAC6)
	fmt.Println("Average distance between the vertices and the plane is : ", PlaneRANSAC6.DistAvrPointPlane(V6.V))
	fmt.Print("\n")
	PlaneLS6 := PlaneFittingLeastSquare(V6.V)
	fmt.Println("Plane fitted by least square for group of dots V6 : ", PlaneLS6)
	fmt.Println("Average distance between the vertices and the plane is : ", PlaneLS6.DistAvrPointPlane(V6.V))
	fmt.Print("\n")
	fmt.Print("\n")

	//fmt.Println("The angle formed by the plane representing V1 and the one representing V3 (LS) is : ", mymath.VectorsAngle(PlaneLS1.A, PlaneLS1.B, PlaneLS1.C, PlaneLS3.A, PlaneLS3.B, PlaneLS3.C))
	//fmt.Print("\N")
	//fmt.Println("The angle formed by the plane representing V1 and the one representing V3 (RANSAC) is : ", mymath.VectorsAngle(PlaneRANSAC3.A, PlaneRANSAC3.B, PlaneRANSAC3.C, PlaneRANSAC1.A, PlaneRANSAC1.B, PlaneRANSAC1.C))
	//fmt.Print("\N")
	//fmt.Print("\N")
	//
	//fmt.Println("The angle formed by the plane representing V3 and the one representing V4 (LS) is : ", mymath.VectorsAngle(PlaneLS3.A, PlaneLS3.B, PlaneLS3.C, PlaneLS4.A, PlaneLS4.B, PlaneLS4.C))
	//fmt.Print("\N")
	//fmt.Println("The angle formed by the plane representing V3 and the one representing V4 (RANSAC) is : ", mymath.VectorsAngle(PlaneRANSAC3.A, PlaneRANSAC3.B, PlaneRANSAC3.C, PlaneRANSAC4.A, PlaneRANSAC4.B, PlaneRANSAC4.C))
	//fmt.Print("\N")
	//fmt.Print("\N")
	//
	fmt.Println("The angle formed by the plane representing V5 and the one representing V6 (LS) is : ", mymath.VectorsAngle(PlaneLS6.A, PlaneLS6.B, PlaneLS6.C, PlaneLS5.A, PlaneLS5.B, PlaneLS5.C))
	fmt.Print("\n")
	fmt.Println("The angle formed by the plane representing V5 and the one representing V6 (RANSAC) is : ", mymath.VectorsAngle(PlaneRANSAC6.A, PlaneRANSAC6.B, PlaneRANSAC6.C, PlaneRANSAC5.A, PlaneRANSAC5.B, PlaneRANSAC5.C))
}





func ChooseSparseVertex(group *ZoneCritical, nmax int) (int, int, int) {
	Dbuff := 0.
	V1, V2, V3 := 0, 0, 0
	for i := 0; i < nmax; i++ {
		p1 := rand.Intn(group.N)
		p2 := rand.Intn(group.N)
		p3 := rand.Intn(group.N)
		Dsum := DistVertex(group.V[p1], group.V[p2]) + DistVertex(group.V[p1], group.V[p3]) + DistVertex(group.V[p2], group.V[p3])
		if Dsum > Dbuff {
			Dbuff = Dsum
			V1, V2, V3 = p1, p2, p3
		}
	}
	return V1, V2, V3
}

func TestAngle3Points(t *testing.T) {
	filename := "../data/cabinet_paperdestroyer_ground/2021-04-15-12:10:03-.ply"

	// Reading the file using adopted library plyfile
	vlist, _ := reader.ReadPLY(filename)

	xmin, xmax, ymin, ymax := 0., 0., 0., 0.

	for _, vertex := range vlist {
		if vertex.Ply_x < xmin {
			xmin = vertex.Ply_x
		} else if vertex.Ply_x > xmax {
			xmax = vertex.Ply_x
		}
		if vertex.Ply_y < ymin{
			ymin = vertex.Ply_y
		} else if vertex.Ply_y > ymax {
			ymax = vertex.Ply_y
		}
	}
	fmt.Println(xmin, xmax, ymin, ymax)

	// V1 : the surface of the cabinet
	V1 := New_z(0.2, 0.31, -0.15, 0, vlist)
	fmt.Println("Number of points in V1 : ", V1.Zone_count())
	// V2 : the surface of the paperdestroyer
	V2 := New_z(-0.37, -0.15, -0.15, 0, vlist)
	fmt.Println("Number of points in V2 : ", V2.Zone_count())
	// V3 : the surface of the floor
	V3 := New_z(-999, 999, 0.35, 0.45, vlist)
	fmt.Println("Number of points in V3 : ", V3.Zone_count())

	rand.Seed(time.Now().Unix())

	Buff1, Buff2, Buff3 := ChooseSparseVertex(V1, 2000)
	P1 := New_p_by_vertices(V1.V[Buff1], V1.V[Buff2], V1.V[Buff3])
	fmt.Println("Plane fitted by RANSAC for group of dots V1 : ", P1)
	fmt.Println("Average distance between the vertices and the plane is : ", P1.DistAvrPointPlane(V1.V))
	fmt.Print("\n")
	fmt.Print("\n")

	Buff1, Buff2, Buff3 = ChooseSparseVertex(V2, 2000)
	P2 := New_p_by_vertices(V2.V[Buff1], V2.V[Buff2], V2.V[Buff3])
	fmt.Println("Plane fitted by RANSAC for group of dots V2 : ", P2)
	fmt.Println("Average distance between the vertices and the plane is : ", P2.DistAvrPointPlane(V2.V))
	fmt.Print("\n")
	fmt.Print("\n")

	Buff1, Buff2, Buff3 = ChooseSparseVertex(V3, 2000)
	P3 := New_p_by_vertices(V3.V[Buff1], V3.V[Buff2], V3.V[Buff3])
	fmt.Println("Plane fitted by RANSAC for group of dots V3 : ", P3)
	fmt.Println("Average distance between the vertices and the plane is : ", P3.DistAvrPointPlane(V3.V))
	fmt.Print("\n")
	fmt.Print("\n")


	fmt.Println("The angle formed by the plane representing V2 and the one representing V1 is : ", mymath.VectorsAngle(P1.A, P1.B, P1.C, P2.A, P2.B, P2.C))
	fmt.Print("\n")
	fmt.Print("\n")

	fmt.Println("The angle formed by the plane representing V3 and the one representing V1 is : ", mymath.VectorsAngle(P1.A, P1.B, P1.C, P3.A, P3.B, P3.C))
	fmt.Print("\n")
	fmt.Print("\n")

	fmt.Println("The angle formed by the plane representing V2 and the one representing V3 is : ", mymath.VectorsAngle(P3.A, P3.B, P3.C, P2.A, P2.B, P2.C))
}
