package plyfile

import (
	"fmt"
	"os"
	"testing"
	"unsafe"
)

/* Exported Fields Note: All struct fields must be exported (capitalized) for use in the plyfile package! */

type Vertex struct {
	X, Y, Z float32
}

type Face struct {
	Intensity byte
	Nverts    byte
	//Verts 		*int32 // ptr to memory location
	Verts [8]byte // maximum size array
}

type VertexIndices [4]int32

func GenerateVertexFaceData() (verts []Vertex, faces []Face, vertex_indices []VertexIndices) {
	verts = make([]Vertex, 8)
	faces = make([]Face, 6)

	verts[0] = Vertex{0.0, 0.0, 0.0}
	verts[1] = Vertex{1.0, 0.0, 0.0}
	verts[2] = Vertex{1.0, 1.0, 0.0}
	verts[3] = Vertex{0.0, 1.0, 0.0}
	verts[4] = Vertex{0.0, 0.0, 1.0}
	verts[5] = Vertex{1.0, 0.0, 1.0}
	verts[6] = Vertex{1.0, 1.0, 1.0}
	verts[7] = Vertex{0.0, 1.0, 1.0}

	/* To support arbitrary size lists, we build two lists: one of the element in question and one of the arbitrary size list we wish to embed in the element.
	Then, we store the memory location of the arbitrary size list into the element in question as a byte array.
	This isn't the greatest implementation from a Go perspective, but it works well enough as long as we keep the two list variables together (otherwise the arbitrary sized list will be garbage collected).
	*/

	vertex_indices = make([]VertexIndices, 6)
	vertex_indices[0] = VertexIndices{0, 1, 2, 3}
	vertex_indices[1] = VertexIndices{7, 6, 5, 4}
	vertex_indices[2] = VertexIndices{0, 4, 5, 1}
	vertex_indices[3] = VertexIndices{1, 5, 6, 2}
	vertex_indices[4] = VertexIndices{2, 6, 7, 3}
	vertex_indices[5] = VertexIndices{3, 7, 4, 0}

	faces[0] = Face{'\001', 4, [8]byte{}}
	faces[1] = Face{'\004', 4, [8]byte{}}
	faces[2] = Face{'\010', 4, [8]byte{}}
	faces[3] = Face{'\020', 4, [8]byte{}}
	faces[4] = Face{'\144', 4, [8]byte{}}
	faces[5] = Face{'\377', 4, [8]byte{}}
	for i := 0; i < 6; i++ {
		copy(faces[i].Verts[:], PointerToByteSlice(uintptr(unsafe.Pointer(&vertex_indices[i]))))
	}

	return verts, faces, vertex_indices
}

/* Testing Functions */

func SetPlyProperties() (vert_props []PlyProperty, face_props []PlyProperty) {
	vert_props = make([]PlyProperty, 3)
	vert_props[0] = PlyProperty{"x", PLY_FLOAT, PLY_FLOAT, int(unsafe.Offsetof(Vertex{}.X)), 0, 0, 0, 0}
	vert_props[1] = PlyProperty{"y", PLY_FLOAT, PLY_FLOAT, int(unsafe.Offsetof(Vertex{}.Y)), 0, 0, 0, 0}
	vert_props[2] = PlyProperty{"z", PLY_FLOAT, PLY_FLOAT, int(unsafe.Offsetof(Vertex{}.Z)), 0, 0, 0, 0}

	face_props = make([]PlyProperty, 2)
	face_props[0] = PlyProperty{"intensity", PLY_UCHAR, PLY_UCHAR, int(unsafe.Offsetof(Face{}.Intensity)), 0, 0, 0, 0}
	face_props[1] = PlyProperty{"vertex_indices", PLY_INT, PLY_INT, int(unsafe.Offsetof(Face{}.Verts)), 1, PLY_UCHAR, PLY_UCHAR, int(unsafe.Offsetof(Face{}.Nverts))}

	// PLT_INT : int16

	return vert_props, face_props

}

