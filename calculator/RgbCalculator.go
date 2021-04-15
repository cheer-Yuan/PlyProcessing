package calculator

type RgbSensor struct {
	lengthPixel, heightPixel int		// resolution of the camera
	pixelReference           int		// numbers of pixels captured by the example object in dimension 1
	pixelsErr                int		// error tolerated of the number of pixels
	distanceReference		 float64	// distance between the camera and the example object
	focusLength              float64	// focus length
	objectReference          float64	// the size of the example object
}

/* NumPixelsOccupy estimate how many pixels will an object occupy, given the size and the distance of the object to evalute and A the structure containing the sensor data */
func NumPixelsOccupy(sizeEstmating float64, distEstmating float64, sensor *RgbSensor) (float64, float64) {
	imageRef := sensor.focusLength * sensor.objectReference / sensor.distanceReference
	pixelDistMin := imageRef / float64((sensor.pixelReference - sensor.pixelsErr))
	pixelDistMax := imageRef / float64((sensor.pixelReference + sensor.pixelsErr))
	sizeImageEstimating := sensor.focusLength * sizeEstmating / distEstmating
	pixelsMin := sizeImageEstimating / pixelDistMin
	pixelsMax := sizeImageEstimating / pixelDistMax
	return pixelsMin, pixelsMax
}