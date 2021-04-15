which defines the//Test if a vertex belongs to a plane
func main() {
	// Routine of the ply file in which we will pick points
	filename := ""

	// Reading the file
	header, vertex, face := reader.Read_Ply(filename)

	fmt.Println(header.Vertex_num, header.Face_num)
	fmt.Println(face[0])

	file, err_F := os.OpenFile("./Test/test.txt", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err_F != nil { log.Fatal(err_F) }
	defer file.Close()
	mw := io.MultiWriter(os.Stdout, file)
	log.SetOutput(mw)

	// Pick a random vertex
	rand.Seed(time.Now().Unix())
	V := rand.Intn(header.Vertex_num)

	// Set the bias
	bias := 0.02

	p := calculator.New_n(99, 99, -1.73, 0, 0, -1)
	dist, belon := p.Ifbelongto(vertex[V].Ply_x, vertex[V].Ply_y, vertex[V].Ply_z, bias)
	log.Println("Testing with the file <", filename, ">")
	log.Println("Testing point: ", vertex[V], ". and the plane ", p, ". Distance tolerated : ", bias, " m.")
	log.Println("Distance between the point and the plane : ", dist , "m\t", "Whether the point belongs to the plane :", belon)
	log.Print("\n\n")
}
