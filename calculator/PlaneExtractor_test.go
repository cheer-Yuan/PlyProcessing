package calculator

import (
	"../reader"
	"fmt"
	"testing"
)

func TestVoxel(t *testing.T)  {
	// this ply file is composed by 3 surfaces : the cabinet, the shredder and the floor, perpendicular one to another (please check the .png file under the same directory)
	filename := "../data/cabinet_paperdestroyer_ground/2021-04-15-12:10:03-.ply"

	// Reading the file using adopted library plyfile
	vlist, _ := reader.ReadPLY(filename)


	newVertices := VoxelDownsample(vlist, 0.005)
	fmt.Println("The number of vertices before down sampling : ", len(vlist))

	fmt.Println("The number of vertices after down sampling : ", len(newVertices))
}
