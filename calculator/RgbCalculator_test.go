package calculator

import (
	"testing"
)

const (
	SIZE_ESTMATE = 5000
)

// in mm, compatible for the color module of D455 since they share the same focal length and sensor resolution
func NewL515() *RgbSensor {
	return &RgbSensor{
		lengthPixel: 1920,
		heightPixel: 1080,
		pixelsErr: 15,
		pixelReference: 710,
		distanceReference: 700,
		focusLength: 0.00188,
		objectReference: 381,
	}
}

func TestL515(t *testing.T) {
	L515 := NewL515()
	var min, max float64

	// calculation A vehicle of 3000 mm from 50, 100, 200, 400 meters
	dist := float64(50000)
	min, max = NumPixelsOccupy(3000, dist, L515)
	println("A vehicle of ", float64(SIZE_ESTMATE) / float64(1000), " meters at the distance of ", float64(dist) / float64(1000), " meters will match from ", min, " to ", max, " pixels.")
	dist = float64(100000)
	min, max = NumPixelsOccupy(3000, dist, L515)
	println("A vehicle of ", float64(SIZE_ESTMATE) / float64(1000), " meters at the distance of ", float64(dist) / float64(1000), " meters will match from ", min, " to ", max, " pixels.")
	dist = float64(200000)
	min, max = NumPixelsOccupy(3000, dist, L515)
	println("A vehicle of ", float64(SIZE_ESTMATE) / float64(1000), " meters at the distance of ", float64(dist) / float64(1000), " meters will match from ", min, " to ", max, " pixels.")
	dist = float64(400000)
	min, max = NumPixelsOccupy(3000, dist, L515)
	println("A vehicle of ", float64(SIZE_ESTMATE) / float64(1000), " meters at the distance of ", float64(dist) / float64(1000), " meters will match from ", min, " to ", max, " pixels.")
}

