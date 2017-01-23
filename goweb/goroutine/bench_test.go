package main

import "testing"

func BenchmarkPrintTask1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		printTask1()
	}
}

func BenchmarkPrintTask2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		printTask2()
	}
}