/* TestWritePly tests writing a PLY file using the cplyfile function for creating a new file and transparently handling the file pointer. */
func TestWritePly(t *testing.T) {
	elem_names := make([]string, 2)
	elem_names[0] = "vertex"
	elem_names[1] = "face"

	var version float32

	fmt.Println("Writing PLY file 'test.ply'...")

	cplyfile := PlyOpenForWriting("./Test/test1.ply", len(elem_names), elem_names, PLY_BINARY_LE, &version)

	/* Note that we don't need a variable for vertex_indices, but we do need to return vertex_indices. Otherwise, the garbage collector will remove them once GenerateVertexFaceData() returns. */
	verts, faces, _ := GenerateVertexFaceData()
	vert_props, face_props := SetPlyProperties()

	// Describe vertex properties
	PlyElementCount(cplyfile, "vertex", len(verts))
	PlyDescribeProperty(cplyfile, "vertex", vert_props[0])
	PlyDescribeProperty(cplyfile, "vertex", vert_props[1])
	PlyDescribeProperty(cplyfile, "vertex", vert_props[2])

	// Describe face properties
	PlyElementCount(cplyfile, "face", len(faces))
	PlyDescribeProperty(cplyfile, "face", face_props[0])
	PlyDescribeProperty(cplyfile, "face", face_props[1])

	// Add a comment and an object information field
	PlyPutComment(cplyfile, "go author: Alex Baden, c author: Greg Turk")
	PlyPutObjInfo(cplyfile, "random information")

	// Finish writing header
	PlyHeaderComplete(cplyfile)

	// Setup and write vertex elements
	PlyPutElementSetup(cplyfile, "vertex")
	for _, vertex := range verts {
		PlyPutElement(cplyfile, vertex)
	}

	// Setup and write face elements
	PlyPutElementSetup(cplyfile, "face")
	for _, face := range faces {
		PlyPutElementFace(cplyfile, face)
	}

	// close the PLY file
	PlyClose(cplyfile)

	fmt.Println("Wrote PLY file.")
}

/* TestReadPLY tests reading a ply file from disk using the cplyfile function for opening the file. */
func TestReadPLY(t *testing.T) {
	fmt.Println("Reading PLY file 'test.ply'...")

	// setup properties
	vert_props, face_props := SetPlyProperties()

	// open the PLY file for reading
	cplyfile, elem_names := PlyOpenForReading("./Test/test1.ply")

	// print what we found out about the file
	fmt.Printf("version: %f\n", cplyfile.version)
	fmt.Printf("file_type: %d\n", cplyfile.file_type)

	// read each element
	for _, name := range elem_names {

		// get element description
		plist, num_elems, num_props := PlyGetElementDescription(cplyfile, name)

		// print the name of the element, for debugging
		fmt.Println("element", name, num_elems)

		if name == "vertex" {

			// create a list to store all vertices
			vlist := make([]Vertex, num_elems)

			/* set up for getting vertex elements
			   specifically, we are ensuring the 3 desirable properties of a vertex (x,,z) are returned.
			*/
			PlyGetProperty(cplyfile, name, vert_props[0])
			PlyGetProperty(cplyfile, name, vert_props[1])
			PlyGetProperty(cplyfile, name, vert_props[2])

			// grab vertex elements
			for i := 0; i < num_elems; i++ {
				PlyGetElement(cplyfile, &vlist[i], unsafe.Sizeof(Vertex{}))
				fmt.Println(i)
				fmt.Println( &vlist[i])
				// print out vertex for debugging
				fmt.Printf("vertex: %g %g %g\n", vlist[i].X, vlist[i].Y, vlist[i].Z)
			}
		} else if name == "face" {
			// create a list to hold all face elements
			flist := make([]Face, num_elems)

			/* set up for getting face elements (See above) */
			PlyGetProperty(cplyfile, name, face_props[0])
			PlyGetProperty(cplyfile, name, face_props[1])

			// grab face elements
			for i := 0; i < num_elems; i++ {
				PlyGetElement(cplyfile, &flist[i], unsafe.Sizeof(Face{}))
				fmt.Println(i)
				fmt.Println( &flist[i])
				// print out faces for debugging
				fmt.Printf("face: %d, list = ", flist[i].Intensity)

				/* Here we handle arbitrary sized arrays. We first convert the byte slice storing the location of the C memory to a pointer. Next, we read from C memory space, creating a byte slice, then convert the byte slice to a int32 slice using the ReadPLYListInt32 function. */
				listptr := ByteSliceToPointer(flist[i].Verts[:])
				list :=
					ReadPLYListInt32(listptr, int(flist[i].Nverts))

				for j := 0; j < int(flist[i].Nverts); j++ {
					fmt.Printf("%d ", list[j])
				}
				fmt.Printf("\n")

			}

		}

		for i := 0; i < num_props; i++ {
			fmt.Println("property", plist[i].Name)
		}

	}

	// grab and print comments in the file
	comments := PlyGetComments(cplyfile)
	for _, comment := range comments {
		fmt.Println("comment =", comment)
	}

	// grab and print object information
	objinfo := PlyGetObjInfo(cplyfile)
	for _, text := range objinfo {
		fmt.Println("obj_info = ", text)
	}

	// close the PLY file
	PlyClose(cplyfile)

}

