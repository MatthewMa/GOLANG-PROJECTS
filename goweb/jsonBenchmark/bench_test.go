package main

import "testing"

func BenchmarkUnmarshall(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Unmarshall("post.json")
	}
}

func BenchmarkDecode(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Decode("post.json")
	}
}
