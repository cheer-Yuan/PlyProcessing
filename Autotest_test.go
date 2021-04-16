package main

import (
	"./calculator"
	"./mymath"
	"./reader"
	"fmt"
	"io"
	"log"
	"os"
	"testing"
)

func TestAuto(t *testing.T) {
	file, err_F := os.OpenFile("./Test/autotest.txt", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err_F != nil { log.Fatal(err_F) }
	defer file.Close()
	mw := io.MultiWriter(os.Stdout, file)
	log.SetOutput(mw)

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


	log.Println(xmin, xmax, ymin, ymax)

	// V1 : the surface of the shredder
	V1 := calculator.New_z(0.2, 0.31, -0.15, 0, vlist)
	log.Println("Number of points in V1 : ", V1.Zone_count())
	// V2 : the surface of the shredder
	V2 := calculator.New_z(-0.37, -0.15, -0.15, 0, vlist)
	log.Println("Number of points in V2 : ", V2.Zone_count())
	// V3 : the surface of the floor
	V3 := calculator.New_z(-999, 999, 0.4, 0.45, vlist)
	log.Println("Number of points in V3 : ", V3.Zone_count())
	// V4 : subset og V1 composed by 2000 elements randomly picked
	// V5 : subset og V2 composed by 2000 elements randomly picked
	V4, V5 := make([]reader.VertexFormat, 2000), make([]reader.VertexFormat, 2000)
	for i := 0; i < 2000; i++ {
		V4 = append(V4, V1.V[i * 10])
		V5 = append(V5, V2.V[i * 21])
	}
	log.Println("Number of points in V4 : ", len(V4))
	log.Println("Number of points in V5 : ", len(V5))
	log.Print("\n")
	log.Print("\n")


	// fit the vertices in V1 to a plane using the least square method
	PlaneLS1 := calculator.PlaneFittingLeastSquare(V1.V)
	log.Println("Plane fitted by least square for group of dots V1 : ", PlaneLS1)
	log.Println("Average distance between the vertices and the plane is : ", PlaneLS1.DistAvrPointPlane(V1.V), "  meters")
	log.Print("\n")
	PlaneLS2 := calculator.PlaneFittingLeastSquare(V2.V)
	log.Println("Plane fitted by least square for group of dots V2 : ", PlaneLS2)
	log.Println("Average distance between the vertices and the plane is : ", PlaneLS2.DistAvrPointPlane(V2.V), "  meters")
	log.Print("\n")
	PlaneLS3 := calculator.PlaneFittingLeastSquare(V3.V)
	log.Println("Plane fitted by least square for group of dots V3 : ", PlaneLS3)
	log.Println("Average distance between the vertices and the plane is : ", PlaneLS3.DistAvrPointPlane(V3.V), "  meters")
	log.Print("\n")
	PlaneLS4 := calculator.PlaneFittingLeastSquare(V4)
	log.Println("Plane fitted by least square for group of dots V4 : ", PlaneLS4)
	log.Println("Average distance between the vertices and the plane is : ", PlaneLS4.DistAvrPointPlane(V4), "  meters")
	log.Print("\n")
	PlaneLS5 := calculator.PlaneFittingLeastSquare(V5)
	log.Println("Plane fitted by least square for group of dots V5 : ", PlaneLS5)
	log.Println("Average distance between the vertices and the plane is : ", PlaneLS5.DistAvrPointPlane(V5), "  meters")
	log.Print("\n")
	log.Print("\n")
	
	
	log.Println("The angle formed by the plane representing V2 and the one representing V1 (LS) is : ", mymath.VectorsAngle(PlaneLS1.A, PlaneLS1.B, PlaneLS1.C, PlaneLS2.A, PlaneLS2.B, PlaneLS2.C))
	log.Print("\n")
	log.Println("The angle formed by the plane representing V3 and the one representing V1 (LS) is : ", mymath.VectorsAngle(PlaneLS1.A, PlaneLS1.B, PlaneLS1.C, PlaneLS3.A, PlaneLS3.B, PlaneLS3.C))
	log.Print("\n")
	log.Println("The angle formed by the plane representing V3 and the one representing V2 (LS) is : ", mymath.VectorsAngle(PlaneLS2.A, PlaneLS2.B, PlaneLS2.C, PlaneLS3.A, PlaneLS3.B, PlaneLS3.C))
	log.Print("\n")
	log.Println("The angle formed by the plane representing V4 and the one representing V5 (LS) is : ", mymath.VectorsAngle(PlaneLS4.A, PlaneLS4.B, PlaneLS4.C, PlaneLS5.A, PlaneLS5.B, PlaneLS5.C))
	log.Print("\n")
	log.Println("The angle formed by the plane representing V3 and the one representing V4 (LS) is : ", mymath.VectorsAngle(PlaneLS4.A, PlaneLS4.B, PlaneLS4.C, PlaneLS3.A, PlaneLS3.B, PlaneLS3.C))
	log.Print("\n")
	log.Println("The angle formed by the plane representing V3 and the one representing V5 (LS) is : ", mymath.VectorsAngle(PlaneLS5.A, PlaneLS5.B, PlaneLS5.C, PlaneLS3.A, PlaneLS3.B, PlaneLS3.C))
	log.Print("\n")
	log.Print("\n")


	for RecMax := 300000; RecMax > 1000; RecMax = RecMax / 2 {
		for TL := 0.01; TL < 0.5; TL = TL * 2 {
			for SM := 0.98; SM > 0.85; SM = SM - 0.3 {
				fmt.Println("RANSAC Parameters : ", TL, SM, 500, RecMax)
				log.Print("\n")
				// fit the vertices in V1 to a plane using the RANSAC method
				PlaneRANSAC1, _, hs1 := calculator.PlaneFittingRANSAC(V1.V , TL, SM,500, RecMax)
				log.Println("Plane fitted by RANSAC for group of dots V1 : ", PlaneRANSAC1)
				log.Println("Average distance between the vertices and the plane is : ", PlaneRANSAC1.DistAvrPointPlane(V1.V), "  meters")
				log.Println("The high score is : ", hs1)
				log.Print("\n")
				
				// fit the vertices in V2 to a plane using the 2 methods
				PlaneRANSAC2, _, hs2 := calculator.PlaneFittingRANSAC(V2.V, TL, SM,500, RecMax)
				log.Println("Plane fitted by RANSAC for group of dots V2 : ", PlaneRANSAC2)
				log.Println("Average distance between the vertices and the plane is : ", PlaneRANSAC2.DistAvrPointPlane(V2.V), "  meters")
				log.Println("The high score is : ", hs2)
				log.Print("\n")

				// fit the vertices in V3 to a plane using the 2 methods
				PlaneRANSAC3, _, hs3 := calculator.PlaneFittingRANSAC(V3.V, TL, SM,500, 4 * len(V3.V))
				log.Println("Plane fitted by RANSAC for group of dots V3 : ", PlaneRANSAC3)
				log.Println("Average distance between the vertices and the plane is : ", PlaneRANSAC3.DistAvrPointPlane(V3.V), "  meters")
				log.Println("The high score is : ", hs3)
				log.Print("\n")

				// compute the angle formed by the plane from V1 and the plane from V2
				log.Println("The angle formed by the plane representing V2 and the one representing V1 (RANSAC) is : ", mymath.VectorsAngle(PlaneRANSAC1.A, PlaneRANSAC1.B, PlaneRANSAC1.C, PlaneRANSAC2.A, PlaneRANSAC2.B, PlaneRANSAC2.C))
				log.Print("\n")
				
				// compute the angle formed by the plane from V1 and the plane from V3
				log.Println("The angle formed by the plane representing V3 and the one representing V1 (RANSAC) is : ", mymath.VectorsAngle(PlaneRANSAC1.A, PlaneRANSAC1.B, PlaneRANSAC1.C, PlaneRANSAC3.A, PlaneRANSAC3.B, PlaneRANSAC3.C))
				log.Print("\n")

				// compute the angle formed by the plane from V2 and the plane from V3
				log.Println("The angle formed by the plane representing V3 and the one representing V2 (RANSAC) is : ", mymath.VectorsAngle(PlaneRANSAC2.A, PlaneRANSAC2.B, PlaneRANSAC2.C, PlaneRANSAC3.A, PlaneRANSAC3.B, PlaneRANSAC3.C))
				log.Print("\n")


			}
		}
	}

	for RecMax := 10000; RecMax > 1000; RecMax = RecMax - 2000 {
		for TL := 0.01; TL < 0.5; TL = TL * 2 {
			for SM := 0.98; SM > 0.85; SM = SM - 0.3 {
				fmt.Println("RANSAC Parameters : ", TL, SM, 500, RecMax)
				log.Print("\n")
				// fit the vertices in V4 to a plane using the RANSAC method
				PlaneRANSAC4, _, hs4 := calculator.PlaneFittingRANSAC(V4 , TL, SM,500, RecMax)
				log.Println("Plane fitted by RANSAC for group of dots V4 : ", PlaneRANSAC4)
				log.Println("Average distance between the vertices and the plane is : ", PlaneRANSAC4.DistAvrPointPlane(V4), "  meters")
				log.Println("The high score is : ", hs4)
				log.Print("\n")

				// fit the vertices in V5 to a plane using the 2 methods
				PlaneRANSAC5, _, hs5 := calculator.PlaneFittingRANSAC(V5, TL, SM,500, RecMax)
				log.Println("Plane fitted by RANSAC for group of dots V5 : ", PlaneRANSAC5)
				log.Println("Average distance between the vertices and the plane is : ", PlaneRANSAC5.DistAvrPointPlane(V5), "  meters")
				log.Println("The high score is : ", hs5)
				log.Print("\n")

				// fit the vertices in V3 to a plane using the 2 methods
				PlaneRANSAC3, _, hs3 := calculator.PlaneFittingRANSAC(V3.V, TL, SM,500, 4 * len(V3.V))
				log.Println("Plane fitted by RANSAC for group of dots V3 : ", PlaneRANSAC3)
				log.Println("Average distance between the vertices and the plane is : ", PlaneRANSAC3.DistAvrPointPlane(V3.V), "  meters")
				log.Println("The high score is : ", hs3)
				log.Print("\n")

				// compute the angle formed by the plane from V4 and the plane from V5
				log.Println("The angle formed by the plane representing V5 and the one representing V4 (RANSAC) is : ", mymath.VectorsAngle(PlaneRANSAC4.A, PlaneRANSAC4.B, PlaneRANSAC4.C, PlaneRANSAC5.A, PlaneRANSAC5.B, PlaneRANSAC5.C))
				log.Print("\n")

				// compute the angle formed by the plane from V4 and the plane from V3
				log.Println("The angle formed by the plane representing V3 and the one representing V4 (RANSAC) is : ", mymath.VectorsAngle(PlaneRANSAC4.A, PlaneRANSAC4.B, PlaneRANSAC4.C, PlaneRANSAC3.A, PlaneRANSAC3.B, PlaneRANSAC3.C))
				log.Print("\n")

				// compute the angle formed by the plane from V5 and the plane from V3
				log.Println("The angle formed by the plane representing V3 and the one representing V5 (RANSAC) is : ", mymath.VectorsAngle(PlaneRANSAC5.A, PlaneRANSAC5.B, PlaneRANSAC5.C, PlaneRANSAC3.A, PlaneRANSAC3.B, PlaneRANSAC3.C))
				log.Print("\n")


			}
		}
	}


	}