package main

import "testing"

func BenchmarkFask(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		fast()
	}
}

//      9	 129528933 ns/op	       0 B/op	       0 allocs/op

func BenchmarkSlow(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		slow()
	}
}

//      5	 243398140 ns/op	       0 B/op	       0 allocs/op

func BenchmarkFast2(b *testing.B) {
	b.ResetTimer()
	b.Run("fast", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			fast()
		}
	})

	b.Run("slow", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			slow()
		}
	})

}
