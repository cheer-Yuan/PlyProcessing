package Compressor

import (
	"log"
	"testing"
	"time"
)

func TestNocomp(t *testing.T) {
	t1 := time.Now()

	files := ListDir("../../../output/")

	for _, file := range files {
		Compress("../gzoutput/" + file, "../../../output/" + file, 0)
	}

	timeused := time.Since(t1)
	log.Println("Time used : ", timeused)
}
func TestBestspeed(t *testing.T) {
	t1 := time.Now()

	files := ListDir("../../../output/")

	for _, file := range files {
		Compress("../gzoutput/" + file, "../../../output/" + file, 1)
	}

	timeused := time.Since(t1)
	log.Println("Time used : ", timeused)
}
func TestDefault(t *testing.T) {
	t1 := time.Now()

	files := ListDir("../../../output/")

	for i := 0; i < 5; i++{
		for _, file := range files {
			Compress("../gzoutput/" + file, "../../../output/" + file, -1)
		}
	}


	timeused := time.Since(t1)
	log.Println("Time used : ", timeused)
}
func TestHuffman(t *testing.T) {
	t1 := time.Now()

	files := ListDir("../../../output/")

	for i := 0; i < 5; i++{
		for _, file := range files {
			Compress("../gzoutput/" + file, "../../../output/" + file, -2)
		}
	}


	timeused := time.Since(t1)
	log.Println("Time used : ", timeused)
}
func TestComp9(t *testing.T) {
	t1 := time.Now()

	files := ListDir("../../../output/")

	for i := 0; i < 5; i++{
		for _, file := range files {
			Compress("../gzoutput/" + file, "../../../output/" + file, 9)
		}
	}


	timeused := time.Since(t1)
	log.Println("Time used : ", timeused)
}
func TestComp5(t *testing.T) {
	t1 := time.Now()

	files := ListDir("../../../output/")

	for i := 0; i < 5; i++{
		for _, file := range files {
			Compress("../gzoutput/" + file, "../../../output/" + file, 5)
		}
	}


	timeused := time.Since(t1)
	log.Println("Time used : ", timeused)
}
func TestComp7(t *testing.T) {
	t1 := time.Now()

	files := ListDir("../../../output/")

	for i := 0; i < 5; i++{
		for _, file := range files {
			Compress("../gzoutput/" + file, "../../../output/" + file, 7)
		}
	}


	timeused := time.Since(t1)
	log.Println("Time used : ", timeused)
}
func TestComp3(t *testing.T) {
	t1 := time.Now()

	files := ListDir("../../../output/")

	for i := 0; i < 5; i++{
		for _, file := range files {
			Compress("../gzoutput/" + file, "../../../output/" + file, 3)
		}
	}


	timeused := time.Since(t1)
	log.Println("Time used : ", timeused)
}
func TestComp2(t *testing.T) {
        t1 := time.Now()

        files := ListDir("../../../output/")

        for i := 0; i < 5; i++{
                for _, file := range files {
                        Compress("../gzoutput/" + file, "../../../output/" + file, 2)
                }
        }


        timeused := time.Since(t1)
        log.Println("Time used : ", timeused)
}
func TestComp4(t *testing.T) {
        t1 := time.Now()

        files := ListDir("../../../output/")

        for i := 0; i < 5; i++{
                for _, file := range files {
                        Compress("../gzoutput/" + file, "../../../output/" + file, 4)
                }
        }


        timeused := time.Since(t1)
        log.Println("Time used : ", timeused)
}
func TestComp6(t *testing.T) {
        t1 := time.Now()

        files := ListDir("../../../output/")

        for i := 0; i < 5; i++{
                for _, file := range files {
                        Compress("../gzoutput/" + file, "../../../output/" + file, 6)
                }
        }


        timeused := time.Since(t1)
        log.Println("Time used : ", timeused)
}
func TestComp8(t *testing.T) {
        t1 := time.Now()

        files := ListDir("../../../output/")

        for i := 0; i < 5; i++{
                for _, file := range files {
                        Compress("../gzoutput/" + file, "../../../output/" + file, 8)
                }
        }


        timeused := time.Since(t1)
        log.Println("Time used : ", timeused)
}