/* TestWritePlyFilePointer tests writing a PLY file using a go generated file pointer. */
func TestWritePlyFilePointer(t *testing.T) {
	elem_names := make([]string, 2)
	elem_names[0] = "vertex"
	elem_names[1] = "face"

	var version float32

	fmt.Println("Writing PLY file 'test.ply'...")

	// open file
	file, err := os.Create("./Test/test2.ply")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	cplyfile := PlyUseExistingForWriting(file, len(elem_names), elem_names, PLY_ASCII, &version)

	/* Note that we don't need a variable for vertex_indices, but we do need to return vertex_indices. Otherwise, the garbage collector will remove them once GenerateVertexFaceData() returns. */
	verts, faces, _ := GenerateVertexFaceData()
	vert_props, face_props := SetPlyProperties()

	// Describe vertex properties
	PlyElementCount(cplyfile, "vertex", len(verts))
	PlyDescribeProperty(cplyfile, "vertex", vert_props[0])
	PlyDescribeProperty(cplyfile, "vertex", vert_props[1])
	PlyDescribeProperty(cplyfile, "vertex", vert_props[2])

	// Describe face properties
	PlyElementCount(cplyfile, "face", len(faces))
	PlyDescribeProperty(cplyfile, "face", face_props[0])
	PlyDescribeProperty(cplyfile, "face", face_props[1])

	// Add a comment and an object information field
	PlyPutComment(cplyfile, "go author: Alex Baden, c author: Greg Turk")
	PlyPutObjInfo(cplyfile, "random information")

	// Finish writing header
	PlyHeaderComplete(cplyfile)

	// Setup and write vertex elements
	PlyPutElementSetup(cplyfile, "vertex")
	for _, vertex := range verts {
		PlyPutElement(cplyfile, vertex)
	}

	// Setup and write face elements
	PlyPutElementSetup(cplyfile, "face")
	for _, face := range faces {
		PlyPutElementFace(cplyfile, face)
	}

	// close the PLY file
	PlyClose(cplyfile)

	fmt.Println("Wrote PLY file.")
}

/* TestReadPLY tests reading a ply file from disk using the cplyfile function for opening the file. */
func TestReadPLYFilePointer(t *testing.T) {
	fmt.Println("Reading PLY file 'test.ply'...")

	// setup properties
	vert_props, face_props := SetPlyProperties()

	// open the PLY file for reading
	cplyfile, elem_names := PlyOpenForReading("./Test/test2.ply")

	// print what we found out about the file
	fmt.Printf("version: %f\n", cplyfile.version)
	fmt.Printf("file_type: %d\n", cplyfile.file_type)

	// read each element
	for _, name := range elem_names {

		// get element description
		plist, num_elems, num_props := PlyGetElementDescription(cplyfile, name)

		// print the name of the element, for debugging
		fmt.Println("element", name, num_elems)

		if name == "vertex" {

			// create a list to store all vertices
			vlist := make([]Vertex, num_elems)

			/* set up for getting vertex elements
			   specifically, we are ensuring the 3 desirable properties of a vertex (x,,z) are returned.
			*/
			PlyGetProperty(cplyfile, name, vert_props[0])
			PlyGetProperty(cplyfile, name, vert_props[1])
			PlyGetProperty(cplyfile, name, vert_props[2])

			// grab vertex elements
			for i := 0; i < num_elems; i++ {
				PlyGetElement(cplyfile, &vlist[i], unsafe.Sizeof(Vertex{}))

				// print out vertex for debugging
				fmt.Printf("vertex: %g %g %g\n", vlist[i].X, vlist[i].Y, vlist[i].Z)

			}
		} else if name == "face" {
			// create a list to hold all face elements
			flist := make([]Face, num_elems)

			/* set up for getting face elements (See above) */
			PlyGetProperty(cplyfile, name, face_props[0])
			PlyGetProperty(cplyfile, name, face_props[1])

			// grab face elements
			for i := 0; i < num_elems; i++ {
				PlyGetElement(cplyfile, &flist[i], unsafe.Sizeof(Face{}))

				// print out faces for debugging
				fmt.Printf("face: %d, list = ", flist[i].Intensity)

				/* Here we handle arbitrary sized arrays. We first convert the byte slice storing the location of the C memory to a pointer. Next, we read from C memory space, creating a byte slice, then convert the byte slice to a int32 slice using the ReadPLYListInt32 function. */
				listptr := ByteSliceToPointer(flist[i].Verts[:])
				list :=
					ReadPLYListInt32(listptr, int(flist[i].Nverts))

				for j := 0; j < int(flist[i].Nverts); j++ {
					fmt.Printf("%d ", list[j])
				}
				fmt.Printf("\n")

			}

		}

		for i := 0; i < num_props; i++ {
			fmt.Println("property", plist[i].Name)
		}

	}

	// grab and print comments in the file
	comments := PlyGetComments(cplyfile)
	for _, comment := range comments {
		fmt.Println("comment =", comment)
	}

	// grab and print object information
	objinfo := PlyGetObjInfo(cplyfile)
	for _, text := range objinfo {
		fmt.Println("obj_info = ", text)
	}

	// close the PLY file
	PlyClose(cplyfile)
}
