package calculator
//
//import (
//	"dataprocessing/mymath"
//	"dataprocessing/reader"
//)
//
///* ColorList Creates and returns a high contrast colors' list according to Paul Green-Armytage, in <A Colour Alphabet and the Limits of Colour Coding>
// * @return The matrix of the colors. Each line contains a group of rgb value of one color
// */
//func ColorList() [26][3]int {
//	list := [26][3]int {
//		{240,163,255},{0,117,220},{153,63,0},{76,0,92},{25,25,25},{0,92,49},{43,206,72},{255,204,153},{128,128,128},{148,255,181},{143,124,0},{157,204,0},{194,0,136},{0,51,128},{255,164,5},{255,168,187},{66,102,0},{255,0,16},{94,241,242},{0,153,143},{224,255,102},{116,10,255},{153,0,0},{255,255,128},{255,255,0},{255,80,5},
//	}
//	return list
//}
//
//func PaintWithIndex(vlist []reader.VertexFormat, color [3]int) {
//	for _, vIndex := range vlist {
//		vIndex.Ply_r = color[0]
//		vIndex.Ply_g = color[1]
//		vIndex.Ply_b = color[2]
//
//	}
//}
//
//func PaintWithLeastDist(vlist []reader.VertexFormat, P []Plane, colors [26][3]int) {
//	for vIndex, vertex := range vlist {
//		distMin := 100.
//		color := 0
//		for colorIndex, plane := range P {
//			distBuff := mymath.DistPointPlane(vertex.Ply_x, vertex.Ply_y, vertex.Ply_z, plane.A, plane.B, plane.C, plane.D)
//			if distBuff < distMin {
//				distMin = distBuff
//				color = colorIndex
//			}
//		}
//		vlist[vIndex].Ply_r = colors[color][0]
//		vlist[vIndex].Ply_g = colors[color][1]
//		vlist[vIndex].Ply_b = colors[color][2]
//	}
//}