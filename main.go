package main

import (
	"./reader"
	"fmt"
	"./calculator"
	"./mymath"
)

func main() {
	// this ply file is composed by 3 surfaces : the cabinet, the shredder and the floor, perpendicular one to another (please check the .png file under the same directory)
	filename := "./data/cabinet_paperdestroyer_ground/2021-04-15-12:10:03-.ply"

	// Reading the file using adopted library plyfile
	vlist, _ := reader.ReadPLY(filename)

	xmin, xmax, ymin, ymax := 0., 0., 0., 0.

	// set the critical zone of the points manually
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
	
	// V1 : the surface of the shredder
	V1 := calculator.New_z(0.2, 0.31, -0.15, 0, vlist)
	fmt.Println("Number of points in V1 : ", V1.Zone_count())
	// V2 : the surface of the shredder
	V2 := calculator.New_z(-0.37, -0.15, -0.15, 0, vlist)
	fmt.Println("Number of points in V2 : ", V2.Zone_count())
	// V3 : the surface of the floor
	V3 := calculator.New_z(-999, 999, 0.4, 0.45, vlist)
	fmt.Println("Number of points in V3 : ", V3.Zone_count())

	// fit the vertices in V1 to a plane using the RANSAC method
	PlaneRANSAC1, _, _ := calculator.PlaneFittingRANSAC(V1.V , 0.02, 0.98,500, 4 * len(V1.V))
	fmt.Println("Plane fitted by RANSAC for group of dots V1 : ", PlaneRANSAC1)
	fmt.Println("Average distance between the vertices and the plane is : ", PlaneRANSAC1.DistAvrPointPlane(V1.V), "  meters")
	fmt.Print("\n")
	// fit the vertices in V1 to a plane using the least square method
	PlaneLS1 := calculator.PlaneFittingLeastSquare(V1.V)
	fmt.Println("Plane fitted by least square for group of dots V1 : ", PlaneLS1)
	fmt.Println("Average distance between the vertices and the plane is : ", PlaneLS1.DistAvrPointPlane(V1.V), "  meters")
	fmt.Print("\n")
	fmt.Print("\n")

	// fit the vertices in V2 to a plane using the 2 methods
	PlaneRANSAC2, _, _ := calculator.PlaneFittingRANSAC(V2.V, 0.02, 0.98,500, 4 * len(V2.V))
	fmt.Println("Plane fitted by RANSAC for group of dots V2 : ", PlaneRANSAC2)
	fmt.Println("Average distance between the vertices and the plane is : ", PlaneRANSAC2.DistAvrPointPlane(V2.V), "  meters")
	fmt.Print("\n")
	PlaneLS2 := calculator.PlaneFittingLeastSquare(V2.V)
	fmt.Println("Plane fitted by least square for group of dots V2 : ", PlaneLS2)
	fmt.Println("Average distance between the vertices and the plane is : ", PlaneLS2.DistAvrPointPlane(V2.V), "  meters")
	fmt.Print("\n")
	fmt.Print("\n")

	// fit the vertices in V3 to a plane using the 2 methods
	PlaneRANSAC3, _, _ := calculator.PlaneFittingRANSAC(V3.V, 0.02, 0.98,500, 4 * len(V3.V))
	fmt.Println("Plane fitted by RANSAC for group of dots V3 : ", PlaneRANSAC3)
	fmt.Println("Average distance between the vertices and the plane is : ", PlaneRANSAC3.DistAvrPointPlane(V3.V), "  meters")
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