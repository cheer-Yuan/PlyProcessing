package test

import (
	"dataprocessing/calculator"
	"dataprocessing/reader"
	"testing"
)

func BenchmarkPlaneMonoFittingRANSAC(b *testing.B) {
	b.StopTimer()
	vlist, _ := reader.ReadPLYMono("../../../../output/2021-05-12-10:38:33:730-Points.ply")


	b.StartTimer()
	for i := 0; i < 20; i++ {
		_, _, _ = calculator.PlaneMonoFittingRANSAC(vlist, 0.01, 0.9, 10, 100)
	}
}

func BenchmarkPlaneMonoFittingRANSAC32(b *testing.B) {
	b.StopTimer()
	vlist, _ := reader.ReadPLYMono32("../../../../output/2021-05-12-10:38:33:730-Points.ply")


	b.StartTimer()
	for i := 0; i < 20; i++ {
		_, _, _ = calculator.PlaneMonoFittingRANSAC32(vlist, 0.01, 0.9, 10, 100)
	}
}