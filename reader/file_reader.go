package reader

// Reader, first version, supports only tythe particular ply file generated by L515 under particular circumstances

import (
	"encoding/binary"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
)

type Header struct {
	Vertex_num, Face_num int
}

// struct of vertices, using float64, for data processing
type VertexFormat struct {
	Ply_x, Ply_y, Ply_z float64
	ply_r, ply_g, ply_b int
}

// struct of vertices, using float32, at the reader's side
type Vertex struct {
	Ply_x, Ply_y, Ply_z float32
	ply_r, ply_g, ply_b int
}

type Face struct {
	vertex_1, vertex_2, vertex_3 int
}

type File struct {
	header *Header
	vertex *[]Vertex
	face *[]Face
}

/** Convert the slice of string into a integer, by the correspondance in the ASCII table
 * @param Slice of string
 * @return An integer
 */
func Conv_uint8_ascii(byte []byte) int {
	converted_int, err := strconv.Atoi(string(byte))
	if err != nil { log.Fatal(err) }
	return converted_int
}

/** Convert the slice of string into a 32-bit-float
 * @param Slice of string
 * @return The converted float
 */
func Conv_float32(byte []byte) float32 {
	bit := binary.LittleEndian.Uint32(byte)
	float := math.Float32frombits(bit)
	return float
}

/** Get the information in the header
 * @return number of vertex. number of faces
 */
func (header *Header) Header_Val() (int, int) {
	return header.Vertex_num, header.Face_num
}


func Read_Ply(filename string) (Header, []Vertex, []Face) {

	// Open the file
	file, err_F := os.Open(filename)
	if err_F != nil { log.Fatal(err_F) }
	defer file.Close()

	// Instantiate the header struct
	var header Header

	// Create a buffer to hold the data
	buff := make([]byte, 6)

	// Location of the first parameter, the number of vertex, at 99th byte from the beginning of the file
	file.Seek(98, 0)

	// Reading
	_, _ = file.Read(buff)
	// Converting
	header.Vertex_num = Conv_uint8_ascii(buff)

	// Location of the second parameter, the number of vertex, at 99th byte from the beginning of the file
	file.Seek(131, 1)

	// Reading
	_, _ = file.Read(buff)
	// Converting
	header.Face_num = Conv_uint8_ascii(buff)




	// Instantiate the body data struct
	vertex := make([]Vertex, header.Vertex_num)


	// Locate the beginning of the vertex data
	file.Seek(51, 1)

	// Count the numbers of vertex read
	num_v_read := 0

	// Create buffer
	buff_f32 := make([]byte, 4)
	buff_uchar := make([]byte, 1)

	// Read the vertex data
	for i := 0; i < header.Vertex_num; i++ {

		_, _ = file.Read(buff_f32)
		vertex[i].Ply_x = Conv_float32(buff_f32)
		_, _ = file.Read(buff_f32)
		vertex[i].Ply_y = Conv_float32(buff_f32)
		_, _ = file.Read(buff_f32)
		vertex[i].Ply_z = Conv_float32(buff_f32)


		_, _ = file.Read(buff_uchar)
		vertex[i].ply_r = int(buff_uchar[0])
		_, _ = file.Read(buff_uchar)
		vertex[i].ply_g = int(buff_uchar[0])
		_, _ = file.Read(buff_uchar)
		vertex[i].ply_b = int(buff_uchar[0])

		num_v_read++
	}

	fmt.Println("Vertex read : ", num_v_read)




	// Count the numbers of surfaces read
	num_f_read := 0

	// Instantiate the surface data struct
	face := make([]Face, header.Face_num)

	// Read the surface data
	for i := 0; i < header.Face_num; i++ {

		_, _ = file.Read(buff_uchar)
		_, _ = file.Read(buff_f32)
		face[i].vertex_1 = int(binary.LittleEndian.Uint32(buff_f32))
		_, _ = file.Read(buff_f32)
		face[i].vertex_2 = int(binary.LittleEndian.Uint32(buff_f32))
		_, _ = file.Read(buff_f32)
		face[i].vertex_3 = int(binary.LittleEndian.Uint32(buff_f32))

		num_f_read++
	}

	fmt.Println("Surfaces read : ", num_f_read)



	return header, vertex, face
}
