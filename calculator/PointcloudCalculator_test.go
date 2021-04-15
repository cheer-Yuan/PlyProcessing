package calculator

import (
	"../mymath"
	"../reader"
	"fmt"
	"testing"
)

func TestNew_n(t *testing.T) {
	testval := New_p_by_normal(0, 0, 0, 0, 0, 0)
	resultval := &Plane{A: 0, B: 0, C: 0, D: -0}
	if *testval == *resultval {
		t.Log("ok")
	} else {
		t.Errorf("expected %+v got %+v  \n", resultval, testval)
	}
}

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

	// V5 : small subset of



	fmt.Println("Number of points in V4 : ", V4.Zone_count())
	fmt.Print("\n")

	PlaneRANSAC1, _ := PlaneFittingRANSAC(V1.v, 0.1, 0.9,500, len(V1.v))
	fmt.Println("Plane fitted by RANSAC for group of dots V1 : ", PlaneRANSAC1)
	fmt.Println("Average distance between the vertices and the plane is : ", PlaneRANSAC1.DistAvrPointPlane(V1.v))
	fmt.Print("\n")

	PlaneLS1 := PlaneFittingLeastSquare(V1.v)
	fmt.Println("Plane fitted by least square for group of dots V1 : ", PlaneLS1)
	fmt.Println("Average distance between the vertices and the plane is : ", PlaneLS1.DistAvrPointPlane(V1.v))
	fmt.Print("\n")
	fmt.Print("\n")

	PlaneRANSAC3, _ := PlaneFittingRANSAC(V3.v, 0.1, 0.9,500, len(V3.v))
	fmt.Println("Plane fitted by RANSAC for group of dots V3 : ", PlaneRANSAC3)
	fmt.Println("Average distance between the vertices and the plane is : ", PlaneRANSAC3.DistAvrPointPlane(V3.v))
	fmt.Print("\n")

	PlaneLS3 := PlaneFittingLeastSquare(V3.v)
	fmt.Println("Plane fitted by least square for group of dots V3 : ", PlaneLS3)
	fmt.Println("Average distance between the vertices and the plane is : ", PlaneLS3.DistAvrPointPlane(V3.v))
	fmt.Print("\n")
	fmt.Print("\n")

	PlaneRANSAC4, _ := PlaneFittingRANSAC(V4.v, 0.1, 0.9,500, len(V4.v))
	fmt.Println("Plane fitted by RANSAC for group of dots V4 : ", PlaneRANSAC4)
	fmt.Println("Average distance between the vertices and the plane is : ", PlaneRANSAC4.DistAvrPointPlane(V4.v))
	fmt.Print("\n")

	PlaneLS4 := PlaneFittingLeastSquare(V4.v)
	fmt.Println("Plane fitted by least square for group of dots V4 : ", PlaneLS4)
	fmt.Println("Average distance between the vertices and the plane is : ", PlaneLS4.DistAvrPointPlane(V4.v))
	fmt.Print("\n")
	fmt.Print("\n")
	
	fmt.Println("The angle formed by the plane representing V1 and the one representing V3 (LS) is : ", mymath.VectorsAngle(PlaneLS1.A, PlaneLS1.B, PlaneLS1.C, PlaneLS3.A, PlaneLS3.B, PlaneLS3.C))
	fmt.Print("\n")
	fmt.Println("The angle formed by the plane representing V1 and the one representing V3 (RANSAC) is : ", mymath.VectorsAngle(PlaneRANSAC3.A, PlaneRANSAC3.B, PlaneRANSAC3.C, PlaneRANSAC1.A, PlaneRANSAC1.B, PlaneRANSAC1.C))
	fmt.Print("\n")
	fmt.Print("\n")

	fmt.Println("The angle formed by the plane representing V3 and the one representing V4 (LS) is : ", mymath.VectorsAngle(PlaneLS3.A, PlaneLS3.B, PlaneLS3.C, PlaneLS4.A, PlaneLS4.B, PlaneLS4.C))
	fmt.Print("\n")
	fmt.Println("The angle formed by the plane representing V3 and the one representing V4 (RANSAC) is : ", mymath.VectorsAngle(PlaneRANSAC3.A, PlaneRANSAC3.B, PlaneRANSAC3.C, PlaneRANSAC4.A, PlaneRANSAC4.B, PlaneRANSAC4.C))

}