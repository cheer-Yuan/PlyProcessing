package reader

// Use the adopted library plyfile

import (
	"../plyfile"
	"encoding/binary"
	"fmt"
	"unsafe"
)

func VertexFormatize(vertex plyfile.Vertex) VertexFormat {
	var VertexBuff VertexFormat
	VertexBuff.Ply_x = float64(vertex.X)
	VertexBuff.Ply_y = float64(vertex.Y)
	VertexBuff.Ply_z = float64(vertex.Z)
	VertexBuff.ply_r = int(vertex.R)
	VertexBuff.ply_g = int(vertex.G)
	VertexBuff.ply_b = int(vertex.B)

	return VertexBuff
}

func ReadPLY(filename string) ([]VertexFormat, []plyfile.Face) {
	fmt.Println("Reading PLY file ...")
	var vertices []VertexFormat
	var faces []plyfile.Face

	// open the PLY file for reading
	cplyfile, elem_names := plyfile.PlyOpenForReading(filename)

	// read each element
	for _, name := range elem_names {

		// get element description
		plist, num_elems, num_props := plyfile.PlyGetElementDescription(cplyfile, name)

		// print the name of the element, for debugging
		fmt.Println("element", name, num_elems)

		// create a list to store all vertices
		vlist := make([]plyfile.Vertex, num_elems)

		// create a list to hold all face elements
		flist := make([]plyfile.FaceReading, num_elems)

		if name == "vertex" {

			// grab vertex elements
			for i := 0; i < num_elems; i++ {
				plyfile.PlyGetElement(cplyfile, &vlist[i], unsafe.Sizeof(plyfile.Vertex{}) - 1)
				vertices = append(vertices, VertexFormatize(vlist[i]))
			}

		} else if name == "face" {

			// grab face elements
			for i := 0; i < num_elems; i++ {
				plyfile.PlyGetElement(cplyfile, &flist[i], unsafe.Sizeof(plyfile.FaceReading{}))

				// extract integers from binary listst and export to a slice of faces
				FaceBuf := plyfile.Face{int(binary.LittleEndian.Uint32(flist[i].Vert1[:])), int(binary.LittleEndian.Uint32(flist[i].Vert2[:])), int(binary.LittleEndian.Uint32(flist[i].Vert3[:]))}
				faces = append(faces, FaceBuf)
			}
		}

		for i := 0; i < num_props; i++ {
			fmt.Println("property", plist[i].Name)
		}
	}
	// close the PLY file
	plyfile.PlyClose(cplyfile)

	return vertices, faces
}