package main

import (
	"datacompressing/Compressor"
	"dataprocessing/calculator"
	"dataprocessing/reader"
	"fmt"
	"io"
	"log"
	"os"
	"runtime"
	"sync"
	"time"
)

// test for parallel structure
func main() {
	// read the configuration
	Config := reader.ReadConfig("./config.json")

	// settings for goroutines in the
	runtime.GOMAXPROCS(Config.NumCore)

	// open the file and set the logger where we export the results
	file, _ := os.OpenFile(Config.OutputAddr + "results.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	defer file.Close()

	// set the log writer
	writer := io.Writer(file)
	log.SetOutput(writer)
	log.SetFlags(log.Lmicroseconds)

	// scan the folder to get the file list, the position in the config file should end with a '/'
	files := calculator.ListDir(Config.InputAddr)

	// start the timer
	t1 := time.Now()

	// containers for the result
	var valAngles []float32		// angles formed by the planes detected, for all files
	var totalPlanes []int		// number of plane detected in each file
	var mutexList sync.Mutex

	// use a channel to limit the number of goroutines
	wg := sync.WaitGroup{}
	wg.Add(len(files))
	controlChannel := make(chan bool, Config.MaxGort)

	// iterate among all files
	for i := 0; i < len(files); i++ {
		controlChannel <- true
		go func(filename string) {
			// read a file
			vlist, _ := reader.ReadPLYMono32(Config.InputAddr + filename)

			// analyse the planes
			Planes, numPlanes, _ := calculator.PlaneMonoConsecRANSAC32(vlist, Config.Parameters.MaxDistance, Config.Parameters.MinScoreRANSAC, Config.Parameters.MinVertexPlane, Config.Parameters.MaxAnglePlane, Config.Parameters.MaxVertexQuit, Config.Parameters.MaxIteration, Config.Parameters.NumBatch)

			// compute angle values
			valangles := calculator.AngleOfPlanes32(Planes)
			valAngles = append(valangles, valangles...)
			totalPlanes = append(totalPlanes, numPlanes)

			// critical area : make sure only 1 goroutine access the log file
			mutexList.Lock()
			log.Println("The number of planes detected : ", numPlanes)
			log.Println("Angles formed by the planes : ", valangles)
			for _, j := range Planes {
				log.Println(j)
			}
			log.Println("\n")
			mutexList.Unlock()

			// compressing
			if Config.CompressLev != 0 {
				_ = Compressor.Compress(Config.OutputAddr + filename, Config.InputAddr + filename, Config.CompressLev)
			}


			wg.Done()
			<- controlChannel
		}(files[i])

	}

	wg.Wait()
	log.Println("\nList of plane numbers : ")
	log.Println(totalPlanes)
	log.Println("\nList of all angles : ")
	log.Println(valAngles)
	fmt.Println("Time used :  ", time.Since(t1))
}