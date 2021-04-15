package main

import (
	"./reader"
)

//Test if a vertex belongs to a plane
func main() {
	//// Reading the file using the original reader
	//header, vertex, face := reader.Read_Ply(filename)

	filename := "./data/angle_of_board_and_paper_card_on_table/2021-03-31-09:56:18-.ply"

	// Reading the file using adopted library plyfile
	vlist, flist := reader.ReadPLY(filename)






/*	// Prepare outputs of the log
	file, err_F := os.OpenFile("./Test/test.txt", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err_F != nil { log.Fatal(err_F) }
	defer file.Close()
	mw := io.MultiWriter(os.Stdout, file)
	log.SetOutput(mw)

	// Create sets of critic zones
	V1 := calculator.New_z(-0.4, -0.2, -0.23, 0.13, vertex, header)
	fmt.Println("Number of points in V1 : ", V1.Zone_count())
	V2 := calculator.New_z(0.15, 0.35, -0.33, 0.13, vertex, header)
	fmt.Println("Number of points in V2 : ", V2.Zone_count())
	V3 := calculator.New_z(-0.2, 0.15, -0.21, 0.11, vertex, header)
	fmt.Println("Number of points in V3 : ", V3.Zone_count())

	// Creating 3 surfaces from the critic zones
	p1 := calculator.New_p_by_points(V1)
	p2 := calculator.New_p_by_points(V2)

	//// Set the bias
	bias := 0.1

	fmt.Println("The percentage of the points in V1 that fit the plane with the bias of ", bias, " m.")
	fmt.Println("\t p1 : ", p1.IfZonebelongto(V1, bias))
	fmt.Println("\t p2 : ", p2.IfZonebelongto(V1, bias))
	fmt.Println("The percentage of the points in V2 that fit the plane with the bias of ", bias, " m.")
	fmt.Println("\t p1 : ", p1.IfZonebelongto(V2, bias))
	fmt.Println("\t p2 : ", p2.IfZonebelongto(V2, bias))
	fmt.Println("The percentage of the points in V3 that fit the plane with the bias of ", bias, " m.")
	fmt.Println("\t p1 : ", p1.IfZonebelongto(V3, bias))
	fmt.Println("\t p2 : ", p2.IfZonebelongto(V3, bias))*/

	//
	//// Create the plane
	//p := calculator.New_n(99, 99, -1.73, 0, 0, -1)


	//dist, belon := p.Ifbelongto(vertex[V].Ply_x, vertex[V].Ply_y, vertex[V].Ply_z, bias)
	//log.Println("Testing with the file <", filename, ">")
	//log.Println("Testing point: ", vertex[V], ". and the plane ", p, ". Distance tolerated : ", bias, " m.")
	//log.Println("Distance between the point and the plane : ", dist , "m\t", "Whether the point belongs to the plane :", belon)
	//log.Print("\n\n")
}